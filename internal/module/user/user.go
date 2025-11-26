package user

import (
	"tron_robot/internal/xtypes"
)

type User struct {
	ID           int64           // 用户ID
	Code         string          // 用户编号
	Nickname     string          // 用户昵称
	Avatar       string          // 用户头像
	Email        string          // 用户邮箱
	UserType     xtypes.UserType // 用户类型
	ChannelCode  string          // 渠道ID
	OnlineStatus string          // 是否在线
}
