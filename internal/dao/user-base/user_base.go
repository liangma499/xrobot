package userbase

import (
	"context"
	"fmt"
	"sync"
	"xbase/cache"
	"xbase/errors"
	"xbase/log"
	"xbase/utils/xconv"
	"xrobot/internal/code"
	mysqlimp "xrobot/internal/component/mysql/mysql-default"
	"xrobot/internal/dao/user-base/internal"
	"xrobot/internal/model"
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

type UserBase struct {
	*internal.UserBase
}

func NewUserBase(db *gorm.DB) *UserBase {
	return &UserBase{UserBase: internal.NewUserBase(db)}
}

var (
	once     sync.Once
	instance *UserBase
)

func Instance() *UserBase {
	once.Do(func() {
		instance = NewUserBase(mysqlimp.Instance())
	})
	return instance
}
func (dao *UserBase) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&model.UserBase{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}

// FindOneByEmail 根据邮箱查询用户
func (dao *UserBase) FindOneByEmail(ctx context.Context, email string) (*model.UserBase, error) {
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
func (dao *UserBase) CountByEmail(ctx context.Context, email string) (int64, error) {
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
func (dao *UserBase) FindOneByUid(ctx context.Context, uid int64) (*model.UserBase, error) {
	user, err := dao.FindOne(ctx, func(cols *internal.Columns) any {
		return map[string]any{
			cols.UID: uid,
		}
	})
	if err != nil {
		log.Errorf("find user failed, id = %s err = %v", uid, err)
		return nil, err
	}

	return user, nil
}

func (dao *UserBase) FindOneByAccount(ctx context.Context, account string) (*model.UserBase, error) {
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

// 获取用户
func (dao *UserBase) GetUserBase(ctx context.Context, uid int64) (*model.UserBase, error) {
	key := fmt.Sprintf(xtypes.CacheUserBaseKey, uid)
	user := &model.UserBase{}

	err := cache.GetSet(ctx, key, func() (any, error) {
		return dao.FindOne(ctx, func(cols *internal.Columns) any {
			return map[string]any{
				cols.UID: uid,
			}
		})
	}).Scan(user)
	if err != nil {
		if errors.Is(err, errors.ErrNil) {
			return nil, nil
		}

		log.Errorf("get user failed, uid = %d err = %v", uid, err)
		return nil, errors.NewError(err, code.InternalError)
	}

	return user, nil
}

// 获取用户UID
func (dao *UserBase) DoGetUserBaseUID(ctx context.Context, ucode string) (int64, error) {
	key := fmt.Sprintf(xtypes.CacheUserBaseCodeToUIDKey, ucode)
	uid := int64(0)

	err := cache.GetSet(ctx, key, func() (any, error) {
		userbase, err := dao.FindOne(ctx, func(cols *internal.Columns) any {
			return map[string]any{
				cols.Code: ucode,
			}
		}, func(cols *internal.Columns) []string {
			return []string{
				cols.UID,
			}
		})
		if err != nil {
			return 0, err
		}
		if userbase == nil {
			return 0, code.NotFound.Err()
		}
		return userbase.UID, nil
	}).Scan(&uid)
	if err != nil {
		if errors.Is(err, errors.ErrNil) {
			return 0, errors.NewError(code.InvalidArgument)
		}

		log.Errorf("get user failed, ucode = %s err = %v", ucode, err)
		return 0, errors.NewError(err, code.InternalError)
	}

	return uid, nil
}

// 根据用户编号获取用户
func (dao *UserBase) DoGetUserBaseByCode(ctx context.Context, ucode string) (*model.UserBase, error) {
	uid, err := dao.DoGetUserBaseUID(ctx, ucode)
	if err != nil {
		return nil, err
	}

	return dao.GetUserBase(ctx, uid)
}

// 更新用户
func (dao *UserBase) DoUpdateUserBase(ctx context.Context, uid int64, values map[string]any) error {
	return dao.DoUpdateBaseUserHandle(ctx, uid, func() error {
		_, err := dao.Update(ctx, func(cols *internal.Columns) any {
			return map[string]any{
				cols.UID: uid,
			}
		}, func(cols *internal.Columns) any {
			return values
		})
		if err != nil {
			log.Errorf("update user failed, uid = %d values = %s err = %v", uid, xconv.Json(values), err)
			return errors.NewError(err, code.InternalError)
		}

		return nil
	})
}

func (dao *UserBase) DoUpdateBaseUserHandle(ctx context.Context, uid int64, fn func() error) error {
	key := fmt.Sprintf(xtypes.CacheUserBaseKey, uid)

	_, err := cache.Delete(ctx, key)
	if err != nil {
		log.Errorf("delete user's cache failed, uid = %d err = %v", uid, err)
		return errors.NewError(err, code.InternalError)
	}

	if err = fn(); err != nil {
		return err
	}

	_, _ = cache.Delete(ctx, key)

	return nil
}
