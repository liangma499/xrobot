package user

import (
	"context"
	"sync"
	"time"
	userevt "tron_robot/internal/event/user"
	usersvc "tron_robot/internal/service/user"
	userpb "tron_robot/internal/service/user/pb"
	"tron_robot/internal/xtypes"
	"xbase/codes"
	"xbase/errors"
	"xbase/transport"
	"xbase/utils/xconv"
	"xbase/utils/xtime"

	"golang.org/x/sync/singleflight"
)

type Manager struct {
	fn         transport.NewMeshClient      // API代理
	users      sync.Map                     // 用户
	rw         sync.RWMutex                 // 锁
	timeWheels map[int64]map[int64]struct{} // 时间轮
	dyingUsers map[int64]int64              // 弥留的玩家（等待被清理）
	sfg        singleflight.Group
}

func NewManager(fn transport.NewMeshClient) *Manager {
	mgr := &Manager{
		fn:         fn,
		timeWheels: make(map[int64]map[int64]struct{}),
		dyingUsers: make(map[int64]int64),
	}

	mgr.init()

	go mgr.clearDyingUser()

	return mgr
}

// 初始化相关事件
func (mgr *Manager) init() {
	// 订阅用户信息变动事件
	userevt.SubscribeInfoChange(func(uid int64) {
		_, _ = mgr.LoadUser(uid, true)
	})

	// 订阅用户上线事件
	userevt.SubscribeLogin(func(payload *userevt.LoginPayload) {
		_, _ = mgr.UserLogin(payload.UID)
	})

	// 订阅用户离线事件
	userevt.SubscribeOffline(func(uid int64) {
		mgr.RemoveUser(uid, 120)
	})
}

// GetUser 获取用户
func (mgr *Manager) GetUser(uid int64) (*User, bool) {
	v, ok := mgr.users.Load(uid)
	if !ok {
		return nil, false
	}

	return v.(*User), true
}

// LoadUser 加载一个用户
func (mgr *Manager) LoadUser(uid int64, isReload ...bool) (*User, error) {
	v, ok := mgr.users.Load(uid)
	if ok {
		if len(isReload) == 0 || !isReload[0] {
			return v.(*User), nil
		}
	} else {
		if len(isReload) > 0 && isReload[0] {
			return nil, nil
		}
	}

	val, err, _ := mgr.sfg.Do(xconv.String(uid), func() (any, error) {
		client, err := usersvc.NewClient(mgr.fn)
		if err != nil {
			return nil, errors.NewError(codes.InternalError, err)
		}

		u, err := client.FetchUser(context.Background(), &userpb.FetchUserArgs{
			UID: uid,
		})
		if err != nil {
			return nil, errors.NewError(codes.Convert(err))
		}

		user := &User{
			ID:          u.User.ID,
			Nickname:    u.User.Nickname,
			Avatar:      u.User.Avatar,
			Email:       u.User.Email,
			UserType:    xtypes.UserType(u.User.UserType),
			ChannelCode: u.User.ChannelCode,
		}

		mgr.users.Store(u.User.ID, user)

		return user, nil
	})
	if err != nil {
		return nil, err
	}

	return val.(*User), nil
}

// RemoveUser 移除用户
func (mgr *Manager) RemoveUser(uid int64, delay ...int64) (*User, bool) {
	mgr.rw.Lock()
	defer mgr.rw.Unlock()
	user, ok := mgr.GetUser(uid)
	if !ok {
		return user, ok
	}

	if len(delay) > 0 && delay[0] > 0 {
		mgr.doAddDyingUser(uid, delay[0])
	} else {
		mgr.users.Delete(user.ID)
	}

	return user, ok
}

// ResumeUser 恢复用户
func (mgr *Manager) ResumeUser(uid int64) (*User, bool) {
	mgr.doRemoveDyingUser(uid)

	return mgr.GetUser(uid)

}

// ResumeUser 恢复用户
func (mgr *Manager) UserLogin(uid int64) (*User, error) {
	mgr.doRemoveDyingUser(uid)
	user, ok := mgr.GetUser(uid)
	if ok {
		return mgr.LoadUser(uid, true)
	}
	return user, nil
}

// 清理弥留用户
func (mgr *Manager) clearDyingUser() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C

		now := xtime.Now().Unix()

		mgr.rw.Lock()
		for clearTime := now - 5; clearTime <= now; clearTime++ {
			timeWheels, ok := mgr.timeWheels[clearTime]
			if !ok {
				continue
			}

			for uid := range timeWheels {
				delete(mgr.dyingUsers, uid)

				mgr.users.Delete(uid)
			}

			delete(mgr.timeWheels, clearTime)
		}
		mgr.rw.Unlock()
	}
}

// 加入弥留用户
func (mgr *Manager) doAddDyingUser(uid, delay int64) {
	mgr.rw.Lock()
	defer mgr.rw.Unlock()

	if clearTime, ok := mgr.dyingUsers[uid]; ok {
		if timeWheels, ok := mgr.timeWheels[clearTime]; ok {
			delete(timeWheels, uid)
		}
	}

	clearTime := xtime.Now().Unix() + delay
	mgr.dyingUsers[uid] = clearTime
	timeWheels, ok := mgr.timeWheels[clearTime]
	if !ok {
		timeWheels = make(map[int64]struct{})
		mgr.timeWheels[clearTime] = timeWheels
	}
	timeWheels[uid] = struct{}{}
}

// 执行移除弥留用户操作
func (mgr *Manager) doRemoveDyingUser(uid int64) {
	mgr.rw.Lock()
	if clearTime, ok := mgr.dyingUsers[uid]; ok {
		delete(mgr.dyingUsers, uid)
		delete(mgr.timeWheels[clearTime], uid)
	}
	mgr.rw.Unlock()
}
