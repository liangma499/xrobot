package cryptocurrency

import (
	"xbase/transport"
	"xrobot/internal/service/cryptocurrency/pb"

	"github.com/smallnest/rpcx/client"
)

const target = "discovery://cryptocurrency"

func NewClient(fn transport.NewMeshClient) (*pb.CryptoCurrencyOneClient, error) {
	c, err := fn(target)
	if err != nil {
		return nil, err
	}

	return pb.NewCryptoCurrencyOneClient(c.Client().(*client.OneClient)), nil
}
