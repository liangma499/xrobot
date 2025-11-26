package internal

import (
	"context"
	"fmt"
	"xbase/log"
	"xrobot/internal/cryptocurrencies/tron/transfer"
	tronscanapi "xrobot/internal/cryptocurrencies/tron/tronscan-api"
	optionChannelDao "xrobot/internal/dao/option-channel"
	"xrobot/internal/event/message"
	"xrobot/internal/model"
	optionCurrencyNetworkCfg "xrobot/internal/option/option-currency-network"
	tgmsg "xrobot/internal/xtelegram/tg-msg"
	tgtypes "xrobot/internal/xtelegram/tg-types"
	waitforinput "xrobot/internal/xtelegram/wait-for-input"
	"xrobot/internal/xtypes"

	"github.com/shopspring/decimal"
)

// "地址详情"
func AddressDetail(userBase *model.UserBase, payload *message.MessageBusiness) {
	if payload.Type != message.MessageType_Private {
		return
	}
	address := payload.WaitforinputMsg.PutMsg
	if address == "" {
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
	accountDetail := doGetAccountDetail(address)
	if accountDetail == nil {
		return
	}
	activated := "未激活"
	if accountDetail.Activated {
		activated = "激活"
	}
	withText := fmt.Sprintf("*查询地址：*\n`%s`\n账户类型：%s \n能量：%s\n带宽：%s / %s\nTRX余额：%s TRX\nUSDT余额：%s USDT\n",
		payload.WaitforinputMsg.PutMsg,
		activated,
		accountDetail.EnergyRemaining.String(),
		accountDetail.NetRemaining.String(),
		accountDetail.NetLimit.String(),
		accountDetail.Balance.String(),
		accountDetail.BalanceUSDT.String(),
	)
	xMsg, err := tgmsg.NewXTelegramMessage(channelCfg.TelegramCfg.MainRobotToken,
		tgmsg.WithText(withText),
		tgmsg.WithDebug(true),
		tgmsg.WithMsgType(tgtypes.RobotMsgTypePhoto),
		tgmsg.WithParseMode(tgtypes.ModeMarkdown))

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

type accountDetail struct {
	Activated       bool
	Balance         decimal.Decimal
	BalanceUSDT     decimal.Decimal
	EnergyRemaining decimal.Decimal
	NetLimit        decimal.Decimal
	NetRemaining    decimal.Decimal
}

func doGetAccountDetail(address string) *accountDetail {
	tfr := transfer.NewTransferNull()

	if !tfr.ValidateAddress(address) {
		log.Warnf("inValidateAddress:%v", address)
	}
	apiCfg := optionCurrencyNetworkCfg.Instance().GetNeedApiByChannelType(xtypes.NetWorkChannelType_TRON, xtypes.APITronscan)

	if apiCfg == nil {
		log.Errorf("config is nil Code:%v")
		return nil
	}
	account, err := tronscanapi.GetAccountDetailV2(apiCfg.Url, apiCfg.AppID, &tronscanapi.AccountDetailV2Req{
		//Address: "TEj8NgQM37dABXxVQRZo2b7nTkoCFM2qCQ",
		Address: address,
	})
	if err != nil {
		log.Warnf("inValidateAddress:%v err:%v", address, err)
		return nil
	}
	//未激活
	if account == nil {
		return nil
	}
	rst := &accountDetail{}
	usdtInfo := account.GetWithPriceTokensByTokenId(xtypes.USDT.Trc20Contract())
	if usdtInfo != nil {
		rst.BalanceUSDT = xtypes.CoefficientToFloat64(usdtInfo.Balance, int8(usdtInfo.TokenDecimal))
	}

	trxInfo := account.GetWithPriceTokensByTokenId(xtypes.TRX.Trc20Contract())
	if trxInfo != nil {
		rst.Balance = xtypes.CoefficientToFloat64(trxInfo.Balance, int8(trxInfo.TokenDecimal))
	}

	rst.Activated = true
	if account.Balance.LessThanOrEqual(decimal.Zero) || !account.Activated {
		rst.Activated = false
	}
	if account.Bandwidth != nil {
		rst.EnergyRemaining = account.Bandwidth.EnergyRemaining
		rst.NetLimit = account.Bandwidth.FreeNetLimit.Add(account.Bandwidth.NetLimit)
		rst.NetRemaining = account.Bandwidth.NetRemaining.Add(account.Bandwidth.FreeNetRemaining)
	}

	return rst
}
