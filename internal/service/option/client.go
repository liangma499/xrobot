package option

import (
	"tron_robot/internal/service/option/pb"
	"xbase/transport"

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
