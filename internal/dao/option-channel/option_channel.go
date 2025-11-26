package optionchannel

import (
	"context"
	"fmt"
	"sync"
	"tron_robot/internal/code"
	mysqlimp "tron_robot/internal/component/mysql/mysql-default"
	redisoption "tron_robot/internal/component/redis/redis-option"
	"tron_robot/internal/dao/option-channel/internal"
	modelpkg "tron_robot/internal/model"
	tgtypes "tron_robot/internal/xtelegram/tg-types"
	"tron_robot/internal/xtypes"
	"xbase/errors"
	"xbase/log"
	"xbase/utils/xtime"

	"github.com/shopspring/decimal"
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

type OptionChannel struct {
	*internal.OptionChannel
}

func NewOptionChannel(db *gorm.DB) *OptionChannel {
	return &OptionChannel{OptionChannel: internal.NewOptionChannel(db)}
}

var (
	once     sync.Once
	instance *OptionChannel
)

func Instance() *OptionChannel {
	once.Do(func() {
		instance = NewOptionChannel(mysqlimp.Instance())
	})
	return instance
}

func (dao *OptionChannel) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.OptionChannel{})
		if err != nil {
			panic(err)
		}
		dao.initOptionChannel()
	}

	return nil

}

func (dao *OptionChannel) initOptionChannel() {
	now := xtime.Now()
	dao.Insert(context.Background(), &modelpkg.OptionChannel{
		ChannelCode: xtypes.OfficialChannelCode,
		Name:        "官方渠道",
		ChannelType: xtypes.ChannelTypeTG,
		TelegramCfg: &xtypes.OptionChannelTelegram{
			MainRobotLink:  "https://t.me/botfather_tron_bot",
			MainRobotToken: "8086135346:AAFFKKBpbdALKPzM2IENCjLw6Wy1-iik5nc",
			PushRobotLink:  "https://t.me/botfather_tron_bot",
			PushRobotToken: "8086135346:AAFFKKBpbdALKPzM2IENCjLw6Wy1-iik5nc",
			CustomerRobot:  "https://t.me/botfather_tron_bot",
			CommunityLink:  "社区链接",
			MainChannel:    "频道",
			MainChannelID:  0,
			GroupID:        0,
		},
		ChannelCfg: &xtypes.OptionChannelCfg{

			BiShuCfg: xtypes.MapComboKindInfo{
				tgtypes.XTelegramButton_NTP_20Bi: {
					Currency: xtypes.TRX.String(),       //笔数
					Price:    decimal.NewFromFloat(4.2), //单笔价格
				},
				tgtypes.XTelegramButton_NTP_30Bi: {
					Currency: xtypes.TRX.String(),       //笔数
					Price:    decimal.NewFromFloat(4.2), //单笔价格
				},
				tgtypes.XTelegramButton_NTP_50Bi: {
					Currency: xtypes.TRX.String(),       //笔数
					Price:    decimal.NewFromFloat(4.2), //单笔价格
				},
				tgtypes.XTelegramButton_NTP_100Bi: {
					Currency: xtypes.TRX.String(),       //笔数
					Price:    decimal.NewFromFloat(4.1), //单笔价格
				},
				tgtypes.XTelegramButton_NTP_200Bi: {
					Currency: xtypes.TRX.String(),       //笔数
					Price:    decimal.NewFromFloat(4.1), //单笔价格
				},
				tgtypes.XTelegramButton_NTP_300Bi: {
					Currency: xtypes.TRX.String(),       //笔数
					Price:    decimal.NewFromFloat(4.0), //单笔价格
				},
				tgtypes.XTelegramButton_NTP_500Bi: {
					Currency: xtypes.TRX.String(),       //笔数
					Price:    decimal.NewFromFloat(3.9), //单笔价格
				},
				tgtypes.XTelegramButton_NTP_1000Bi: {
					Currency: xtypes.TRX.String(),       //笔数
					Price:    decimal.NewFromFloat(3.8), //单笔价格
				},
				tgtypes.XTelegramButton_NTP_2000Bi: {
					Currency: xtypes.TRX.String(),       //笔数
					Price:    decimal.NewFromFloat(3.7), //单笔价格
				},
			}, //笔数配置

			PriceBiShuMax: 10,
			Customer:      "@GGGGGGGGGGG", //客服配置
			EnergySavings: "80%",
			ComboKindEnergyFlashRental: xtypes.ComboKindInfo{
				ComboKind: xtypes.ComboKind_Hour,
				Duration:  1,
			},

			ActivationFee: decimal.NewFromFloat(1.2), //激活费用

		},
		PriceCustomize: nil,
		PriceDefault: &xtypes.Price{
			EnergyPricesU:    decimal.NewFromFloat(64200),  //有U交易需要的能量
			EnergyPricesNou:  decimal.NewFromFloat(131000), //无U交易需要的能量
			TrxPriceU:        decimal.NewFromFloat(2.4),    //有U的转账价格
			TrxPriceNoU:      decimal.NewFromFloat(4.8),    //无U的转账价格
			CustomizeBalance: decimal.NewFromFloat(200),
		},
		OfficialType: xtypes.Official,
		Status:       xtypes.OptionStatus_Normal,
		Memo:         "初始渠道信息",
		OperateUid:   -1,
		OperateUser:  "init",
		CreateAt:     now,
		UpdateAt:     now,
	})

}

