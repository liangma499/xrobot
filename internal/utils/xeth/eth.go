package xeth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const url = "https://cloudflare-eth.com"

func Header(ctx context.Context, number ...int64) (*types.Header, error) {
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return nil, err
	}

	if len(number) > 0 {
		return client.HeaderByNumber(ctx, big.NewInt(number[0]))
	} else {
		return client.HeaderByNumber(ctx, nil)
	}
}

func HeaderNumber(ctx context.Context, number ...int64) (int64, error) {
	header, err := Header(ctx, number...)
	if err != nil {
		return 0, err
	}

	return header.Number.Int64(), nil
}

func HeaderHash(ctx context.Context, number ...int64) (string, error) {
	header, err := Header(ctx, number...)
	if err != nil {
		return "", err
	}

	return header.Hash().Hex(), nil
}
