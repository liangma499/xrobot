package transfer

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"regexp"
	"strings"
	"tron_robot/internal/code"
	"tron_robot/internal/cryptocurrencies/tron/gotron-sdk/pkg/account"
	"tron_robot/internal/cryptocurrencies/tron/gotron-sdk/pkg/address"
	"tron_robot/internal/cryptocurrencies/tron/gotron-sdk/pkg/client"
	"tron_robot/internal/cryptocurrencies/tron/gotron-sdk/pkg/proto/api"
	"tron_robot/internal/cryptocurrencies/tron/gotron-sdk/pkg/proto/core"

	"xbase/encoding/proto"
	"xbase/errors"
	"xbase/log"

	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
	"github.com/shopspring/decimal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TransferInfo struct {
	url    string
	APIKey string
}

func NewTransferNull() *TransferInfo {
	return &TransferInfo{}
}
func NewTransfer(url string, APIKey string) *TransferInfo {
	return &TransferInfo{
		url:    url,
		APIKey: APIKey,
	}
}
func (cr *TransferInfo) Clone() *TransferInfo {
	if cr == nil {
		return nil
	}
	return &TransferInfo{
		url:    cr.url,
		APIKey: cr.APIKey,
	}
}
func (cr *TransferInfo) Client() (*client.GrpcClient, error) {
	c := client.NewGrpcClient(cr.url)
	c.SetAPIKey(cr.APIKey)
	c.SetTimeout(timeOut)

	err := c.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Warnf("%v", err)
		return nil, err
	}
	return c, nil
}

func (cr *TransferInfo) GetAccountDetailed(walletAddRess string) (*account.Account, error) {
	c, err := cr.Client()
	if err != nil {
		return nil, err
	}

	acc, err := c.GetAccountDetailed(walletAddRess)

	if err != nil {
		if err.Error() == "account not found" {
			log.Warnf("%v addrss[%v]acc is nil", walletAddRess)
			return &account.Account{}, nil
		}
		//log.Warnf("err:%v addrss[%v]", err, walletAddRess)
		return nil, err
	}
	if acc == nil {
		log.Warnf("%v addrss[%v]acc is nil", walletAddRess, walletAddRess)
		return &account.Account{}, nil
	}

	return acc, nil
}

// isValidTronAddress 检查 TRON 地址的合理性
func (cr *TransferInfo) isValidTronAddress(address string) bool {
	// TRON 地址正则：以 T 开头，后跟 33 个字符（可为字母和数字）
	var tronAddressPattern = regexp.MustCompile(`^T[a-zA-Z0-9]{33}$`)
	return tronAddressPattern.MatchString(address)
}

// decodeBase58 解码 Base58 地址
func (cr *TransferInfo) decodeBase58(address string) ([]byte, error) {
	return base58.Decode(address)
}

// validateAddress 验证地址的合法性
func (cr *TransferInfo) ValidateAddress(address string) bool {
	if !cr.isValidTronAddress(address) {
		return false
	}

	decoded, err := cr.decodeBase58(address)
	if err != nil {
		return false
	}

	// 对解码后的数据进行 SHA256 哈希
	checksum := decoded[len(decoded)-4:] // 获取最后4个字节作为校验和
	body := decoded[:len(decoded)-4]     // 取出有效数据部分

	// 计算校验和
	hash := sha256.Sum256(body)
	hash = sha256.Sum256(hash[:])

	// 验证校验和
	calculatedChecksum := hash[:4]
	return hex.EncodeToString(calculatedChecksum) == hex.EncodeToString(checksum)
}
func (cr *TransferInfo) GetWalletAccount(walletAddRess string) (*core.Account, error) {

	c, err := cr.Client()
	if err != nil {
		return nil, err
	}

	acc, err := c.GetAccount(walletAddRess)

	if err != nil {
		//log.Warnf("err:%v addrss[%v]", err, walletAddRess)
		return nil, err
	}
	if acc == nil {
		//log.Warnf("addrss[%v]acc is nil", walletAddRess)
		return nil, nil
	}

	return acc, nil

}

