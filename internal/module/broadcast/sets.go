package broadcast

import (
	"context"
	"sync"
	"xbase/cluster"
	"xbase/log"
	"xbase/session"
	"xbase/task"
)

const maxTargetNum = 500 // 最大目标数

type sets struct {
	proxy   proxy
	rw      sync.RWMutex
	loaded  bool
	members map[int64]struct{}
}

// 添加成员
func (s *sets) add(uid int64) {
	s.rw.Lock()
	s.members[uid] = struct{}{}
	s.rw.Unlock()
}

// 删除成员
func (s *sets) del(uid int64) bool {
	s.rw.Lock()
	_, ok := s.members[uid]
	if ok {
		delete(s.members, uid)
	}
	s.rw.Unlock()

	return ok
}

// 广播消息
func (s *sets) broadcast(ctx context.Context, message *cluster.Message) {
	s.rw.RLock()
	targets := make([]int64, 0, len(s.members))
	for uid := range s.members {
		targets = append(targets, uid)
	}
	s.rw.RUnlock()

	count := len(targets)

	if count == 0 {
		return
	}

	for i := 0; i < count; i += maxTargetNum {
		func(start int) {
			end := start + maxTargetNum
			if end > count {
				end = count
			}

			task.AddTask(func() {
				_, err := s.proxy.Multicast(ctx, &cluster.MulticastArgs{
					Kind:    session.User,
					Targets: targets[start:end],
					Message: message,
				})
				if err != nil {
					log.Errorf("multicast message failed: %v", err)
				}
			})
		}(i)
	}
}
