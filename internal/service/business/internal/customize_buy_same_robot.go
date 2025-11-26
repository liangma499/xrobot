package internal

import (
	"context"
	"xbase/log"
	optionChannelDao "xrobot/internal/dao/option-channel"
	optionListenerAddressDao "xrobot/internal/dao/option-listener-address"
	paymentAmountUserDao "xrobot/internal/dao/payment-amount-user"
	"xrobot/internal/event/message"
	"xrobot/internal/model"
	optiontelegramcmd "xrobot/internal/option/option-telegram-cmd"
	tgmsg "xrobot/internal/xtelegram/tg-msg"
	tgtemplate "xrobot/internal/xtelegram/tg-template"
	tgtypes "xrobot/internal/xtelegram/tg-types"
	"xrobot/internal/xtypes"

	"github.com/shopspring/decimal"
)

// "🏆定制同款机器人"
func CustomizeBuySameRobot(userBase *model.UserBase, payload *message.MessageBusiness) {

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

	cmdMsg := optiontelegramcmd.GetChanCodeCmd(payload.ChannelCode, tgtypes.XTelegramCmd_BuySameRobot)
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

	currency := xtypes.USDT
	trc20Address := optionListenerAddressDao.Instance().GetAddressByChannelCode(payload.ChannelCode, xtypes.TRON)
	if trc20Address == "" {
		log.Errorf("trc20Address is nil :%v", payload.ChannelCode)
		return
	}
	payAmountNew, err := paymentAmountUserDao.Instance().GetUniqueAmount(currency, customizeBalance)
	if err != nil {
		return
	}

	expandMap := map[string]string{
		tgtemplate.PayAmountKey:     payAmountNew.String(),
		tgtemplate.Tron20AddressKey: trc20Address,
	}
	doSendCreateOrder(channelCfg,
		userBase,
		payload,
		cmdMsg,
		xtypes.Usage_BuySameRobot,
		payAmountNew,
		expandMap,
		xtypes.ComboKind_Hour.ExpirationTime(1),
		&model.AmountUserInfo{
			AddressInfo: nil,
			MessageID:   0,
			BiShu:       0,
		},
		currency)

}
