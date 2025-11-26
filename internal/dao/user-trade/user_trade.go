package usertrade

import (
	"gorm.io/gorm"
	"sync"
	mysqlimp "tron_robot/internal/component/mysql/mysql-default"
	"tron_robot/internal/dao/user-trade/internal"
	modelpkg "tron_robot/internal/model"
)

type (
	Columns    = internal.Columns
	OrderBy    = internal.OrderBy
	FilterFunc = internal.FilterFunc
	UpdateFunc = internal.UpdateFunc
	ColumnFunc = internal.ColumnFunc
	OrderFunc  = internal.OrderFunc
)

type UserTrade struct {
	*internal.UserTrade
}

func NewUserTrade(db *gorm.DB) *UserTrade {
	return &UserTrade{UserTrade: internal.NewUserTrade(db)}
}

var (
	once     sync.Once
	instance *UserTrade
)

func Instance() *UserTrade {
	once.Do(func() {
		instance = NewUserTrade(mysqlimp.Instance())
	})
	return instance
}
func (dao *UserTrade) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.UserTrade{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}
