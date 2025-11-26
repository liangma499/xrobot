package optioncurrency

import (
	"context"
	"sync"
	mysqlimp "tron_robot/internal/component/mysql/mysql-default"
	"tron_robot/internal/dao/option-currency/internal"
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
