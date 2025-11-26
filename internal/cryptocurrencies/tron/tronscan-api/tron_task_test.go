package tronscanapi_test

import (
	"testing"
	"xbase/log"
	"xbase/utils/xconv"
	"xbase/utils/xtime"
	tronscanapi "xrobot/internal/cryptocurrencies/tron/tronscan-api"
)

const (
	app_key = "97a52859-f212-46b0-8c01-bbf41ceddb41"
)

func Test_InitTask(t *testing.T) {

	resp, err := tronscanapi.GetTransactionsTrc10("https://apilist.tronscanapi.com", app_key, &tronscanapi.TransactionsTrc10Req{
		Start: 0,
		Limit: 50,
		//StartTimestamp: xtime.Now().UnixMilli() - 300000,
		//EndTimestamp:   xtime.Now().UnixMilli() - 3000,
		Address: "eeeeeeeeeeeeeeeeeeeeeeeeemm",
	})
	if err != nil {
		log.Warnf("%v", err)
		return
	}
	for _, item := range resp.Data {
		log.Warnf("%s", xconv.Json(item))
	}

	//tronscanapi.Instance()
}

func Test_GetTransactionsTrc20And721(t *testing.T) {

	resp, err := tronscanapi.GetTransactionsTrc20And721("https://apilist.tronscanapi.com", app_key, &tronscanapi.TransactionsTrc20And721Req{
		Start:          0,
		Limit:          50,
		StartTimestamp: xtime.Now().UnixMilli() - 300000,
		EndTimestamp:   xtime.Now().UnixMilli() - 3000,
		//FromAddress:    "THfVXfjsHuJ24eeVUdcPs15pZhaKbp6uQo",
		//ToAddress: "TJagHTdonvXVR7e4An2ByBWQMnWUWMDCVG",
		//FilterTokenValue: 1,
	})
	if err != nil {
		log.Warnf("%v", err)
		return
	}
	for _, item := range resp.TokenTransfers {
		log.Warnf("%#v", xconv.Json(item))
	}

}

func Test_GetTransactionsAccount(t *testing.T) {

	resp, err := tronscanapi.GetTransfersWithStatus("https://apilist.tronscanapi.com", app_key, &tronscanapi.TransfersWithStatusReq{
		Start:     0,
		Limit:     50,
		Address:   "eeeeeeeeeeeeeeeeeeeeeeeeemm",
		Direction: 0,
		DbVersion: 0,
		Trc20Id:   "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
		//Trc20Id: "_",
		Reverse: true,
	})
	if err != nil {
		log.Warnf("%v", err)
		return
	}
	for _, item := range resp.Data {
		log.Warnf("%#v", xconv.Json(item))
	}

}

func Test_GetAccountDetailV2(t *testing.T) {

	resp, err := tronscanapi.GetAccountDetailV2("https://apilist.tronscanapi.com", app_key, &tronscanapi.AccountDetailV2Req{
		Address: "THtUkvjiWwH52MBgdesCL393NK1NXWc8Bi",
		//Address: "ggggggggggggggg",
	})
	if err != nil {
		log.Warnf("%v", err)
		return
	}
	log.Warnf(xconv.Json(resp))

}
func Test_GetTransaction(t *testing.T) {
	//https://apilist.tronscanapi.com/api/transaction?sort=+timestamp&count=true&limit=20&start=0&start_timestamp=1529856000000&end_timestamp=1734968443000&toAddress=eeeeeeeeeeeeeeeeeeeeeeeeemm
	resp, err := tronscanapi.GetTransaction("https://apilist.tronscanapi.com", app_key, &tronscanapi.TransactionReq{
		Count: true,
		Start: 0,
		Limit: 50,
		//Start_timestamp: 1529856000000,
		//End_timestamp:   xtime.Now().UnixMilli(),
	})
	if err != nil {
		log.Warnf("%v", err)
		return
	}
	for key, item := range resp.Data {
		log.Warnf("index[%d] %#v", key, xconv.Json(item))
	}

}
func Test_GetTransactionInfo(t *testing.T) {
	//https://apilist.tronscanapi.com/api/transaction?sort=+timestamp&count=true&limit=20&start=0&start_timestamp=1529856000000&end_timestamp=1734968443000&toAddress=eeeeeeeeeeeeeeeeeeeeeeeeemm
	resp, err := tronscanapi.GetTransactionInfo("https://apilist.tronscanapi.com", app_key, &tronscanapi.TransactionInfoReq{
		Hash: "24c0a980d919bffbce847a91de46ac5997325852f25f1b90858085ba8111283f",
	})
	if err != nil {
		log.Warnf("%v", err)
		return
	}
	log.Warnf("%v", xconv.Json(resp))

}

//https://apilist.tronscanapi.com/api/token_trc20/transfers-with-status?limit=10&start=0&trc20Id=TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t&address=THfVXfjsHuJ24eeVUdcPs15pZhaKbp6uQo
