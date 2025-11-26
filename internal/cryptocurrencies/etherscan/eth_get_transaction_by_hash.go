package etherscan

type EthGetTransactionByHashReq struct {
	EthCommon
	Txhash string `json:"txhash"`
}
type EthGetTransactionByHashResp struct {
	ID      int64        `json:"id,omitempty"`     //"id": "getblock.io",
	JsonRpc string       `json:"jsonrpc"`          //	"error": null,
	Result  *Transaction `json:"result,omitempty"` //"result": 684634
}

func (gb *ethInfo) EthGetTransactionByHash(txhash string) (*Transaction, error) {
	req := &EthGetTransactionByHashReq{
		EthCommon: EthCommon{
			Module: "proxy",
			Action: "eth_getTransactionByHash",
			Apikey: gb.Apikey,
		},
		Txhash: txhash,
	}

	resp := &EthGetTransactionByHashResp{}
	err := gb.Client.Get("/api", req, resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}
