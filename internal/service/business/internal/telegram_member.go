package internal

import (
	"context"
	"xbase/log"
	optionChannelDao "xrobot/internal/dao/option-channel"
	"xrobot/internal/event/message"
	"xrobot/internal/model"
	tgmsg "xrobot/internal/xtelegram/tg-msg"
	tgtypes "xrobot/internal/xtelegram/tg-types"
	waitforinput "xrobot/internal/xtelegram/wait-for-input"
)

// "👑飞机会员"
func TelegramMember(userBase *model.UserBase, payload *message.MessageBusiness) {
	if payload.Type != message.MessageType_Private {
		return
	}
	ctx := context.Background()
	channelCfg, err := optionChannelDao.Instance().GetChannel(ctx, payload.ChannelCode)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	if channelCfg == nil {
		log.Errorf("channelCfg is nil")
		return
	}
	xMsg, err := tgmsg.NewXTelegramMessage(channelCfg.TelegramCfg.MainRobotToken,
		tgmsg.WithText("功能开发中..."),
		tgmsg.WithDebug(true),
		tgmsg.WithMsgType(tgtypes.RobotMsgTypePhoto),
		tgmsg.WithParseMode(tgtypes.ModeNone))

	/*
		cmdMsg := optiontelegramcmd.GetChanCodeCmd(payload.ChannelCode, tgtypes.XTelegramCmd_Start)
		if cmdMsg == nil {
			return
		}
		xMsg, err := tgmsg.NewXTelegramMessage(channelCfg.TelegramCfg.MainRobotToken, cmdMsg.Text,
			tgmsg.WithDebug(true),
			tgmsg.WithCmd(cmdMsg.Cmd),
			tgmsg.WithMsgType(cmdMsg.Type),
			tgmsg.WithParseMode(cmdMsg.ParseMode),
			tgmsg.WithKeyboard(cmdMsg.Keyboard))
	*/
	if err != nil {
		return
	}
	if xMsg == nil {
		return
	}
	if _, err := xMsg.SendMessage(payload.ChatID, nil); err != nil {
		log.Warnf("sendMessage:%v", err)
	} else {
		waitforinput.SetWaitForinputKey(&waitforinput.WaitforinputInfo{
			UserID:        payload.UserID, //用户TG
			InPutMsg:      payload.WaitforinputMsg.InPutMsg,
			Button:        payload.Button,
			UserBottonKey: payload.Button.WaitForInputKey(),
			Extended:      nil,
		})
	}

}
