package middleware

import (
	"net/url"
	"tron_robot/internal/code"
	casbincomp "tron_robot/internal/http/casbin"
	"tron_robot/internal/http/helper"
	"tron_robot/internal/http/response"
	"xbase/codes"
	"xbase/utils/xconv"

	"github.com/gin-gonic/gin"
)

type QueryPermissionNodeFunc func(method, path string) (int, error)

func Permission(query QueryPermissionNodeFunc) gin.HandlerFunc {
	enforcer := casbincomp.Instance()

	return func(ctx *gin.Context) {
		uid := helper.GetUID(ctx)
		rid := helper.GetRID(ctx)

		if uid == 0 {
			response.Fail(ctx, codes.Unauthorized)
		}

		// super admin role
		if rid == 1 {
			ctx.Next()
			return
		}

		// TODO：放开权限, 后续删除
		if uid > 0 {
			ctx.Next()
			return
		}

		u, err := url.Parse(ctx.Request.RequestURI)
		if err != nil {
			response.Fail(ctx, codes.InternalError)
		}

		// query route node
		nid, err := query(ctx.Request.Method, u.Path)
		if err != nil {
			response.Fail(ctx, codes.InternalError)
		}

		if nid == 0 {
			response.Fail(ctx, code.NoPermission)
		}

		// check access permission of user
		if ok, _ := enforcer.Enforce(xconv.String(uid), xconv.String(nid)); !ok {
			response.Fail(ctx, code.NoPermission)
		}

		ctx.Next()
	}
}
