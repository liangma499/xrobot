package model

import (
	"time"
	"tron_robot/internal/xtypes"

	"github.com/shopspring/decimal"
)

// 平台待处理订单 需要转账 激活 的订单全问放这里面统一处理
//
//go:generate xgorm-dao-generator -model-dir=. -model-names=PlatformPendingOrder -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type PlatformPendingOrder struct {
	ID           int64              `gorm:"column:id;size:64;primarykey;autoIncrement"` // 主键
	UID          int64              `gorm:"column:uid;type:bigint;not null;"`           // 主键
	Code         string             `gorm:"column:code;size:16;not null;"`              // 编号
	ChannelCode  string             `gorm:"column:channel_code;size:32;not null;"`      // 渠道编码
	Address      string             `gorm:"column:address;size:128;not null;"`
	ToCurrency   xtypes.Currency    `gorm:"column:to_currency;size:32;default:'';comment:转出币种"` // 币种
	ToAmount     decimal.Decimal    `gorm:"column:to_amount;type:decimal(32,9);default:0;comment:转出金额"`
	FromCurrency xtypes.Currency    `gorm:"column:from_currency;size:32;default:'';comment:转入币种"` // 币种
	FromAmount   decimal.Decimal    `gorm:"column:from_amount;type:decimal(32,9);default:0;comment:转入金额"`
	Energy       int64              `gorm:"column:energy;type:64;default:0;comment:代理能量"`
	Memo         string             `gorm:"column:memo;size:512;default:'';comment:说明"` // 币种
	Stauts       xtypes.OrderStatus `gorm:"column:stauts;size:32;default:0;index:idx_ns;comment:0无 1待激活 2激活转账完成，待验证激活是否成功 3已激活，代理能量 4兑换订单，99处理完成"`
	CreateAt     time.Time          `gorm:"column:created_at;type:timestamp;comment:创建时间戳" json:"-"`
	UpdateAt     time.Time          `gorm:"column:updated_at;type:timestamp;comment:修改时间戳" json:"-"`
}

func (u *PlatformPendingOrder) TableName() string {
	return "platform_pending_order"
}
