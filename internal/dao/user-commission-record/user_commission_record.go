package usercommissionrecord

import (
	"sync"
	mysqlimp "tron_robot/internal/component/mysql/mysql-default"
	"tron_robot/internal/dao/user-commission-record/internal"
	modelpkg "tron_robot/internal/model"

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

type UserCommissionRecord struct {
	*internal.UserCommissionRecord
}

func NewUserCommissionRecord(db *gorm.DB) *UserCommissionRecord {
	return &UserCommissionRecord{UserCommissionRecord: internal.NewUserCommissionRecord(db)}
}

var (
	once     sync.Once
	instance *UserCommissionRecord
)

func Instance() *UserCommissionRecord {
	once.Do(func() {
		instance = NewUserCommissionRecord(mysqlimp.Instance())
	})
	return instance
}
func (dao *UserCommissionRecord) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.UserCommissionRecord{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}
