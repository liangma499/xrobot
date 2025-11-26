package main

import (
	"fmt"
	mysqldefault "xrobot/internal/component/mysql/mysql-default"

	"xbase/config"
	"xbase/etc"
	"xbase/log"

	"xbase/config/etcd"
)

func main() {
	// 设置配置中心
	config.SetConfigurator(config.NewConfigurator(config.WithSources(etcd.NewSource())))
	db := mysqldefault.Instance()

	ignoreTables := etc.Get("etc.truncate.ignoreTables").Strings()
	tables := make([]string, 0)
	checks := make(map[string]struct{}, len(ignoreTables))

	for _, table := range ignoreTables {
		checks[table] = struct{}{}
	}

	err := db.Raw("SELECT TABLE_NAME FROM information_schema.`TABLES` WHERE TABLE_SCHEMA LIKE ?", "tg.platform").Scan(&tables).Error
	if err != nil {
		log.Fatalf("exec sql failed: %v", err)
	}

	for _, table := range tables {

		_, ok := checks[table]
		if ok {
			continue
		}

		db.Exec(fmt.Sprintf("TRUNCATE `%s`", table))
	}

	fmt.Println("truncate ok")
}
