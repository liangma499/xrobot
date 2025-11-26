package solana

import (
	"context"
	"xbase/log"
	"xbase/utils/xconv"

	"fmt"
	"xrobot/internal/cryptocurrencies/solana/internal"
	associatedtokenaccount "xrobot/internal/cryptocurrencies/solana/internal/programs/associated-token-account"
	"xrobot/internal/cryptocurrencies/solana/internal/programs/system"
	"xrobot/internal/cryptocurrencies/solana/internal/programs/token"
	"xrobot/internal/cryptocurrencies/solana/internal/rpc"
	"xrobot/internal/xtypes"

	"github.com/shopspring/decimal"
)

const (
	accountBalance   = 0.001
	gasFeelimit      = uint64(5100)
	gasMinitFeelimit = uint64(5100)
	gasCretelimit    = uint64(2100000)
)

func (sol *solanaInfo) SolCreateAndSendTrx(privateKeyStr string, toAddress, contractAddress string, amount uint64, places int8) (*TransactionRet, error) {
	if sol.client == nil {
		return nil, fmt.Errorf("sol.client is err")
	}
	if contractAddress == "" || contractAddress == "-" || contractAddress == "11111111111111111111111111111111" {
		return sol.doSolCreateAndSendTrxSystem(privateKeyStr, toAddress, amount, places)
	} else {
		return sol.doSolCreateAndSendTrxMint(privateKeyStr, toAddress, contractAddress, amount, places)
	}
}
func (sol *solanaInfo) doSolCreateAndSendTrxSystem(privateKeyStr string, toAddress string, amount uint64, places int8) (*TransactionRet, error) {
	if sol.client == nil {
		return nil, fmt.Errorf("sol.client is err")
	}
	// 设置发送者地址和私钥
	senderPrivateKey := internal.MustPrivateKeyFromBase58(privateKeyStr)

	fromPublicKey := senderPrivateKey.PublicKey()
	// 检查发送者账户余额
	balanceResp, err := sol.client.GetBalance(context.Background(), fromPublicKey, rpc.CommitmentConfirmed)
	if err != nil {
		return nil, fmt.Errorf("failed to getBalance: %v", err)
	}
	//用户余额
	balance := balanceResp.Value
	fromADdressStr := fromPublicKey.String()
	availableBalance := xtypes.CoefficientInt64(decimal.NewFromFloat(accountBalance), places).BigInt().Uint64()
	ret := &TransactionRet{
		From:     fromADdressStr,
		To:       toAddress,
		Contract: "-",
		IsCreate: false,
	}
	if availableBalance > balance {
		return ret, fmt.Errorf("insufficient funds:fromAddress:%s balance %d,availableBalance:%d lamports", fromADdressStr, balance, availableBalance)
	}
	// 确保余额足够支付转账和费用
	if balance < amount+availableBalance+gasFeelimit {
		return ret, fmt.Errorf("insufficient funds:fromAddress:%s balance %d,availableBalance:%d, required %d lamports", fromADdressStr, balance, availableBalance, amount+gasFeelimit)
	}
	recentBlockhash, err := sol.client.GetLatestBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return ret, fmt.Errorf("failed getRecentBlockhash: %v", err)
	}
	// 接收人地址
	recipientPublicKey := internal.MustPublicKeyFromBase58(toAddress)
	ret.To = toAddress
	// 创建转账指令
	transferInstruction := system.NewTransferInstruction(
		amount,
		fromPublicKey,
		recipientPublicKey,
	).Build()

	// 创建交易
	tx, err := internal.NewTransaction(
		[]internal.Instruction{transferInstruction},
		recentBlockhash.Value.Blockhash,
		internal.TransactionPayer(fromPublicKey),
	)
	if err != nil {
		return ret, fmt.Errorf("failed to create transaction: %v", err)
	}

	// 使用发起人的私钥签名交易
	_, err = tx.Sign(
		func(key internal.PublicKey) *internal.PrivateKey {
			if fromPublicKey.Equals(key) {
				return &senderPrivateKey
			}
			return nil
		},
	)
	if err != nil {
		return ret, fmt.Errorf("failed to sign transaction: %v", err)
	}

	// 广播交易
	txSignature, err := sol.client.SendTransaction(context.Background(), tx)
	if err != nil {
		return ret, fmt.Errorf("failed to send transaction: %v", err)
	}
	ret.TxID = txSignature.String()
	return ret, nil
}
func (sol *solanaInfo) GetTokenAccountBalance(address, contractAddress string) (uint64, error) {
	publicKey := internal.MustPublicKeyFromBase58(address)
	mintAddress := internal.MustPublicKeyFromBase58(contractAddress)
	// 接收人的 Token 账户地址（ATA）
	tokenAccount, err := sol.doGetAssociatedTokenAddress(publicKey, mintAddress)
	if err != nil {
		return 0, fmt.Errorf("failed to get receiver token account: %v", err)
	}
	return sol.doGetTokenAccountBalanceAccount(tokenAccount)
}
func (sol *solanaInfo) doGetTokenAccountBalanceAccount(account internal.PublicKey) (uint64, error) {
	if sol.client == nil {
		return 0, fmt.Errorf("sol.client is err")
	}
	// 检查发送者账户余额
	balanceResp, err := sol.client.GetTokenAccountBalance(context.Background(), account, rpc.CommitmentConfirmed)
	if err != nil {
		return 0, fmt.Errorf("failed to getBalance: %v", err)
	}
	if balanceResp.Value == nil {
		return 0, fmt.Errorf("balanceResp.Value is nil")
	}
	amount, err := decimal.NewFromString(balanceResp.Value.Amount)
	if err != nil {
		return 0, err
	}
	//decimals := xconv.Int8(balanceResp.Value.Decimals)
	return amount.BigInt().Uint64(), nil
}
func (sol *solanaInfo) doSolCreateAndSendTrxMint(privateKeyStr string, toAddress, contractAddress string, amount uint64, places int8) (*TransactionRet, error) {
	if sol.client == nil {
		return nil, fmt.Errorf("sol.client is err")
	}
	// 设置发送者地址和私钥
	senderPrivateKey := internal.MustPrivateKeyFromBase58(privateKeyStr)
	// SPL Token 的 Mint 地址（例如 USDT）
	mintAddress := internal.MustPublicKeyFromBase58(contractAddress)

	fromPublicKey := senderPrivateKey.PublicKey()

	// 发起人的 Token 账户地址（ATA - Associated Token Account）
	senderTokenAccount, err := sol.doGetAssociatedTokenAddress(fromPublicKey, mintAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get sender token account: %v", err)
	}
	// 接收人地址
	recipientPublicKey := internal.MustPublicKeyFromBase58(toAddress)

	// 接收人的 Token 账户地址（ATA）
	receiverTokenAccount, err := sol.doGetAssociatedTokenAddress(recipientPublicKey, mintAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get receiver token account: %v", err)
	}
	// 检查发送者账户余额
	balanceResp, err := sol.client.GetBalance(context.Background(), fromPublicKey, rpc.CommitmentConfirmed)
	if err != nil {
		return nil, fmt.Errorf("failed to getBalance: %v", err)
	}
	//用户余额
	balanceSol := balanceResp.Value

	//用户余额
	balance, err := sol.doGetTokenAccountBalanceAccount(senderTokenAccount)
	if err != nil {
		return nil, fmt.Errorf("failed to getBalance: %v", err)
	}

	fromADdressStr := fromPublicKey.String()
	ret := &TransactionRet{
		From:     fromADdressStr,
		To:       toAddress,
		Contract: "-",
		IsCreate: false,
	}
	availableBalance := xtypes.CoefficientInt64(decimal.NewFromFloat(accountBalance), places).BigInt().Uint64()
	if availableBalance > balance {
		return ret, fmt.Errorf("insufficient funds:fromAddress:%s balance %d,availableBalance:%d lamports", fromADdressStr, balance, availableBalance)
	}
	if balanceSol < gasMinitFeelimit {
		return ret, fmt.Errorf("insufficient funds:fromAddress:%s balanceSol %d,availableBalance:%d lamports", fromADdressStr, balanceSol, availableBalance)

	}

	haveAccount, err := sol.doGetTokenAccountsByOwner(recipientPublicKey, mintAddress)
	if err != nil {
		return ret, fmt.Errorf("failed doGetTokenAccountsByOwner: %v", err)
	}
	trxs := make([]internal.Instruction, 0)

	if !haveAccount {
		//创建代币账号
		createInstruction := associatedtokenaccount.NewCreateInstruction(fromPublicKey,
			recipientPublicKey,
			mintAddress).Build()
		trxs = append(trxs, createInstruction)
		// 确保余额足够支付转账和费用
		if balanceSol < gasMinitFeelimit+gasCretelimit {
			return ret, fmt.Errorf("insufficient funds:fromAddress:%s balanceSol %d,availableBalance:%d, required %d lamports", fromADdressStr, balanceSol, availableBalance, gasMinitFeelimit+gasCretelimit)
		}
	}
	//balanceSol
	recentBlockhash, err := sol.client.GetLatestBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return ret, fmt.Errorf("failed getRecentBlockhash: %v", err)
	}

	// 创建转账指令
	transferInstruction := token.NewTransferInstruction(
		amount,
		senderTokenAccount,
		receiverTokenAccount,
		fromPublicKey,
		nil, // 多签账户（如果有的话）
	).Build()

	// 创建交易
	trxs = append(trxs, transferInstruction)
	tx, err := internal.NewTransaction(
		trxs,
		recentBlockhash.Value.Blockhash,
		internal.TransactionPayer(fromPublicKey),
	)
	if err != nil {
		return ret, fmt.Errorf("failed to create transaction: %v", err)
	}

	// 使用发起人的私钥签名交易
	_, err = tx.Sign(
		func(key internal.PublicKey) *internal.PrivateKey {
			if fromPublicKey.Equals(key) {
				return &senderPrivateKey
			}
			return nil
		},
	)
	if err != nil {
		return ret, fmt.Errorf("failed to sign transaction: %v", err)
	}

	if false {
		//模拟交易
		simulationResult, err := sol.client.SimulateTransaction(context.Background(), tx)
		if err != nil {
			return ret, fmt.Errorf("failed to send transaction: %v", err)
		}
		log.Warnf("%s", xconv.Json(simulationResult))
		return ret, nil

	} else {
		// 广播交易
		txSignature, err := sol.client.SendTransaction(context.Background(), tx)
		if err != nil {
			return ret, fmt.Errorf("failed to send transaction: %v", err)
		}
		ret.IsCreate = false
		ret.TxID = txSignature.String()
		return ret, nil
	}
}
func (sol *solanaInfo) doGetAssociatedTokenAddress(walletPubKey, mintAddress internal.PublicKey) (internal.PublicKey, error) {
	// Associated Token Program ID
	associatedTokenProgramID := internal.MustPublicKeyFromBase58("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL")

	// SPL Token Program ID
	tokenProgramID := internal.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")

	// Derive Associated Token Address
	ata, _, err := internal.FindProgramAddress(
		[][]byte{
			walletPubKey.Bytes(),
			tokenProgramID.Bytes(),
			mintAddress.Bytes(),
		},
		associatedTokenProgramID,
	)
	return ata, err
}
func (sol *solanaInfo) GetTokenAccountsByOwner(owner, mintAddress string) (bool, error) {

	// 获取该钱包地址下的所有 SPL 代币账户
	ownerPublic := internal.MustPublicKeyFromBase58(owner)
	mintAddressPublic := internal.MustPublicKeyFromBase58(mintAddress)

	return sol.doGetTokenAccountsByOwner(ownerPublic, mintAddressPublic)
}

