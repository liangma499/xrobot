package user

import (
	"context"
	"sync"
	"xbase/errors"
	"xbase/log"
	"xrobot/internal/code"
	mysqlimp "xrobot/internal/component/mysql/mysql-default"
	"xrobot/internal/dao/user/internal"
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

type User struct {
	*internal.User
}

func NewUser(db *gorm.DB) *User {
	return &User{User: internal.NewUser(db)}
}

var (
	once     sync.Once
	instance *User
)

func Instance() *User {
	once.Do(func() {
		instance = NewUser(mysqlimp.Instance())
	})
	return instance
}
func (dao *User) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&model.User{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}

// FindOneByEmail 根据邮箱查询用户
func (dao *User) FindOneByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := dao.FindOne(ctx, func(cols *internal.Columns) any {
		return map[string]any{
			cols.Email: email,
		}
	})
	if err != nil {
		log.Errorf("find user failed, email = %s err = %v", email, err)
		return nil, errors.NewError(err, code.InternalError)
	}

	return user, nil
}

// CountByEmail 根据邮箱统计用户
func (dao *User) CountByEmail(ctx context.Context, email string) (int64, error) {
	count, err := dao.Count(ctx, func(cols *internal.Columns) any {
		return map[string]any{
			cols.Email: email,
		}
	})
	if err != nil {
		log.Errorf("count user failed, email = %s err = %v", email, err)
		return 0, errors.NewError(err, code.InternalError)
	}

	return count, nil
}

// FindOneByUid 根据id询用户
func (dao *User) FindOneByUid(ctx context.Context, uid int64) (*model.User, error) {
	user, err := dao.FindOne(ctx, func(cols *internal.Columns) any {
		return map[string]any{
			cols.ID: uid,
		}
	})
	if err != nil {
		log.Errorf("find user failed, id = %s err = %v", uid, err)
		return nil, err
	}

	return user, nil
}

func (dao *User) FindOneByAccount(ctx context.Context, account string) (*model.User, error) {
	user, err := dao.FindOne(ctx, func(cols *internal.Columns) any {
		return map[string]any{
			cols.Account: account,
		}
	})
	if err != nil {
		log.Errorf("find user failed, email = %s err = %v", account, err)
		return nil, errors.NewError(err, code.InternalError)
	}

	return user, nil
}
func (dao *User) FindOneByTelegramUserID(ctx context.Context, telegramUserID int64) (*model.User, error) {
	user, err := dao.FindOne(ctx, func(cols *internal.Columns) any {
		return map[string]any{
			cols.TelegramUserID: telegramUserID,
		}
	})
	if err != nil {
		log.Errorf("find user failed, email = %s err = %v", telegramUserID, err)
		return nil, errors.NewError(err, code.InternalError)
	}

	return user, nil
}
