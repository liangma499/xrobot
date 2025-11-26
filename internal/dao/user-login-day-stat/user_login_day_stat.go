package userlogindaystat

import (
	"context"
	"fmt"
	"sync"
	mysqlimp "xrobot/internal/component/mysql/mysql-default"
	"xrobot/internal/dao/user-login-day-stat/internal"
	"xrobot/internal/model"

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

type UserLoginDayStat struct {
	*internal.UserLoginDayStat
}

func NewUserLoginDayStat(db *gorm.DB) *UserLoginDayStat {
	return &UserLoginDayStat{UserLoginDayStat: internal.NewUserLoginDayStat(db)}
}

var (
	once     sync.Once
	instance *UserLoginDayStat
)

func Instance() *UserLoginDayStat {
	once.Do(func() {
		instance = NewUserLoginDayStat(mysqlimp.Instance())
	})
	return instance
}
func (dao *UserLoginDayStat) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&model.UserLoginDayStat{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}
func (dao *UserLoginDayStat) InsertOrUpdate(ctx context.Context, uid, zeroTime int64) error {
	sql := fmt.Sprintf("INSERT INTO `%s` (`%s`, `%s`, `%s`) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE `%s` = `%s` + 1",
		dao.TableName,
		dao.Columns.UID,
		dao.Columns.ZeroTime,
		dao.Columns.LoginTime,

		dao.Columns.LoginTime,
		dao.Columns.LoginTime,
	)

	return dao.Database.WithContext(ctx).Exec(sql,
		uid,
		zeroTime,
		1,
	).Error
}
