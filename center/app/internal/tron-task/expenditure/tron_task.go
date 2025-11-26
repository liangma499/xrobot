package expenditure

import (
	"context"
	"sync"
	"time"
	"tron_robot/internal/xtypes"
	"xbase/cluster/node"
	"xbase/log"
	"xbase/task"
)

type TronTask struct {
	bInitTimer      bool
	cycleTime       time.Duration
	cycleTimeVerify time.Duration
	ctx             context.Context
	proxy           *node.Proxy
	mu              sync.Mutex
}

var (
	trc20Task *TronTask
	once      sync.Once
)

const (
	cycleTime          = 2    //订单任务间隔时间
	OrderTaskRunCount  = 30   // 一个订单最多查询次数
	MaxTaskUserInfo    = 1024 //一次处理多少个订单
	netWorkChannelType = xtypes.NetWorkChannelType_TRON
	apiKind            = xtypes.APITrongrid
)

func Instance(proxy *node.Proxy) *TronTask {

	once.Do(func() {
		trc20Task = &TronTask{
			cycleTime:       cycleTime,
			cycleTimeVerify: 1,
			ctx:             context.Background(),
			bInitTimer:      false,
			proxy:           proxy,
		}
		trc20Task.initTask()
	})

	return trc20Task
}

func (s *TronTask) initTask() {
	//定时任务只初始化一次
	if s.bInitTimer {
		return
	}
	s.bInitTimer = true
	task.AddTask(func() {
		s.initFetchTransaction()
	})

}

func (s *TronTask) initFetchTransaction() {
	time.Sleep(10 * time.Second)
	for {
		s.timeAfter()
		s.fetchPlatformPendingOrder()
		log.Warnf("拉取待支付订单")
	}
}

func (s *TronTask) timeAfter() {
	tmr := time.After(s.cycleTime * time.Second)
	<-tmr
}

// 转出数据
func (s *TronTask) fetchPlatformPendingOrder() {
	s.mu.Lock()
	defer s.mu.Unlock()
}
