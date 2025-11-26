package commission

import (
	"context"

	"xbase/cluster/mesh"

	"tron_robot/internal/service/commission/pb"
)

const (
	serviceName = "commission" // 服务名称
	servicePath = "Commission" // 服务路径要与pb中的服务路径保持一致
)

var _ pb.CommissionAble = &Server{}

type Server struct {
	proxy *mesh.Proxy
	//mu    sync.Mutex
}

func NewServer(proxy *mesh.Proxy) *Server {
	return &Server{
		proxy: proxy,
	}
}

func (s *Server) Init() {
	s.proxy.AddServiceProvider(serviceName, servicePath, s)
}

// 拉取用户佣金
func (s *Server) FetchUserCommission(ctx context.Context, args *pb.FetchUserCommissionArgs, reply *pb.FetchUserCommissionReply) error {
	return nil
}
