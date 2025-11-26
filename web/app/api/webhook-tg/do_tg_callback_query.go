package webhooktg

import (
	"xbase/log"
	"xbase/utils/xconv"
	"xrobot/internal/event/message"
	"xrobot/internal/xtelegram/telegram/types"
	tgtypes "xrobot/internal/xtelegram/tg-types"
)

func (a *API) botCallBackMsg(callbackQuery *types.CallbackQuery, ip, channelCode string) error {
	if callbackQuery == nil {
		return nil
	}

	msg := callbackQuery.Message
	if msg == nil {
		return nil
	}
	log.Warnf("botCommand callbackQuery:%s  msg :%s \n\r\n", callbackQuery.Data, xconv.Json(msg))
	userID := callbackQuery.From.ID
	button, inputMsg, _ := a.doGetUserMsg(callbackQuery.Data, userID)

	if button == tgtypes.XTelegramButton_None {
		log.Warnf("userbotRester:%v channelCode:%v", callbackQuery.Data, channelCode)
		return nil
	}

	commonMsg := message.MessageCommon{
		Button:          button,
		Type:            message.MessageType(msg.Chat.Type),
		ChatID:          msg.Chat.ID, //用户Telegram ID
		UserID:          userID,
		ChannelCode:     channelCode, //渠道码
		ClientIP:        ip,
		ChatInstance:    callbackQuery.ChatInstance,
		WaitforinputMsg: inputMsg,
		OrderID:         button.GetOrderID(callbackQuery.Data),
	}

	message.PublishMessageBusiness(&message.MessageBusiness{
		MessageCommon: commonMsg,
		UserName:      msg.Chat.UserName,
		InviteCode:    "",                 //邀请码
		FirstName:     msg.Chat.FirstName, //姓
		LastName:      msg.Chat.LastName,  //名

	})
	return nil

}
