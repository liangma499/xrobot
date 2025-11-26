package redisoption

import (
	"sync"
	"time"

	"github.com/go-redis/redis/v8"

	cacheRedis "xbase/cache/redis"
	rediscomp "xrobot/internal/component/redis"
)

var (
	once      sync.Once
	instance  redis.UniversalClient
	onceCache sync.Once
	cache     *cacheRedis.Cache
)

func newNewInstance() redis.UniversalClient {
	return rediscomp.NewInstance("etc.redis.option")
}

// Instance 获取单例
func Instance() redis.UniversalClient {
	once.Do(func() {
		instance = newNewInstance()
	})

	return instance
}
func InstanceCache() *cacheRedis.Cache {
	onceCache.Do(func() {
		cache = cacheRedis.NewCache(cacheRedis.WithClient(newNewInstance()),
			cacheRedis.WithPrefix("cache"),
			cacheRedis.WithNilExpiration(10*time.Second),
			cacheRedis.WithMinExpiration(15*time.Second),
			cacheRedis.WithMaxExpiration(1*time.Hour),
		)
	})
	return cache
}
