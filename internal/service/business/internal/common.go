package internal

import (
	"context"
	"fmt"
	redisCryptoCurrencies "tron_robot/internal/component/redis/redis-crypto-currencies"
	paymentAmountUserDao "tron_robot/internal/dao/payment-amount-user"
	"tron_robot/internal/event/message"
	"tron_robot/internal/model"
	optiontelegramcmd "tron_robot/internal/option/option-telegram-cmd"
	"tron_robot/internal/utils/xstr"
	tgmsg "tron_robot/internal/xtelegram/tg-msg"
	waitforinput "tron_robot/internal/xtelegram/wait-for-input"
	"tron_robot/internal/xtypes"
	"xbase/log"
	"xbase/utils/xconv"
	"xbase/utils/xtime"

	"github.com/shopspring/decimal"
	"gorm.io/gorm/clause"
)

func doSetLastMessgeID(uid int64, cmd string, messID int64) error {
	key := fmt.Sprintf(xtypes.LastMessgeIDKey, uid, cmd)

	_, err := redisCryptoCurrencies.Instance().Set(context.Background(), key, messID, xtypes.LastMessgeIDExpiration).Result()
	return err
}
func doGetLastMessgeID(uid int64, cmd string) int64 {
	key := fmt.Sprintf(xtypes.LastMessgeIDKey, uid, cmd)

	rst, _ := redisCryptoCurrencies.Instance().Get(context.Background(), key).Int64()
	return rst
}
func doSendCreateOrder(channelCfg *model.OptionChannel,
	userBase *model.UserBase,
	payload *message.MessageBusiness,
	cmdMsg *optiontelegramcmd.OptionTelegramCmdCfg,
	usage xtypes.Usage,
	payAmount decimal.Decimal,
	expandMap map[string]string,
	expirationTime int64,
	amountUserInfo *model.AmountUserInfo,
	currency xtypes.Currency) {

	payAmountStr := payAmount.String()
	xMsg, err := tgmsg.NewXTelegramMessage(channelCfg.TelegramCfg.MainRobotToken,
		tgmsg.WithText(cmdMsg.Text),
		tgmsg.WithDebug(true),
		tgmsg.WithCmd(cmdMsg.Cmd),
		tgmsg.WithMsgType(cmdMsg.Type),
		tgmsg.WithParseMode(cmdMsg.ParseMode),
		tgmsg.WithKeyboard(cmdMsg.Keyboard))
	if err != nil {
		paymentAmountUserDao.Instance().ClearAmount(currency, payAmountStr, "")
		return
	}

	if xMsg == nil {
		paymentAmountUserDao.Instance().ClearAmount(currency, payAmountStr, "")
		return
	}
	orderID := xstr.SerialNO()
	if err := xMsg.SetExtraCallBackData(orderID); err != nil {
		log.Warnf("%v", err)
		paymentAmountUserDao.Instance().ClearAmount(currency, payAmountStr, "")
		return
	}

	now := xtime.Now()
	//userbaseInfo, err := userbase.Instance().DoGetUserBaseByCode(context.Background(), payload.UserID)
	oldOrderInfo, _ := paymentAmountUserDao.Instance().FindOne(context.Background(), func(cols *paymentAmountUserDao.Columns) any {
		return clause.And(clause.Eq{
			Column: cols.Usage,
			Value:  usage,
		}, clause.Eq{
			Column: cols.UID,
			Value:  userBase.UID,
		}, clause.Eq{
			Column: cols.Currency,
			Value:  userBase.UID,
		})
	})
	lastMsgID := 0
	if oldOrderInfo != nil {
		if oldOrderInfo.Extend != nil {
			lastMsgID = oldOrderInfo.Extend.MessageID
		}
		paymentAmountUserDao.Instance().ClearAmountByOrderID(oldOrderInfo.OrderID)
	}

	if lastMsgID == 0 {
		if rstMsg, err := xMsg.SendMessage(payload.ChatID, expandMap); err != nil {
			log.Warnf("sendMessage:%v", err)
			paymentAmountUserDao.Instance().ClearAmount(currency, payAmountStr, "")
		} else {
			if amountUserInfo == nil {
				amountUserInfo = &model.AmountUserInfo{
					MessageID: rstMsg.MessageID,
				}
			} else {
				amountUserInfo.MessageID = rstMsg.MessageID
			}
			userAmount := &model.PaymentAmountUser{
				OrderID:        orderID,
				Usage:          usage,
				UID:            userBase.UID,
				TelegramUid:    payload.UserID,       // 账号
				ChannelName:    userBase.ChannelName, // 渠道名
				ChannelCode:    userBase.ChannelCode,
				Amount:         payAmountStr, // 编号
				Currency:       currency,
				Extend:         amountUserInfo,
				ExpirationTime: expirationTime,
				CreateAt:       now,
				UpdateAt:       now,
			}

			_, err = paymentAmountUserDao.Instance().Insert(context.Background(), userAmount)
			if err != nil {
				log.Warnf("sendMessage:%v", err, xconv.Json(userAmount))
				paymentAmountUserDao.Instance().ClearAmount(currency, payAmountStr, "")
				if _, err = xMsg.DeleteMessage(payload.ChatID, rstMsg.MessageID); err != nil {
					log.Warnf("deleteMessage:%v", err)
				}
				return
			}
			waitforinput.SetWaitForinputKey(&waitforinput.WaitforinputInfo{
				UserID:        payload.UserID, //用户TG
				InPutMsg:      payload.WaitforinputMsg.InPutMsg,
				Button:        payload.Button,
				UserBottonKey: payload.Button.WaitForInputKey(),
				Extended:      nil,
			})

		}
	} else {
		if _, err := xMsg.EditMessageText(payload.ChatID, int64(lastMsgID), expandMap); err != nil {
			log.Warnf("sendMessage:%v", err)
			paymentAmountUserDao.Instance().ClearAmount(currency, payAmountStr, "")
			return
		} else {
			if amountUserInfo == nil {
				amountUserInfo = &model.AmountUserInfo{
					MessageID: lastMsgID,
				}
			} else {
				amountUserInfo.MessageID = lastMsgID
			}
			userAmount := &model.PaymentAmountUser{
				OrderID:        orderID,
				Usage:          usage,
				UID:            userBase.UID,
				TelegramUid:    payload.UserID,       // 账号
				ChannelName:    userBase.ChannelName, // 渠道名
				ChannelCode:    userBase.ChannelCode,
				Amount:         payAmountStr, // 编号
				Currency:       currency,
				Extend:         amountUserInfo,
				ExpirationTime: expirationTime,
				CreateAt:       now,
				UpdateAt:       now,
			}
			_, err = paymentAmountUserDao.Instance().Insert(context.Background(), userAmount)
			if err != nil {
				if _, err = xMsg.DeleteMessage(payload.ChatID, lastMsgID); err != nil {
					log.Warnf("deleteMessage:%v", err)
				}
				log.Warnf("sendMessage:%v", err, xconv.Json(userAmount))
				paymentAmountUserDao.Instance().ClearAmount(currency, payAmountStr, "")
				return
			}
		}

	}

}
