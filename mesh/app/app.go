package app

import (
	"tron_robot/internal/dao"
	"tron_robot/internal/service/basic"
	"tron_robot/internal/service/business"
	"tron_robot/internal/service/cryptocurrency"
	"tron_robot/internal/service/option"
	"tron_robot/internal/service/user"
	"tron_robot/internal/service/wallet"
	"xbase/cluster/mesh"
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
