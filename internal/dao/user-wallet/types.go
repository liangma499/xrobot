package userwallet

import (
	"tron_robot/internal/xtypes"

	"github.com/shopspring/decimal"
)

type InitBalanceArgs struct {
	UID      int64                      // 用户ID
	Currency xtypes.Currency            // 货币
	Cash     decimal.Decimal            // 现金
	Used     decimal.Decimal            // 奖金
	Def      xtypes.WalletDefaultStatus // 是否默认账户
}
type BalanceInfo struct {
	Currency   xtypes.Currency
	Cash       float64
	Used       float64
	IsDefault  bool
	AmountKind xtypes.WalletAmountKind
}

type ChangeInfo struct {
	BeforeCash float64
	AfterCash  float64
	ChangeCash float64
	Def        int
	Used       float64
}
