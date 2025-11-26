package optionwithdrawcurrency

import (
	"context"
	"sync"
	mysqlimp "tron_robot/internal/component/mysql/mysql-default"
	"tron_robot/internal/dao/option-withdraw-currency/internal"
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

type OptionWithdrawCurrency struct {
	*internal.OptionWithdrawCurrency
}

func NewOptionWithdrawCurrency(db *gorm.DB) *OptionWithdrawCurrency {
	return &OptionWithdrawCurrency{OptionWithdrawCurrency: internal.NewOptionWithdrawCurrency(db)}
}

var (
	once     sync.Once
	instance *OptionWithdrawCurrency
)

func Instance() *OptionWithdrawCurrency {
	once.Do(func() {
		instance = NewOptionWithdrawCurrency(mysqlimp.Instance())
	})
	return instance
}


func (dao *OptionWithdrawCurrency) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.OptionWithdrawCurrency{})
		if err != nil {
			panic(err)
		}
		now := xtime.Now()
		intdata := []*modelpkg.OptionWithdrawCurrency{

			{
				WithdrawType: xtypes.WithdrawType_Card,
				Currency:     "USDT",
				Channel:      "TRC20",
				ThirdKey:     "tron",
				GasFee:       decimal.NewFromFloat(2),
				Min:          decimal.NewFromFloat(5),
				Max:          decimal.NewFromFloat(99999),
				Premium:      decimal.NewFromFloat(0.018),
				RemainderMin: decimal.NewFromFloat(1),
				Sort:         1,
				Memo:         "",
				Status:       xtypes.OptionStatus_Normal,
				OperateUid:   0,
				OperateUser:  "",
				CreateAt:     now,
				UpdateAt:     now,
			},

			{
				WithdrawType: xtypes.WithdrawType_Commission,
				Currency:     "USDT",
				Channel:      "TRC20",
				ThirdKey:     "tron",
				GasFee:       decimal.NewFromFloat(2),
				Min:          decimal.NewFromFloat(5),
				Max:          decimal.NewFromFloat(99999),
				Premium:      decimal.NewFromFloat(0.018),
				RemainderMin: decimal.NewFromFloat(0),
				Sort:         1,
				Memo:         "",
				Status:       xtypes.OptionStatus_Normal,
				OperateUid:   0,
				OperateUser:  "",
				CreateAt:     now,
				UpdateAt:     now,
			},
		}
		dao.Insert(context.Background(), intdata...)
	}
	return nil
}
