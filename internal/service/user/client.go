package user

import (
	"tron_robot/internal/service/user/pb"
	"xbase/transport"

	"github.com/smallnest/rpcx/client"
)

const target = "discovery://user"

func NewClient(fn transport.NewMeshClient) (*pb.UserOneClient, error) {
	c, err := fn(target)
	if err != nil {
		return nil, err
	}

	return pb.NewUserOneClient(c.Client().(*client.OneClient)), nil
}
