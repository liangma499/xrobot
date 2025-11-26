package tronscanapi

import (
	"xrobot/internal/cryptocurrencies/tron/internal"
)

type TransactionInfoReq struct {
	Hash string `json:"hash,omitempty"`
}

func (t *TransactionInfoReq) ToMap() map[string]any {
	if t == nil {
		return nil
	}
	return map[string]any{
		"hash": t.Hash,
	}
}

// https://apilist.tronscanapi.com/api/accountv2?address=TEj8NgQM37dABXxVQRZo2b7nTkoCFM2qCQ
// https://apilist.tronscanapi.com/api/token_trc20/transfers-with-status?limit=10&start=0&trc20Id=TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t&address=THfVXfjsHuJ24eeVUdcPs15pZhaKbp6uQo

type TransactionInfoResp struct {
	Block                int64           `json:"block,omitempty"`
	Hash                 string          `json:"hash,omitempty"`
	Timestamp            int64           `json:"timestamp,omitempty"`
	OwnerAddress         string          `json:"ownerAddress,omitempty"`
	SignatureAddresses   any             `json:"signature_addresses,omitempty"`
	ContractType         int32           `json:"contractType,omitempty"`
	ToAddress            string          `json:"toAddress,omitempty"`
	Confirmations        int32           `json:"confirmations,omitempty"`
	Confirmed            bool            `json:"confirmed,omitempty"`
	Revert               bool            `json:"revert,omitempty"`
	ContractRet          string          `json:"contractRet,omitempty"`
	ContractData         *ContractData   `json:"contractData,omitempty"`
	Cost                 *Cost           `json:"cost,omitempty"`
	Data                 string          `json:"data,omitempty"`
	TriggerInfo          any             `json:"trigger_info,omitempty"`
	InternalTransactions any             `json:"internal_transactions,omitempty"`
	SrConfirmList        []*ConfirmList  `json:"srConfirmList,omitempty"`
	Info                 any             `json:"info,omitempty"`
	AddressTag           any             `json:"addressTag,omitempty"`
	ContractInfo         any             `json:"contractInfo,omitempty"`
	ContractMap          map[string]bool `json:"contract_map,omitempty"`
}

func GetTransactionInfo(baseUrl, apiKey string, args *TransactionInfoReq) (*TransactionInfoResp, error) {
	c := internal.NewClient(baseUrl, apiKey, false)
	resp := &TransactionInfoResp{
		ContractMap:   make(map[string]bool),
		SrConfirmList: make([]*ConfirmList, 0),
	}
	err := c.Get("/api/transaction-info", args.ToMap(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
