package common

import (
	"fmt"
	"xbase/network/http"
)

/*
const url = `${apiUrl}/${walletAddress}/transactions/trc20`;

	{
	   transaction_id: '60e97cbaa581b70febeccf155c41534bfe39cab273eb6d1f1967e9a71a200f5c',
	   token_info: {
		 symbol: 'USDT',
		 address: 'TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t',
		 decimals: 6,
		 name: 'Tether USD'
	   },
	   block_timestamp: 1716886755000,
	   from: 'TJDENsfBJs4RFETt1X1W8wMDc8M5XnJhCe',
	   to: 'cccccccccccccccccccc',
	   type: 'Transfer',
	   value: '9000000'
	 }

const apiUrl = 'https://api.trongrid.io/v1/accounts';
const walletAddress = 'cccccccccccccccccccc';
*/
type TokenInfo struct {
	Symbol   string `json:"symbol"`
	Address  string `json:"address"`
	Decimals int32  `json:"decimals"`
	Name     string `json:"name"`
}

type TransactionTrc20 struct {
	TransactionID  string    `json:"transaction_id"`
	TokenInfo      TokenInfo `json:"token_info"`
	BlockTimestamp int64     `json:"block_timestamp"`
	From           string    `json:"from"`
	To             string    `json:"to"`
	Type           string    `json:"type"`
	Value          string    `json:"value"`
}

func (t *TransactionTrc20) Clone() *TransactionTrc20 {
	return &TransactionTrc20{
		TransactionID: t.TransactionID,
		TokenInfo: TokenInfo{
			Symbol:   t.TokenInfo.Symbol,
			Address:  t.TokenInfo.Address,
			Decimals: t.TokenInfo.Decimals,
			Name:     t.TokenInfo.Name,
		},
		BlockTimestamp: t.BlockTimestamp,
		From:           t.From,
		To:             t.To,
		Type:           t.Type,
		Value:          t.Value,
	}
}

type TransactionResponse struct {
	Data []*TransactionTrc20 `json:"data"`
}

type TransactionReq struct {

	//true | false. If false, it returns both confirmed and unconfirmed transactions. If no param is specified, it returns both confirmed and unconfirmed transactions. Cannot be used at the same time with only_unconfirmed param.
	OnlyConfirmed bool `json:"only_confirmed"`
	//true | false. If false, it returns both confirmed and unconfirmed transactions. If no param is specified, it returns both confirmed and unconfirmed transactions. Cannot be used at the same time with only_confirmed param.
	OnlyUnconfirmed bool `json:"only_unconfirmed"`
	//number of transactions per page, default 20, max 200
	Limit int32 `json:"limit"`
	//fingerprint of the last transaction returned by the previous page; when using it, the other parameters and filters should remain the same
	Fingerprint string `json:"fingerprint"`
	//block_timestamp,asc | block_timestamp,desc (default)
	OrderBy string `json:"order_by"`
	//minimum block_timestamp, default 0
	MinTimestamp int64 `json:"min_timestamp"`
	//maximum block_timestamp, default now
	MaxTimestamp int64 `json:"max_timestamp"`
	//contract address in base58 or hex
	ContractAddress string `json:"contract_address"`
	//true | false. If true, only transactions to this address, default: false
	OnlyTo bool `json:"only_to"`
	//true | false. If true, only transactions from this address, default: false
	OnlyFrom bool `json:"only_from"`
}

const (
	base_main_url       = "base_main_url"
	transactions_api    = "/accounts/%s/transactions/trc20"
	contractAddress_api = "/contracts/%s/transactions"
)

//const walletAddress = 'cccccccccccccccccccc'

type clientBase struct {
	client *http.Client
}

func newClientGet() *clientBase {
	c := &clientBase{client: http.NewClient()}

	c.client.SetBaseUrl("https://api.trongrid.io/v1")
	c.client.SetHeaders(map[string]string{
		"Accept": "application/json",
	})
	return c
}

// Get 执行Get请求
func (c *clientBase) get(url string, req, resp any) error {
	return c.request(http.MethodGet, url, req, resp)
}

// 执行请求
func (c *clientBase) request(method string, url string, req, resp any) error {
	res, err := c.client.Request(method, url, req)
	if err != nil {
		return err
	}

	return res.ScanBody(resp)
}

