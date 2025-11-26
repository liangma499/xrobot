package platformpendingorder

import (
	"sync"
	"tron_robot/internal/dao/platform-pending-order/internal"
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