/*
AddrKind_None      AddrKind = 0 // 无
AddrKind_Official  AddrKind = 1 // 官方
AddrKind_Customize AddrKind = 2 // 自定义
AddrKind_Pay       AddrKind = 3 // 付款方是自己的
*/

func (dao *OptionChannel) CreateChannel(ctx context.Context, channelCfg *modelpkg.OptionChannel) error {
	if channelCfg == nil {
		return errors.NewError(code.InvalidArgument)
	}

	return nil
}
func (dao *OptionChannel) GetChannel(ctx context.Context, channelCode string) (*modelpkg.OptionChannel, error) {
	key := fmt.Sprintf(xtypes.CacheChannelCfgKey, channelCode)
	channelCfg := &modelpkg.OptionChannel{}

	err := redisoption.InstanceCache().GetSet(ctx, key, func() (any, error) {
		cfg, err := dao.FindOne(ctx, func(cols *internal.Columns) any {
			return map[string]any{
				cols.ChannelCode: channelCode,
			}
		})
		if err != nil {
			return nil, err
		}
		if cfg == nil {
			return nil, nil
		}
		return cfg, nil
	}, xtypes.CacheChannelCfgExpiration).Scan(channelCfg)
	if err != nil {
		if errors.Is(err, errors.ErrNil) {
			return dao.GetOfficialChannel(ctx)
		}

		log.Errorf("get channel failed, code = %d err = %v", channelCode, err)
		return nil, errors.NewError(err, code.InternalError)
	}
	if channelCfg.ChannelCode == "" {
		return dao.GetOfficialChannel(ctx)
	} else {
		return channelCfg, nil
	}

}
func (dao *OptionChannel) GetOfficialChannel(ctx context.Context) (*modelpkg.OptionChannel, error) {
	key := fmt.Sprintf(xtypes.CacheChannelCfgKey, xtypes.OfficialChannelCode)
	channelCfg := &modelpkg.OptionChannel{}

	err := redisoption.InstanceCache().GetSet(ctx, key, func() (any, error) {
		cfg, err := dao.FindOne(ctx, func(cols *internal.Columns) any {
			return map[string]any{
				cols.ChannelCode: xtypes.OfficialChannelCode,
			}
		})
		if err != nil {
			return nil, err
		}
		if cfg == nil {
			return nil, nil
		}
		return cfg, nil
	}, xtypes.CacheChannelCfgExpiration).Scan(channelCfg)
	if err != nil {
		if errors.Is(err, errors.ErrNil) {
			return nil, nil
		}

		log.Errorf("get channel failed, code = %d err = %v", xtypes.OfficialChannelCode, err)
		return nil, errors.NewError(err, code.InternalError)
	}
	return channelCfg, nil
}
