package tronscanapi

import (
	"xrobot/internal/cryptocurrencies/tron/internal"

	"github.com/shopspring/decimal"
)

type TransactionsTrc10Req struct {
	Sort             string `json:"sort"`
	Count            bool   `json:"count"`
	Start            int64  `json:"start"`
	Limit            int64  `json:"limit"`
	Address          string `json:"address"`
	FilterTokenValue int    `json:"filterTokenValue"`
}

//sort=-timestamp&count=true&limit=20&start=0&address=eeeeeeeeeeeeeeeeeeeeeeeeemm&filterTokenValue=1

func (t *TransactionsTrc10Req) ToMap() map[string]any {
	if t == nil {
		return nil
	}
	return map[string]any{
		"sort":             "+timestamp",
		"count":            true,
		"start":            t.Start,
		"limit":            t.Limit,
		"address":          t.Address,
		"filterTokenValue": 1,
	}
}

type Transactions struct {
	Block         int64           `json:"block"`         //"block": 49963474,
	Hash          string          `json:"hash"`          //"hash": "a211f0728aa129c17195322a38c3ca0917ab04af75c5f0934215e6447f4ef6d4",
	TimesTamp     int64           `json:"timestamp"`     //"timestamp": 1680503190000,
	OwnerAddress  string          `json:"ownerAddress"`  //"ownerAddress": "TWbdVwjHTNn2PXDPbtSNvvESDd8PpApFmX",
	ToAddressList []string        `json:"toAddressList"` //"ownerAddress": "TWbdVwjHTNn2PXDPbtSNvvESDd8PpApFmX",
	ToAddress     string          `json:"toAddress"`
	ContractType  int64           `json:"contractType"`
	Confirmed     bool            `json:"confirmed"`
	Revert        bool            `json:"revert"`
	ContractData  *ContractData   `json:"contractData"`
	SmartCalls    string          `json:"SmartCalls"`  //"SmartCalls": "",
	Events        string          `json:"Events"`      //"Events": "",
	Id            string          `json:"id"`          //"id": "",
	Data          string          `json:"data"`        //"data": "",
	Fee           string          `json:"fee"`         //"fee": "",
	ContractRet   string          `json:"contractRet"` //"contractRet": "SUCCESS",
	Result        string          `json:"result"`      //"result": "SUCCESS",
	Amount        decimal.Decimal `json:"amount"`      //"amount": "93000000",
	CostData      *Cost           `json:"cost"`
	TokenInfoData *TokenInfo      `json:"tokenInfo"`
	TokenType     string          `json:"tokenType"`
}
type TransactionsTrc10Resp struct {
	Total             int64           `json:"total"`
	RangeTotal        int64           `json:"rangeTotal"`
	WholeChainTxCount int64           `json:"wholeChainTxCount"`
	ContractMap       map[string]bool `json:"contractMap"`
	Data              []*Transactions `json:"data"`
}

/*
{
  "total": 1,
  "data": [
    {
      "contractRet": "SUCCESS",
      "amount": 154844000,
      "data": "",
      "tokenName": "_",
      "confirmed": true,
      "transactionHash": "a13ea910a37e693162772fcda42a6f1d9ada6c0b4928f44b49450fe445facb65",
      "tokenInfo": {
        "tokenId": "_",
        "tokenAbbr": "trx",
        "tokenName": "trx",
        "tokenDecimal": 6,
        "tokenCanShow": 1,
        "tokenType": "trc10",
        "tokenLogo": "https://static.tronscan.org/production/logo/trx.png",
        "tokenLevel": "2",
        "vip": true
      },
      "transferFromAddress": "TNXoiAJ3dct8Fjg4M9fkLFh9S2v9TXc32G",
      "transferToAddress": "eeeeeeeeeeeeeeeeeeeeeeeeemm",
      "block": 63144140,
      "id": "",
      "cheatStatus": false,
      "riskTransaction": false,
      "transferFromTag": "Binance-Hot 4",
      "timestamp": 1720077252000
    }
  ],
  "contractMap": {
    "TNXoiAJ3dct8Fjg4M9fkLFh9S2v9TXc32G": false,
    "eeeeeeeeeeeeeeeeeeeeeeeeemm": false
  },
  "rangeTotal": 1,
  "contractInfo": {

  },
  "timeInterval": -1,
  "normalAddressInfo": {
    "TNXoiAJ3dct8Fjg4M9fkLFh9S2v9TXc32G": {
      "risk": false
    },
    "eeeeeeeeeeeeeeeeeeeeeeeeemm": {
      "risk": false
    }
  }
}
*/

func GetTransactionsTrc10(baseUrl, apiKey string, args *TransactionsTrc10Req) (*TransactionsTrc10Resp, error) {
	c := internal.NewClient(baseUrl, apiKey, false)
	resp := &TransactionsTrc10Resp{
		ContractMap: make(map[string]bool),
		Data:        make([]*Transactions, 0),
	}
	err := c.Get("/api/new/transfer", args.ToMap(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
