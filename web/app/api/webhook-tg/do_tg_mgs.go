package webhooktg

import (
	"strings"
	"tron_robot/internal/event/message"
	"tron_robot/internal/xtelegram/telegram/types"
	tgtypes "tron_robot/internal/xtelegram/tg-types"
	"xbase/log"
	"xbase/utils/xconv"
)

func (a *API) botMsg(msg *types.Message, ip, channelCode string) error {
	if msg == nil {
		return nil
	}
	userID := int64(0)
	if msg.From != nil {
		userID = msg.From.ID
	}
	log.Warnf("botCommand:%s", xconv.Json(msg))
	text := strings.TrimSpace(msg.Text)
	button, inputMsg, inviteCode := a.doGetUserMsg(text, userID)
	if button == tgtypes.XTelegramButton_None {
		return nil
	}

	commonMsg := message.MessageCommon{
		Button:          button,
		Type:            message.MessageType(msg.Chat.Type),
		ChatID:          msg.Chat.ID, //用户Telegram ID
		UserID:          userID,
		ChannelCode:     channelCode, //渠道码
		ClientIP:        ip,
		WaitforinputMsg: inputMsg,
		OrderID:         button.GetOrderID(text),
	}
	switch button {
	case tgtypes.XTelegramButton_Start:

		message.PublishMessageStart(&message.MessageStart{
			MessageCommon: commonMsg,
			UserName:      msg.Chat.UserName,
			InviteCode:    inviteCode,         //邀请码
			FirstName:     msg.Chat.FirstName, //姓
			LastName:      msg.Chat.LastName,  //名

		})
		return nil
	default:
		message.PublishMessageBusiness(&message.MessageBusiness{
			MessageCommon: commonMsg,
			UserName:      msg.Chat.UserName,
			InviteCode:    inviteCode,         //邀请码
			FirstName:     msg.Chat.FirstName, //姓
			LastName:      msg.Chat.LastName,  //名

		})
		return nil
	}

}
