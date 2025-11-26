package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// 用户佣金表
//
//go:generate xgorm-dao-generator -model-dir=. -model-names=UserCommission -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type UserCommission struct {
	UID                 int64           `gorm:"column:uid;primaryKey;type:bigint;"` // 主键
	CommissionTotal     decimal.Decimal `gorm:"column:commission_total;type:decimal(32,4);default:0;comment:总佣金"`
	CommissionAvailable decimal.Decimal `gorm:"column:commission_available;type:decimal(32,4);default:0;comment:可用佣金"`
	CreateAt            time.Time       `gorm:"column:created_at;type:timestamp;comment:创建时间戳"`
	UpdateAt            time.Time       `gorm:"column:updated_at;type:timestamp;comment:修改时间戳"` // 上次登录IP
}

func (u *UserCommission) TableName() string {
	return "user_commission"
}