func (cr *TransferInfo) GetTrxWalletAccountBalance(walletAddRess string) int64 {

	c, err := cr.Client()
	if err != nil {
		return 0
	}

	acc, err := c.GetAccount(walletAddRess)

	if err != nil {
		log.Warnf("%v addrss[%v]", err, walletAddRess)
		return 0
	}
	if acc == nil {
		log.Warnf("%v addrss[%v]acc is nil", walletAddRess, walletAddRess)
		return 0
	}

	return acc.Balance

}

func (cr *TransferInfo) TRXSend(amount int64, fromTrc20Address, fromPrivateKey, toAddress string) (string, error) {

	c, err := cr.Client()
	if err != nil {
		return "", err
	}

	tx, err := c.Transfer(fromTrc20Address, toAddress, amount)
	if err != nil {
		log.Warnf("%v", err)
		return "", errors.NewError(code.Trc20RpcErr, err)
	}
	if !tx.Result.Result {
		return "", errors.NewError(code.Trc20RpcErr, tx.Result.Code.String())
	}

	return cr.doBroadcast(c, fromPrivateKey, hex.EncodeToString(tx.Txid), tx.Transaction)
}

func (cr *TransferInfo) TRC20ContractBalance(address, contract string) int64 {

	c, err := cr.Client()
	if err != nil {
		return 0
	}

	balance, err := c.TRC20ContractBalance(address, contract)
	if err != nil {
		log.Warnf("%v", err)
		return 0
	}

	return balance.Int64()

}
func (cr *TransferInfo) TRC20Send(amount int64, fromTrc20Address, fromPrivateKey, toAddress, contract string, feeLimit int64) (string, error) {

	c, err := cr.Client()
	if err != nil {
		return "", err
	}

	//toAddress := "TTXgY2jzdNoNrZtbDQrBzL8XnjL5wu3oR5"
	tx, err := c.TRC20Send(fromTrc20Address, toAddress, contract, big.NewInt(amount), feeLimit)
	if err != nil {
		log.Warnf("%v", err)
		return "", errors.NewError(code.Trc20RpcErr, err)
	}
	if !tx.Result.Result {
		return "", errors.NewError(code.Trc20RpcErr, tx.Result.Code.String())
	}
	return cr.doBroadcast(c, fromPrivateKey, hex.EncodeToString(tx.Txid), tx.Transaction)
}

func (cr *TransferInfo) GetAccountResource(address string) (*api.AccountResourceMessage, error) {

	c, err := cr.Client()
	if err != nil {
		return nil, err
	}

	rst, err := c.GetAccountResource(address)
	if err != nil {
		log.Warnf("%v", err)
		return nil, err
	}

	return rst, nil

}
func (cr *TransferInfo) DelegateResourceEnegy(from, fromPrivateKey, to string, delegateBalance int64) (string, error) {

	c, err := cr.Client()
	if err != nil {
		return "", err
	}

	tx, err := c.DelegateResource(from, to, core.ResourceCode_ENERGY, delegateBalance, true, 28800)
	if err != nil {
		log.Warnf("%v", err)
		return "", err
	}
	if !tx.Result.Result {
		return "", errors.NewError(code.Trc20RpcErr, tx.Result.Code.String())
	}

	return cr.doBroadcast(c, fromPrivateKey, hex.EncodeToString(tx.Txid), tx.Transaction)
}
func (cr *TransferInfo) UnDelegateResource(from, fromPrivateKey, to string, delegateBalance int64) (string, error) {

	c, err := cr.Client()
	if err != nil {
		return "", err
	}

	tx, err := c.UnDelegateResource(from, to, core.ResourceCode_ENERGY, delegateBalance, true)
	if err != nil {
		log.Warnf("%v", err)
		return "", err
	}
	if !tx.Result.Result {
		return "", errors.NewError(code.Trc20RpcErr, tx.Result.Code.String())
	}

	return cr.doBroadcast(c, fromPrivateKey, hex.EncodeToString(tx.Txid), tx.Transaction)
}

