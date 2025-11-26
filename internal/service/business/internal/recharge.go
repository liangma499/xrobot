package internal

import (
	"context"
	"xbase/log"
	optionChannelDao "xrobot/internal/dao/option-channel"
	optionListenerAddressDao "xrobot/internal/dao/option-listener-address"
	"xrobot/internal/event/message"
	"xrobot/internal/model"
	"xrobot/internal/xtypes"
)

func Recharge(userBase *model.UserBase, payload *message.MessageBusiness) {
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
	trxValue := payload.Button.Value()
	if trxValue <= 0 {

	}

	if payload.WaitforinputMsg.InPutMsg {
		doRechargeOtherAddressesInPutMsg(channelCfg, userBase, payload, trc20Address)
	} else {
		doRechargeOtherAddresses(channelCfg, payload)
	}

}
