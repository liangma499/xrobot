package business

import (
	"context"
	"sync"

	"xbase/cluster/mesh"

	"tron_robot/internal/event/message"
	"tron_robot/internal/service/business/pb"
)

const (
	serviceName = "business" // 服务名称
	servicePath = "Business" // 服务路径要与pb中的服务路径保持一致
)

var _ pb.BusinessAble = &Server{}

type Server struct {
	proxy *mesh.Proxy
	mu    sync.Mutex
}

func NewServer(proxy *mesh.Proxy) *Server {
	return &Server{
		proxy: proxy,
	}
}

func (s *Server) Init() {
	s.proxy.AddServiceProvider(serviceName, servicePath, s)
	message.SubscribeMessageBusiness(s.doSubscribeMessageBusiness)
}

// 拉取用户佣金
func (s *Server) Business(ctx context.Context, args *pb.BusinessArgs, reply *pb.BusinessReply) error {
	return nil
}
