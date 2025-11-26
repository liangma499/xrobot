package middleware

import (
	"xbase/codes"
	"xbase/utils/xconv"
	"xrobot/internal/code"
	jwtcomp "xrobot/internal/component/jwt"
	"xrobot/internal/http/helper"
	"xrobot/internal/http/response"

	"xbase/utils/jwt"

	"github.com/gin-gonic/gin"
)

func Auth(isMustAuth ...bool) gin.HandlerFunc {
	http := jwtcomp.Instance().Http()

	return func(ctx *gin.Context) {
		r, err := http.Middleware(ctx.Request)

		if err != nil {
			if len(isMustAuth) > 0 && isMustAuth[0] {
				switch {
				case jwt.IsMissingToken(err):
					response.Fail(ctx, code.Unauthorized)
				case jwt.IsInvalidToken(err):
					response.Fail(ctx, code.Unauthorized)
				case jwt.IsIdentityMissing(err):
					response.Fail(ctx, code.Unauthorized)
				case jwt.IsExpiredToken(err):
					response.Fail(ctx, code.AuthorizationExpired)
				case jwt.IsAuthElsewhere(err):
					response.Fail(ctx, code.AuthorizationElsewhere)
				default:
					response.Fail(ctx, codes.Unauthorized)
				}
			}
		} else {
			identity, err := http.ExtractIdentity(r)
			if err != nil {
				if len(isMustAuth) > 0 && isMustAuth[0] {
					response.Fail(ctx, codes.Unauthorized)
				}
			} else {
				helper.SetUID(ctx, xconv.Int64(identity))
			}
			payload, err := http.ExtractPayload(r)
			if err != nil {
				if len(isMustAuth) > 0 && isMustAuth[0] {
					response.Fail(ctx, codes.Unauthorized)
				}
			} else {
				//账号
				helper.SetACC(ctx, xconv.String(payload[helper.ACCKey]))
				//TGUID
				helper.SetTgUidKey(ctx, xconv.String(payload[helper.TgUidKey]))
				// 邮箱账号
				helper.SetUserMail(ctx, xconv.String(payload[helper.UserMailKey]))
				//渠道编号
				helper.SetChannelCode(ctx, xconv.String(payload[helper.ChannelCodeKey]))
				//渠道名
				helper.SetChannelName(ctx, xconv.String(payload[helper.ChannelNameKey]))
			}
		}

		ctx.Next()
	}
}
