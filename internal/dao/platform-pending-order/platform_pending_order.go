package platformpendingorder

import (
	"sync"
	mysqlimp "xrobot/internal/component/mysql/mysql-default"
	"xrobot/internal/dao/platform-pending-order/internal"
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

type PlatformPendingOrder struct {
	*internal.PlatformPendingOrder
}

func NewPlatformPendingOrder(db *gorm.DB) *PlatformPendingOrder {
	return &PlatformPendingOrder{PlatformPendingOrder: internal.NewPlatformPendingOrder(db)}
}

var (
	once     sync.Once
	instance *PlatformPendingOrder
)

func Instance() *PlatformPendingOrder {
	once.Do(func() {
		instance = NewPlatformPendingOrder(mysqlimp.Instance())
	})
	return instance
}
func (dao *PlatformPendingOrder) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.PlatformPendingOrder{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}
