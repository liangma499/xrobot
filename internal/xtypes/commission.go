package xtypes

type CommissionType string

const (
	CommissionType_Commission = "commission" //返佣
	CommissionType_WithDraw   = "withdraw"   //提现
)

type WithdrawType string

const (
	WithdrawType_Commission = "commission" //佣金提现
	WithdrawType_Card       = "card"       //银行卡提现
)
