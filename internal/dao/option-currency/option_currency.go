package optioncurrency

import (
	"context"
	"sync"
	"xbase/utils/xtime"
	mysqlimp "xrobot/internal/component/mysql/mysql-default"
	"xrobot/internal/dao/option-currency/internal"
	modelpkg "xrobot/internal/model"
	"xrobot/internal/xtypes"

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

type OptionCurrency struct {
	*internal.OptionCurrency
}

func NewOptionCurrency(db *gorm.DB) *OptionCurrency {
	return &OptionCurrency{OptionCurrency: internal.NewOptionCurrency(db)}
}

var (
	once     sync.Once
	instance *OptionCurrency
)

func Instance() *OptionCurrency {
	once.Do(func() {
		instance = NewOptionCurrency(mysqlimp.Instance())
	})
	return instance
}
func (dao *OptionCurrency) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.OptionCurrency{})
		if err != nil {
			panic(err)
		}
		now := xtime.Now()
		initData := []*modelpkg.OptionCurrency{
			{
				Currency:    xtypes.TRX,
				Url:         "www.baidu.com",
				Sort:        1,
				Memo:        "",
				Status:      xtypes.OptionStatus_Normal,
				OperateUid:  0,
				OperateUser: "",
				CreateAt:    now,
				UpdateAt:    now,
			},
			{
				Currency:    xtypes.USDT,
				Url:         "www.baidu.com",
				Sort:        1,
				Memo:        "",
				Status:      xtypes.OptionStatus_Normal,
				OperateUid:  0,
				OperateUser: "",
				CreateAt:    now,
				UpdateAt:    now,
			},
			{
				Currency:    xtypes.ENERGY,
				Url:         "www.baidu.com",
				Sort:        1,
				Memo:        "",
				Status:      xtypes.OptionStatus_Normal,
				OperateUid:  0,
				OperateUser: "",
				CreateAt:    now,
				UpdateAt:    now,
			},
			{
				Currency:    xtypes.BISHU,
				Url:         "www.baidu.com",
				Sort:        1,
				Memo:        "",
				Status:      xtypes.OptionStatus_Normal,
				OperateUid:  0,
				OperateUser: "",
				CreateAt:    now,
				UpdateAt:    now,
			},
		}

		dao.Insert(context.Background(), initData...)

	}
	return nil
}
