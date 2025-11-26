package model

import (
	"time"
	"tron_robot/internal/xtypes"
)

// 卡片配置表 用户类型
//
//go:generate xgorm-dao-generator -model-dir=. -model-names=OptionCurrency -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type OptionCurrency struct {
	Currency    xtypes.Currency     `gorm:"column:currency;primarykey;size:32;comment:大写"`
	Url         string              `gorm:"column:url;size:256;comment:链接地址"`
	Sort        int                 `gorm:"column:sort;size:64;comment:排序ID,大的排到前面"` //
	Memo        string              `gorm:"column:memo;size:512;comment:说明"`         //
	Status      xtypes.OptionStatus `gorm:"column:status;size:32;comment:状态( 1启用,2禁用)"`
	OperateUid  int64               `gorm:"column:operate_uid;size:64;comment:操作用户ID"`      //
	OperateUser string              `gorm:"column:operate_user;size:64;comment:操作用户名"`      //
	CreateAt    time.Time           `gorm:"column:created_at;type:timestamp;comment:创建时间戳"` //
	UpdateAt    time.Time           `gorm:"column:updated_at;type:timestamp;comment:修改时间戳"` //
}

// `gorm:"column:login_at;size:64"`
func (c *OptionCurrency) TableName() string {
	return "option_currency"
}
