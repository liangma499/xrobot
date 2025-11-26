package internal

import (
	"context"
	"fmt"
	optionChannelDao "tron_robot/internal/dao/option-channel"
	paymentAmountUserDao "tron_robot/internal/dao/payment-amount-user"
	platformpendingorderdao "tron_robot/internal/dao/platform-pending-order"

	"tron_robot/internal/event/message"
	"tron_robot/internal/model"
	optiontelegramcmd "tron_robot/internal/option/option-telegram-cmd"
	walletsvc "tron_robot/internal/service/wallet"
	walletpb "tron_robot/internal/service/wallet/pb"
	tgmsg "tron_robot/internal/xtelegram/tg-msg"
	tgtemplate "tron_robot/internal/xtelegram/tg-template"
	tgtypes "tron_robot/internal/xtelegram/tg-types"
	waitforinput "tron_robot/internal/xtelegram/wait-for-input"
	"tron_robot/internal/xtypes"
	"xbase/cluster/mesh"
	"xbase/log"
	"xbase/utils/xconv"
	"xbase/utils/xtime"

	"github.com/shopspring/decimal"
	"gorm.io/gorm/clause"
)

// "余额支付"
func RechargeOtherAddressesBalancePayment(userBase *model.UserBase, payload *message.MessageBusiness, proxy *mesh.Proxy) {
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
		return clause.And(
			clause.Eq{
				Column: cols.OrderID,
				Value:  payload.OrderID,
			},
			clause.Eq{
				Column: cols.Currency,
				Value:  xtypes.TRX.String(),
			})
	})
	if err != nil {
		log.Errorf("find order :%s err:%v", payload.OrderID, err)
		return
	}
	if amountUser == nil {
		log.Errorf("amountUser is nil order :%s", payload.OrderID)
		return
	}

	client, err := walletsvc.NewClient(proxy.NewMeshClient)
	if err != nil {
		log.Errorf("new client orderID:%s UID :%d err:%v", payload.OrderID, userBase.UID, err)
		return
	}
	decrBalanceRst, err := client.DecrBalance(ctx, &walletpb.DecrBalanceArgs{
		UID:        amountUser.UID,
		Currency:   amountUser.Currency.String(),
		Cash:       xconv.Float64(amountUser.Amount),
		Type:       int32(xtypes.TradeTypeRechargeOtherAddresses),
		AmountKind: xtypes.WalletAmountKindCash.Int32(),
	})
	if err != nil {
		fetchBalanceRst, err := client.FetchBalance(ctx, &walletpb.FetchBalanceArgs{
			UID:      amountUser.UID,
			Currency: amountUser.Currency.String(),
		})
		if err != nil {
			log.Errorf("fetchBalance orderID:%s UID :%d err:%v", payload.OrderID, userBase.UID, err)
			return
		}
		fetchBalance, ok := fetchBalanceRst.Balance[xtypes.WalletAmountKindCash.Int32()]
		if !ok {
			log.Errorf("fetchBalance orderID:%s UID :%d err:%v", payload.OrderID, userBase.UID, err)
			return
		}
		//处理余额不足
		cmdMsg := optiontelegramcmd.GetChanCodeCmd(payload.ChannelCode, tgtypes.XTelegramCmd_Button_Recharge)

		if cmdMsg == nil {
			log.Errorf("cmdMsg is nil")
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
		expandMap := map[string]string{
			tgtemplate.Balance: fmt.Sprintf("%f", fetchBalance.Cash),
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
		return
	} else {
		if amountUser.Extend == nil {
			return
		}
		if amountUser.Extend.AddressInfo == nil {
			return
		}
		addressInfoLen := len(amountUser.Extend.AddressInfo)

		if addressInfoLen == 0 {
			return
		}
		platformPendingOrder := make([]*model.PlatformPendingOrder, 0)
		now := xtime.Now()
		for i := 0; i < addressInfoLen; i++ {
			item := amountUser.Extend.AddressInfo[i]
			if item == nil {
				continue
			}
			stauts := xtypes.OrderStatus_Activated
			if !item.Activated {
				stauts = xtypes.OrderStatus_ToBeActivated
			}

			platformPendingOrder = append(platformPendingOrder, &model.PlatformPendingOrder{
				UID:          userBase.UID,         // 主键
				Code:         userBase.Code,        // 编号
				ChannelCode:  userBase.ChannelCode, // 渠道编码
				Address:      item.Address,
				ToCurrency:   "", // 币种
				ToAmount:     decimal.Zero,
				FromCurrency: "", // 币种
				FromAmount:   decimal.Zero,
				Energy:       0,
				Stauts:       stauts,
				CreateAt:     now,
				UpdateAt:     now,
			})
		}
		if len(platformPendingOrder) > 0 {
			platformpendingorderdao.Instance().Insert(ctx, platformPendingOrder...)
		}

		xMsg, err := tgmsg.NewXTelegramMessage(channelCfg.TelegramCfg.MainRobotToken,
			tgmsg.WithText(fmt.Sprintf("余额：%f, \n \r \n 消耗 ：%s TRX 为他人充值成功，处理中....",
				decrBalanceRst.Balance.Cash,
				amountUser.Amount)),
			tgmsg.WithDebug(true),
			tgmsg.WithMsgType(tgtypes.RobotMsgTypePhoto),
			tgmsg.WithParseMode(tgtypes.ModeNone))
		if err != nil {
			return
		}
		if xMsg == nil {
			return
		}
		xMsg.DeleteMessage(payload.ChatID, amountUser.Extend.MessageID)
		paymentAmountUserDao.Instance().ClearAmountByOrderID(payload.OrderID)
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

}
