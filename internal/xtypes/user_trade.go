package xtypes

// TradeType 交易类型
type TradeType int

const (
	TradeTypeRecharge               TradeType = 1 // 充值
	TradeTypeWithdraw               TradeType = 2 // 提现
	TradeTypeGameBet                TradeType = 3 // 游戏押注
	TradeTypeGameWin                TradeType = 4 // 游戏赢钱
	TradeTypeGameReward             TradeType = 5 // 游戏奖励
	TradeTypeAffiliate              TradeType = 6 // 代理收益
	TradeTypeRechargeOtherAddresses TradeType = 7 // 为其他用户购买能量
	TradeTypeRechargeBiShu          TradeType = 8 // 购买笔数
)

// TradeStatus 交易状态
type TradeStatus int

const (
	TradeStatusSuccess TradeStatus = 1 // 交易成功
	TradeStatusFailed  TradeStatus = 2 // 交易失败
	TradeStatusTrading TradeStatus = 3 // 正在交易
)

// TradeChannel 交易渠道
type TradeChannel string

const (
	TradeChannelRecharge  TradeChannel = "recharge"
	TradeChannelWithdraw  TradeChannel = "withdraw"
	TradeChannelGame      TradeChannel = "game"
	TradeChannelAffiliate TradeChannel = "affiliate"
)

func (t TradeType) TradeChannel() TradeChannel {
	switch t {
	case TradeTypeRecharge:
		return TradeChannelRecharge
	case TradeTypeWithdraw:
		{
			return TradeChannelWithdraw
		}
	case TradeTypeGameBet, TradeTypeGameWin, TradeTypeGameReward:
		return TradeChannelGame
	case TradeTypeAffiliate:
		return TradeChannelAffiliate

	}

	return ""
}
