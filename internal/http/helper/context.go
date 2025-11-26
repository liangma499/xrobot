package helper

import "github.com/gin-gonic/gin"

const (
	UIDKey         = "_uid" // 用户ID
	RIDKey         = "_rid"
	ACCKey         = "_acc"       // 用户账号
	TgUidKey       = "_tguserid"  // TGUID
	UserMailKey    = "_useremail" // 邮箱账号
	ChannelCodeKey = "_c_code"
	ChannelNameKey = "_c_name"
)

// SetUID 设置用户ID
func SetUID(ctx *gin.Context, uid int64) {
	ctx.Set(UIDKey, uid)
}

// GetUID 获取用户ID
func GetUID(ctx *gin.Context) int64 {
	return ctx.GetInt64(UIDKey)
}

// SetUID 设置用户ID
func SetRID(ctx *gin.Context, rid int64) {
	ctx.Set(RIDKey, rid)
}

// GetUID 获取用户ID
func GetRID(ctx *gin.Context) int64 {
	return ctx.GetInt64(RIDKey)
}

// SetACC 设置账号
func SetACC(ctx *gin.Context, acc string) {
	ctx.Set(ACCKey, acc)
}

// GetACC 获取账号
func GetACC(ctx *gin.Context) string {
	return ctx.GetString(ACCKey)
}

// SetTGUIDKey 设置TGID
func SetTgUidKey(ctx *gin.Context, acc string) {
	ctx.Set(TgUidKey, acc)
}

// GetTGUIDKey 获取TG用户ID
func GetTgUidKey(ctx *gin.Context) string {
	return ctx.GetString(TgUidKey)
}

// SetUserMail 设置用户邮箱
func SetUserMail(ctx *gin.Context, acc string) {
	ctx.Set(UserMailKey, acc)
}

// GetUserMail 获取用户邮箱
func GetUserMail(ctx *gin.Context) string {
	return ctx.GetString(UserMailKey)
}

// SetChannelCode 设置渠道码
func SetChannelCode(ctx *gin.Context, channelCode string) {
	ctx.Set(ChannelCodeKey, channelCode)
}

// GetChannelCode 获取渠道码
func GetChannelCode(ctx *gin.Context) string {
	return ctx.GetString(ChannelCodeKey)
}

// SetChannelName 设置渠道名
func SetChannelName(ctx *gin.Context, channelName string) {
	ctx.Set(ChannelNameKey, channelName)
}

// GetChannelName 获取渠道名
func GetChannelName(ctx *gin.Context) string {
	return ctx.GetString(ChannelNameKey)
}
