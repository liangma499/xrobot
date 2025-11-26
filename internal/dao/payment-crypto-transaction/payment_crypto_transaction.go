package paymentcryptotransaction

import (
	"sync"
	mysqlimp "xrobot/internal/component/mysql/mysql-default"
	"xrobot/internal/dao/payment-crypto-transaction/internal"
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

type PaymentCryptoTransaction struct {
	*internal.PaymentCryptoTransaction
}

func NewPaymentCryptoTransaction(db *gorm.DB) *PaymentCryptoTransaction {
	return &PaymentCryptoTransaction{PaymentCryptoTransaction: internal.NewPaymentCryptoTransaction(db)}
}

var (
	once     sync.Once
	instance *PaymentCryptoTransaction
)

func Instance() *PaymentCryptoTransaction {
	once.Do(func() {
		instance = NewPaymentCryptoTransaction(mysqlimp.Instance())
	})
	return instance
}
func (dao *PaymentCryptoTransaction) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.PaymentCryptoTransaction{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}
