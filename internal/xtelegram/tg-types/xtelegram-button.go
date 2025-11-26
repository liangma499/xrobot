package tgtypes

import (
	"strings"
)

type XTelegramButton string

const (
	XTelegramButton_None                  XTelegramButton = ""
	XTelegramButton_Test                  XTelegramButton = "test"
	XTelegramButton_Start                 XTelegramButton = "/start"
	XTelegramButton_EFR                   XTelegramButton = "🔋能量闪租"
	XTelegramButton_TRXConvert            XTelegramButton = "✅TRX闪兑"
	XTelegramButton_EnergyAdvances        XTelegramButton = "🆘能量预支"
	XTelegramButton_NTP                   XTelegramButton = "🔥笔数套餐"
	XTelegramButton_TelegramMember        XTelegramButton = "👑飞机会员"
	XTelegramButton_PromoteMakeMoney      XTelegramButton = "💰推广赚钱"
	XTelegramButton_GoodAddress           XTelegramButton = "💎靓号地址"
	XTelegramButton_ListeningAddress      XTelegramButton = "🔔监听地址"
	XTelegramButton_PersonalCenter        XTelegramButton = "👤个人中心"
	XTelegramButton_CustomizeTheSameRobot XTelegramButton = "🏆定制同款机器人"

	XTelegramButton_AddressDetail XTelegramButton = "XB_AddressDetail" //查询地址详情

	XTelegramButton_NTP_20Bi   XTelegramButton = "20笔"
	XTelegramButton_NTP_30Bi   XTelegramButton = "30笔"
	XTelegramButton_NTP_50Bi   XTelegramButton = "50笔"
	XTelegramButton_NTP_100Bi  XTelegramButton = "100笔"
	XTelegramButton_NTP_200Bi  XTelegramButton = "200笔"
	XTelegramButton_NTP_300Bi  XTelegramButton = "300笔"
	XTelegramButton_NTP_500Bi  XTelegramButton = "500笔"
	XTelegramButton_NTP_1000Bi XTelegramButton = "1000笔"
	XTelegramButton_NTP_2000Bi XTelegramButton = "2000笔"

	XTelegramButton_EFR_ExpirationDate XTelegramButton = "👇有效期1小时👇"
	XTelegramButton_EFR_1Bi            XTelegramButton = "1笔"
	XTelegramButton_EFR_2Bi            XTelegramButton = "2笔"
	XTelegramButton_EFR_3Bi            XTelegramButton = "3笔"
	XTelegramButton_EFR_5Bi            XTelegramButton = "5笔"
	XTelegramButton_EFR_10Bi           XTelegramButton = "10笔"

	XTelegramButton_EFR_RechargeOtherAddresses               XTelegramButton = "为其他地址充值（支持批量归集+激活）"
	XTelegramButton_EFR_RechargeOtherAddressesBalancePayment XTelegramButton = "余额支付"
	XTelegramButton_EFR_RechargeOtherAddressesCancelOrder    XTelegramButton = "取消订单"

	XTelegramButton_EFR_Recharge50TRX   XTelegramButton = "50 TRX"
	XTelegramButton_EFR_Recharge100TRX  XTelegramButton = "100 TRX"
	XTelegramButton_EFR_Recharge300TRX  XTelegramButton = "300 TRX"
	XTelegramButton_EFR_Recharge500TRX  XTelegramButton = "500 TRX"
	XTelegramButton_EFR_Recharge1000TRX XTelegramButton = "1000 TRX"
	XTelegramButton_EFR_Recharge2000TRX XTelegramButton = "2000 TRX"
	XTelegramButton_EFR_Recharge3000TRX XTelegramButton = "3000 TRX"
	XTelegramButton_EFR_Recharge5000TRX XTelegramButton = "5000 TRX"

	XTelegramButton_BuySameRobot XTelegramButton = "💎购买机器人"
)

