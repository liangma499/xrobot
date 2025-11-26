package basic

import (
	"xbase/transport"
	"xrobot/internal/service/basic/pb"

	"github.com/smallnest/rpcx/client"
)

const target = "discovery://basic"

func NewClient(fn transport.NewMeshClient) (*pb.BasicOneClient, error) {
	c, err := fn(target)
	if err != nil {
		return nil, err
	}

	return pb.NewBasicOneClient(c.Client().(*client.OneClient)), nil
}
