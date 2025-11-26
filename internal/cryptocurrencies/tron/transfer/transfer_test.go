package transfer_test

import (
	"encoding/hex"
	"testing"
	"xbase/log"
	"xbase/utils/xconv"
	"xrobot/internal/cryptocurrencies/tron/gotron-sdk/pkg/address"
	"xrobot/internal/cryptocurrencies/tron/gotron-sdk/pkg/common"
	"xrobot/internal/cryptocurrencies/tron/transfer"
	"xrobot/internal/xtypes"
)

var (
	transferIn = transfer.NewTransfer("grpc.trongrid.io:50051", "869bf846-01b8-4d06-9ad3-b5a1375e4d39")
)

func TestClient_GetWalletAccount(t *testing.T) {

	trxAount, err := transferIn.GetAccountDetailed("ggggggggggggggg")
	if err != nil {
		log.Warnf("%v", err)
		return
	}
	log.Warnf("account:%s", xconv.String(trxAount))
}
func TestClient_GetTrxWalletAccount(t *testing.T) {
	trxAount := transferIn.GetTrxWalletAccountBalance("eeeeeeeeeeeeeeeeeeeeeeeeemm")
	log.Warnf("%v", trxAount)
}
func TestClient_TRC20ContractBalance(t *testing.T) {
	balance := transferIn.TRC20ContractBalance("eeeeeeeeeeeeeeeeeeeeeeeeemm", xtypes.ENERGY.Trc20Contract())
	log.Warnf("%v", balance)
}
func TestClient_GetAccountResource(t *testing.T) {
	//balance, err := transferIn.GetAccountResource("eeeeeeeeeeeeeeeeeeeeeeeeemm")
	balance, err := transferIn.GetAccountResource("TEj8NgQM37dABXxVQRZo2b7nTkoCFM2qCQ")
	if err != nil {
		log.Warnf("%v", err)
		return
	}
	log.Warnf("%s", xconv.Json(balance))
}
func TestClient_ValidateAddress(t *testing.T) {
	//balance, err := transferIn.GetAccountResource("eeeeeeeeeeeeeeeeeeeeeeeeemm")
	isValidate := transferIn.ValidateAddress("ccccccccccccccccccccccccccc")

	log.Warnf("%v", isValidate)
}

func TestClient_TRXSend(t *testing.T) {
	txid, err := transferIn.TRXSend(100000, "cccccccccccccccccccc", "", "aaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		log.Warnf("%v", err)
		return
	}
	log.Warnf("%v", txid)
}

func TestClient_DelegateResourceEnegy(t *testing.T) {
	txid, err := transferIn.DelegateResourceEnegy("eeeeeeeeeeeeeeeeeeeeeeeeemm", "pravateKey", "bbbbbbbbbbbbbbbb", 2000000)
	if err != nil {
		log.Warnf("%v", err)
		return
	}

	log.Warnf("%s", txid)

}

func TestClient_GetBlockInfoByNum(t *testing.T) {
	resp, err := transferIn.GetBlockInfoByNum(65887871)
	if err != nil {
		log.Warnf("%v", err)
		return
	}
	if resp == nil {
		return
	}
	for _, item := range resp.TransactionInfo {
		if item.ContractAddress == nil {
			continue
		}

		//dest := common.BytesToHash(item.Id) 交易hash
		//log.Warnf("%s", xconv.Json(item))

		addr := address.HexToAddress(hex.EncodeToString(item.ContractAddress))
		log.Warnf("%s", addr.String())

		for _, info := range item.ContractResult {
			rst := address.HexToAddress(hex.EncodeToString(info))
			log.Warnf("11:%v", rst.String())
		}
		orderId := address.HexToAddress(hex.EncodeToString(item.OrderId))
		log.Warnf("orderId:%v", orderId.String())
		if item.Log != nil {
			for _, logItem := range item.Log {
				logAddress := address.HexToAddress(hex.EncodeToString(logItem.Address))
				log.Warnf("logAddress:%v", logAddress.String())
			}

		}
		for _, itrx := range item.InternalTransactions {
			iHash := common.BytesToHash(itrx.Hash)
			callerAddress := address.HexToAddress(hex.EncodeToString(itrx.CallerAddress))
			transferToAddress := address.HexToAddress(hex.EncodeToString(itrx.TransferToAddress))
			log.Warnf("iHash:%s callerAddress:%s transferToAddress:%s", iHash.String(), callerAddress.String(), transferToAddress.String())
		}

		/*
			addr, err := address.Base64ToAddress(dest)
			if err != nil {
				log.Warnf("%v", err)
				continue
			}
		*/

		/*

			log.Warnf("%s", rst.String())
		*/
	}

}
func TestClient_GetBlockByID(t *testing.T) {
	resp, err := transferIn.GetBlockByID("ffffffffffffffffff")
	if err != nil {
		log.Warnf("%v", err)
		return
	}
	for _, item := range resp.Transactions {
		//log.Warnf("%s", xconv.Json(item))
		//sorce, byte := json.Marshal(item)
		for _, contractItem := range item.RawData.Contract {
			if contractItem == nil {
				continue
			}
			_ = transferIn.UnmarshalTo(contractItem)
			/*
				if contractItem.Type == core.Transaction_Contract_TransferContract {
					trxInfo := transferIn.UnmarshalTo(contractItem)
					if trxInfo == nil {
						continue
					}
					if trx, ok := trxInfo.(*core.TransferContract); !ok {
						continue
					} else {
						toAddress := address.HexToAddress(hex.EncodeToString(trx.ToAddress))
						formAddress := address.HexToAddress(hex.EncodeToString(trx.OwnerAddress))
						log.Warnf("TransferContract toAddress:%s,formAddress:%s amount:%d", toAddress, formAddress, trx.Amount)
					}
				} else if contractItem.Type == core.Transaction_Contract_TransferAssetContract {
					trxInfo := transferIn.UnmarshalTo(contractItem)
					if trxInfo == nil {
						continue
					}
					if trx, ok := trxInfo.(*core.TransferAssetContract); !ok {
						continue
					} else {
						toAddress := address.HexToAddress(hex.EncodeToString(trx.ToAddress))
						formAddress := address.HexToAddress(hex.EncodeToString(trx.OwnerAddress))
						log.Warnf("TransferAssetContract toAddress:%s,formAddress:%s amount:%d", toAddress, formAddress, trx.Amount)
					}
				}
			*/

		}
	}

}
func TestClient_GetNowBlock(t *testing.T) {
	//balance, err := transferIn.GetAccountResource("eeeeeeeeeeeeeeeeeeeeeeeeemm")
	blockNum, err := transferIn.GetNowBlockNum()
	if err != nil {
		return
	}
	log.Warnf("blockNum:%d", blockNum)
}