// 按钮响应
func StringToXTelegramButton(cmd string) XTelegramButton {
	//cmd = strings.ToLower(cmd)
	if strings.Contains(cmd, "start") || cmd == "XB_Start" {
		return XTelegramButton_Start
	} else if strings.Contains(cmd, "能量闪租") || cmd == "XB_EFR" {
		return XTelegramButton_EFR
	} else if strings.Contains(cmd, "TRX闪兑") || cmd == "XB_TRXConvert" {
		return XTelegramButton_TRXConvert
	} else if strings.Contains(cmd, "能量预支") || cmd == "XB_EnergyAdvances" {
		return XTelegramButton_EnergyAdvances
	} else if strings.Contains(cmd, "笔数套餐") || cmd == "XB_NTP" {
		return XTelegramButton_NTP
	} else if strings.Contains(cmd, "飞机会员") || cmd == "XB_TelegramMember" {
		return XTelegramButton_TelegramMember
	} else if strings.Contains(cmd, "推广赚钱") || cmd == "XB_PromoteMakeMoney" {
		return XTelegramButton_PromoteMakeMoney
	} else if strings.Contains(cmd, "靓号地址") || cmd == "XB_GoodAddress" {
		return XTelegramButton_GoodAddress
	} else if strings.Contains(cmd, "监听地址") || cmd == "XB_ListeningAddress" {
		return XTelegramButton_ListeningAddress
	} else if strings.Contains(cmd, "个人中心") || cmd == "XB_PersonalCenter" {
		return XTelegramButton_PersonalCenter
	} else if strings.Contains(cmd, "定制同款机器人") || cmd == "XB_CustomizeTheSameRobot" {
		return XTelegramButton_CustomizeTheSameRobot
	}
	//以前是固定的
	switch cmd {

	case "XB_NTP_20Bi":
		{
			return XTelegramButton_NTP_20Bi
		}
	case "XB_NTP_30Bi":
		{
			return XTelegramButton_NTP_30Bi
		}
	case "XB_NTP_50Bi":
		{
			return XTelegramButton_NTP_50Bi
		}
	case "XB_NTP_100Bi":
		{
			return XTelegramButton_NTP_100Bi
		}
	case "XB_NTP_200Bi":
		{
			return XTelegramButton_NTP_200Bi
		}
	case "XB_NTP_300Bi":
		{
			return XTelegramButton_NTP_300Bi
		}
	case "XB_NTP_500Bi":
		{
			return XTelegramButton_NTP_500Bi
		}
	case "XB_NTP_1000Bi":
		{
			return XTelegramButton_NTP_1000Bi
		}
	case "XB_NTP_2000Bi":
		{
			return XTelegramButton_NTP_2000Bi
		}

	case "XB_EFR_1Bi":
		{
			return XTelegramButton_EFR_1Bi
		}
	case "XB_EFR_2Bi":
		{
			return XTelegramButton_EFR_2Bi
		}
	case "XB_EFR_3Bi":
		{
			return XTelegramButton_EFR_3Bi
		}
	case "XB_EFR_5Bi":
		{
			return XTelegramButton_EFR_5Bi
		}
	case "XB_EFR_10Bi":
		{
			return XTelegramButton_EFR_10Bi
		}
	case "XB_EFR_ROA":
		{
			return XTelegramButton_EFR_RechargeOtherAddresses
		}
	case "XB_EFR_50TRX":
		{
			return XTelegramButton_EFR_Recharge50TRX
		}
	case "XB_EFR_100TRX":
		{
			return XTelegramButton_EFR_Recharge100TRX
		}
	case "XB_EFR_300TRX":
		{
			return XTelegramButton_EFR_Recharge300TRX
		}
	case "XB_EFR_500TRX":
		{
			return XTelegramButton_EFR_Recharge500TRX
		}
	case "XB_EFR_1000TRX":
		{
			return XTelegramButton_EFR_Recharge1000TRX
		}
	case "XB_EFR_2000TRX":
		{
			return XTelegramButton_EFR_Recharge2000TRX
		}
	case "XB_EFR_3000TRX":
		{
			return XTelegramButton_EFR_Recharge3000TRX
		}
	case "XB_EFR_5000TRX":
		{
			return XTelegramButton_EFR_Recharge5000TRX
		}
	case "XB_BSR":
		{
			return XTelegramButton_BuySameRobot
		}
	}
	if strings.HasPrefix(cmd, "XB_EFR_ROABP") {
		return XTelegramButton_EFR_RechargeOtherAddressesBalancePayment
	}
	if strings.HasPrefix(cmd, "XB_EFR_ROACO") {
		return XTelegramButton_EFR_RechargeOtherAddressesCancelOrder
	}

	if strings.HasPrefix(cmd, "test") {
		return XTelegramButton_Test
	}
	return XTelegramButton_None
}

