package internal

import (
	"context"
	"xbase/errors"
	"xbase/log"
	"xbase/utils/xconv"
	"xrobot/internal/cryptocurrencies/tron/transfer"
	tronscanapi "xrobot/internal/cryptocurrencies/tron/tronscan-api"
	optionChannelDao "xrobot/internal/dao/option-channel"
	optionListenerAddressDao "xrobot/internal/dao/option-listener-address"
	paymentAmountUserDao "xrobot/internal/dao/payment-amount-user"
	"xrobot/internal/event/message"
	"xrobot/internal/model"
	optionCurrencyNetworkCfg "xrobot/internal/option/option-currency-network"
	optiontelegramcmd "xrobot/internal/option/option-telegram-cmd"
	parseInPutMsg "xrobot/internal/xtelegram/parse_in_put_msg"
	tgmsg "xrobot/internal/xtelegram/tg-msg"
	tgtemplate "xrobot/internal/xtelegram/tg-template"
	tgtypes "xrobot/internal/xtelegram/tg-types"
	waitforinput "xrobot/internal/xtelegram/wait-for-input"
	"xrobot/internal/xtypes"

	"xrobot/internal/code"

	"github.com/shopspring/decimal"
)

func RechargeOtherAddresses(userBase *model.UserBase, payload *message.MessageBusiness) {
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
	trc20Address := optionListenerAddressDao.Instance().GetAddressByChannelCode(payload.ChannelCode, xtypes.TRON)
	if trc20Address == "" {
		log.Errorf("trc20Address is nil :%v", payload.ChannelCode)
		return
	}
	if payload.WaitforinputMsg.InPutMsg {
		doRechargeOtherAddressesInPutMsg(channelCfg, userBase, payload, trc20Address)
	} else {
		doRechargeOtherAddresses(channelCfg, payload)
	}

}
func doRechargeOtherAddressesInPutMsg(channelCfg *model.OptionChannel, userBase *model.UserBase, payload *message.MessageBusiness, trc20Address string) {
	if channelCfg.PriceDefault == nil {
		log.Warnf("channelCfg.priceDefault is nil")
		return
	}
	if channelCfg.ChannelCfg == nil {
		log.Warnf("channelCfg.ChannelCfg is nil")
		return
	}
	//defer waitforinput.DelWaitForinputKey(payload.UserID)
	lastCallbackData := ""
	if payload.WaitforinputMsg.Extended != nil {
		lastCallbackData = payload.WaitforinputMsg.Extended.LastCallbackData
	}
	if lastCallbackData == "" {
		log.Warnf("no callback data")
		return
	}
	cmdMsg := optiontelegramcmd.GetChanCodeCmd(payload.ChannelCode, tgtypes.XTelegramCmd_Button_RechargeOtherAddresses)
	if cmdMsg == nil {
		return
	}

	// 解析callbackData 获取需要操作的地址
	lastButton := tgtypes.StringToXTelegramButton(lastCallbackData)
	if lastButton == tgtypes.XTelegramButton_None {
		log.Warnf("lastButton is None")
		return
	}
	count := lastButton.Value()
	if count < 1 {
		log.Warnf("count is not in range")
		return
	}
	countDec := decimal.NewFromInt(xconv.Int64(count))
	addresses := parseInPutMsg.ParseInPutMsg(payload.WaitforinputMsg.PutMsg)
	if addresses == nil {

		return
	}
	countAddress := len(addresses)
	if countAddress == 0 {
		log.Warnf("addresses len is zero")
		return
	}
	countAddressDec := decimal.NewFromInt32(xconv.Int32(countAddress))
	//这里需要检测地址的合理性
	addressInfo, err := doAddressesInfo(channelCfg, addresses)
	if err != nil {
		return
	}
	//能量费用
	energyFee := channelCfg.PriceDefault.TrxPriceU.Mul(countAddressDec).Mul(countDec)
	//激活费用
	activationfee := channelCfg.ChannelCfg.ActivationFee.Mul(addressInfo.notActivated)
	//支付金额
	payAmount := energyFee.Add(activationfee)

	currency := xtypes.TRX
	payAmountNew, err := paymentAmountUserDao.Instance().GetUniqueAmount(currency, payAmount)
	if err != nil {
		return
	}
	comboKind := channelCfg.ChannelCfg.ComboKindEnergyFlashRental.ComboKind
	duration := channelCfg.ChannelCfg.ComboKindEnergyFlashRental.Duration
	expandMap := map[string]string{
		tgtemplate.ComboKindEnergyFlashRentalNumKey:  xconv.String(duration),
		tgtemplate.ComboKindEnergyFlashRentalNameKey: comboKind.Name(),
		tgtemplate.PriceNumKey:                       countDec.String(),
		tgtemplate.NotActivatedAddressCountKey:       addressInfo.notActivated.String(),
		tgtemplate.ReceivingAddressCountKey:          countAddressDec.String(),
		tgtemplate.EnergyFeeKey:                      energyFee.String(),
		tgtemplate.ActivationfeeKey:                  activationfee.String(),
		tgtemplate.PayAmountKey:                      payAmountNew.String(),
		tgtemplate.Tron20AddressKey:                  trc20Address,
	}

	doSendCreateOrder(channelCfg,
		userBase,
		payload,
		cmdMsg,
		xtypes.Usage_Recharge_OtherAddress,
		payAmountNew,
		expandMap,
		comboKind.ExpirationTime(duration),
		&model.AmountUserInfo{
			AddressInfo: addressInfo.addressInfo.Clone(),
			MessageID:   0,
			BiShu:       count,
		},
		currency)
}

