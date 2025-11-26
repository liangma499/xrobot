package internal

import (
	"context"
	optionChannelDao "tron_robot/internal/dao/option-channel"
	"tron_robot/internal/event/message"
	"tron_robot/internal/model"
	optiontelegramcmd "tron_robot/internal/option/option-telegram-cmd"
	tgmsg "tron_robot/internal/xtelegram/tg-msg"
	tgtemplate "tron_robot/internal/xtelegram/tg-template"
	tgtypes "tron_robot/internal/xtelegram/tg-types"
	waitforinput "tron_robot/internal/xtelegram/wait-for-input"
	"xbase/log"

	"github.com/shopspring/decimal"
)

// "🏆定制同款机器人"
func CustomizeTheSameRobot(userBase *model.UserBase, payload *message.MessageBusiness) {

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

	cmdMsg := optiontelegramcmd.GetChanCodeCmd(payload.ChannelCode, tgtypes.XTelegramCmd_CustomizeTheSameRobot)
	if cmdMsg == nil {
		return
	}
	xMsg, err := tgmsg.NewXTelegramMessage(channelCfg.TelegramCfg.MainRobotToken,
		tgmsg.WithDebug(true),
		tgmsg.WithText(cmdMsg.Text),
		tgmsg.WithCmd(cmdMsg.Cmd),
		tgmsg.WithMsgType(cmdMsg.Type),
		tgmsg.WithParseMode(cmdMsg.ParseMode),
		tgmsg.WithKeyboard(cmdMsg.Keyboard))

	if err != nil {
		return
	}
	if xMsg == nil {
		return
	}

	customizeBalance := decimal.Zero
	if channelCfg.PriceCustomize != nil {
		customizeBalance = channelCfg.PriceCustomize.CustomizeBalance
	} else if channelCfg.PriceDefault != nil {

		customizeBalance = channelCfg.PriceDefault.CustomizeBalance
	}
	if customizeBalance.LessThanOrEqual(decimal.Zero) {
		return
	}
	expandMap := map[string]string{
		tgtemplate.CustomizeBalance: customizeBalance.String(),
		tgtemplate.CustomerKey:      channelCfg.ChannelCfg.Customer,
	}
	if _, err := xMsg.SendMessage(payload.ChatID, expandMap); err != nil {
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
