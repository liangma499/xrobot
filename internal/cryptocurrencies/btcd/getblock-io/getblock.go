package getblockio

type GetblockResp struct {
	Err    any        `json:"error"`            //	"error": null,
	ID     string     `json:"id,omitempty"`     //"id": "getblock.io",
	Result *BlockInfo `json:"result,omitempty"` //"result": 684634
}

func (gb *getBlockIO) Getblock(hash string) (*GetblockResp, error) {
	req := &CommonReq{
		ID:      gb.ID,
		JsonRpc: gb.JsonRpc,
		Method:  "getblock",
		Params: []any{
			hash,
			2,
		},
	}
	resp := &GetblockResp{}
	err := gb.Client.Post(gb.Method, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
