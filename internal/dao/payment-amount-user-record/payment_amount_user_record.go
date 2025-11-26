package paymentamountuserrecord

import (
	"sync"
	"tron_robot/internal/dao/payment-amount-user-record/internal"
	modelpkg "tron_robot/internal/model"
	mysqlimp "tron_robot/internal/component/mysql/mysql-default"
	"gorm.io/gorm"
)

type (
	Columns = internal.Columns
	OrderBy = internal.OrderBy
	FilterFunc = internal.FilterFunc
	UpdateFunc = internal.UpdateFunc
	ColumnFunc = internal.ColumnFunc
	OrderFunc = internal.OrderFunc
)

type PaymentAmountUserRecord struct {
	*internal.PaymentAmountUserRecord
}

func NewPaymentAmountUserRecord(db *gorm.DB) *PaymentAmountUserRecord {
	return &PaymentAmountUserRecord{PaymentAmountUserRecord: internal.NewPaymentAmountUserRecord(db)}
}
var (
	once     sync.Once
	instance *PaymentAmountUserRecord
)

func Instance() *PaymentAmountUserRecord {
	once.Do(func() {
		instance = NewPaymentAmountUserRecord(mysqlimp.Instance())
	})
	return instance
}
func (dao *PaymentAmountUserRecord) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.PaymentAmountUserRecord{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}
