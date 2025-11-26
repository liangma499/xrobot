package bnbscan

type BnbGetTransactionReceiptReq struct {
	BnbCommon
	Txhash string `json:"txhash"`
}
type BnbGetTransactionReceiptResp struct {
	ID      int64               `json:"id,omitempty"`     //"id": "getblock.io",
	JsonRpc string              `json:"jsonrpc"`          //	"error": null,
	Result  *TransactionReceipt `json:"result,omitempty"` //"result": 684634
}

func (gb *ethInfo) BnbGetTransactionReceipt(txhash string) (*TransactionReceipt, error) {
	req := &BnbGetTransactionReceiptReq{
		BnbCommon: BnbCommon{
			Module: "proxy",
			Action: "eth_getTransactionReceipt",
			Apikey: gb.Apikey,
		},
		Txhash: txhash,
	}

	resp := &BnbGetTransactionReceiptResp{}
	err := gb.Client.Get("/api", req, resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}
