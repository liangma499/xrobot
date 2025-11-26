package wallet

import (
	"tron_robot/internal/xtypes"

	"github.com/shopspring/decimal"
)

type BalanceChangePayload struct {
	UID          int64              `json:"uid"`         // 用户ID
	TradeType    xtypes.TradeType   `json:"tradeType"`   // 交易类型
	Currency     xtypes.Currency    `json:"currency"`    // 币种ID
	TradeStatus  xtypes.TradeStatus `json:"tradeStatus"` // 交易状态
	Change       *ChangeInfo        `json:"change"`      // 变动金额
	RelatedID    string             `json:"relatedID"`   // 关联ID
	BetAmount    decimal.Decimal    `json:"bet_amount"`
	Taxation     decimal.Decimal    `json:"taxation"`   //税收
	Rebate       bool               `json:"rebate"`     //是否返佣金
	UserType     xtypes.UserType    `json:"user_type"`  //是否返佣金
	BeforeCash   decimal.Decimal    `json:"beforeCash"` //交易前现金
	AfterCash    decimal.Decimal    `json:"afterCash"`  //交易后现金
	RegisterTime int64              `json:"registerTime"`
	ChannelCode  string             `json:"channelCode"`
}

type ChangeInfo struct {
	Currency        xtypes.Currency         `json:"currency"`        // 币种ID
	Cash            decimal.Decimal         `json:"cash"`            // 变动的现金
	AmountKind      xtypes.WalletAmountKind `json:"amountKind"`      // 1真金 2奖金
	UserControlKind xtypes.UserControlKind  `json:"userControlKind"` // 1真金 2奖金
	PlayCurrency    xtypes.Currency         `json:"play_currency"`   // 支付币种
	PlayCash        decimal.Decimal         `json:"play_cash"`       // 变动的现金

}

type WithdrawWaterEventPayload struct {
	UID        int64           `json:"uid"`
	Currency   xtypes.Currency `json:"currency"`    // 币种ID
	CurrencyID int64           `json:"currency_id"` // 币种
	Amount     decimal.Decimal `json:"amount"`      // 金额
	IsFail     bool            `json:"is_fail"`
}
