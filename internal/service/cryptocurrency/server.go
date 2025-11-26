package cryptocurrency

import (
	"context"
	"sync"
	"xbase/cluster/mesh"
	"xrobot/internal/event/cryptocurrencyevent"
	"xrobot/internal/service/cryptocurrency/pb"
)

// 配置中心数据 后台过来所有的配置走这个服务
const (
	serviceName = "cryptocurrency" // 服务名称
	servicePath = "CryptoCurrency" // 服务路径要与pb中的服务路径保持一致
)

var _ pb.CryptoCurrencyAble = &Server{}

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
	cryptocurrencyevent.SubscribeCryptoCurrency(s.doSubscribeCryptoCurrency)
	s.doBalanceTimer()
}

// 拉取用户佣金
func (s *Server) CryptoCurrency(ctx context.Context, args *pb.CryptoCurrencyArgs, reply *pb.CryptoCurrencyReply) error {
	return nil
}
