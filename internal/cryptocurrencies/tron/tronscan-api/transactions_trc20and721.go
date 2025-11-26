package tronscanapi

import "tron_robot/internal/cryptocurrencies/tron/internal"

type TransactionsTrc20And721Req struct {
	Start            int64  `json:"start"`            //start Start number. Default 0
	Limit            int64  `json:"limit"`            //limit Number of items per page. Default 10
	StartTimestamp   int64  `json:"start_timestamp"`  //毫秒 start_timestamp Start time
	EndTimestamp     int64  `json:"end_timestamp"`    //毫秒 end_timestamp End time
	ContractAddress  string `json:"contract_address"` //contract_address Contract address
	Confirm          bool   `json:"confirm"`          //confirm Whether to return confirmed transfers only. Default: true
	RelatedAddress   string `json:"relatedAddress"`   //relatedAddress Account address
	FromAddress      string `json:"fromAddress"`      //fromAddress Sender's address
	ToAddress        string `json:"toAddress"`        //toAddress Recipient's address
	FilterTokenValue int    `json:"filterTokenValue"` //filterTokenValue=1
}

func (t *TransactionsTrc20And721Req) ToMap() map[string]any {
	if t == nil {
		return nil
	}
	return map[string]any{
		"start":            t.Start,
		"limit":            t.Limit,
		"start_timestamp":  t.StartTimestamp,
		"end_timestamp":    t.EndTimestamp,
		"contract_address": t.ContractAddress,
		"confirm":          t.Confirm,
		"relatedAddress":   t.RelatedAddress,
		"fromAddress":      t.FromAddress,
		"toAddress":        t.ToAddress,
		"filterTokenValue": t.FilterTokenValue,
	}
}

type ContractInfoTrc20And721 struct {
	Tag1        string `json:"tag1"`         //"tag1":"jUSDJ Token",
	Tag1Url     string `json:"tag1Url"`      //"tag1Url":"justlend.just.network",
	Name        string `json:"name"`         //"name":"CErc20Delegator",
	Vip         bool   `json:"vip"`          //"vip":false,
	OpenSource  bool   `json:"open_source"`  //"open_source":false,
	ProjectLogo string `json:"project_logo"` //"project_logo":"https://coin.top/production/upload/logo/default.png"
}
type TokenTransfersTrc20And721 struct {
	TransactionId         string     `json:"transaction_id"`        //"transaction_id":"cc74c25b36ac29c0c22ae07955ebd93d2c58654ed976cd39af3224b4292dd4d1",
	TransacstatustionId   int        `json:"status"`                //"status":0,
	BlockTs               int64      `json:"block_ts"`              //"block_ts":1680513705000,
	FromAddress           string     `json:"from_address"`          //"from_address":"TH3N6kYXow3FUP8Giyjm344qpDgjpChQx7",
	FromAddressTag        any        `json:"from_address_tag"`      //"from_address_tag":{},
	ToAddress             string     `json:"to_address"`            //"to_address":"TL5x9MtSnDy537FXKx53yAaHRRNdg9TkkA",
	ToAddressTag          any        `json:"to_address_tag"`        //"to_address_tag":{},
	Block                 int64      `json:"block"`                 //"block":49966979,
	ContractAddress       string     `json:"contract_address"`      //"contract_address":"TMwFHYXLJaRUPeW6421aqXL4ZEzPRFGkGT",
	Quant                 string     `json:"quant"`                 //"quant":"11258727923280441931",
	Confirmed             bool       `json:"confirmed"`             //"confirmed":true,
	ContractRet           string     `json:"contractRet"`           //"contractRet":"SUCCESS",
	FinalResult           string     `json:"finalResult"`           //"finalResult":"SUCCESS",
	Revert                bool       `json:"revert"`                //"revert":false,
	ContractType          string     `json:"contract_type"`         //"contract_type":"trc20",
	FromAddressIsContract bool       `json:"fromAddressIsContract"` //"fromAddressIsContract":false,
	ToAddressIsContract   bool       `json:"toAddressIsContract"`   //"toAddressIsContract":true
	TokenInfoData         *TokenInfo `json:"tokenInfo"`
}
type TransactionsTrc20And721Resp struct {
	Total             int64                               `json:"total"`
	RangeTotal        int64                               `json:"rangeTotal"`
	WholeChainTxCount int64                               `json:"wholeChainTxCount"`
	ContractInfo      map[string]*ContractInfoTrc20And721 `json:"contractInfo"`
	TokenTransfers    []*TokenTransfersTrc20And721        `json:"token_transfers"`
}

func GetTransactionsTrc20And721(baseUrl, apiKey string, args *TransactionsTrc20And721Req) (*TransactionsTrc20And721Resp, error) {
	c := internal.NewClient(baseUrl, apiKey, true)
	resp := &TransactionsTrc20And721Resp{
		ContractInfo:   make(map[string]*ContractInfoTrc20And721),
		TokenTransfers: make([]*TokenTransfersTrc20And721, 0),
	}
	err := c.Get("/api/token_trc20/transfers", args.ToMap(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
