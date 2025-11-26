package internal

import (
	"context"
	"xbase/log"
	optionChannelDao "xrobot/internal/dao/option-channel"
	optionListenerAddressDao "xrobot/internal/dao/option-listener-address"
	"xrobot/internal/event/message"
	"xrobot/internal/model"
	optiontelegramcmd "xrobot/internal/option/option-telegram-cmd"
	tgmsg "xrobot/internal/xtelegram/tg-msg"
	tgtemplate "xrobot/internal/xtelegram/tg-template"
	tgtypes "xrobot/internal/xtelegram/tg-types"
	waitforinput "xrobot/internal/xtelegram/wait-for-input"
	"xrobot/internal/xtypes"

	"github.com/shopspring/decimal"
)

// "🔋能量闪租笔数"
func EnergyFlashRentalBiShu(userBase *model.UserBase, payload *message.MessageBusiness) {
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
		log.Errorf("channelCfg is nil :%v", payload.ChannelCode)
		return
	}
	trc20Address := optionListenerAddressDao.Instance().GetAddressByChannelCode(payload.ChannelCode, xtypes.TRON)
	if trc20Address == "" {
		log.Errorf("trc20Address is nil :%v", payload.ChannelCode)
		return
	}
	cmdMsg := optiontelegramcmd.GetChanCodeCmd(payload.ChannelCode, tgtypes.XTelegramCmd_Button_EnergyFlashRental_BiSu)
	if cmdMsg == nil {
		return
	}
	xMsg, err := tgmsg.NewXTelegramMessage(channelCfg.TelegramCfg.MainRobotToken,
		tgmsg.WithText(cmdMsg.Text),
		tgmsg.WithDebug(true),
		tgmsg.WithCmd(cmdMsg.Cmd),
		tgmsg.WithMsgType(cmdMsg.Type),
		tgmsg.WithParseMode(cmdMsg.ParseMode),
		tgmsg.WithKeyboard(cmdMsg.Keyboard))

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
	priceNum := decimal.NewFromInt(payload.Button.Value())
	if priceNum.LessThanOrEqual(decimal.Zero) {
		return
	}
	expandMap := map[string]string{
		tgtemplate.PriceNumKey:      priceNum.String(),
		tgtemplate.PriceKey:         channelCfg.PriceDefault.TrxPriceU.Mul(priceNum).String(),
		tgtemplate.Tron20AddressKey: trc20Address,
	}
	lastMsgID := doGetLastMessgeID(userBase.UID, cmdMsg.Cmd.String())
	if lastMsgID == 0 {
		if msg, err := xMsg.SendMessage(payload.ChatID, expandMap); err != nil {
			log.Warnf("sendMessage:%v", err)
		} else {
			doSetLastMessgeID(userBase.UID, cmdMsg.Cmd.String(), int64(msg.MessageID))
			waitforinput.SetWaitForinputKey(&waitforinput.WaitforinputInfo{
				UserID:        payload.UserID, //用户TG
				InPutMsg:      payload.WaitforinputMsg.InPutMsg,
				Button:        payload.Button,
				UserBottonKey: payload.Button.WaitForInputKey(),
				Extended:      nil,
			})
		}
	} else {
		if _, err := xMsg.EditMessageText(payload.ChatID, lastMsgID, expandMap); err != nil {
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

}
