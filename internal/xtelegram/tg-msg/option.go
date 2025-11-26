package tgmsg

import (
	tgbutton "xrobot/internal/xtelegram/tg-button"
	tgtypes "xrobot/internal/xtelegram/tg-types"
)

type Option func(o *telegramMessage)

// 按钮绑定参数密码
func WithTelegramPwd(telegramPwd string) Option {
	return func(o *telegramMessage) { o.telegramPwd = telegramPwd }
}

// 文本
func WithText(text string) Option {
	return func(o *telegramMessage) { o.text = text }
}

// 推送方式
func WithCmd(cmd tgtypes.XTelegramCmd) Option {
	return func(o *telegramMessage) { o.cmd = cmd }
}

// 按钮
func WithKeyboard(keyboard *tgbutton.TelegramButton) Option {
	return func(o *telegramMessage) { o.keyboard = keyboard.Clone() }
}

// 链接地址
func WithURL(url string) Option {
	return func(o *telegramMessage) { o.url = url }
}

// 文件类型(1=图片,2=视频)
func WithMsgType(msgType tgtypes.RobotMsgType) Option {
	return func(o *telegramMessage) { o.msgType = msgType }
}

// 是否开调式
func WithDebug(debug bool) Option {
	return func(o *telegramMessage) { o.debug = debug }
}

// 是否要发送到群里
func WithMessageThreadID(messageThreadID int64) Option {
	return func(o *telegramMessage) { o.messageThreadID = messageThreadID }
}

// 是否要发送到群里
func WithParseMode(parseMode tgtypes.ParseMode) Option {
	return func(o *telegramMessage) { o.parseMode = parseMode.String() }
}
