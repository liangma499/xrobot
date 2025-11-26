package main

import (
	"tron_robot/mesh/app"
	"xbase"
	"xbase/cache"
	credis "xbase/cache/redis"
	"xbase/cluster/mesh"
	"xbase/config"
	"xbase/config/etcd"
	"xbase/config/file"
	"xbase/eventbus"
	"xbase/eventbus/nats"
	"xbase/locate/redis"
	"xbase/log"
	"xbase/log/zap"
	"xbase/registry/consul"
	"xbase/transport/rpcx"
)

func main() {
	// 设置日志
	log.SetLogger(zap.NewLogger(zap.WithCallerSkip(2)))
	// 设置缓存
	cache.SetCache(credis.NewCache())
	// 初始化事件总线
	eventbus.SetEventbus(nats.NewEventbus())
	// 设置配置中心
	config.SetConfigurator(config.NewConfigurator(config.WithSources(file.NewSource(), etcd.NewSource())))
	// 创建容器
	container := xbase.NewContainer()
	// 创建用户定位器
	locator := redis.NewLocator()
	// 创建服务发现
	registry := consul.NewRegistry()
	// 创建RPC传输器
	transporter := rpcx.NewTransporter()
	// 创建网格组件
	component := mesh.NewMesh(
		mesh.WithLocator(locator),
		mesh.WithRegistry(registry),
		mesh.WithTransporter(transporter),
	)
	// 初始化应用
	app.Init(component.Proxy())
	// 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
