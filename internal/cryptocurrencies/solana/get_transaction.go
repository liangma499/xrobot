package solana

import (
	"context"
	"fmt"

	"xrobot/internal/cryptocurrencies/solana/internal"
	"xrobot/internal/cryptocurrencies/solana/internal/rpc"
)

type TransactionResp struct {
	Err    any                             `json:"error"`            //	"error": null,
	ID     string                          `json:"id,omitempty"`     //"id": "getblock.io",
	Result *rpc.GetParsedTransactionResult `json:"result,omitempty"` //"result": 684634
}

func (sol *solanaInfo) SolGetTransaction(txSig string) (*rpc.GetParsedTransactionResult, error) {
	if sol.client != nil {
		return sol.client.GetParsedTransaction(context.Background(),
			internal.MustSignatureFromBase58(txSig),
			&rpc.GetParsedTransactionOpts{
				Commitment:                     rpc.CommitmentFinalized,
				MaxSupportedTransactionVersion: &rpc.MaxSupportedTransactionVersion0,
			})
	} else if sol.httpClient != nil {
		req := &CommonReq{
			ID:      GetBlockIo,
			JsonRpc: JsonRpcVersion,
			Method:  "getTransaction",
			Params: []any{
				txSig,
				&GetParsedTransactionOpts{
					Encoding:                       internal.EncodingJSONParsed,
					Commitment:                     rpc.CommitmentFinalized,
					MaxSupportedTransactionVersion: &rpc.MaxSupportedTransactionVersion1,
				},
			},
		}
		resp := &TransactionResp{}
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
