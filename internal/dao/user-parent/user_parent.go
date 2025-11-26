package userparent

import (
	"sync"
	mysqlimp "tron_robot/internal/component/mysql/mysql-default"
	"tron_robot/internal/dao/user-parent/internal"
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

type UserParent struct {
	*internal.UserParent
}

func NewUserParent(db *gorm.DB) *UserParent {
	return &UserParent{UserParent: internal.NewUserParent(db)}
}

var (
	once     sync.Once
	instance *UserParent
)

func Instance() *UserParent {
	once.Do(func() {
		instance = NewUserParent(mysqlimp.Instance())
	})
	return instance
}
func (dao *UserParent) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&model.UserParent{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}