func TransactionsTrc20(walletAddress string, req *TransactionReq) *TransactionResponse {
	resData := &TransactionResponse{
		Data: make([]*TransactionTrc20, 0),
	}
	c := newClientGet()
	c.get(fmt.Sprintf(transactions_api, walletAddress), req, resData)

	return resData
}

type ContractsTransactionReq struct {

	//true | false. If false, it returns both confirmed and unconfirmed transactions. If no param is specified, it returns both confirmed and unconfirmed transactions. Cannot be used at the same time with only_unconfirmed param.
	OnlyConfirmed bool `json:"only_confirmed"`
	//true | false. If false, it returns both confirmed and unconfirmed transactions. If no param is specified, it returns both confirmed and unconfirmed transactions. Cannot be used at the same time with only_confirmed param.
	OnlyUnconfirmed bool `json:"only_unconfirmed"`
	//Minimal block timestamp
	MinBlockTimestamp int64 `json:"min_block_timestamp"`
	//Maximal block timestamp
	MaxBlockTimestamp int64 `json:"max_block_timestamp"`
	//block_timestamp,asc | block_timestamp,desc (default)
	OrderBy string `json:"order_by"`
	//fingerprint of the last transaction returned by the previous page; when using it, the other parameters and filters should remain the same
	Fingerprint string `json:"fingerprint"`
	//number of transactions per page, default 20, max 200
	Limit int32 `json:"limit"`
	//true (default) | false. If true, query params applied to both normal and internal transactions. If false, query params only applied to normal transactions.
	SearchInternal bool `json:"search_internal"`
}

/*
	{
		"ret": [
		  {
			"contractRet": "SUCCESS"
		  }
		],
		"signature": [
		  "2c527348c6c982d4aeb1fd075b35cd1a4b09226dc5ebe402546061d5a8c0bb53d90ce89d4fadd52e323746d65f81311d4d2f77ab5a28f6ebef7ebd4ee2d0b48101"
		],
		"txID": "5e8e3cdf239f31c678b6119c95c76c402a71f50947d960c3905825ee648cd75a",
		"net_usage": 345,
		"raw_data_hex": "0a0285372208f373eda9e68b4dc14098a2f1d582325aae01081f12a9010a31747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e54726967676572536d617274436f6e747261637412740a15418ca132af72e7864eee48eac163a1fb4028662091121541a614f803b6fd780986a42c78ec9c7f77e6ded13c2244a9059cbb000000000000000000000041838eed07fa72a857058e4f9f98bf0e6a2f1af473000000000000000000000000000000000000000000000000000000000000465170d5dbedd58232900180c2d72f",
		"net_fee": 0,
		"energy_usage": 31895,
		"block_timestamp": "1718703846000",
		"blockNumber": "62686538",
		"energy_fee": 0,
		"energy_usage_total": 31895,
		"raw_data": {
		  "contract": [
			{
			  "parameter": {
				"value": {
				  "data": "a9059cbb000000000000000000000041838eed07fa72a857058e4f9f98bf0e6a2f1af4730000000000000000000000000000000000000000000000000000000000004651",
				  "owner_address": "418ca132af72e7864eee48eac163a1fb4028662091",
				  "contract_address": "41a614f803b6fd780986a42c78ec9c7f77e6ded13c"
				},
				"type_url": "type.googleapis.com/protocol.TriggerSmartContract"
			  },
			  "type": "TriggerSmartContract"
			}
		  ],
		  "ref_block_bytes": "8537",
		  "ref_block_hash": "f373eda9e68b4dc1",
		  "expiration": 1718703903000,
		  "fee_limit": 100000000,
		  "timestamp": 1718703844821
		},
		"internal_transactions": []
	  }

],
"success": true,

	"meta": {
	  "at": 1718703849933,
	  "fingerprint": "TmGrm87pwxo5LxaKFHALctkQmHPKAfAhHAZu35eVXLtAAgTTmkS8pF1PMShfc7JTX97eD1DqBfRhubRVixt8WhBi1f1cbqHEcYFjfBDRG1ihq1ccT9L9DK6UrLZ2rxmyzPi4XEVJgvivg5ELDR7XcxCzrZ6eeu3rzChLEx6ekcnaLH6t8e83rUaMAAGrmnWAw5jZ395CRxgkUkf7wej9MtKeCJXaA",
	  "links": {
		"next": "https://api.trongrid.io/v1/contracts/TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t/transactions?fingerprint=TmGrm87pwxo5LxaKFHALctkQmHPKAfAhHAZu35eVXLtAAgTTmkS8pF1PMShfc7JTX97eD1DqBfRhubRVixt8WhBi1f1cbqHEcYFjfBDRG1ihq1ccT9L9DK6UrLZ2rxmyzPi4XEVJgvivg5ELDR7XcxCzrZ6eeu3rzChLEx6ekcnaLH6t8e83rUaMAAGrmnWAw5jZ395CRxgkUkf7wej9MtKeCJXaA"
	  },
	  "page_size": 20
	}
*/
type Ret struct {
	ContractRet string `json:"contractRet"`
}
type ValueData struct {
	Data            string `json:"data"`
	OwnerAddress    string `json:"owner_address"`
	ContractAddress string `json:"contract_address"`
}
type Parameter struct {
	ValueData ValueData `json:"value"`
	TypeUrl   string    `json:"type_url"`
}
type Contract struct {
	Parameter Parameter `json:"parameter"`
	Type      string    `json:"type"`
}
type RawData struct {
	Contract      []Contract `json:"contract"`
	RefBlockBytes string     `json:"ref_block_bytes"`
	RefBlockHash  string     `json:"ref_block_hash"`
	Expiration    int64      `json:"expiration"`
	FeeLimit      int64      `json:"fee_limit"`
	Timestamp     int64      `json:"timestamp"`
}
type Links struct {
	Next string `json:"next"`
}
type Meta struct {
	At          int64  `json:"at"`
	Fingerprint string `json:"fingerprint"`
	Links       Links  `json:"links"`
}