func (xb XTelegramButton) String() string {
	return string(xb)
}

func (xb XTelegramButton) Value() int64 {
	switch xb {
	case XTelegramButton_NTP_20Bi:
		{
			return 20
		}
	case XTelegramButton_NTP_30Bi:
		{
			return 30
		}
	case XTelegramButton_NTP_50Bi:
		{
			return 50
		}
	case XTelegramButton_NTP_100Bi:
		{
			return 100
		}
	case XTelegramButton_NTP_200Bi:
		{
			return 200
		}
	case XTelegramButton_NTP_300Bi:
		{
			return 300
		}
	case XTelegramButton_NTP_500Bi:
		{
			return 500
		}
	case XTelegramButton_NTP_1000Bi:
		{
			return 1000
		}
	case XTelegramButton_NTP_2000Bi:
		{
			return 2000
		}
	case XTelegramButton_EFR_1Bi:
		{
			return 1
		}
	case XTelegramButton_EFR_2Bi:
		{
			return 2
		}
	case XTelegramButton_EFR_3Bi:
		{
			return 3
		}
	case XTelegramButton_EFR_5Bi:
		{
			return 5
		}
	case XTelegramButton_EFR_10Bi:
		{
			return 10
		}

	case XTelegramButton_EFR_Recharge50TRX:
		{
			return 50
		}
	case XTelegramButton_EFR_Recharge100TRX:
		{
			return 100
		}
	case XTelegramButton_EFR_Recharge300TRX:
		{
			return 300
		}
	case XTelegramButton_EFR_Recharge500TRX:
		{
			return 500
		}
	case XTelegramButton_EFR_Recharge1000TRX:
		{
			return 1000
		}
	case XTelegramButton_EFR_Recharge2000TRX:
		{
			return 2000
		}
	case XTelegramButton_EFR_Recharge3000TRX:
		{
			return 3000
		}
	case XTelegramButton_EFR_Recharge5000TRX:
		{
			return 5000
		}

	}
	return 0
}

