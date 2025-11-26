package model

import (
	"time"
	"xrobot/internal/xtypes"
)

type AmountUserRecordStauts int32

const (
	AmountUserRecordStauts_Processed AmountUserRecordStauts = 1
	AmountUserRecordStauts_Overdue   AmountUserRecordStauts = 2
)

//go:generate xgorm-dao-generator -model-dir=. -model-names=PaymentAmountUserRecord -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type PaymentAmountUserRecord struct {
	ID          int64           `gorm:"column:uid;primaryKey;type:bigint;" json:"uid" redis:"uid"` // 主键
	OrderID     string          `gorm:"column:order_id;unique;not null;size:32"`
	UserID      int64           `gorm:"column:user_id;not null;size:64;index"`                                // 账号
	ChannelName string          `gorm:"column:channel_name;size:64" json:"channel_name" redis:"channel_name"` // 渠道名
	ChannelCode string          `gorm:"column:channel_code;size:32" json:"channel_code" redis:"channel_code"`
	Amount      string          `gorm:"column:amount;size:64;uniqueIndex:uin_amount_currency;" json:"amount" redis:"amount"` // 编号
	Currency    string          `gorm:"column:currency;uniqueIndex:uin_amount_currency;size:32;comment:币种,USDT,USDC,大写"`
	Extend      *AmountUserInfo `gorm:"column:extend;type:json"`
	Usage       xtypes.Usage    `gorm:"column:usage;size:32"`
	Status      int32           `gorm:"column:status;not null;size:32;comment:1已处理,2过期"`
	CreateAt    time.Time       `gorm:"column:created_at;type:timestamp;comment:创建时间戳" json:"-"`
	UpdateAt    time.Time       `gorm:"column:updated_at;type:timestamp;comment:修改时间戳" json:"-"`
}

func (u *PaymentAmountUserRecord) TableName() string {
	return "payment_amount_user_record"
}
