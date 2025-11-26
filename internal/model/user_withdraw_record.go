package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// 用户佣金表记录表
//
//go:generate xgorm-dao-generator -model-dir=. -model-names=UserWithdrawRecord -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type UserWithdrawRecord struct {
	ID           int64           `gorm:"column:id;primaryKey;size:64,autoIncrement"`
	UID          int64           `gorm:"column:uid;type:bigint;index:index_uid_tuid;comment:用户UId"`    //
	ThirdUID     string          `gorm:"column:third_uid;size:64;index:index_uid_tuid;comment:用户用UID"` // 三方用户ID
	Amount       decimal.Decimal `gorm:"column:amount;type:decimal(32,4);default:0;comment:提现金额"`
	BeforeAmount decimal.Decimal `gorm:"column:before_amount;type:decimal(32,4);default:0;comment:变化前"`
	AfterAmount  decimal.Decimal `gorm:"column:after_amount;type:decimal(32,4);default:0;comment:变化后"`
	GasFee       decimal.Decimal `gorm:"column:gas_fee;type:decimal(16,8);comment:GasFee"` // 佣金比例
	Premium      decimal.Decimal `gorm:"column:premium;size:64;comment:手续费用 只认USTD"`       //
	WithdrawType string          `gorm:"column:withdraw_type;size:64;comment:commission card"`
	DayZeroTime  int64           `gorm:"column:day_zero_time;size:64;index;"` //记录的当天
	CreateAt     time.Time       `gorm:"column:created_at;type:timestamp;comment:创建时间戳"`
	UpdateAt     time.Time       `gorm:"column:updated_at;type:timestamp;comment:修改时间戳"` // 上次登录IP
}

func (u *UserWithdrawRecord) TableName() string {
	return "user_withdraw_record"
}
