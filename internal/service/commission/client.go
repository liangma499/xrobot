package commission

import (
	"tron_robot/internal/service/commission/pb"
	"xbase/transport"

	"github.com/smallnest/rpcx/client"
)

const target = "discovery://commission"

func NewClient(fn transport.NewMeshClient) (*pb.CommissionOneClient, error) {
	c, err := fn(target)
	if err != nil {
		return nil, err
	}

	return pb.NewCommissionOneClient(c.Client().(*client.OneClient)), nil
}
