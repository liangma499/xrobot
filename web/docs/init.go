package docs

import (
	"xbase/etc"
	"xbase/log"
	"xbase/mode"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
)

func Init(engine *gin.Engine) {
	if !mode.IsDebugMode() {
		return
	}

	doc := &swag.Spec{}

	err := etc.Get("etc.doc").Scan(doc)
	if err != nil {
		log.Fatalf("load doc config failed: %v", err)
	}

	// 初始化文档信息
	SwaggerInfo.Version = doc.Version
	SwaggerInfo.Host = doc.Host
	SwaggerInfo.BasePath = doc.BasePath
	SwaggerInfo.Schemes = doc.Schemes
	SwaggerInfo.Title = doc.Title
	SwaggerInfo.Description = doc.Description

	// 初始化文档
	engine.GET("/swagger/*any", ginswagger.WrapHandler(swaggerfiles.Handler))
}
