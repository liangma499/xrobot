package bnbscan

type BnbGetBlockByNumberReq struct {
	BnbCommon
	Tag     string `json:"tag"`
	Boolean bool   `json:"boolean"`
}
type BnbGetBlockByNumberReqResp struct {
	ID      int64      `json:"id,omitempty"`     //"id": "getblock.io",
	JsonRpc string     `json:"jsonrpc"`          //	"error": null,
	Result  *BlockInfo `json:"result,omitempty"` //"result": 684634
}

func (gb *ethInfo) BnbGetBlockByNumber(blockNum int64) (*BlockInfo, error) {
	req := &BnbGetBlockByNumberReq{
		BnbCommon: BnbCommon{
			Module: "proxy",
			Action: "eth_getBlockByNumber",
			Apikey: gb.Apikey,
		},
		Tag:     gb.Init64Hex(blockNum),
		Boolean: true,
	}

	resp := &BnbGetBlockByNumberReqResp{}
	err := gb.Client.Get("/api", req, resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}
