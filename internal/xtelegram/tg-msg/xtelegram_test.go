package tgmsg_test

import (
	"context"
	"fmt"
	"testing"
	optionChannelDao "tron_robot/internal/dao/option-channel"
	optionListenerAddressDao "tron_robot/internal/dao/option-listener-address"
	optiontelegramcmd "tron_robot/internal/option/option-telegram-cmd"
	"tron_robot/internal/utils/xstr"
	tgmsg "tron_robot/internal/xtelegram/tg-msg"
	tgtemplate "tron_robot/internal/xtelegram/tg-template"
	tgtypes "tron_robot/internal/xtelegram/tg-types"
	"tron_robot/internal/xtypes"
	"xbase/config"
	"xbase/config/etcd"
	"xbase/config/file"
	"xbase/log"
	"xbase/utils/xconv"
)

const (
	bot_token = "6867997452:AAFYZXHAC_TDvcfBiYto2ShutRSiUcboa04"
)

func init() {
	// 设置配置中心
	config.SetConfigurator(config.NewConfigurator(config.WithSources(file.NewSource(), etcd.NewSource())))
}
func TestClient_SendMessage(t *testing.T) {

	msg := `🎉 *Congratulations,${winner_name}!* 🎉
	You won with a *${multiplier}x* multiplier!
		This will earn you a total of *${win_amount}${currency}*.
	You're on fire! 🔥 Keep it up, aim for the moon!`
	replaces := make(map[string]string)

	replaces["winner_name"] = "黄老师"
	replaces["win_amount"] = fmt.Sprintf("%.4f", 1.0)
	replaces["multiplier"] = fmt.Sprintf("%.4f", 1.1)
	replaces["game_name"] = "黄老师"
	replaces["currency"] = "csd"

	xMsg, err := tgmsg.NewXTelegramMessage(bot_token,
		tgmsg.WithText(msg),
		tgmsg.WithDebug(true),
		tgmsg.WithMessageThreadID(-1002066585210),
		tgmsg.WithMsgType(tgtypes.RobotMsgTypePhoto),
		tgmsg.WithParseMode(tgtypes.ModeMarkdown))
	if err == nil {
		log.Errorf("%v", err)
		return
	}
	xMsg.SendMessage(33, replaces)

}

func TestClient_SendMessageCmd(t *testing.T) {

	ctx := context.Background()
	channelCfg, err := optionChannelDao.Instance().GetChannel(ctx, xtypes.OfficialChannelCode)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	if channelCfg == nil {
		log.Errorf("channelCfg is nil")
		return
	}
	trc20Address := optionListenerAddressDao.Instance().GetAddressByChannelCode(xtypes.OfficialChannelCode, xtypes.TRON)
	if trc20Address == "" {
		log.Errorf("trc20Address is nil :%v", xtypes.OfficialChannelCode)
		return
	}
	cmdMsg := optiontelegramcmd.GetChanCodeCmd(xtypes.OfficialChannelCode, tgtypes.XTelegramCmd_Button_RechargeOtherAddresses)
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
	if err != nil {
		return
	}

	if xMsg == nil {
		return
	}
	orderID := xstr.SerialNO()
	if err := xMsg.SetExtraCallBackData(orderID); err != nil {
		log.Warnf("%v", err)
		return
	}
	expandMap := map[string]string{
		tgtemplate.ComboKindEnergyFlashRentalNumKey:  xconv.String(channelCfg.ChannelCfg.ComboKindEnergyFlashRental.Duration),
		tgtemplate.ComboKindEnergyFlashRentalNameKey: channelCfg.ChannelCfg.ComboKindEnergyFlashRental.ComboKind.Name(),
		tgtemplate.PriceNumKey:                       "1",
		tgtemplate.NotActivatedAddressCountKey:       "2",
		tgtemplate.ReceivingAddressCountKey:          "3",
		tgtemplate.EnergyFeeKey:                      "4",
		tgtemplate.ActivationfeeKey:                  "5",
		tgtemplate.PayAmountKey:                      "6",
		tgtemplate.Tron20AddressKey:                  trc20Address,
	}
	if _, err := xMsg.SendMessage(7026994919, expandMap); err != nil {
		log.Warnf("sendMessage:%v", err)
	}

}
