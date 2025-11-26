package route

import (
	"xrobot/internal/http/middleware"

	"xbase/transport"
	webhooktg "xrobot/web/app/api/webhook-tg"

	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine, transporter transport.Transporter) {

	engine.Use(middleware.Cors())
	base := engine.Group("/", middleware.Auth())
	//auth := engine.Group("/", middleware.Auth(true))

	//tg 机器人模块
	{
		// webhook监听
		api := webhooktg.NewAPI(transporter)
		// 路由组
		group := base.Group("/tg")
		group.POST("/webhook", api.TelegramWebhook)
	}

}
