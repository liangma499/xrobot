package model

import (
	"time"
	"xrobot/internal/xtypes"

	"github.com/shopspring/decimal"
)

//go:generate xgorm-dao-generator -model-dir=. -model-names=UserTrade -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type UserTrade struct {
	ID          int64                   `gorm:"column:id" json:"id"`                                      // 交易ID
	NO          string                  `gorm:"column:no;unique;size:32;" json:"no"`                      // 交易单号
	UID         int64                   `gorm:"column:uid;index:index_u_c;size:64" json:"uid"`            // 用户ID
	Currency    xtypes.Currency         `gorm:"column:currency;size:32;index:index_u_c;" json:"currency"` // 交易的货币类型
	Type        xtypes.TradeType        `gorm:"column:type;size:32;index:index_g_t" json:"type"`          // 交易类型
	Channel     xtypes.TradeChannel     `gorm:"column:channel;size:32;index" json:"channel"`              // 交易渠道
	Status      xtypes.TradeStatus      `gorm:"column:status;size:32" json:"status"`                      // 交易状态
	BeforeCash  decimal.Decimal         `gorm:"column:before_cash;type:decimal(32,10)" json:"beforeCash"` // 交易前的现金
	AfterCash   decimal.Decimal         `gorm:"column:after_cash;type:decimal(32,10)" json:"afterCash"`   // 交易后的现金
	ChangeCash  decimal.Decimal         `gorm:"column:change_cash;type:decimal(32,10)" json:"changeCash"` // 交易变动的现金
	BetCash     decimal.Decimal         `gorm:"column:bet_cash;type:decimal(32,10)" json:"bet_cash"`      // 唯一关联ID
	ChannelCode string                  `gorm:"column:channel_code;size:32" json:"channel_code"`          // 用户渠道码
	RelatedID   string                  `gorm:"column:related_id;index;size:32" json:"relatedId"`         // 唯一关联ID
	Taxation    decimal.Decimal         `gorm:"column:taxation;type:decimal(32,10);" json:"taxation"`     // 唯一关联ID
	UserType    xtypes.UserType         `gorm:"column:user_type" json:"user_type"`                        // 用户类型
	AmountKind  xtypes.WalletAmountKind `gorm:"column:amount_kind" json:"amountKind"`                     // 1真金 2奖金
	CreatedAt   time.Time               `gorm:"column:created_at" json:"createdAt"`                       // 创建时间
	UpdatedAt   time.Time               `gorm:"column:updated_at" json:"updatedAt"`                       // 更新时间
}

func (u *UserTrade) TableName() string {
	return "user_trade"
}
