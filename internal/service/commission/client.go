package commission

import (
	"xbase/transport"
	"xrobot/internal/service/commission/pb"

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
