package tronscanapi

import (
	"tron_robot/internal/cryptocurrencies/tron/internal"
)

type TransactionReq struct {
	Count           bool  `json:"count,omitempty"`
	Start           int64 `json:"start,omitempty"`           //Start number. Default 0
	Limit           int64 `json:"limit,omitempty"`           //Number of items per page. Default 10
	Start_timestamp int64 `json:"start_timestamp,omitempty"` //Start time
	End_timestamp   int64 `json:"end_timestamp,omitempty"`   //End time
}

func (t *TransactionReq) ToMap() map[string]any {
	if t == nil {
		return nil
	}
	return map[string]any{
		"sort":            "+timestamp",
		"count":           t.Count,
		"start":           t.Start,
		"limit":           t.Limit,
		"start_timestamp": t.Start_timestamp,
		"end_timestamp":   t.End_timestamp,
	}
}

// https://apilist.tronscanapi.com/api/accountv2?address=TEj8NgQM37dABXxVQRZo2b7nTkoCFM2qCQ
// https://apilist.tronscanapi.com/api/token_trc20/transfers-with-status?limit=10&start=0&trc20Id=TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t&address=THfVXfjsHuJ24eeVUdcPs15pZhaKbp6uQo

type TransactionResp struct {
	Total             int64              `json:"total,omitempty"`
	RangeTotal        int64              `json:"rangeTotal,omitempty"`
	Data              []*TransactionData `json:"data,omitempty"`
	WholeChainTxCount int64              `json:"wholeChainTxCount,omitempty"`
	ContractMap       map[string]bool    `json:"contractMap,omitempty"`
	ContractInfo      any                `json:"contractInfo,omitempty"`
}

func GetTransaction(baseUrl, apiKey string, args *TransactionReq) (*TransactionResp, error) {
	c := internal.NewClient(baseUrl, apiKey, true)
	resp := &TransactionResp{
		ContractMap: make(map[string]bool),
		Data:        make([]*TransactionData, 0),
	}
	err := c.Get("/api/transaction", args.ToMap(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
