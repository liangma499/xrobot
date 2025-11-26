package solana

import (
	"context"
	"fmt"
	"xbase/utils/xconv"

	"tron_robot/internal/cryptocurrencies/solana/internal"
	"tron_robot/internal/cryptocurrencies/solana/internal/rpc"
)

type GetblockResp struct {
	Err    any                       `json:"error"`            //	"error": null,
	ID     string                    `json:"id,omitempty"`     //"id": "getblock.io",
	Result *rpc.GetParsedBlockResult `json:"result,omitempty"` //"result": 684634
}

func (sol *solanaInfo) SolGetBlock(slot int64) (*rpc.GetParsedBlockResult, error) {
	if sol.client != nil {
		return sol.client.GetParsedBlockWithOpts(context.Background(), xconv.Uint64(slot), &rpc.GetBlockOpts{
			Encoding:                       internal.EncodingJSONParsed,
			TransactionDetails:             rpc.TransactionDetailsFull,
			Rewards:                        rpc.NewBoolean(true),
			Commitment:                     rpc.CommitmentFinalized,
			MaxSupportedTransactionVersion: &rpc.MaxSupportedTransactionVersion0,
		})
	} else if sol.httpClient != nil {
		req := &CommonReq{
			ID:      GetBlockIo,
			JsonRpc: JsonRpcVersion,
			Method:  "getBlock",
			Params: []any{
				slot,
				&GetBlockOpts{
					Encoding:                       internal.EncodingJSONParsed,
					TransactionDetails:             rpc.TransactionDetailsFull,
					Rewards:                        rpc.NewBoolean(true),
					Commitment:                     rpc.CommitmentFinalized,
					MaxSupportedTransactionVersion: &rpc.MaxSupportedTransactionVersion1,
				},
			},
		}
		resp := &GetblockResp{}
		err := sol.httpClient.Post(fmt.Sprintf("/%s", sol.httpToken), req, resp)
		if err != nil {
			return nil, err
		}
		if resp.Err != nil {
			return nil, fmt.Errorf("faild:%v", resp.Err)
		}
		return resp.Result, nil
	}
	return nil, fmt.Errorf("config is err")
}