type ContractsTransactionsData struct {
	Ret              []Ret    `json:"ret"`
	Signature        []string `json:"signature"`
	TxID             string   `json:"txID"`
	NetUsage         int64    `json:"net_usage"`
	RawDataHex       string   `json:"raw_data_hex"`
	NetFee           int64    `json:"type"`
	Energy_usage     int64    `json:"energy_usage"`
	Block_timestamp  string   `json:"block_timestamp"`
	BlockNumber      string   `json:"blockNumber"`
	EnergyFee        int64    `json:"energy_fee"`
	EnergyUsageTotal int64    `json:"energy_usage_total"`
	RawData          RawData  `json:"raw_data"`
}

func (c *ContractsTransactionsData) Clone() *ContractsTransactionsData {
	data := &ContractsTransactionsData{
		TxID:             c.TxID,
		NetUsage:         c.NetUsage,
		RawDataHex:       c.RawDataHex,
		NetFee:           c.NetFee,
		Energy_usage:     c.Energy_usage,
		Block_timestamp:  c.Block_timestamp,
		BlockNumber:      c.BlockNumber,
		EnergyFee:        c.EnergyFee,
		EnergyUsageTotal: c.EnergyUsageTotal,
	}
	data.Ret = make([]Ret, 0)
	for _, item := range c.Ret {
		data.Ret = append(data.Ret, Ret{
			ContractRet: item.ContractRet,
		})
	}
	data.Signature = make([]string, 0)

	data.Signature = append(data.Signature, c.Signature...)

	data.RawData = RawData{
		Contract:      make([]Contract, 0),
		RefBlockBytes: c.RawData.RefBlockBytes,
		RefBlockHash:  c.RawData.RefBlockHash,
		Expiration:    c.RawData.Expiration,
		FeeLimit:      c.RawData.FeeLimit,
		Timestamp:     c.RawData.Timestamp,
	}
	for _, item := range c.RawData.Contract {
		data.RawData.Contract = append(data.RawData.Contract, Contract{
			Parameter: Parameter{
				ValueData: ValueData{
					Data:            item.Parameter.ValueData.Data,
					OwnerAddress:    item.Parameter.ValueData.OwnerAddress,
					ContractAddress: item.Parameter.ValueData.ContractAddress,
				},
				TypeUrl: item.Parameter.TypeUrl,
			},
			Type: item.Type,
		})
	}
	return data
}

type ContractsTransactionsResponse struct {
	Data     []*ContractsTransactionsData `json:"data"`
	Success  bool                         `json:"success"`
	Meta     Meta                         `json:"meta"`
	PageSize int32                        `json:"page_size"`
}

func ContractsTransactions(contractAddress string, req *ContractsTransactionReq) *ContractsTransactionsResponse {
	resData := &ContractsTransactionsResponse{
		Data: make([]*ContractsTransactionsData, 0),
	}
	c := newClientGet()
	c.get(fmt.Sprintf(contractAddress_api, contractAddress), req, resData)

	return resData
}