type addressesInfoRst struct {
	addressInfo  model.AddressInfoType
	notActivated decimal.Decimal
}

func doAddressesInfo(channelCfg *model.OptionChannel, addresses []string) (*addressesInfoRst, error) {

	apiCfg := optionCurrencyNetworkCfg.Instance().GetNeedApiByChannelType(xtypes.NetWorkChannelType_TRON, xtypes.APITronscan)

	if apiCfg == nil {
		log.Errorf("config is nil Code:%v", channelCfg.ChannelCode)
		return nil, errors.NewError(code.NotFound)
	}
	var (
		notActivated = int64(0)
	)

	tfr := transfer.NewTransferNull()
	rst := &addressesInfoRst{
		addressInfo:  make(model.AddressInfoType, 0),
		notActivated: decimal.Zero,
	}
	for _, item := range addresses {
		if !tfr.ValidateAddress(item) {
			log.Warnf("inValidateAddress:%v", item)
			return nil, errors.NewError(code.InValidateAddress)
		}
		account, err := tronscanapi.GetAccountDetailV2(apiCfg.Url, apiCfg.AppID, &tronscanapi.AccountDetailV2Req{
			//Address: "TEj8NgQM37dABXxVQRZo2b7nTkoCFM2qCQ",
			Address: item,
		})
		if err != nil {
			log.Warnf("inValidateAddress:%v err:%v", item, err)
			return nil, errors.NewError(code.InValidateAddress, err)
		}
		//未激活
		if account == nil {
			return nil, errors.NewError(code.InValidateAddress, err)
		}
		activated := true
		if account.Balance.LessThanOrEqual(decimal.Zero) || !account.Activated {
			notActivated++
			activated = false
		}
		rst.addressInfo = append(rst.addressInfo, &model.AddressInfo{
			Address:   item,
			Activated: activated,
		})
	}
	rst.notActivated = decimal.NewFromInt(notActivated)
	return rst, nil
}

func doRechargeOtherAddresses(channelCfg *model.OptionChannel, payload *message.MessageBusiness) {
	lastCallBack := waitforinput.Instance().GetLastCallBack(payload.UserID)
	lastButton := tgtypes.StringToXTelegramButton(lastCallBack)
	if lastButton == tgtypes.XTelegramButton_None {
		log.Warnf("lastButton is None")
		return
	}
	count := lastButton.Value()
	if count < 1 {
		log.Warnf("count is not in range")
		return
	}

	xMsg, err := tgmsg.NewXTelegramMessage(channelCfg.TelegramCfg.MainRobotToken,
		tgmsg.WithText(tgtemplate.RechargeOtherAddresses()),
		tgmsg.WithDebug(true),
		tgmsg.WithMsgType(tgtypes.RobotMsgTypeText),
		tgmsg.WithParseMode(tgtypes.ModeMarkdown))

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
			Extended: &message.WaitforinputExtended{
				LastCallbackData: lastCallBack,
			},
		})
	}
}
