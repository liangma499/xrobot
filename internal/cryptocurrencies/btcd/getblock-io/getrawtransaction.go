package getblockio

type GetRawTransactionResp struct {
	Err    any                 `json:"error"`            //	"error": null,
	ID     string              `json:"id,omitempty"`     //"id": "getblock.io",
	Result *TransactionRawInfo `json:"result,omitempty"` //"result": 684634
}

func (gb *getBlockIO) GetRawTransaction(blockHash, txHash string) (*GetRawTransactionResp, error) {
	req := &CommonReq{
		ID:      gb.ID,
		JsonRpc: gb.JsonRpc,
		Method:  "getrawtransaction",
		Params: []any{
			txHash,
			true,
			blockHash,
		},
	}
	resp := &GetRawTransactionResp{}
	err := gb.Client.Post(gb.Method, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
