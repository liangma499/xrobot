package snowflake

import (
	"sync"
	"xbase/etc"
	"xbase/log"

	"github.com/bwmarrin/snowflake"
)

var (
	once     sync.Once
	instance *snowflake.Node
)

type Config struct {
	Node int64 `json:"node"`
}

func Instance() *snowflake.Node {
	once.Do(func() {
		instance = NewInstance("etc.snowflake")
	})

	return instance
}

// NewInstance 新建实例
func NewInstance[T string | Config | *Config](config T) *snowflake.Node {
	var (
		conf *Config
		v    any = config
	)

	switch c := v.(type) {
	case string:
		conf = &Config{}
		if err := etc.Get(c).Scan(conf); err != nil {
			log.Fatalf("load snowflake config failed: %v", err)
		}
	case Config:
		conf = &c
	case *Config:
		conf = c
	}

	node, err := snowflake.NewNode(conf.Node)
	if err != nil {
		log.Fatalf("create snowflake's node failed: %v", err)
	}

	return node
}
