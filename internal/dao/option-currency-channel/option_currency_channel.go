package optioncurrencychannel

import (
	"context"
	"sync"
	mysqlimp "tron_robot/internal/component/mysql/mysql-default"
	"tron_robot/internal/dao/option-currency-channel/internal"
	modelpkg "tron_robot/internal/model"
	"tron_robot/internal/xtypes"
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

type OptionCurrencyChannel struct {
	*internal.OptionCurrencyChannel
}

func NewOptionCurrencyChannel(db *gorm.DB) *OptionCurrencyChannel {
	return &OptionCurrencyChannel{OptionCurrencyChannel: internal.NewOptionCurrencyChannel(db)}
}

var (
	once     sync.Once
	instance *OptionCurrencyChannel
)

func Instance() *OptionCurrencyChannel {
	once.Do(func() {
		instance = NewOptionCurrencyChannel(mysqlimp.Instance())
	})
	return instance
}
func (dao *OptionCurrencyChannel) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.OptionCurrencyChannel{})
		if err != nil {
			panic(err)
		}
		now := xtime.Now()
		initData := []*modelpkg.OptionCurrencyChannel{
			{
				Currency:     "USDT",
				Channel:      "Ethereum",
				ThirdKey:     "ethereum",
				Minutes:      "8",
				Confirmation: "200",
				Second:       "1",
				CollectFee:   decimal.NewFromFloat(10.0),
				GasFee:       decimal.NewFromFloat(2),
				Sort:         1,
				Memo:         "",
				Status:       xtypes.OptionStatus_Normal,
				OperateUid:   0,
				OperateUser:  "",
				CreateAt:     now,
				UpdateAt:     now,
			},
			{
				Currency:     "USDT",
				Channel:      "TRC20",
				ThirdKey:     "tron",
				Minutes:      "2",
				Confirmation: "1",
				Second:       "1",
				CollectFee:   decimal.NewFromFloat(10.0),
				GasFee:       decimal.NewFromFloat(2),
				Sort:         1,
				Memo:         "",
				Status:       xtypes.OptionStatus_Normal,
				OperateUid:   0,
				OperateUser:  "",
				CreateAt:     now,
				UpdateAt:     now,
			},
			{
				Currency:     "USDT",
				Channel:      "Polygon",
				ThirdKey:     "polygon",
				Minutes:      "3",
				Confirmation: "15",
				Second:       "1",
				CollectFee:   decimal.NewFromFloat(10.0),
				GasFee:       decimal.NewFromFloat(2),
				Sort:         1,
				Memo:         "",
				Status:       xtypes.OptionStatus_Normal,
				OperateUid:   0,
				OperateUser:  "",
				CreateAt:     now,
				UpdateAt:     now,
			},

			{
				Currency:     "USDC",
				Channel:      "Ethereum",
				ThirdKey:     "ethereum",
				Minutes:      "8",
				Confirmation: "200",
				Second:       "1",
				CollectFee:   decimal.NewFromFloat(10.0),
				GasFee:       decimal.NewFromFloat(2),
				Sort:         1,
				Memo:         "",
				Status:       xtypes.OptionStatus_Normal,
				OperateUid:   0,
				OperateUser:  "",
				CreateAt:     now,
				UpdateAt:     now,
			},
			{
				Currency:     "USDC",
				Channel:      "TRC20",
				ThirdKey:     "tron",
				Minutes:      "2",
				Confirmation: "1",
				Second:       "1",
				CollectFee:   decimal.NewFromFloat(10.0),
				GasFee:       decimal.NewFromFloat(2),
				Sort:         1,
				Memo:         "",
				Status:       xtypes.OptionStatus_Normal,
				OperateUid:   0,
				OperateUser:  "",
				CreateAt:     now,
				UpdateAt:     now,
			},
			{
				Currency:     "USDC",
				Channel:      "Polygon",
				ThirdKey:     "polygon",
				Minutes:      "3",
				Confirmation: "15",
				Second:       "1",
				CollectFee:   decimal.NewFromFloat(10.0),
				GasFee:       decimal.NewFromFloat(2),
				Sort:         1,
				Memo:         "",
				Status:       xtypes.OptionStatus_Normal,
				OperateUid:   0,
				OperateUser:  "",
				CreateAt:     now,
				UpdateAt:     now,
			},
		}

		dao.Insert(context.Background(), initData...)
	}
	return nil
}
