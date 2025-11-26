package optionbaseconfig

import (
	"context"
	"sync"
	mysqlimp "tron_robot/internal/component/mysql/mysql-default"
	"tron_robot/internal/dao/option-base-config/internal"
	modelpkg "tron_robot/internal/model"
	"tron_robot/internal/xtypes"
	"xbase/utils/xtime"

	"gorm.io/gorm"
)

type (
	Columns    = internal.Columns
	OrderBy    = internal.OrderBy
	FilterFunc = internal.FilterFunc
	UpdateFunc = internal.UpdateFunc
	ColumnFunc = internal.ColumnFunc
	OrderFunc  = internal.OrderFunc
)

type OptionBaseConfig struct {
	*internal.OptionBaseConfig
}

func NewOptionBaseConfig(db *gorm.DB) *OptionBaseConfig {
	return &OptionBaseConfig{OptionBaseConfig: internal.NewOptionBaseConfig(db)}
}

var (
	once     sync.Once
	instance *OptionBaseConfig
)

func Instance() *OptionBaseConfig {
	once.Do(func() {
		instance = NewOptionBaseConfig(mysqlimp.Instance())
	})
	return instance
}
func (dao *OptionBaseConfig) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.OptionBaseConfig{})
		if err != nil {
			panic(err)
		}

		now := xtime.Now()
		initData := []*modelpkg.OptionBaseConfig{
			{
				Key:         xtypes.BaseUrlKey,
				Value:       "https://127.0.0.1",
				Memo:        "基础图片地址",
				OperateUid:  0,
				OperateUser: "",
				Status:      xtypes.OptionStatus_Normal,
				CreateAt:    now,
				UpdateAt:    now,
			},
			{
				Key:         xtypes.AvatarUrlKey,
				Value:       "http://127.0.0.1:8086",
				Memo:        "头像地址",
				OperateUid:  0,
				OperateUser: "",
				Status:      xtypes.OptionStatus_Normal,
				CreateAt:    now,
				UpdateAt:    now,
			},

			{
				Key:         xtypes.InviteCodeUrlKey,
				Value:       "http://127.0.0.1:6600/?channelCode=${channelCode}&inviteCode=${code}",
				Memo:        "分享地址",
				OperateUid:  0,
				OperateUser: "",
				Status:      xtypes.OptionStatus_Normal,
				CreateAt:    now,
				UpdateAt:    now,
			},

			{
				Key:         xtypes.WebLobbyUrlKey,
				Value:       "http://127.0.0.1:6600/",
				Memo:        "大厅地址",
				OperateUid:  0,
				OperateUser: "",
				Status:      xtypes.OptionStatus_Normal,
				CreateAt:    now,
				UpdateAt:    now,
			},
			{
				Key:         xtypes.TelegramWebhookUrlKey,
				Value:       "https://robot.77bgame.cn/tg/webhook",
				Memo:        "telegram webhook回调地址",
				OperateUid:  0,
				OperateUser: "",
				Status:      xtypes.OptionStatus_Normal,
				CreateAt:    now,
				UpdateAt:    now,
			},
			{
				Key:         xtypes.TelegramApiUrlKey,
				Value:       "https://api.telegram.org/bot${token}",
				Memo:        "telegram 回调",
				OperateUid:  0,
				OperateUser: "",
				Status:      xtypes.OptionStatus_Normal,
				CreateAt:    now,
				UpdateAt:    now,
			},
			{
				Key:         xtypes.CommissionRatioKey,
				Value:       "0.2",
				Memo:        "佣金比例",
				OperateUid:  0,
				OperateUser: "",
				Status:      xtypes.OptionStatus_Normal,
				CreateAt:    now,
				UpdateAt:    now,
			},
			{
				Key:         xtypes.CommissionCreateCardKey,
				Value:       "1",
				Memo:        "开卡奖励",
				OperateUid:  0,
				OperateUser: "",
				Status:      xtypes.OptionStatus_Normal,
				CreateAt:    now,
				UpdateAt:    now,
			},
			{
				Key:         xtypes.PointsMineRatioKey,
				Value:       "50",
				Memo:        "自己充值的获取积分比例",
				OperateUid:  0,
				OperateUser: "",
				Status:      xtypes.OptionStatus_Normal,
				CreateAt:    now,
				UpdateAt:    now,
			},
			{
				Key:         xtypes.PointsChildRatioKey,
				Value:       "500",
				Memo:        "子级充值获取的积分比例",
				OperateUid:  0,
				OperateUser: "",
				Status:      xtypes.OptionStatus_Normal,
				CreateAt:    now,
				UpdateAt:    now,
			},
			{
				Key:         xtypes.ServerStatusKey,
				Value:       "1",
				Memo:        "1:Opened(正常);2:Maintain(维护);3:Closed(关闭);4:Error(错误)",
				OperateUid:  0,
				OperateUser: "",
				Status:      xtypes.OptionStatus_Normal,
				CreateAt:    now,
				UpdateAt:    now,
			},
			{
				Key:         xtypes.WhiteListKey,
				Value:       ";1;2;3;",
				Memo:        "维护白名单必须要用英文的;开始和结束",
				OperateUid:  0,
				OperateUser: "",
				Status:      xtypes.OptionStatus_Normal,
				CreateAt:    now,
				UpdateAt:    now,
			},
		}
		dao.Insert(context.Background(), initData...)
	}
	return nil
}
