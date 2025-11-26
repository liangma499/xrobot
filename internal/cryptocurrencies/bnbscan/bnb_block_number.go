package bnbscan

type BnbBlockNumberReq struct {
	BnbCommon
}
type BnbBlockNumberResp struct {
	ID      int64  `json:"id,omitempty"`     //"id": "getblock.io",
	JsonRpc string `json:"jsonrpc"`          //	"error": null,
	Result  string `json:"result,omitempty"` //"result": 684634
}

func (gb *ethInfo) BnbBlockNumber() (int64, error) {
	req := &BnbBlockNumberReq{
		BnbCommon: BnbCommon{
			Module: "proxy",
			Action: "eth_blockNumber",
			Apikey: gb.Apikey,
		},
	}

	resp := &BnbBlockNumberResp{}
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