func (cr *TransferInfo) doBroadcast(c *client.GrpcClient, fromPrivateKey, txid string, transaction *core.Transaction) (string, error) {
	rawData, err := proto.Marshal(transaction.GetRawData())
	if err != nil {
		return "", errors.NewError(code.Trc20RpcErr, err)
	}

	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)

	if strings.HasPrefix(fromPrivateKey, "0x") {
		fromPrivateKey = strings.TrimPrefix(fromPrivateKey, "0x")
	} else if strings.HasPrefix(fromPrivateKey, "0X") {
		fromPrivateKey = strings.TrimPrefix(fromPrivateKey, "0X")
	}

	priKey, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		log.Warnf("%v", err)
		return "", errors.NewError(code.Trc20RpcErr, err)
	}
	signature, err := crypto.Sign(hash, priKey)
	if err != nil {
		log.Warnf("%v", err)
		return "", errors.NewError(code.Trc20RpcErr, err)
	}
	transaction.Signature = append(transaction.Signature, signature)
	result, err := c.Broadcast(transaction)
	if err != nil {
		log.Warnf("%v", err)
		return "", errors.NewError(code.Trc20RpcErr, err)
	}
	if !result.Result {
		return "", errors.NewError(code.Trc20RpcErr, result.Code.String())
	}
	log.Warnf("%#v", result)

	return txid, nil
}

func (cr *TransferInfo) GetBlockInfoByNum(blockID int64) (*api.TransactionInfoList, error) {

	c, err := cr.Client()
	if err != nil {
		return nil, err
	}

	tx, err := c.GetBlockInfoByNum(blockID)
	if err != nil {
		log.Warnf("%v", err)
		return nil, err
	}

	return tx, nil
}

