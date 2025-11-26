package model

import (
	"time"
	"tron_robot/internal/xtypes"

	"github.com/shopspring/decimal"
)

// 用户佣金表记录表
//
//go:generate xgorm-dao-generator -model-dir=. -model-names=UserCommissionRecord -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type UserCommissionRecord struct {
	ID                    int64                 `gorm:"column:id;primaryKey;size:64,autoIncrement"`
	UID                   int64                 `gorm:"column:uid;type:bigint;index:index_uid_cid;comment:用户UId"` //
	CID                   int64                 `gorm:"column:cid;type:bigint;index:index_uid_cid;comment:子级ID"`
	DirectCID             int64                 `gorm:"column:direct_cid;type:bigint;comment:直属ID"`
	ChildLevel            int64                 `gorm:"column:child_level;type:bigint;comment:子级等级"`
	Commission            decimal.Decimal       `gorm:"column:commission;type:decimal(32,4);default:0;comment:佣金"`
	CommissionRatio       decimal.Decimal       `gorm:"column:commission_ratio;type:decimal(32,4);comment:佣金比例"`          // 佣金比例
	CommissionRealRatio   decimal.Decimal       `gorm:"column:commission_real_ratio;type:decimal(32,4);comment:实际获取比例"`   // 佣金比例
	CommissionDirectRatio decimal.Decimal       `gorm:"column:commission_direct_ratio;type:decimal(32,4);comment:直属获取比例"` // 佣金比例
	CommissionType        xtypes.CommissionType `gorm:"column:commission_type;size:32;comment:commission card"`
	DayZeroTime           int64                 `gorm:"column:day_zero_time;size:64;index;"` //记录的当天
	CreateAt              time.Time             `gorm:"column:created_at;type:timestamp;comment:创建时间戳"`
	UpdateAt              time.Time             `gorm:"column:updated_at;type:timestamp;comment:修改时间戳"` // 上次登录IP
}

func (u *UserCommissionRecord) TableName() string {
	return "user_commission_record"
}
