package tgmsg

import (
	tgbutton "xrobot/internal/xtelegram/tg-button"
	tgtypes "xrobot/internal/xtelegram/tg-types"
)

type XTelegramMessage struct {
	botToken string //渠道code
	msg      *telegramMessage
}
type telegramMessage struct {
	telegramPwd     string
	cmd             tgtypes.XTelegramCmd     //推送方式
	keyboard        *tgbutton.TelegramButton //按钮
	url             string                   //图片
	msgType         tgtypes.RobotMsgType     //文件类型(1=图片,2=视频)
	text            string                   //文本内容
	debug           bool
	messageThreadID int64  //是否要发送到群里
	parseMode       string //编码方式
}
