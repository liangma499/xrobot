package wallet

import (
	"xbase/transport"
	"xrobot/internal/service/wallet/pb"

	"github.com/smallnest/rpcx/client"
)

const target = "discovery://wallet"

func NewClient(fn transport.NewMeshClient) (*pb.WalletOneClient, error) {
	c, err := fn(target)
	if err != nil {
		return nil, err
	}

	return pb.NewWalletOneClient(c.Client().(*client.OneClient)), nil
}