func (xb XTelegramButton) CallbackData() string {
	switch xb {

	case XTelegramButton_None:
		{
			return "XB_None"
		}
	case XTelegramButton_Test:
		{
			return "XB_Test"
		}
	case XTelegramButton_Start:
		{
			return "XB_Start"
		}
	case XTelegramButton_EFR:
		{
			return "XB_EFR"
		}
	case XTelegramButton_TRXConvert:
		{
			return "XB_TRXConvert"
		}
	case XTelegramButton_EnergyAdvances:
		{
			return "XB_EnergyAdvances"
		}
	case XTelegramButton_NTP:
		{
			return "XB_NTP"
		}
	case XTelegramButton_TelegramMember:
		{
			return "XB_TelegramMember"
		}
	case XTelegramButton_PromoteMakeMoney:
		{
			return "XB_PromoteMakeMoney"
		}
	case XTelegramButton_GoodAddress:
		{
			return "XB_GoodAddress"
		}
	case XTelegramButton_ListeningAddress:
		{
			return "XB_ListeningAddress"
		}
	case XTelegramButton_PersonalCenter:
		{
			return "XB_PersonalCenter"
		}
	case XTelegramButton_CustomizeTheSameRobot:
		{
			return "XB_CustomizeTheSameRobot"
		}
	case XTelegramButton_NTP_20Bi:
		{
			return "XB_NTP_20Bi"
		}
	case XTelegramButton_NTP_30Bi:
		{
			return "XB_NTP_30Bi"
		}
	case XTelegramButton_NTP_50Bi:
		{
			return "XB_NTP_50Bi"
		}
	case XTelegramButton_NTP_100Bi:
		{
			return "XB_NTP_100Bi"
		}
	case XTelegramButton_NTP_200Bi:
		{
			return "XB_NTP_200Bi"
		}
	case XTelegramButton_NTP_300Bi:
		{
			return "XB_NTP_300Bi"
		}
	case XTelegramButton_NTP_500Bi:
		{
			return "XB_NTP_500Bi"
		}
	case XTelegramButton_NTP_1000Bi:
		{
			return "XB_NTP_1000Bi"
		}
	case XTelegramButton_NTP_2000Bi:
		{
			return "XB_NTP_2000Bi"
		}
	case XTelegramButton_EFR_ExpirationDate:
		{
			return "XB_EFR_ExpirationDate"
		}
	case XTelegramButton_EFR_1Bi:
		{
			return "XB_EFR_1Bi"
		}
	case XTelegramButton_EFR_2Bi:
		{
			return "XB_EFR_2Bi"
		}
	case XTelegramButton_EFR_3Bi:
		{
			return "XB_EFR_3Bi"
		}
	case XTelegramButton_EFR_5Bi:
		{
			return "XB_EFR_5Bi"
		}
	case XTelegramButton_EFR_10Bi:
		{
			return "XB_EFR_10Bi"
		}
	case XTelegramButton_EFR_RechargeOtherAddresses:
		{
			return "XB_EFR_ROA"
		}
	case XTelegramButton_EFR_RechargeOtherAddressesBalancePayment:
		{
			return "XB_EFR_ROABP"
		}
	case XTelegramButton_EFR_RechargeOtherAddressesCancelOrder:
		{
			return "XB_EFR_ROACO"
		}
	case XTelegramButton_EFR_Recharge50TRX:
		{
			return "XB_EFR_50TRX"
		}
	case XTelegramButton_EFR_Recharge100TRX:
		{
			return "XB_EFR_100TRX"
		}
	case XTelegramButton_EFR_Recharge300TRX:
		{
			return "XB_EFR_300TRX"
		}
	case XTelegramButton_EFR_Recharge500TRX:
		{
			return "XB_EFR_500TRX"
		}
	case XTelegramButton_EFR_Recharge1000TRX:
		{
			return "XB_EFR_1000TRX"
		}
	case XTelegramButton_EFR_Recharge2000TRX:
		{
			return "XB_EFR_2000TRX"
		}
	case XTelegramButton_EFR_Recharge3000TRX:
		{
			return "XB_EFR_3000TRX"
		}
	case XTelegramButton_EFR_Recharge5000TRX:
		{
			return "XB_EFR_5000TRX"
		}
	case XTelegramButton_BuySameRobot:
		{
			return "XB_BSR"
		}

	}
	return xb.String()
}

func (xb XTelegramButton) WaitForInputKey() string {
	switch xb {

	case XTelegramButton_EFR_RechargeOtherAddresses:
		{
			return xb.CallbackData()
		}
	}
	return ""
}
func (xb XTelegramButton) IsAddOrder() bool {
	switch xb {

	case XTelegramButton_EFR_RechargeOtherAddressesBalancePayment,
		XTelegramButton_EFR_RechargeOtherAddressesCancelOrder:
		{
			return true
		}
	}
	return false
}
func (xb XTelegramButton) GetOrderID(callBackStr string) string {
	prefix := xb.CallbackData()
	if strings.HasPrefix(callBackStr, prefix) {
		return strings.TrimPrefix(callBackStr, prefix)
	}
	return ""
}
func (xb XTelegramButton) IsSaveLastButton() bool {
	switch xb {

	case XTelegramButton_EFR_1Bi,
		XTelegramButton_EFR_2Bi,
		XTelegramButton_EFR_3Bi,
		XTelegramButton_EFR_5Bi,
		XTelegramButton_EFR_10Bi,
		XTelegramButton_NTP_20Bi,
		XTelegramButton_NTP_30Bi,
		XTelegramButton_NTP_50Bi,
		XTelegramButton_NTP_100Bi,
		XTelegramButton_NTP_200Bi,
		XTelegramButton_NTP_300Bi,
		XTelegramButton_NTP_500Bi,
		XTelegramButton_NTP_1000Bi,
		XTelegramButton_NTP_2000Bi:
		{
			return true
		}
	}
	return false
}
