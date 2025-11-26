package user

import "tron_robot/internal/xtypes"

type registerArgs struct {
	account           string // 账号
	email             string // 邮箱
	password          string // 密码
	avatar            string // 头像
	birthday          string // 生日
	channelName       string // 注册名
	channelCode       string // 渠道码
	clientIP          string // 客户端IP
	language          string
	telegramUserID    int64
	telegramuserName  string
	telegramPwd       string
	inviteCode        string
	registerLoginType xtypes.RegisterLoginType
	userType          xtypes.UserType // 用户类型
	isVerifiedEmail   xtypes.UserVerifiedEmailStatus
}
