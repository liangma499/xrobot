package app

import (
	"xbase/cluster/mesh"
	"xrobot/internal/dao"
	"xrobot/internal/service/basic"
	"xrobot/internal/service/business"
	"xrobot/internal/service/cryptocurrency"
	"xrobot/internal/service/option"
	"xrobot/internal/service/user"
	"xrobot/internal/service/wallet"
)

func Init(proxy *mesh.Proxy) {

	//初始化表
	dao.InitTableOption()
	dao.InitTableUser()
	// 初始化后台配置相关服务服务
	option.NewServer(proxy).Init()
	// 初始化基础服务
	basic.NewServer(proxy).Init()
	// 初始化用户服务
	user.NewServer(proxy).Init()
	// 初始化业务服务 TG机器消息相关
	business.NewServer(proxy).Init()
	// 初始化钱包服务
	wallet.NewServer(proxy).Init()
	// 初始化区块信息服务
	cryptocurrency.NewServer(proxy).Init()
}
