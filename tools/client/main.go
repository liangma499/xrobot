package main

import (
	"tron_robot/tools/client/app"

	"xbase"
	"xbase/cluster/client"
	"xbase/eventbus"
	"xbase/eventbus/nats"
	"xbase/log"
	"xbase/log/zap"
	"xbase/network/ws"
)

func main() {
	// 设置日志
	log.SetLogger(zap.NewLogger(zap.WithCallerSkip(2)))
	// 初始化事件总线
	eventbus.SetEventbus(nats.NewEventbus())
	// 创建容器
	container := xbase.NewContainer()
	// 创建客户端组件
	component := client.NewClient(
		client.WithClient(ws.NewClient()),
	)
	// 初始化事件和路由
	app.Init(component.Proxy())
	// 添加客户端组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
