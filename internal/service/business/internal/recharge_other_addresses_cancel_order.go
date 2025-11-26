package internal

import (
	"context"
	"xbase/log"
	optionChannelDao "xrobot/internal/dao/option-channel"
	paymentAmountUserDao "xrobot/internal/dao/payment-amount-user"
	"xrobot/internal/event/message"
	"xrobot/internal/model"
	tgmsg "xrobot/internal/xtelegram/tg-msg"
	tgtypes "xrobot/internal/xtelegram/tg-types"
	waitforinput "xrobot/internal/xtelegram/wait-for-input"

	"gorm.io/gorm/clause"
)

// "取消订单"
func RechargeOtherAddressesCancelOrder(userBase *model.UserBase, payload *message.MessageBusiness) {
	if payload.Type != message.MessageType_Private {
		return
	}
	if payload.OrderID == "" {
		log.Errorf("orderID is nil")
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

	amountUser, err := paymentAmountUserDao.Instance().FindOne(ctx, func(cols *paymentAmountUserDao.Columns) any {

		return clause.Eq{
			Column: cols.OrderID,
			Value:  payload.OrderID,
		}

	})
	if err != nil {
		log.Errorf("find order :%s err:%v", payload.OrderID, err)
		return
	}
	if amountUser == nil {
		log.Errorf("amountUser is nil order :%s", payload.OrderID)
		return
	}

	xMsg, err := tgmsg.NewXTelegramMessage(channelCfg.TelegramCfg.MainRobotToken,
		tgmsg.WithText("delete msg"),
		tgmsg.WithDebug(true),
		tgmsg.WithMsgType(tgtypes.RobotMsgTypePhoto),
		tgmsg.WithParseMode(tgtypes.ModeNone))

	if err != nil {
		return
	}
	if xMsg == nil {
		return
	}
	if _, err = xMsg.DeleteMessage(payload.ChatID, amountUser.Extend.MessageID); err != nil {
		log.Warnf("deleteMessage:%v", err)
	} else {
		paymentAmountUserDao.Instance().ClearAmountByOrderID(payload.OrderID)
		waitforinput.SetWaitForinputKey(&waitforinput.WaitforinputInfo{
			UserID:        payload.UserID, //用户TG
			InPutMsg:      payload.WaitforinputMsg.InPutMsg,
			Button:        payload.Button,
			UserBottonKey: payload.Button.WaitForInputKey(),
			Extended:      nil,
		})
	}

}
