package solana

import (
	"context"
	"fmt"
	"xbase/utils/xconv"

	"xrobot/internal/cryptocurrencies/solana/internal/rpc"
)

func (sol *solanaInfo) SolBlockNumber() (int64, error) {
	if sol.client != nil {
		slot, err := sol.client.GetSlot(context.Background(), rpc.CommitmentFinalized)
		if err != nil {
			return 0, err
		}
		return xconv.Int64(slot), nil
	}
	return 0, fmt.Errorf("config is err")
}
