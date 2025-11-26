package userbalancedeposittransaction

import (
	"sync"
	mysqlimp "tron_robot/internal/component/mysql/mysql-default"
	"tron_robot/internal/dao/user-balance-deposit-transaction/internal"
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
