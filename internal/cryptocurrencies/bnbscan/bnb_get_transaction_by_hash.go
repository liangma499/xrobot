package bnbscan

type BnbGetTransactionByHashReq struct {
	BnbCommon
	Txhash string `json:"txhash"`
}
type BnbGetTransactionByHashResp struct {
	ID      int64        `json:"id,omitempty"`     //"id": "getblock.io",
	JsonRpc string       `json:"jsonrpc"`          //	"error": null,
	Result  *Transaction `json:"result,omitempty"` //"result": 684634
}

func (gb *ethInfo) BnbGetTransactionByHash(txhash string) (*Transaction, error) {
	req := &BnbGetTransactionByHashReq{
		BnbCommon: BnbCommon{
			Module: "proxy",
			Action: "eth_getTransactionByHash",
			Apikey: gb.Apikey,
		},
		Txhash: txhash,
	}

	resp := &BnbGetTransactionByHashResp{}
	err := gb.Client.Get("/api", req, resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}
