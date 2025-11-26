package tronscanapi

import "xrobot/internal/cryptocurrencies/tron/internal"

type TransfersWithStatusReq struct {
	Start     int64  `json:"start"`      //Start index，default is 0
	Limit     int64  `json:"limit"`      //Number of transfers per page
	Trc20Id   string `json:"trc20Id"`    //TRC20 token address Contract Address  必须传
	Address   string `json:"address"`    //Account address
	Direction int    `json:"direction"`  //0: all.  1: transfer-out. 2: transfer-in
	DbVersion int    `json:"db_version"` //Whether to include approval transfers.  1: include. 0: exclude
	Reverse   bool   `json:"reverse"`    //Sort by creation time. Valid values: true/false*/

}

func (t *TransfersWithStatusReq) ToMap() map[string]any {
	if t == nil {
		return nil
	}
	return map[string]any{
		"start":      t.Start,
		"limit":      t.Limit,
		"trc20Id":    t.Trc20Id,
		"address":    t.Address,
		"direction":  t.Direction,
		"db_version": t.DbVersion,
		"reverse":    t.Reverse,
	}
}

//https://apilist.tronscanapi.com/api/accountv2?address=TEj8NgQM37dABXxVQRZo2b7nTkoCFM2qCQ
//https://apilist.tronscanapi.com/api/token_trc20/transfers-with-status?limit=10&start=0&trc20Id=TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t&address=THfVXfjsHuJ24eeVUdcPs15pZhaKbp6uQo

type TransfersWithStatus struct {
	Amount           string `json:"amount"`           //"amount":"2531319816832000000000000",
	Status           int    `json:"status"`           //"status":0,
	Approval_amount  string `json:"approval_amount"`  //"approval_amount":"0",
	BlockTimestamp   int64  `json:"block_timestamp"`  //"block_timestamp":1680520824000,
	Block            int64  `json:"block"`            //"block":49969352,
	From             string `json:"from"`             //"from":"TV6MuMXfmLbBqPZvBHdwFsDnQeVfnmiuSi",
	To               string `json:"to"`               //"to":"TNXoiAJ3dct8Fjg4M9fkLFh9S2v9TXc32G",
	Hash             string `json:"hash"`             //"hash":"09e76a0a358c205cfaaf366f340982024bdc9184c46fbad340003fdd34b8178d",
	Contract_address string `json:"contract_address"` //"contract_address":"TUpMhErZL2fhh4sVNULAbNKLokS4GjC1F4",
	Confirmed        int    `json:"confirmed"`        //"confirmed":0,
	Contract_type    string `json:"contract_type"`    //"contract_type":"TriggerSmartContract",
	ContractType     int    `json:"contractType"`     //"contractType":31,
	Revert           int    `json:"revert"`           //"revert":0,
	Contract_ret     string `json:"contract_ret"`     //"contract_ret":"SUCCESS",
	Final_result     string `json:"final_result"`     //"final_result":"SUCCESS",
	Event_type       string `json:"event_type"`       //"event_type":"Transfer",
	Issue_address    string `json:"issue_address"`    //"issue_address":"TRX6Q82wMqWNbCCiLqejbZe43wk1h1zJHm",
	Decimals         int    `json:"decimals"`         //"decimals":18,
	Token_name       string `json:"token_name"`       //"token_name":"TrueUSD",
	Id               string `json:"id"`               //"id":"TUpMhErZL2fhh4sVNULAbNKLokS4GjC1F4",
	Direction        int    `json:"direction"`        //"direction":1
}

type TransfersWithStatusResp struct {
	Total             int64                  `json:"total"`
	RangeTotal        int64                  `json:"rangeTotal"`
	WholeChainTxCount int64                  `json:"wholeChainTxCount"`
	ContractMap       map[string]bool        `json:"contractMap"`
	TokenInfoData     *TokenInfo             `json:"tokenInfo"`
	Data              []*TransfersWithStatus `json:"data"`
}

func GetTransfersWithStatus(baseUrl, apiKey string, args *TransfersWithStatusReq) (*TransfersWithStatusResp, error) {
	c := internal.NewClient(baseUrl, apiKey, false)
	resp := &TransfersWithStatusResp{
		ContractMap: make(map[string]bool),
		Data:        make([]*TransfersWithStatus, 0),
	}
	err := c.Get("/api/token_trc20/transfers-with-status", args.ToMap(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
