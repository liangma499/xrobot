package business

import (
	"tron_robot/internal/service/business/pb"
	"xbase/transport"

	"github.com/smallnest/rpcx/client"
)

const target = "discovery://business"

func NewClient(fn transport.NewMeshClient) (*pb.BusinessOneClient, error) {
	c, err := fn(target)
	if err != nil {
		return nil, err
	}

	return pb.NewBusinessOneClient(c.Client().(*client.OneClient)), nil
}
