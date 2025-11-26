package model

import (
	"time"
	"tron_robot/internal/xtypes"

	"github.com/shopspring/decimal"
)

// 卡片配置表 用户类型
//
//go:generate xgorm-dao-generator -model-dir=. -model-names=OptionCurrencyChannel -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type OptionCurrencyChannel struct {
	ID           int64               `gorm:"column:id;size:64;primarykey;autoIncrement"` // 主键
	Currency     xtypes.Currency     `gorm:"column:currency;uniqueIndex:uin_currency_channel_net;size:32;comment:币种,USDT,USDC,大写"`
	Channel      string              `gorm:"column:channel;uniqueIndex:uin_currency_channel_net;size:32;comment:小写"`
	ThirdKey     string              `gorm:"column:third_key;size:32;comment:三方Key"`
	Minutes      string              `gorm:"column:minutes;size:32;comment:到账时间"`
	Confirmation string              `gorm:"column:confirmation;size:32;comment:到账时间"`
	Second       string              `gorm:"column:second;size:32;comment:到账时间"`
	CollectFee   decimal.Decimal     `gorm:"column:collect_fee;type:decimal(32,4);default:0;comment:归集费用"`
	GasFee       decimal.Decimal     `gorm:"column:gas_fee;type:decimal(32,4);default:0;comment:gas_fee"`
	Sort         int                 `gorm:"column:sort;size:64;comment:排序ID,大的排到前面"`
	Memo         string              `gorm:"column:memo;size:512;comment:说明"`
	Status       xtypes.OptionStatus `gorm:"column:status;size:32;comment:状态( 1启用,2禁用)"`
	OperateUid   int64               `gorm:"column:operate_uid;size:64;comment:操作用户ID"`
	OperateUser  string              `gorm:"column:operate_user;size:64;comment:操作用户名"`
	CreateAt     time.Time           `gorm:"column:created_at;type:timestamp;comment:创建时间戳"`
	UpdateAt     time.Time           `gorm:"column:updated_at;type:timestamp;comment:修改时间戳"`
}

// `gorm:"column:login_at;size:64"`
func (c *OptionCurrencyChannel) TableName() string {
	return "option_currency_channel"
}
