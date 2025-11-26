package model

import (
	"time"
	"xrobot/internal/xtypes"
)

// 卡片配置表 用户类型
//
//go:generate xgorm-dao-generator -model-dir=. -model-names=OptionListenerAddress -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type OptionListenerAddress struct {
	NetWork     xtypes.NetWork      `gorm:"column:net_work;size:32;primarykey;comment:类型" json:"netWork,omitempty"`
	Address     string              `gorm:"column:address;size:128;primarykey;comment:地址" json:"address,omitempty"`
	ChannelCode string              `gorm:"column:channel_code;index;size:100;comment:渠道名称" json:"channelCode"`
	AddressKind xtypes.AddressKind  `gorm:"column:address_kind;size:32;comment:地址类型" json:"addressKind,omitempty"`
	Status      xtypes.OptionStatus `gorm:"column:status;size:32;comment:状态( 1启用,2禁用)" json:"-"`
	OperateUid  int64               `gorm:"column:operate_uid;size:64;comment:操作用户ID" json:"-"`
	OperateUser string              `gorm:"column:operate_user;size:64;comment:操作用户名" json:"-"`
	CreateAt    time.Time           `gorm:"column:created_at;type:timestamp;comment:创建时间戳" json:"-"`
	UpdateAt    time.Time           `gorm:"column:updated_at;type:timestamp;comment:修改时间戳" json:"-"`
}

func (c *OptionListenerAddress) TableName() string {
	return "option_listener_address"
}
