package wallet

import (
	"database/sql"
	"time"
	"xrobot/internal/xtypes"
)

type incrBalanceArgs struct {
	UID             int64                   `json:"uid"`            // 用户ID
	Currency        xtypes.Currency         `json:"currency"`       // 货币
	Cash            float64                 `json:"cash"`           // 现金
	RelatedID       string                  `json:"relatedID"`      // 关联ID
	TradeType       xtypes.TradeType        `json:"tradeType"`      // 交易类型
	TradeStatus     xtypes.TradeStatus      `json:"tradeStatus"`    // 交易状态
	Taxation        float64                 `json:"taxation"`       // 税收
	BetAmount       float64                 `json:"betAmount"`      // 下注金额
	Rebate          bool                    `json:"rebate"`         // 是否返佣金
	UserType        xtypes.UserType         `json:"userType"`       // 是否返佣金
	ChannelCode     string                  `json:"channel"`        // 渠道
	NickName        string                  `json:"nickName"`       // 用户昵称
	GameName        string                  `json:"gameName"`       // 游戏名
	AmountKind      xtypes.WalletAmountKind `json:"amountKind"`     // 1真金 2奖金
	PlayCurrencyID  int64                   `json:"playCurrencyID"` // 币种
	PlayCurrency    string                  `json:"playCurrency"`   // 币种
	PlayCash        float64                 `json:"playCash"`       // 变动的现金
	TransactionID   string                  `json:"transactionID"`
	UserControlKind xtypes.UserControlKind  `json:"userControlKind"`
}

type decrBalanceArgs struct {
	UID             int64                   `json:"uid"`         // 用户ID
	Currency        xtypes.Currency         `json:"currency"`    // 货币
	Cash            float64                 `json:"cash"`        // 现金
	RelatedID       string                  `json:"relatedID"`   // 关联ID
	TradeType       xtypes.TradeType        `json:"tradeType"`   // 交易类型
	TradeStatus     xtypes.TradeStatus      `json:"tradeStatus"` // 交易状态
	Taxation        float64                 `json:"taxation"`    // 税收
	BetAmount       float64                 `json:"bet_amount"`
	Rebate          bool                    `json:"rebate"`         //是否返佣金
	UserType        xtypes.UserType         `json:"userType"`       //是否返佣金
	ChannelCode     string                  `json:"channel"`        //渠道
	NickName        string                  `json:"nickName"`       //用户昵称
	GameName        string                  `json:"gameName"`       //游戏名
	AmountKind      xtypes.WalletAmountKind `json:"amountKind"`     // 1真金 2奖金
	PlayCurrencyID  int64                   `json:"playCurrencyID"` // 币种
	PlayCurrency    string                  `json:"playCurrency"`   // 币种
	PlayCash        float64                 `json:"playCash"`       // 变动的现金
	TransactionID   string                  `json:"transactionID"`
	UserControlKind xtypes.UserControlKind  `json:"userControlKind"`
}

type queryTradeArgs struct {
	TradeNO   string `json:"tradeNO"`   // 交易号
	RelatedID string `json:"relatedID"` // 关联ID
}

type TradeList struct {
	ID          int64                   `gorm:"column:id" json:"id"`                                     // 交易ID
	NO          string                  `gorm:"column:no" json:"no"`                                     // 交易单号
	UID         int64                   `gorm:"column:uid" json:"uid"`                                   // 用户ID
	Currency    xtypes.Currency         `json:"currency"`                                                // 货币
	Type        xtypes.TradeType        `gorm:"column:type" json:"type"`                                 // 交易类型
	Channel     xtypes.TradeChannel     `gorm:"column:channel" json:"channel"`                           // 交易渠道
	ChannelCode string                  `gorm:"column:channel_code" json:"channel_code"`                 // 用户渠道码
	Status      xtypes.TradeStatus      `gorm:"column:status" json:"status"`                             // 交易状态
	BeforeCash  float64                 `gorm:"column:before_cash" json:"beforeCash"`                    // 交易前的现金
	AfterCash   float64                 `gorm:"column:after_cash" json:"afterCash"`                      // 交易后的现金
	ChangeCash  float64                 `gorm:"column:change_cash" json:"changeCash"`                    // 交易变动的现金
	BetCash     float64                 `gorm:"column:bet_cash" json:"bet_cash"`                         // 唯一关联ID
	RelatedID   string                  `gorm:"column:related_id" json:"relatedId"`                      // 唯一关联ID
	Taxation    float64                 `gorm:"column:taxation" json:"taxation"`                         // 唯一关联ID
	CreatedAt   time.Time               `gorm:"column:created_at" json:"createdAt"`                      // 创建时间
	UpdatedAt   time.Time               `gorm:"column:updated_at" json:"updatedAt"`                      // 更新时间
	DeletedAt   sql.NullTime            `gorm:"column:deleted_at" json:"deletedAt"`                      // 软删除时间
	Code        string                  `gorm:"column:code" json:"code" redis:"code"`                    // 编号
	Nickname    string                  `gorm:"column:nickname" json:"nickname" redis:"nickname"`        // 昵称
	AmountKind  xtypes.WalletAmountKind `gorm:"column:amount_kind" json:"amountKind" redis:"amountKind"` // 1真金 2奖金
}
