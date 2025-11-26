package common

import (
	"xbase/errors"
	"xbase/transport"
	"xrobot/internal/code"
	"xrobot/internal/http/helper"

	usersvc "xrobot/internal/service/user"

	userpb "xrobot/internal/service/user/pb"

	"github.com/gin-gonic/gin"
)

// 获取用户信息
func GetUserInfo(ctx *gin.Context, transporter transport.Transporter) (*userpb.UserOneClient, *userpb.FetchUserReply, error) {
	uid := helper.GetUID(ctx)
	if uid == 0 {
		return nil, nil, errors.NewError(code.Unauthorized)
	}
	client, err := usersvc.NewClient(transporter.NewClient)
	if err != nil {
		return nil, nil, errors.NewError(code.InternalError)
	}
	reply, err := client.FetchUser(ctx, &userpb.FetchUserArgs{
		UID: uid,
	})
	if err != nil {
		return nil, nil, errors.NewError(code.InternalError)
	}
	return client, reply, nil
}
