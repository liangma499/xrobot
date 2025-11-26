package userloginlog

import (
	"sync"
	mysqlimp "tron_robot/internal/component/mysql/mysql-default"
	"tron_robot/internal/dao/user-login-log/internal"
	"tron_robot/internal/model"

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

type UserLoginLog struct {
	*internal.UserLoginLog
}

func NewUserLoginLog(db *gorm.DB) *UserLoginLog {
	return &UserLoginLog{UserLoginLog: internal.NewUserLoginLog(db)}
}

var (
	once     sync.Once
	instance *UserLoginLog
)

func Instance() *UserLoginLog {
	once.Do(func() {
		instance = NewUserLoginLog(mysqlimp.Instance())
	})
	return instance
}
func (dao *UserLoginLog) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&model.UserLoginLog{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}
