package userbalancedeposittransaction

import (
	"sync"
	mysqlimp "xrobot/internal/component/mysql/mysql-default"
	"xrobot/internal/dao/user-balance-deposit-transaction/internal"
	modelpkg "xrobot/internal/model"

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

type UserBalanceDepositTransaction struct {
	*internal.UserBalanceDepositTransaction
}

func NewUserBalanceDepositTransaction(db *gorm.DB) *UserBalanceDepositTransaction {
	return &UserBalanceDepositTransaction{UserBalanceDepositTransaction: internal.NewUserBalanceDepositTransaction(db)}
}

var (
	once     sync.Once
	instance *UserBalanceDepositTransaction
)

func Instance() *UserBalanceDepositTransaction {
	once.Do(func() {
		instance = NewUserBalanceDepositTransaction(mysqlimp.Instance())
	})
	return instance
}
func (dao *UserBalanceDepositTransaction) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.UserBalanceDepositTransaction{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}
