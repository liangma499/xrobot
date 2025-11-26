package main

import (
	"tron_robot/internal/http"
	"tron_robot/web/app/route"
	"tron_robot/web/docs"
	"xbase"
	"xbase/config"
	"xbase/config/etcd"
	"xbase/config/file"
	"xbase/eventbus"
	"xbase/eventbus/nats"
	"xbase/log"
	"xbase/log/zap"
	"xbase/registry/consul"
	"xbase/transport/rpcx"

	"github.com/gin-gonic/gin"
)

// //go:embed all:resource/dist
// var Static embed.FS

func main() {
	// 设置日志
	log.SetLogger(zap.NewLogger(zap.WithCallerSkip(2)))
	// 创建容器
	container := xbase.NewContainer()
	// 初始化事件总线
	eventbus.SetEventbus(nats.NewEventbus())
	// 创建服务发现
	registry := consul.NewRegistry()
	// 设置配置中心
	config.SetConfigurator(config.NewConfigurator(config.WithSources(file.NewSource(), etcd.NewSource())))
	// 创建RPC传输器
	transporter := rpcx.NewTransporter(rpcx.WithClientDiscovery(registry))
	// 创建HTTP组件
	component := http.NewHttp(func(engine *gin.Engine) {
		route.InitStatic(engine)

		route.Init(engine, transporter)

		docs.Init(engine)
	})
	// 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
