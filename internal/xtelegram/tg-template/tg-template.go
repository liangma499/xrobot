package tgtemplate

import "strings"

const (
	CustomerKey                       = "customer"
	EnergySavingsKey                  = "energySavings"
	PriceKey                          = "price"
	PriceNoUKey                       = "priceNou"
	PriceNumKey                       = "priceNum"
	PriceNumIndexKey                  = "price%02d"
	PriceBiShuKey                     = "pricebishu%02d"
	PriceBiShuMaxKey                  = "pricebishuMax"
	Tron20AddressKey                  = "tron20AddressKey"
	ComboKindEnergyFlashRentalNumKey  = "comboKindEnergyFlashRentalNumKey"
	ComboKindEnergyFlashRentalNameKey = "comboKindEnergyFlashRentalNameKey"
	NotActivatedAddressCountKey       = "notActivatedAddressCount" //未激活个数
	ReceivingAddressCountKey          = "receivingAddressCount"    //接收地址个数
	EnergyFeeKey                      = "energyFee"                //能量费用
	ActivationfeeKey                  = "activationfee"            //激活费用
	PayAmountKey                      = "payAmount"                //支付金额
	Balance                           = "balance"                  //支付金额
	CustomizeBalance                  = "customizeBalance"         //定制金额
)

var (
	Start = func() string {
		var builder strings.Builder
		builder.WriteString("欢迎使用能量闪租、TRX兑换机器人\n")
		builder.WriteString("🌈使用能量可节省 ${energySavings} 转U手续费\n\r\n")
		builder.WriteString("💥转 U 兑TRX，转TRX兑能量\n")
		builder.WriteString("✅全自动到账，默认返原地址\n\r\n")
		builder.WriteString("🚫请勿使用交易所或中心化钱包转账\n")
		builder.WriteString("⚡️如有问题，请联系客服 ${customer}")
		return builder.String()
	}

	EnergyFlashRental = func() string {
		var builder strings.Builder
		builder.WriteString("*🔋能量闪租\n➖➖➖➖➖➖➖➖➖➖*\n")
		builder.WriteString("🌈使用能量可节省 *${energySavings}* 转U手续费\n\r\n")
		builder.WriteString("🔹1笔对方地址*【有U】* ${price} TRX  (${comboKindEnergyFlashRentalNumKey}${comboKindEnergyFlashRentalNameKey}有效)\n")
		builder.WriteString("🔹1笔对方地址*【无U】* ${priceNou} TRX  (${comboKindEnergyFlashRentalNumKey}${comboKindEnergyFlashRentalNameKey}有效)\n\r\n")
		builder.WriteString("🔋*小时套餐（${comboKindEnergyFlashRentalNumKey}${comboKindEnergyFlashRentalNameKey}有效）*\n")
		builder.WriteString("🔸转账 ${price01} TRX = 免费${pricebishu01}笔转账\n")
		builder.WriteString("🔸转账 ${price02} TRX = 免费${pricebishu02}笔转账\n")
		builder.WriteString("🔸转账 ${price03} TRX = 免费${pricebishu03}笔转账\n")
		builder.WriteString("🔸以此类推 ${price}×笔数，单次${pricebishuMax}笔封顶\n\r\n")
		builder.WriteString("📣转 TRX 到下方地址，能量自动到账\n")
		builder.WriteString("`${tron20AddressKey}`\n")
		builder.WriteString("(点击地址复制)\n\r\n")
		builder.WriteString("✅全自动到账，默认返回原地址\n")
		builder.WriteString("🚫请勿使用交易所或中心化钱包转账")
		return builder.String()
	}

	EnergyFlashRentalBiShu = func() string {
		var builder strings.Builder
		builder.WriteString("*⚠️↓↓请按金额支付，否则无法到账↓↓*\n")
		builder.WriteString("---------------------------------\n")
		builder.WriteString("🔸套餐模式：${comboKindEnergyFlashRentalNumKey}${comboKindEnergyFlashRentalNameKey}${priceNum}笔\n")
		builder.WriteString("🔸支付金额：${price} TRX\n")

		builder.WriteString("🔸收款地址：`${tron20AddressKey}`\n")
		builder.WriteString("（点击地址复制）\n---------------------------------\n\r\n")
		builder.WriteString("*✅全自动到账，能量即回原地址*\n")
		builder.WriteString("*🚫请勿使用交易所或中心化钱包转账*")
		return builder.String()
	}

	RechargeOtherAddresses = func() string {
		var builder strings.Builder
		builder.WriteString("请输入接收能量的地址（支持多个）：\n\r\n")
		builder.WriteString("▫️例如：\n▫️`Txxxxx...001`\n▫️")
		builder.WriteString("`Txxxxx...002`\n")
		builder.WriteString("▫️`Txxxxx...003`")
		return builder.String()
	}

	RechargeOtherAddressesRet = func() string {
		var builder strings.Builder
		builder.WriteString("⚠️↓↓请按金额支付，否则无法到账↓↓\n")
		builder.WriteString("---------------------------------\n")
		builder.WriteString("🔸套餐模式：${comboKindEnergyFlashRentalNumKey}${comboKindEnergyFlashRentalNameKey}${priceNum}笔\n")
		builder.WriteString("🔸接收地址：${receivingAddressCount}个\n")
		builder.WriteString("🔸未激活数：${notActivatedAddressCount}个\n")
		builder.WriteString("🔸能量费用：${energyFee} TRX\n")
		builder.WriteString("🔸激活费用：${activationfee} TRX\n")
		builder.WriteString("🔸支付金额：${payAmount} TRX\n")
		builder.WriteString("🔸收款地址：`${tron20AddressKey}`\n")
		builder.WriteString("（点击地址复制）\n---------------------------------\n\r\n")
		builder.WriteString("⚠️⚠️请务必核对金额尾数，金额带小数\n")
		builder.WriteString("🚫请勿使用交易所或中心化钱包转账")
		return builder.String()
	}

	RechargeRet = func() string {
		var builder strings.Builder
		builder.WriteString("⚠️余额不足，请先充值\n\r\n")
		builder.WriteString("💰账户余额：${balance} TRX\n")
		builder.WriteString("👇请在下方选择要充值的金额")
		return builder.String()
	}

	CustomizeTheSameRobot = func() string {
		var builder strings.Builder
		builder.WriteString("*🔋推广赚钱，自用省钱*\n")
		builder.WriteString("*➖➖➖➖➖➖➖➖➖➖*\n")
		builder.WriteString("*❣️诚招代理，只需花 ${customizeBalance}U定制同款机器人可成为代理。*\n\r\n")
		builder.WriteString("*👑代理权益*\n")
		builder.WriteString("🔺拥有专属机器人\n")
		builder.WriteString("🔺最低的供货成本\n")
		builder.WriteString("🔺最低的自用成本\n")
		builder.WriteString("🔺推广机器人收益\n")
		builder.WriteString("🔺下级机器人收益\n\r\n")
		builder.WriteString("*💎专属服务*\n")
		builder.WriteString("✅提供 7x24 售后服务\n")
		builder.WriteString("✅提供专属靓号收款地址\n")
		builder.WriteString("✅提供全套落地获客方案\n")
		builder.WriteString("✅提供全方位订单推送扶持\n")
		builder.WriteString("✅提供永久多功能机器人技术服务\n\r\n")
		builder.WriteString("⚡️如需帮助，请联系客服 ${customer}\n")
		builder.WriteString("💬机器人代理可绑定个人收款地址，进行外网推广。")
		return builder.String()
	}

	CustomizeBuySameRobot = func() string {
		var builder strings.Builder
		builder.WriteString("*⚠️↓↓请按金额支付，否则无法到账↓↓*\n")
		builder.WriteString("---------------------------------\n")
		builder.WriteString("🔸支付商品：机器人\n")
		builder.WriteString("🔸支付金额：`${payAmount}` USDT\n")
		builder.WriteString("🔸收款地址：`${tron20AddressKey}`\n")
		builder.WriteString("（点击地址复制）\n")
		builder.WriteString("‼️请务必核对金额尾数，金额带小数\n")
		builder.WriteString("🚫请勿使用交易所或中心化钱包转账")

		return builder.String()
	}

	RechargeText = func() string {
		var builder strings.Builder
		builder.WriteString("*⚠️↓↓请按金额支付，否则无法到账↓↓*\n")
		builder.WriteString("---------------------------------\n")
		builder.WriteString("🔸支付金额：50.07 TRX 或 12.07 USDT\n")
		builder.WriteString("🔸收款地址：`${tron20AddressKey}`\n")
		builder.WriteString("（点击地址复制）\n")
		builder.WriteString("‼️请务必核对金额尾数，金额带小数\n")
		builder.WriteString("🚫请勿使用交易所或中心化钱包转账")

		return builder.String()
	}
)
