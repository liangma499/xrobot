package etherscan

type EthGetTransactionReceiptReq struct {
	EthCommon
	Txhash string `json:"txhash"`
}
type EthGetTransactionReceiptResp struct {
	ID      int64               `json:"id,omitempty"`     //"id": "getblock.io",
	JsonRpc string              `json:"jsonrpc"`          //	"error": null,
	Result  *TransactionReceipt `json:"result,omitempty"` //"result": 684634
}

func (gb *ethInfo) EthGetTransactionReceipt(txhash string) (*TransactionReceipt, error) {
	req := &EthGetTransactionReceiptReq{
		EthCommon: EthCommon{
			Module: "proxy",
			Action: "eth_getTransactionReceipt",
			Apikey: gb.Apikey,
		},
		Txhash: txhash,
	}

	resp := &EthGetTransactionReceiptResp{}
	err := gb.Client.Get("/api", req, resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}
