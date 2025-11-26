package jwt

import (
	"context"
	"time"
	"xbase/utils/xconv"

	"github.com/go-redis/redis/v8"
)

type store struct {
	redis redis.UniversalClient
}

func (s *store) Get(ctx context.Context, key any) (any, error) {
	return s.redis.Get(ctx, xconv.String(key)).Result()
}

func (s *store) Set(ctx context.Context, key any, value any, duration time.Duration) error {
	return s.redis.Set(ctx, xconv.String(key), value, duration).Err()
}

func (s *store) Remove(ctx context.Context, keys ...any) (value any, err error) {
	list := make([]string, 0, len(keys))
	for _, key := range keys {
		list = append(list, xconv.String(key))
	}

	return s.redis.Del(ctx, list...).Result()
}