func (cr *TransferInfo) GetBlockByID(blockID string) (*core.Block, error) {

	c, err := cr.Client()
	if err != nil {
		return nil, err
	}

	tx, err := c.GetBlockByID(blockID)
	if err != nil {
		log.Warnf("%v", err)
		return nil, err
	}

	return tx, nil
}
func (cr *TransferInfo) UnmarshalTo(trContract *core.Transaction_Contract) *transferContract {
	if trContract == nil {
		return nil
	}
	if trContract.Parameter == nil {
		return nil
	}

	switch trContract.Type {
	case core.Transaction_Contract_AccountCreateContract:
		{
			if true {
				return nil
			}
			rst := &core.AccountCreateContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			toAddress := cr.generateTronAddress(rst.AccountAddress)
			fromAddress := cr.generateTronAddress(rst.OwnerAddress)

			return &transferContract{
				Protocol:    "AccountCreateContract",
				ToAddress:   toAddress,
				FromAddress: fromAddress,
				IsEngergy:   false,
			}
		}
	case core.Transaction_Contract_TransferContract:
		{
			rst := &core.TransferContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			toAddress := cr.generateTronAddress(rst.ToAddress)
			fromAddress := cr.generateTronAddress(rst.OwnerAddress)
			realAmount := decimal.NewFromInt(rst.Amount)
			return &transferContract{
				Protocol:    "TransferContract",
				ToAddress:   toAddress,
				FromAddress: fromAddress,
				RealAmount:  realAmount.Copy(),
				Contract:    "-",
				IsEngergy:   false,
			}
		}
	case core.Transaction_Contract_TransferAssetContract:
		{
			if true {
				return nil
			}
			rst := &core.TransferAssetContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			toAddress := cr.generateTronAddress(rst.ToAddress)
			fromAddress := cr.generateTronAddress(rst.OwnerAddress)
			//log.Warnf("TransferAssetContract: fromAddress:%s toAddress:%s", fromAddress, toAddress)
			return &transferContract{
				Protocol:    "TransferAssetContract",
				ToAddress:   toAddress,
				FromAddress: fromAddress,
				IsEngergy:   false,
			}
		}
	case core.Transaction_Contract_VoteAssetContract:
		{
			if true {
				return nil
			}
			rst := &core.VoteAssetContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			//toAddress := address.HexToAddress(hex.EncodeToString(rst.VoteAddress[]))
			fromAddress := cr.generateTronAddress(rst.OwnerAddress)
			return &transferContract{
				Protocol:    "VoteAssetContract",
				FromAddress: fromAddress,
				IsEngergy:   false,
			}
		}
	case core.Transaction_Contract_VoteWitnessContract:
		{
			if true {
				return nil
			}
			rst := &core.VoteWitnessContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "VoteWitnessContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_WitnessCreateContract:
		{
			if true {
				return nil
			}
			rst := &core.WitnessCreateContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "WitnessCreateContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_AssetIssueContract:
		{
			if true {
				return nil
			}
			rst := &core.AssetIssueContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "AssetIssueContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_WitnessUpdateContract:
		{
			if true {
				return nil
			}
			rst := &core.WitnessUpdateContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "WitnessUpdateContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_ParticipateAssetIssueContract:
		{
			if true {
				return nil
			}
			rst := &core.ParticipateAssetIssueContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "ParticipateAssetIssueContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_AccountUpdateContract:
		{
			if true {
				return nil
			}
			rst := &core.AccountUpdateContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "AccountUpdateContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_FreezeBalanceContract:
		{
			if true {
				return nil
			}
			rst := &core.FreezeBalanceContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			toAddress := cr.generateTronAddress(rst.ReceiverAddress)
			fromAddress := cr.generateTronAddress(rst.OwnerAddress)
			return &transferContract{
				Protocol:    "FreezeBalanceContract",
				ToAddress:   toAddress,
				FromAddress: fromAddress,
				IsEngergy:   false,
			}
		}
	case core.Transaction_Contract_UnfreezeBalanceContract:
		{
			if true {
				return nil
			}
			rst := &core.UnfreezeBalanceContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			toAddress := cr.generateTronAddress(rst.ReceiverAddress)
			fromAddress := cr.generateTronAddress(rst.OwnerAddress)
			return &transferContract{
				Protocol:    "UnfreezeBalanceContract",
				ToAddress:   toAddress,
				FromAddress: fromAddress,
				IsEngergy:   false,
			}
		}
	case core.Transaction_Contract_WithdrawBalanceContract:
		{
			if true {
				return nil
			}
			rst := &core.WithdrawBalanceContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "WithdrawBalanceContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_UnfreezeAssetContract:
		{
			if true {
				return nil
			}
			rst := &core.UnfreezeAssetContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "UnfreezeAssetContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_UpdateAssetContract:
		{
			if true {
				return nil
			}
			rst := &core.UpdateAssetContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "UpdateAssetContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_ProposalCreateContract:
		{
			if true {
				return nil
			}
			rst := &core.ProposalCreateContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "ProposalCreateContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_ProposalApproveContract:
		{
			if true {
				return nil
			}
			rst := &core.ProposalApproveContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "ProposalApproveContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_ProposalDeleteContract:
		{
			if true {
				return nil
			}
			rst := &core.ProposalDeleteContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "ProposalDeleteContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_SetAccountIdContract:
		{
			if true {
				return nil
			}
			rst := &core.SetAccountIdContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "SetAccountIdContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_CustomContract:
		{
			return nil
		}
	case core.Transaction_Contract_CreateSmartContract:
		{
			if true {
				return nil
			}
			rst := &core.CreateSmartContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "CreateSmartContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_TriggerSmartContract:
		{
			rst := &core.TriggerSmartContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			contractAddress := address.HexToAddress(hex.EncodeToString(rst.ContractAddress)).String()
			//现在只监听USDT

			// USDT 转入地址
			lenData := len(rst.Data)
			if lenData < 36 {
				return nil
			}
			toAddress := cr.generateTronAddress(rst.Data[15:36])

			fromAddress := cr.generateTronAddress(rst.OwnerAddress)

			byteSlice := make([]byte, 0)

			for i := 63; i < lenData; i++ {
				byteSlice = append(byteSlice, rst.Data[i])
			}
			realAmount := cr.doBytesDecimal(byteSlice)
			return &transferContract{
				Protocol:    "TriggerSmartContract",
				ToAddress:   toAddress,
				FromAddress: fromAddress,
				Contract:    contractAddress,
				RealAmount:  realAmount.Copy(),
				IsEngergy:   false,
			}
		}
	case core.Transaction_Contract_GetContract:
		{
			return nil
		}
	case core.Transaction_Contract_UpdateSettingContract:
		{
			if true {
				return nil
			}
			rst := &core.UpdateSettingContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "UpdateSettingContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_ExchangeCreateContract:
		{
			if true {
				return nil
			}
			rst := &core.ExchangeCreateContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			toAddress := cr.generateTronAddress(rst.FirstTokenId)
			fromAddress := cr.generateTronAddress(rst.OwnerAddress)
			return &transferContract{
				Protocol:    "ExchangeCreateContract",
				ToAddress:   toAddress,
				FromAddress: fromAddress,
				IsEngergy:   false,
			}
		}
	case core.Transaction_Contract_ExchangeInjectContract:
		{
			if true {
				return nil
			}
			rst := &core.ExchangeInjectContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "ExchangeInjectContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_ExchangeWithdrawContract:
		{
			if true {
				return nil
			}
			rst := &core.ExchangeWithdrawContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "ExchangeWithdrawContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_ExchangeTransactionContract:
		{
			if true {
				return nil
			}
			rst := &core.ExchangeTransactionContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			toAddress := cr.generateTronAddress(rst.TokenId)
			fromAddress := cr.generateTronAddress(rst.OwnerAddress)
			return &transferContract{
				Protocol:    "ExchangeTransactionContract",
				ToAddress:   toAddress,
				FromAddress: fromAddress,
				IsEngergy:   false,
			}
		}
	case core.Transaction_Contract_UpdateEnergyLimitContract:
		{
			if true {
				return nil
			}
			rst := &core.UpdateEnergyLimitContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			toAddress := cr.generateTronAddress(rst.ContractAddress)
			fromAddress := cr.generateTronAddress(rst.OwnerAddress)

			return &transferContract{
				Protocol:    "UpdateEnergyLimitContract",
				ToAddress:   toAddress,
				FromAddress: fromAddress,
				IsEngergy:   false,
			}
		}
	case core.Transaction_Contract_AccountPermissionUpdateContract:
		{
			if true {
				return nil
			}
			rst := &core.AccountPermissionUpdateContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			//toAddress := address.HexToAddress(hex.EncodeToString(rst.ContractAddress))
			fromAddress := cr.generateTronAddress(rst.OwnerAddress)
			return &transferContract{
				Protocol:    "AccountPermissionUpdateContract",
				FromAddress: fromAddress,
				IsEngergy:   false,
			}
		}
	case core.Transaction_Contract_ClearABIContract:
		{
			if true {
				return nil
			}
			rst := &core.ClearABIContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "ClearABIContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_UpdateBrokerageContract:
		{
			if true {
				return nil
			}
			rst := &core.UpdateBrokerageContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "UpdateBrokerageContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_ShieldedTransferContract:
		{
			if true {
				return nil
			}
			rst := &core.ShieldedTransferContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "ShieldedTransferContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_MarketSellAssetContract:
		{
			if true {
				return nil
			}
			rst := &core.MarketSellAssetContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "MarketSellAssetContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_MarketCancelOrderContract:
		{
			if true {
				return nil
			}
			rst := &core.MarketCancelOrderContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			return &transferContract{
				Protocol:  "MarketCancelOrderContract",
				IsEngergy: false,
			}
		}
	case core.Transaction_Contract_FreezeBalanceV2Contract:
		{
			if true {
				return nil
			}
			rst := &core.FreezeBalanceV2Contract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			//toAddress := address.HexToAddress(hex.EncodeToString(rst.ContractAddress))
			fromAddress := cr.generateTronAddress(rst.OwnerAddress)
			return &transferContract{
				Protocol:    "FreezeBalanceV2Contract",
				FromAddress: fromAddress,
				IsEngergy:   false,
			}
		}
	case core.Transaction_Contract_UnfreezeBalanceV2Contract:
		{
			if true {
				return nil
			}
			rst := &core.UnfreezeBalanceV2Contract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			//toAddress := address.HexToAddress(hex.EncodeToString(rst.ContractAddress))
			fromAddress := cr.generateTronAddress(rst.OwnerAddress)
			return &transferContract{
				Protocol:    "UnfreezeBalanceV2Contract",
				FromAddress: fromAddress,
				IsEngergy:   false,
			}
		}
	case core.Transaction_Contract_WithdrawExpireUnfreezeContract:
		{
			if true {
				return nil
			}
			rst := &core.WithdrawExpireUnfreezeContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			//toAddress := address.HexToAddress(hex.EncodeToString(rst.ContractAddress))
			fromAddress := cr.generateTronAddress(rst.OwnerAddress)
			return &transferContract{
				Protocol:    "WithdrawExpireUnfreezeContract",
				FromAddress: fromAddress,
				IsEngergy:   false,
			}
		}
	case core.Transaction_Contract_DelegateResourceContract:
		{
			rst := &core.DelegateResourceContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			toAddress := cr.generateTronAddress(rst.ReceiverAddress)
			fromAddress := cr.generateTronAddress(rst.OwnerAddress)
			amonut := decimal.NewFromInt(rst.Balance)
			return &transferContract{
				Protocol:    "DelegateResourceContract",
				FromAddress: fromAddress,
				ToAddress:   toAddress,
				RealAmount:  amonut,
				Contract:    "TU2MJ5Veik1LRAgjeSzEdvmDYx7mefJZvd",
				IsEngergy:   true,
			}
		}
	case core.Transaction_Contract_UnDelegateResourceContract:
		{
			if true {
				return nil
			}
			rst := &core.UnDelegateResourceContract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			toAddress := cr.generateTronAddress(rst.ReceiverAddress)
			fromAddress := cr.generateTronAddress(rst.OwnerAddress)
			return &transferContract{
				Protocol:    "UnDelegateResourceContract",
				FromAddress: fromAddress,
				ToAddress:   toAddress,
				IsEngergy:   true,
			}
		}
	case core.Transaction_Contract_CancelAllUnfreezeV2Contract:
		{
			if true {
				return nil
			}
			rst := &core.CancelAllUnfreezeV2Contract{}
			err := trContract.Parameter.UnmarshalTo(rst)
			if err != nil {
				log.Warnf("Write:%v", err)
				return nil
			}
			//toAddress := address.HexToAddress(hex.EncodeToString(rst.ReceiverAddress))
			fromAddress := cr.generateTronAddress(rst.OwnerAddress)
			return &transferContract{
				Protocol:    "CancelAllUnfreezeV2Contract",
				FromAddress: fromAddress,
				IsEngergy:   false,
			}
		}

	}
	log.Warnf("default: %d", trContract.Type)
	return nil
}

// 生成 TRON 地址
func (cr *TransferInfo) generateTronAddress(srcAddress []byte) string {

	toAddress := address.HexToAddress(hex.EncodeToString(srcAddress)).String()
	if strings.HasPrefix(toAddress, "T") {
		return toAddress
	} else {
		srcAddress[0] = 0x41
		// TRON 主网地址的前缀是 0x41
		return address.HexToAddress(hex.EncodeToString(srcAddress)).String()

	}

}

func (cr *TransferInfo) doBytesDecimal(buf []byte) decimal.Decimal {

	endPos := len(buf)
	if endPos > 8 {
		endPos = 8
	}
	byteSlice := make([]byte, 8)

	copy(byteSlice[8-endPos:], buf[0:endPos])
	return decimal.NewFromUint64(binary.BigEndian.Uint64(byteSlice))
}

func (cr *TransferInfo) GetNowBlockNum() (int64, error) {

	c, err := cr.Client()
	if err != nil {
		return 0, err
	}

	tx, err := c.GetNowBlock()
	if err != nil {
		log.Warnf("%v", err)
		return 0, err
	}

	return tx.BlockHeader.RawData.Number, nil
}
func (cr *TransferInfo) GetBlockByNum(num int64) (*api.BlockExtention, error) {

	c, err := cr.Client()
	if err != nil {
		return nil, err
	}

	tx, err := c.GetBlockByNum(num)
	if err != nil {
		log.Warnf("%v", err)
		return nil, err
	}

	return tx, nil
}
