package model

import (
	"time"
	"tron_robot/internal/xtypes"
)

//go:generate xgorm-dao-generator -model-dir=. -model-names=OptionChannel -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type OptionChannel struct {
	ID             int64                         `gorm:"column:id;size:64;primarykey;autoIncrement" json:"id"` // 主键
	ChannelCode    string                        `gorm:"column:channel_code;unique;size:100;comment:渠道名称" json:"channelCode"`
	SupervisorInfo xtypes.SupervisorChannel      `gorm:"column:supervisor_info;type:json" json:"supervisorInfo" redis:"supervisorInfo"` // 上级ID
	Name           string                        `gorm:"column:name;unique;size:100;comment:渠道名称" json:"name"`
	ChannelType    xtypes.ChannelType            `gorm:"column:channel_type;comment:1 普通;2:tg机器人" json:"channel_type"`
	TelegramCfg    *xtypes.OptionChannelTelegram `gorm:"column:telegram_cfg;type:json;comment:TG机器人相关配置" json:"telegram_cfg"`
	ChannelCfg     *xtypes.OptionChannelCfg      `gorm:"column:channel_cfg;type:json;comment:渠道相关配置" json:"channel_cfg"`
	PriceCustomize *xtypes.Price                 `gorm:"column:price_customize;type:json;comment:价格相关配置自定义" json:"price_customize"`
	PriceDefault   *xtypes.Price                 `gorm:"column:price_default;type:json;comment:价格相关配置默认" json:"price_default"`
	OfficialType   xtypes.OfficialTypeKind       `gorm:"column:official_type;comment:状态( 1官方,2非官方)" json:"official_type"`
	Status         xtypes.OptionStatus           `gorm:"column:status;size:32;comment:状态( 1启用,2禁用)" json:"status"`
	Memo           string                        `gorm:"column:memo;size:512;comment:备注" json:"-"`           //
	OperateUid     int64                         `gorm:"column:operate_uid;size:64;comment:操作用户ID" json:"-"` //
	OperateUser    string                        `gorm:"column:operate_user;size:64;comment:操作用户名" json:"-"` //
	CreateAt       time.Time                     `gorm:"column:created_at;type:timestamp;comment:创建时间戳" json:"-"`
	UpdateAt       time.Time                     `gorm:"column:updated_at;type:timestamp;comment:修改时间戳" json:"-"`
}

func (c *OptionChannel) TableName() string {
	return "option_channel"
}
func (c *OptionChannel) Clone() *OptionChannel {
	if c == nil {
		return nil
	}
	return &OptionChannel{
		ID:             c.ID,
		ChannelCode:    c.ChannelCode,
		SupervisorInfo: c.SupervisorInfo.Clone(),
		Name:           c.Name,
		ChannelType:    c.ChannelType,
		TelegramCfg:    c.TelegramCfg.Clone(),
		ChannelCfg:     c.ChannelCfg.Clone(),
		PriceCustomize: c.PriceCustomize.Clone(),
		PriceDefault:   c.PriceDefault.Clone(),
		OfficialType:   c.OfficialType,
		Status:         c.Status,
		Memo:           c.Memo,        //
		OperateUid:     c.OperateUid,  //
		OperateUser:    c.OperateUser, //
		CreateAt:       c.CreateAt,
		UpdateAt:       c.UpdateAt,
	}
}
