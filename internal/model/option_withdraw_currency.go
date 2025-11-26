package model

import (
	"time"
	"tron_robot/internal/xtypes"

	"github.com/shopspring/decimal"
)

// 卡片配置表 用户类型
//
//go:generate xgorm-dao-generator -model-dir=. -model-names=OptionWithdrawCurrency -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type OptionWithdrawCurrency struct {
	ID           int64               `gorm:"column:id;size:64;primarykey;autoIncrement"  json:"id"`
	WithdrawType string              `gorm:"column:withdraw_type;uniqueIndex:uin_currency_channel_net_type;size:64;comment:commission card"  json:"withdraw_type"`
	Currency     xtypes.Currency     `gorm:"column:currency;uniqueIndex:uin_currency_channel_net_type;size:32;comment:币种,USDT,USDC,大写" json:"currency"`
	Channel      string              `gorm:"column:channel;uniqueIndex:uin_currency_channel_net_type;size:32;comment:小写" json:"channel"`
	ThirdKey     string              `gorm:"column:third_key;size:32;comment:三方Key" json:"third_key"`
	GasFee       decimal.Decimal     `gorm:"column:gas_fee;type:decimal(32,4);default:0;comment:gas_fee"  json:"gas_fee"`
	Min          decimal.Decimal     `gorm:"column:min;type:decimal(32,4);default:0;comment:单笔最小提现" json:"min"`
	Max          decimal.Decimal     `gorm:"column:max;type:decimal(32,4);default:0;comment:单笔最大提现" json:"max"`
	Premium      decimal.Decimal     `gorm:"column:premium;size:64;default:0;comment:手续费用 只认USTD" json:"premium"` //
	RemainderMin decimal.Decimal     `gorm:"column:remainder_min;type:decimal(32,4);default:0;comment:剩余" json:"remainderMin"`
	Sort         int                 `gorm:"column:sort;size:64;comment:排序ID,大的排到前面" json:"-"`
	Memo         string              `gorm:"column:memo;size:512;comment:说明" json:"-"`
	Status       xtypes.OptionStatus `gorm:"column:status;size:32;comment:状态( 1启用,2禁用)"  json:"status"`
	OperateUid   int64               `gorm:"column:operate_uid;size:64;comment:操作用户ID"  json:"-"`
	OperateUser  string              `gorm:"column:operate_user;size:64;comment:操作用户名"  json:"-"`
	CreateAt     time.Time           `gorm:"column:created_at;type:timestamp;comment:创建时间戳" json:"-"`
	UpdateAt     time.Time           `gorm:"column:updated_at;type:timestamp;comment:修改时间戳"  json:"-"`
}

// `gorm:"column:login_at;size:64"`
func (c *OptionWithdrawCurrency) TableName() string {
	return "option_withdraw_currency"
}
