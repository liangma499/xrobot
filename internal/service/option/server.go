package option

import (
	"context"
	optionBaseConfigCfg "xrobot/internal/option/option-base-config"
	optionCurrencyCfg "xrobot/internal/option/option-currency"
	optionCurrencyChannelCfg "xrobot/internal/option/option-currency-channel"
	optionCurrencyNetworkCfg "xrobot/internal/option/option-currency-network"
	optionTelegramCmdCfg "xrobot/internal/option/option-telegram-cmd"
	optionWithdrawCurrencyCfg "xrobot/internal/option/option-withdraw-currency"

	"fmt"
	"xbase/cluster/mesh"
	"xrobot/internal/service/option/pb"
	"xrobot/internal/xtypes"
)

// 配置中心数据 后台过来所有的配置走这个服务
const (
	serviceName = "option" // 服务名称
	servicePath = "Option" // 服务路径要与pb中的服务路径保持一致
)

var _ pb.OptionAble = &Server{}

type Server struct {
	proxy *mesh.Proxy
}

func NewServer(proxy *mesh.Proxy) *Server {
	return &Server{
		proxy: proxy,
	}
}

func (s *Server) Init() {
	s.proxy.AddServiceProvider(serviceName, servicePath, s)
	s.doLoadAllOption()

}
func (s *Server) doLoadAllOption() {
	//加载基础配置
	if err := optionBaseConfigCfg.SetOpts(context.Background(), xtypes.OptionOperate_LoadAll); err != nil {
		panic(fmt.Sprintf("optionBaseConfigCfg err:%v", err))
	}
	//加载币种配置
	if err := optionCurrencyCfg.SetOpts(context.Background(), xtypes.OptionOperate_LoadAll); err != nil {
		panic(fmt.Sprintf("optionCurrencyCfg err:%v", err))
	}
	//加载归集费用配置
	if err := optionCurrencyChannelCfg.SetOpts(context.Background()); err != nil {
		panic(fmt.Sprintf("optionCurrencyChannelCfg err:%v", err))
	}
	//加载渠道tg命令配置
	if err := optionTelegramCmdCfg.SetOpts(context.Background(), xtypes.OptionOperate_LoadAll); err != nil {
		panic(fmt.Sprintf("optionTelegramCmdCfg err:%v", err))
	}
	//加载渠道tg命令配置
	if err := optionWithdrawCurrencyCfg.SetOpts(context.Background()); err != nil {
		panic(fmt.Sprintf("optionTelegramCmdCfg err:%v", err))
	}
	if err := optionCurrencyNetworkCfg.SetOpts(context.Background(), xtypes.OptionOperate_LoadAll); err != nil {
		panic(fmt.Sprintf("optionCurrencyNetworkCfg err:%v", err))
	}

	s.initWebhook()
}

// 开卡
func (s *Server) OptionBaseConfig(ctx context.Context, args *pb.OptionBaseConfigArgs, reply *pb.OptionBaseConfigReply) error {
	return nil
}
