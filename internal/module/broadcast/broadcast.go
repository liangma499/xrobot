package broadcast

import (
	"context"
	"sync"
	userevt "tron_robot/internal/event/user"
	"xbase/cluster"
)

type proxy interface {
	// Multicast 推送组播消息
	Multicast(ctx context.Context, args *cluster.MulticastArgs) (int64, error)
}

type Broadcast struct {
	proxy proxy
	users *sets
	rw    sync.RWMutex
}

func NewBroadcast(proxy proxy) *Broadcast {
	f := &Broadcast{
		proxy: proxy,
		users: &sets{
			proxy:   proxy,
			members: make(map[int64]struct{}),
		},
	}

	f.init()

	return f
}

func (b *Broadcast) init() {
	userevt.SubscribeLogin(func(payload *userevt.LoginPayload) {
		if payload.UID == 0 {
			return
		}

		b.rw.RLock()
		b.users.add(payload.UID)
		b.rw.RUnlock()
	})

	userevt.SubscribeOffline(func(uid int64) {
		if uid == 0 {
			return
		}
		b.rw.RLock()
		b.users.del(uid)
		b.rw.RUnlock()
	})
}

// Broadcast 推送广播消息给某个代理的所有玩家
func (b *Broadcast) Broadcast(ctx context.Context, message *cluster.Message) {
	b.users.broadcast(ctx, message)
}
