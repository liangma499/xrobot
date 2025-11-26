package casbin

import (
	"sync"
	mysqldefault "tron_robot/internal/component/mysql/mysql-default"
	"xbase/etc"
	"xbase/log"
	casbin "xbase/utils/gorm-casbin"
)

var (
	once     sync.Once
	instance *casbin.Enforcer
)

type Config struct {
	Model string `json:"model"`
	DSN   string `json:"dsn"`
	Table string `json:"table"`
}

func Instance() *casbin.Enforcer {
	once.Do(func() {
		instance = NewInstance("etc.casbin.default")
	})

	return instance
}

// NewInstance 新建实例
func NewInstance[T string | Config | *Config](config T) *casbin.Enforcer {
	var (
		conf     *Config
		v        any = config
		database any
	)

	switch c := v.(type) {
	case string:
		conf = &Config{}
		if err := etc.Get(c).Scan(conf); err != nil {
			log.Fatalf("load casbin config failed: %v", err)
		}
	case Config:
		conf = &c
	case *Config:
		conf = c
	}

	if conf.DSN != "" {
		database = conf.DSN
	} else {
		database = mysqldefault.Instance()
	}

	enforcer, err := casbin.NewEnforcer(&casbin.Options{
		Model:    conf.Model,
		Enable:   true,
		Autoload: true,
		Table:    conf.Table,
		Database: database,
	})
	if err != nil {
		log.Fatalf("new a casbin enforcer instance failed: %v", err)
	}

	return enforcer
}
