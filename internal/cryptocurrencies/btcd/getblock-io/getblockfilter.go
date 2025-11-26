package getblockio

type GetblockfilterResp struct {
	Err    any        `json:"error"`            //	"error": null,
	ID     string     `json:"id,omitempty"`     //"id": "getblock.io",
	Result *BlockInfo `json:"result,omitempty"` //"result": 684634
}

func (gb *getBlockIO) Getblockfilter(hash string) (*GetblockfilterResp, error) {
	req := &CommonReq{
		ID:      gb.ID,
		JsonRpc: gb.JsonRpc,
		Method:  "getblockfilter",
		Params: []any{
			hash,
			"2",
		},
	}
	resp := &GetblockfilterResp{}
	err := gb.Client.Post(gb.Method, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