func (sol *solanaInfo) doGetTokenAccountsByOwner(owner, mintAddress internal.PublicKey) (bool, error) {
	if sol.client == nil {
		return false, fmt.Errorf("sol.client is err")
	}
	// 获取该钱包地址下的所有 SPL 代币账户
	accounts, err := sol.client.GetTokenAccountsByOwner(context.Background(), owner, &rpc.GetTokenAccountsConfig{
		Mint: &mintAddress,
	}, nil)
	if err != nil {
		return false, err
	}
	if accounts == nil {
		return false, nil
	} else {
		if len(accounts.Value) == 0 {
			return false, nil
		}
	}
	return true, nil
}
func (sol *solanaInfo) CreateInstruction(privateKeystr, owner, splTokenMintAddress string) (string, error) {
	walletAddress := internal.MustPublicKeyFromBase58(owner)
	senderPrivateKey := internal.MustPrivateKeyFromBase58(privateKeystr)
	splTokenMintAddressPublic := internal.MustPublicKeyFromBase58(splTokenMintAddress)

	return sol.doCreateInstruction(senderPrivateKey, walletAddress, splTokenMintAddressPublic)
}
func (sol *solanaInfo) doCreateInstruction(senderPrivateKey internal.PrivateKey, walletAddress, splTokenMintAddress internal.PublicKey) (string, error) {
	// 设置发送者地址和私钥
	if sol.client == nil {
		return "", fmt.Errorf("sol.client is err")
	}
	fromPublicKey := senderPrivateKey.PublicKey()
	// 获取该钱包地址下的所有 SPL 代币账户
	createInstruction := associatedtokenaccount.NewCreateInstruction(fromPublicKey,
		walletAddress,
		splTokenMintAddress).Build()
	recentBlockhash, err := sol.client.GetLatestBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return "", fmt.Errorf("failed getRecentBlockhash: %v", err)
	}

	tx, err := internal.NewTransaction(
		[]internal.Instruction{createInstruction},
		recentBlockhash.Value.Blockhash,
		internal.TransactionPayer(fromPublicKey),
	)
	if err != nil {
		return "", fmt.Errorf("failed to create transaction: %v", err)
	}

	// 使用发起人的私钥签名交易
	_, err = tx.Sign(
		func(key internal.PublicKey) *internal.PrivateKey {
			if fromPublicKey.Equals(key) {
				return &senderPrivateKey
			}
			return nil
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %v", err)
	}
	txSignature, err := sol.client.SendTransaction(context.Background(), tx)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %v", err)
	}
	log.Warnf("doCreateInstruction:%v", txSignature.String())
	return txSignature.String(), nil
}
