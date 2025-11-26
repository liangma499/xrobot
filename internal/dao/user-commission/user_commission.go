package usercommission

import (
	"context"
	"fmt"
	"sync"
	"xbase/utils/xtime"
	mysqlimp "xrobot/internal/component/mysql/mysql-default"
	"xrobot/internal/dao/user-commission/internal"
	modelpkg "xrobot/internal/model"

	"github.com/shopspring/decimal"
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

type UserCommission struct {
	*internal.UserCommission
}

func NewUserCommission(db *gorm.DB) *UserCommission {
	return &UserCommission{UserCommission: internal.NewUserCommission(db)}
}

var (
	once     sync.Once
	instance *UserCommission
)

func Instance() *UserCommission {
	once.Do(func() {
		instance = NewUserCommission(mysqlimp.Instance())
	})
	return instance
}
func (dao *UserCommission) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.UserCommission{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (dao *UserCommission) UpdateCommissionSql(uid int64, commission decimal.Decimal) string {
	now := xtime.Now()
	commissionStr := commission.RoundFloor(4).String()

	return fmt.Sprintf("INSERT INTO `%s` ( `%s`, `%s`, `%s`, `%s`, `%s`) VALUES (%d, %s, %s, '%v', '%v') ON DUPLICATE KEY UPDATE `%s` = `%s` + %s,`%s` = `%s` + %s,`%s` = '%v'",
		dao.TableName,
		dao.Columns.UID,
		dao.Columns.CommissionTotal,
		dao.Columns.CommissionAvailable,
		dao.Columns.CreateAt,
		dao.Columns.UpdateAt, //上面是intsert

		uid,
		commissionStr,
		commissionStr,
		now,
		now,

		dao.Columns.CommissionTotal,
		dao.Columns.CommissionTotal,
		commissionStr,
		dao.Columns.CommissionAvailable,
		dao.Columns.CommissionAvailable,
		commissionStr,
		dao.Columns.UpdateAt,
		now,
	)

}
func (dao *UserCommission) UpdateCommissionWithDraw(ctx context.Context, uid int64, commission float64, bFail bool) error {
	if bFail {
		_, err := dao.Update(ctx, func(cols *internal.Columns) any {
			return map[string]any{
				cols.UID: uid,
			}
		}, func(cols *internal.Columns) any {
			return map[string]any{
				cols.CommissionAvailable: gorm.Expr(fmt.Sprintf("%s + ?", cols.CommissionAvailable), commission),
			}
		})
		return err
	} else {
		_, err := dao.Update(ctx, func(cols *internal.Columns) any {
			return map[string]any{
				cols.UID: uid,
			}
		}, func(cols *internal.Columns) any {
			return map[string]any{
				cols.CommissionAvailable: gorm.Expr(fmt.Sprintf("%s - ?", cols.CommissionAvailable), commission),
			}
		})
		return err
	}

}
