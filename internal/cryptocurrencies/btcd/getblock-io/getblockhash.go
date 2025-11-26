package getblockio

type GetBlockHashResp struct {
	Err    any    `json:"error"`            //	"error": null,
	ID     string `json:"id,omitempty"`     //"id": "getblock.io",
	Result string `json:"result,omitempty"` //"result": 684634
}

func (gb *getBlockIO) GetBlockHash(hashID int64) (string, error) {
	req := &CommonReq{
		ID:      gb.ID,
		JsonRpc: gb.JsonRpc,
		Method:  "getblockhash",
		Params: []any{
			hashID,
		},
	}
	resp := &GetBlockHashResp{}
	err := gb.Client.Post(gb.Method, req, resp)
	if err != nil {
		return "", err
	}
	return resp.Result, nil
}
