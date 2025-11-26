package business

import (
	"context"
	userBaseDao "tron_robot/internal/dao/user-base"
	"tron_robot/internal/event/message"
	"tron_robot/internal/service/business/internal"
	tgtypes "tron_robot/internal/xtelegram/tg-types"
	"xbase/task"
	"xbase/utils/xconv"
)

func (s *Server) doSubscribeMessageBusiness(uuid string, payload *message.MessageBusiness) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if payload == nil {
		return
	}

	userBase, err := userBaseDao.Instance().DoGetUserBaseByCode(context.Background(), xconv.String(payload.UserID))
	if err != nil {
		return
	}
	if userBase == nil {
		return
	}
	switch payload.Button {

	case tgtypes.XTelegramButton_EFR:
		{
			// "🔋能量闪租"
			task.AddTask(func() {
				internal.EnergyFlashRental(userBase.Clone(), payload.Clone())
			})
			return
		}
	case tgtypes.XTelegramButton_TRXConvert:
		{
			// "✅TRX闪兑"
			task.AddTask(func() {
				internal.TRXConvert(userBase.Clone(), payload.Clone())
			})
			return
		}
	case tgtypes.XTelegramButton_EnergyAdvances:
		{
			// "🆘能量预支"
			task.AddTask(func() {
				internal.EnergyAdvances(userBase.Clone(), payload.Clone())
			})
			return
		}
	case tgtypes.XTelegramButton_NTP:
		{
			// "🔥笔数套餐"
			task.AddTask(func() {
				internal.NumberOfTransactionPackages(userBase.Clone(), payload.Clone())
			})
			return
		}
	case tgtypes.XTelegramButton_TelegramMember:
		{
			// "👑飞机会员"
			task.AddTask(func() {
				internal.TelegramMember(userBase.Clone(), payload.Clone())
			})
			return
		}
	case tgtypes.XTelegramButton_PromoteMakeMoney:
		{
			// "💰推广赚钱"
			task.AddTask(func() {
				internal.PromoteMakeMoney(userBase.Clone(), payload.Clone())
			})
			return
		}
	case tgtypes.XTelegramButton_GoodAddress:
		{
			// "💎靓号地址"
			task.AddTask(func() {
				internal.GoodAddress(userBase.Clone(), payload.Clone())
			})
			return
		}
	case tgtypes.XTelegramButton_ListeningAddress:
		{
			// "🔔监听地址"
			task.AddTask(func() {
				internal.ListeningAddress(userBase.Clone(), payload.Clone())
			})
			return
		}
	case tgtypes.XTelegramButton_PersonalCenter:
		{
			// "👤个人中心"
			task.AddTask(func() {
				internal.PersonalCenter(userBase.Clone(), payload.Clone())
			})
			return
		}
	case tgtypes.XTelegramButton_CustomizeTheSameRobot:
		{
			// "🏆定制同款机器人"
			task.AddTask(func() {
				internal.CustomizeTheSameRobot(userBase.Clone(), payload.Clone())
			})
			return

		}
	case tgtypes.XTelegramButton_BuySameRobot:
		{
			//购买机器人
			task.AddTask(func() {
				internal.CustomizeBuySameRobot(userBase.Clone(), payload.Clone())
			})
			return

		}
	case tgtypes.XTelegramButton_NTP_20Bi,
		tgtypes.XTelegramButton_NTP_30Bi,
		tgtypes.XTelegramButton_NTP_50Bi,
		tgtypes.XTelegramButton_NTP_100Bi,
		tgtypes.XTelegramButton_NTP_200Bi,
		tgtypes.XTelegramButton_NTP_300Bi,
		tgtypes.XTelegramButton_NTP_500Bi,
		tgtypes.XTelegramButton_NTP_1000Bi,
		tgtypes.XTelegramButton_NTP_2000Bi:
		{
			// 笔数套餐消息
			task.AddTask(func() {
				internal.NumberOfTransaction(userBase.Clone(), payload.Clone())
			})
			return
		}
	case tgtypes.XTelegramButton_EFR_1Bi,
		tgtypes.XTelegramButton_EFR_2Bi,
		tgtypes.XTelegramButton_EFR_3Bi,
		tgtypes.XTelegramButton_EFR_5Bi,
		tgtypes.XTelegramButton_EFR_10Bi:
		{
			// 能量闪租笔数
			task.AddTask(func() {
				internal.EnergyFlashRentalBiShu(userBase.Clone(), payload.Clone())
			})
			return
		}
	case tgtypes.XTelegramButton_EFR_RechargeOtherAddresses:
		{
			//为其他地址充值
			task.AddTask(func() {
				internal.RechargeOtherAddresses(userBase.Clone(), payload.Clone())
			})
			return

		}
	case tgtypes.XTelegramButton_AddressDetail:
		{
			//地址数据详情
			task.AddTask(func() {
				internal.AddressDetail(userBase.Clone(), payload.Clone())
			})
			return

		}
	case tgtypes.XTelegramButton_EFR_RechargeOtherAddressesBalancePayment:
		{
			//余额支付
			task.AddTask(func() {
				internal.RechargeOtherAddressesBalancePayment(userBase.Clone(), payload.Clone(), s.proxy)
			})
			return

		}
	case tgtypes.XTelegramButton_EFR_RechargeOtherAddressesCancelOrder:
		{
			//取消订单
			task.AddTask(func() {
				internal.RechargeOtherAddressesCancelOrder(userBase.Clone(), payload.Clone())
			})
			return

		}
	case tgtypes.XTelegramButton_EFR_Recharge50TRX,
		tgtypes.XTelegramButton_EFR_Recharge100TRX,
		tgtypes.XTelegramButton_EFR_Recharge300TRX,
		tgtypes.XTelegramButton_EFR_Recharge500TRX,
		tgtypes.XTelegramButton_EFR_Recharge1000TRX,
		tgtypes.XTelegramButton_EFR_Recharge2000TRX,
		tgtypes.XTelegramButton_EFR_Recharge3000TRX,
		tgtypes.XTelegramButton_EFR_Recharge5000TRX:
		{
			//取消订单
			task.AddTask(func() {
				internal.Recharge(userBase.Clone(), payload.Clone())
			})
			return

		}

	}
}
