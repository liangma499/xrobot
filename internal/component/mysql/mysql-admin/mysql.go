package mysqladmin

import (
	"sync"
	"tron_robot/internal/component/mysql"

	"gorm.io/gorm"
)

var (
	once     sync.Once
	instance *gorm.DB
)

func Instance() *gorm.DB {
	once.Do(func() {
		instance = mysql.NewInstance("etc.mysql.admin")
	})

	return instance
}
