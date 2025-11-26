package etherscan

type EthBlockNumberReq struct {
	EthCommon
}
type EthBlockNumberResp struct {
	ID      int64  `json:"id,omitempty"`     //"id": "getblock.io",
	JsonRpc string `json:"jsonrpc"`          //	"error": null,
	Result  string `json:"result,omitempty"` //"result": 684634
}

func (gb *ethInfo) EthBlockNumber() (int64, error) {
	req := &EthBlockNumberReq{
		EthCommon: EthCommon{
			Module: "proxy",
			Action: "eth_blockNumber",
			Apikey: gb.Apikey,
		},
	}

	resp := &EthBlockNumberResp{}
	err := gb.Client.Get("/api", req, resp)
	if err != nil {
		return -1, err
	}
	blockNumber, err := gb.HexToInit64(resp.Result)
	if err != nil {
		return -1, err
	}

	return blockNumber.IntPart(), nil
}
