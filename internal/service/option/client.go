package option

import (
	"xbase/transport"
	"xrobot/internal/service/option/pb"

	"github.com/smallnest/rpcx/client"
)

const target = "discovery://option"

func NewClient(fn transport.NewMeshClient) (*pb.OptionOneClient, error) {
	c, err := fn(target)
	if err != nil {
		return nil, err
	}

	return pb.NewOptionOneClient(c.Client().(*client.OneClient)), nil
}
