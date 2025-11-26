package userwithdrawrecord

import (
	"sync"
	mysqlimp "tron_robot/internal/component/mysql/mysql-default"
	"tron_robot/internal/dao/user-withdraw-record/internal"
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

type UserWithdrawRecord struct {
	*internal.UserWithdrawRecord
}

func NewUserWithdrawRecord(db *gorm.DB) *UserWithdrawRecord {
	return &UserWithdrawRecord{UserWithdrawRecord: internal.NewUserWithdrawRecord(db)}
}

var (
	once     sync.Once
	instance *UserWithdrawRecord
)

func Instance() *UserWithdrawRecord {
	once.Do(func() {
		instance = NewUserWithdrawRecord(mysqlimp.Instance())
	})
	return instance
}
func (dao *UserWithdrawRecord) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.UserWithdrawRecord{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}
