package response

import (
	"net/http"
	"xbase/codes"
	"xrobot/internal/code"
	"xrobot/internal/http/common"

	"github.com/gin-gonic/gin"
)

const ExitPanic = "http response and exit"

// Fail 返回错误
func Fail(ctx *gin.Context, code *codes.Code) {
	Response(ctx, code)
}

// Success 返回成功
func Success(ctx *gin.Context, data ...any) {
	Response(ctx, code.OK, data...)
}

// Response 响应消息
func Response(ctx *gin.Context, code *codes.Code, data ...any) {
	//log.Debug(code.String())

	if len(data) > 0 {
		ctx.JSON(http.StatusOK, &common.Res{Code: code.Code(), Data: data[0]})
	} else {
		ctx.JSON(http.StatusOK, &common.Res{Code: code.Code()})
	}

	panic(ExitPanic)
}

// Redirect 重定向
func Redirect(ctx *gin.Context, location string, code ...int) {
	if len(code) > 0 {
		ctx.Redirect(code[0], location)
	} else {
		ctx.Redirect(http.StatusFound, location)
	}

	panic(ExitPanic)
}
