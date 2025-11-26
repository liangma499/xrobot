package getblockio

type GetBlockCountResp struct {
	Err    any    `json:"error"`            //	"error": null,
	ID     string `json:"id,omitempty"`     //"id": "getblock.io",
	Result int64  `json:"result,omitempty"` //"result": 684634
}

func (gb *getBlockIO) GetBlockCount() (int64, error) {

	req := &CommonReq{
		ID:      gb.ID,
		JsonRpc: gb.JsonRpc,
		Method:  "getblockcount",
		Params:  make([]any, 0),
	}
	resp := &GetBlockCountResp{}
	err := gb.Client.Post(gb.Method, req, resp)
	if err != nil {
		return 0, err
	}
	return resp.Result, nil
}
