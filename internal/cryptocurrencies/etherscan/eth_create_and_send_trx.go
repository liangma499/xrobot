package etherscan

import (
	"context"
	"xrobot/internal/xtypes"

	"fmt"
	"math/big"
	"strings"
	"xbase/log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

const (
	gasLimit          = uint64(23000)
	gasLimitMint      = uint64(230000)
	accountBalance    = 0.001
	erc20TransferABI  = `[{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`
	erc20BalanceOfABI = `[{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`
	erc20AllowanceABI = `[{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`
)

func (gb *ethInfo) Balance(client *ethclient.Client, accountAddress string, contractAddress string) (*big.Int, error) {
	accountAddressHex := common.HexToAddress(accountAddress)
	return gb.doBalance(client, accountAddressHex, contractAddress, true)
}
func (gb *ethInfo) doBalance(client *ethclient.Client, accountAddress common.Address, contractAddress string, bllowance bool) (*big.Int, error) {
	if contractAddress == "-" || contractAddress == "" {
		balance, err := client.BalanceAt(context.Background(), accountAddress, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get balance: %v", err)
		}
		return balance, nil
	} else {
		if bllowance {
			// 创建合约实例
			tokenABI, err := abi.JSON(strings.NewReader(erc20AllowanceABI))
			if err != nil {
				return nil, fmt.Errorf("failed to parse token ABI: %v", err)
			}
			contractAddressHex := common.HexToAddress(contractAddress)
			// 构造调用数据
			data, err := tokenABI.Pack("allowance", accountAddress, accountAddress)
			if err != nil {
				return nil, fmt.Errorf("failed to pack balanceOf call: %v", err)
			}

			// 使用 eth_call 获取余额
			var result *big.Int
			callMsg := ethereum.CallMsg{
				To:   &contractAddressHex,
				Data: data,
			}
			resultBytes, err := client.CallContract(context.Background(), callMsg, nil)
			if err != nil {
				return nil, fmt.Errorf("failed to get balance: %v", err)
			}
			// 解析结果
			err = tokenABI.UnpackIntoInterface(&result, "allowance", resultBytes)
			if err != nil {
				return nil, fmt.Errorf("failed to unpack result: %v", err)
			}
			return decimal.NewFromBigInt(result, 0).BigInt(), nil
		} else {
			// 创建合约实例
			tokenABI, err := abi.JSON(strings.NewReader(erc20BalanceOfABI))
			//tokenABI, err := abi.JSON(strings.NewReader(erc20AllowanceABI))
			if err != nil {
				return nil, fmt.Errorf("failed to parse token ABI: %v", err)
			}
			contractAddressHex := common.HexToAddress(contractAddress)
			// 构造调用数据
			data, err := tokenABI.Pack("balanceOf", accountAddress)
			//data, err := tokenABI.Pack("allowance", accountAddress, accountAddress)
			if err != nil {
				return nil, fmt.Errorf("failed to pack balanceOf call: %v", err)
			}

			// 使用 eth_call 获取余额
			var result *big.Int
			callMsg := ethereum.CallMsg{
				To:   &contractAddressHex,
				Data: data,
			}
			resultBytes, err := client.CallContract(context.Background(), callMsg, nil)
			if err != nil {
				return nil, fmt.Errorf("failed to get balance: %v", err)
			}
			// 解析结果
			err = tokenABI.UnpackIntoInterface(&result, "balanceOf", resultBytes)
			//err = tokenABI.UnpackIntoInterface(&result, "allowance", resultBytes)
			if err != nil {
				return nil, fmt.Errorf("failed to unpack result: %v", err)
			}
			return decimal.NewFromBigInt(result, 0).BigInt(), nil
		}

	}
}
func (gb *ethInfo) EthCreateAndSendTrx(privateKeyStr string, toAddress, contractAddress string, amount int64, gasPlaces, contractPlaces int8) (string, string, error) {

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/ecbb62bf29164de1b4c8fcbe9260cb8f")

	if err != nil {
		return "", "", fmt.Errorf("ethclient.dial: %v", err)
	}
	//如果有0x去掉0x
	privateKeyStr = gb.removePrefix(privateKeyStr)

	// 设置发送者地址和私钥
	privateKey, err := crypto.HexToECDSA(privateKeyStr) // 使用你的私钥
	if err != nil {
		return "", "", fmt.Errorf("failed to convert private key: %v", err)
	}

	// 获取地址
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	fromADdressStr := fromAddress.String()

	// 获取当前 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fromADdressStr, fmt.Errorf("failed to convert private key: %v", err)
	}
	value := big.NewInt(amount)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fromADdressStr, fmt.Errorf("failed to suggest gas price: %v", err)
	}

	minGasBalance := xtypes.CoefficientInt64(decimal.NewFromFloat(accountBalance), gasPlaces).BigInt()
	if contractAddress == "-" || contractAddress == "" {

		// 检查余额
		balance, err := gb.doBalance(client, fromAddress, contractAddress, false)
		if err != nil {
			return "", fromADdressStr, fmt.Errorf("failed to get balance: %v", err)
		}
		//可用余额
		minNeed := new(big.Int).Set(balance).Sub(balance, minGasBalance)
		// 检查是否有足够的余额
		txCost := new(big.Int).SetUint64(gasLimit).Mul(gasPrice, new(big.Int).SetUint64(gasLimit))
		totalCost := new(big.Int).Set(txCost).Add(txCost, value)
		log.Warnf("BRP_20 fromAddress:%s,balance:%s  minNeed:%s txCost:%s totalCost:%s wei\n",
			fromADdressStr,
			balance.String(),
			minNeed.String(),
			txCost.String(),
			totalCost.String(),
		)

		if minNeed.Cmp(totalCost) < 0 {
			return "", fromADdressStr, fmt.Errorf("insufficient funds:fromAddress:%s balance %s,minNeed:%s, required %s wei", fromADdressStr, balance.String(), minNeed.String(), totalCost.String())
		}
	} else {
		minBalance := big.NewInt(1)
		if contractPlaces > 3 {
			minBalance = xtypes.CoefficientInt64(decimal.NewFromFloat(accountBalance), gasPlaces).BigInt()
		}

		availableBalance, err := gb.doBalance(client, fromAddress, contractAddress, false)
		if err != nil {
			return "", fromADdressStr, fmt.Errorf("failed to get availableBalance: %v", err)
		}
		minNeed := new(big.Int).Set(minBalance).Add(minBalance, value)
		if availableBalance.Cmp(minNeed) < 0 {
			return "", fromADdressStr, fmt.Errorf("insufficient funds:fromAddress:%s balance %s,minNeed:%s, required %s wei", fromADdressStr, availableBalance.String(), minNeed.String(), minNeed.String())
		}

		balance, err := gb.doBalance(client, fromAddress, "-", false)
		if err != nil {
			return "", fromADdressStr, fmt.Errorf("failed to get balance: %v", err)
		}

		// 检查是否有足够的余额
		txCost := new(big.Int).SetUint64(gasLimitMint).Mul(gasPrice, new(big.Int).SetUint64(gasLimitMint))
		totalCost := new(big.Int).Set(txCost).Add(txCost, minGasBalance)
		log.Warnf("BRP_20 contractAddress:%s ,fromAddress:%s,balance: %s  spl :%s minNeed:%s txCost:%s totalCost:%s  wei\n",
			contractAddress,
			fromADdressStr,
			balance.String(),
			availableBalance.String(),
			minNeed.String(),
			txCost.String(),
			totalCost.String())
		if balance.Cmp(totalCost) < 0 {
			return "", fromADdressStr, fmt.Errorf("insufficient funds:fromAddress:%s balance %s,minNeed:%s, txCost %s wei", fromADdressStr, balance.String(), totalCost.String(), txCost.String())
		}
	}

	// 签名交易
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", fromADdressStr, fmt.Errorf("failed to get network ID: %v", err)
	}
	recipientAddress := common.HexToAddress(toAddress)
	var tx *types.Transaction
	if contractAddress == "-" || contractAddress == "" {
		tx = types.NewTransaction(nonce, recipientAddress, value, gasLimit, gasPrice, nil)
	} else {
		tokenAddress := common.HexToAddress(contractAddress)
		tokenABI, err := abi.JSON(strings.NewReader(erc20TransferABI))
		if err != nil {
			return "", fromADdressStr, fmt.Errorf("failed to parse token ABI: %v", err)
		}
		data, err := tokenABI.Pack("transfer", recipientAddress, value)
		if err != nil {
			log.Fatalf("Failed to pack transfer call: %v", err)
		}
		tx = types.NewTransaction(nonce, tokenAddress, big.NewInt(0), gasLimitMint, gasPrice, data)

	}
	if tx == nil {
		return "", fromADdressStr, fmt.Errorf("tx is nil: %v", err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", fromADdressStr, fmt.Errorf("failed to sign transaction: %v", err)
	}
	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fromADdressStr, fmt.Errorf("failed to send transaction: %v", err)
	}

	return signedTx.Hash().String(), fromADdressStr, nil
}
