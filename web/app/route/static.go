package route

import (
	"github.com/gin-gonic/gin"
)

func InitStatic(engine *gin.Engine) {
	//engine.StaticFS("/resource/upload/avatar", gin.Dir("../resource/upload/avatar", false))
	engine.StaticFS("/resource/upload", gin.Dir("../resource/upload", false))
}
