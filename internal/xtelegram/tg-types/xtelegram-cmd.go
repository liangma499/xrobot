package tgtypes

import "strings"

type XTelegramCmd string

const (
	XTelegramCmd_None                               XTelegramCmd = ""
	XTelegramCmd_Test                               XTelegramCmd = "/test"
	XTelegramCmd_Start                              XTelegramCmd = "/start"
	XTelegramCmd_Button_EnergyFlashRental           XTelegramCmd = "🔋能量闪租"
	XTelegramCmd_Button_TRXConvert                  XTelegramCmd = "✅TRX闪兑"
	XTelegramCmd_Button_EnergyAdvances              XTelegramCmd = "🆘能量预支"
	XTelegramCmd_Button_NumberOfTransactionPackages XTelegramCmd = "🔥笔数套餐"
	XTelegramCmd_Button_TelegramMember              XTelegramCmd = "👑飞机会员"
	XTelegramCmd_Button_PromoteMakeMoney            XTelegramCmd = "💰推广赚钱"
	XTelegramCmd_Button_GoodAddress                 XTelegramCmd = "💎靓号地址"
	XTelegramCmd_Button_ListeningAddress            XTelegramCmd = "🔔监听地址"
	XTelegramCmd_Button_PersonalCenter              XTelegramCmd = "👤个人中心"
	XTelegramCmd_Button_EnergyFlashRental_BiSu      XTelegramCmd = "能量闪租笔数"
	XTelegramCmd_Button_RechargeOtherAddresses      XTelegramCmd = "rechargeOtherAddresses"
	XTelegramCmd_Button_Recharge                    XTelegramCmd = "recharge"
	XTelegramCmd_CustomizeTheSameRobot              XTelegramCmd = "customizeTheSameRobot"
	XTelegramCmd_BuySameRobot                       XTelegramCmd = "buySameRobot "
)

func StringToXTelegramCmd(cmd string) XTelegramCmd {
	cmd = strings.ToLower(cmd)
	switch cmd {
	case "/start", "start":
		{
			return XTelegramCmd_Start
		}

	case "/🔋能量闪租", "🔋能量闪租", "能量闪租":
		{
			return XTelegramCmd_Button_EnergyFlashRental
		}
	case "/✅TRX闪兑", "✅TRX闪兑", "TRX闪兑":
		{
			return XTelegramCmd_Button_TRXConvert
		}
	case "/🆘能量预支", "🆘能量预支", "能量预支":
		{
			return XTelegramCmd_Button_EnergyAdvances
		}
	case "/🔥笔数套餐", "🔥笔数套餐", "笔数套餐":
		{
			return XTelegramCmd_Button_NumberOfTransactionPackages
		}
	case "/👑飞机会员", "👑飞机会员", "飞机会员":
		{
			return XTelegramCmd_Button_TelegramMember
		}
	case "/💰推广赚钱", "💰推广赚钱", "推广赚钱":
		{
			return XTelegramCmd_Button_PromoteMakeMoney
		}
	case "/💎靓号地址", "💎靓号地址", "靓号地址":
		{
			return XTelegramCmd_Button_GoodAddress
		}
	case "/🔔监听地址", "🔔监听地址", "监听地址":
		{
			return XTelegramCmd_Button_ListeningAddress
		}
	case "/👤个人中心", "👤个人中心", "个人中心":
		{
			return XTelegramCmd_Button_PersonalCenter
		}
	case "/能量闪租笔数", "能量闪租笔数":
		{
			return XTelegramCmd_Button_EnergyFlashRental_BiSu
		}
	case "/rechargeOtherAddresses", "rechargeOtherAddresses":
		{
			return XTelegramCmd_Button_RechargeOtherAddresses
		}
	case "/recharge", "recharge":
		{
			return XTelegramCmd_Button_Recharge
		}
	case "/customizeTheSameRobot", "customizeTheSameRobot":
		{
			return XTelegramCmd_Button_Recharge
		}
	case "/buySameRobot", "buySameRobot":
		{
			return XTelegramCmd_BuySameRobot
		}
	}

	if strings.Contains(cmd, "test") {
		return XTelegramCmd_Test
	}
	return XTelegramCmd_None
}
func (xc XTelegramCmd) String() string {
	return string(xc)
}
func (xc XTelegramCmd) Description() string {
	switch xc {
	case XTelegramCmd_Test:
		{
			return "测试"
		}
	case XTelegramCmd_Start:
		{
			return "菜单"
		}
	case XTelegramCmd_Button_EnergyFlashRental:
		{
			return "能量闪租"
		}
	case XTelegramCmd_Button_TRXConvert:
		{
			return "TRX闪兑"
		}
	case XTelegramCmd_Button_EnergyAdvances:
		{
			return "能量预支兑"
		}
	case XTelegramCmd_Button_NumberOfTransactionPackages:
		{
			return "笔数套餐"
		}
	case XTelegramCmd_Button_TelegramMember:
		{
			return "飞机会员"
		}
	case XTelegramCmd_Button_PromoteMakeMoney:
		{
			return "推广赚钱"
		}
	case XTelegramCmd_Button_GoodAddress:
		{
			return "靓号地址"
		}
	case XTelegramCmd_Button_ListeningAddress:
		{
			return "监听地址"
		}
	case XTelegramCmd_Button_PersonalCenter:
		{
			return "个人中心"
		}
	case XTelegramCmd_Button_EnergyFlashRental_BiSu:
		{
			return "能量闪租笔数"
		}
	case XTelegramCmd_Button_RechargeOtherAddresses:
		{
			return "能量闪租笔数"
		}
	case XTelegramCmd_Button_Recharge:
		{
			return "充值"
		}
	}
	return ""
}
