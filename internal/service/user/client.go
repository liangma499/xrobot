package user

import (
	"xbase/transport"
	"xrobot/internal/service/user/pb"

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
