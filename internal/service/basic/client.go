package basic

import (
	"tron_robot/internal/service/basic/pb"
	"xbase/transport"

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
