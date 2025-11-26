package paymentnetwork

import (
	"sync"
	"tron_robot/internal/dao/payment-network/internal"
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

type PaymentNetwork struct {
	*internal.PaymentNetwork
}

func NewPaymentNetwork(db *gorm.DB) *PaymentNetwork {
	return &PaymentNetwork{PaymentNetwork: internal.NewPaymentNetwork(db)}
}
var (
	once     sync.Once
	instance *PaymentNetwork
)

func Instance() *PaymentNetwork {
	once.Do(func() {
		instance = NewPaymentNetwork(mysqlimp.Instance())
	})
	return instance
}
func (dao *PaymentNetwork) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.PaymentNetwork{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}
