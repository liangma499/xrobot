package internal

import (
	"context"
	"fmt"
	"xbase/log"
	"xbase/utils/xconv"
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

// "🔋能量闪租"
func EnergyFlashRental(userBase *model.UserBase, payload *message.MessageBusiness) {
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
	if channelCfg.PriceDefault == nil {
		log.Errorf("channelCfg priceDefault is nil")
		return
	}
	trc20Address := optionListenerAddressDao.Instance().GetAddressByChannelCode(payload.ChannelCode, xtypes.TRON)
	if trc20Address == "" {
		log.Errorf("trc20Address is nil :%v", payload.ChannelCode)
		return
	}
	cmdMsg := optiontelegramcmd.GetChanCodeCmd(payload.ChannelCode, tgtypes.XTelegramCmd_Button_EnergyFlashRental)
	if cmdMsg == nil {
		return
	}
	xMsg, err := tgmsg.NewXTelegramMessage(channelCfg.TelegramCfg.MainRobotToken,
		tgmsg.WithText(cmdMsg.Text),
		tgmsg.WithDebug(true),
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

	expandMap := map[string]string{
		tgtemplate.EnergySavingsKey:                  channelCfg.ChannelCfg.EnergySavings,
		tgtemplate.PriceKey:                          channelCfg.PriceDefault.TrxPriceU.String(),
		tgtemplate.PriceNoUKey:                       channelCfg.PriceDefault.TrxPriceNoU.String(),
		tgtemplate.PriceBiShuMaxKey:                  xconv.String(channelCfg.ChannelCfg.PriceBiShuMax),
		tgtemplate.Tron20AddressKey:                  trc20Address,
		tgtemplate.ComboKindEnergyFlashRentalNumKey:  xconv.String(channelCfg.ChannelCfg.ComboKindEnergyFlashRental.Duration),
		tgtemplate.ComboKindEnergyFlashRentalNameKey: channelCfg.ChannelCfg.ComboKindEnergyFlashRental.ComboKind.Name(),
	}

	for i := 1; i <= 3; i++ {
		priceNumKey := fmt.Sprintf(tgtemplate.PriceNumIndexKey, i)
		price := channelCfg.PriceDefault.TrxPriceU.Mul(decimal.NewFromInt(int64(i)))
		expandMap[priceNumKey] = price.String()

		priceBiShuKey := fmt.Sprintf(tgtemplate.PriceBiShuKey, i)
		expandMap[priceBiShuKey] = xconv.String(i)

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
