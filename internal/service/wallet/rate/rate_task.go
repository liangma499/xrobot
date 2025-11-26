package rate

import (
	"context"
	"strings"
	"sync"
	"time"
	optionCurrencyRateOpt "tron_robot/internal/option/option-currency-rate"
	"xbase/log"
	"xbase/task"
	"xbase/utils/xconv"

	"github.com/shopspring/decimal"
)

type RechargeTask struct {
	bInitTimer         bool
	timeMillisecond    time.Duration
	runRunBolckTask    task.Pool
	runRunRechargeTask task.Pool
}

var (
	rechargeTask *RechargeTask
	once         sync.Once
)

const (
	normalTime  = 60000
	rate_tokens = "usd,usdt,trx"
	currencies  = "usd,usdt,trx"
)
const (
	runRunBolckTaskSize        = 8
	runRunRechargeTaskSize     = 64
	runRunBolckTaskNonblocking = false
)

func Instance() *RechargeTask {

	once.Do(func() {
		rechargeTask = &RechargeTask{
			bInitTimer:         false,
			timeMillisecond:    normalTime,
			runRunBolckTask:    task.NewPool(task.WithSize(runRunBolckTaskSize), task.WithNonblocking(runRunBolckTaskNonblocking)),
			runRunRechargeTask: task.NewPool(task.WithSize(runRunRechargeTaskSize)),
		}

	})

	return rechargeTask
}

func (s *RechargeTask) InitTimer() {
	//定时任务只初始化一次

	if s.bInitTimer {
		return
	}
	s.bInitTimer = true
	s.runRunBolckTask.AddTask(func() {
		s.initTimer()
	})
}

func (s *RechargeTask) initTimer() {
	time.Sleep(10 * time.Second)

	log.Warnf("进入拉取税率")

	for {
		s.doFetchRate()
		s.timeAfterMillisecond()

	}
}

func (s *RechargeTask) timeAfterMillisecond() {
	tmr := time.After(s.timeMillisecond * time.Millisecond)
	<-tmr
}

func (s *RechargeTask) doFetchRate() {

	sourceRst := strings.Split(rate_tokens, ",")
	optCfg := optionCurrencyRateOpt.GetOpts()
	if optCfg.Opts == nil {
		optCfg.Opts = map[string]optionCurrencyRateOpt.OptionRate{}
	}
	for _, item := range sourceRst {
		fsym := strings.ToUpper(item)
		tsyms := strings.ToUpper(currencies)
		rst, err := s.doCryptocompareRates(&CryptoCompareRateReq{
			Fsym:  fsym,
			Tsyms: tsyms,
		})
		if err != nil {
			log.Errorf("doFetchRate:%v", err)
			continue
		}
		if _, ok := optCfg.Opts[fsym]; !ok {
			optCfg.Opts[fsym] = optionCurrencyRateOpt.OptionRate{
				Prices: make(map[string]decimal.Decimal),
			}
		}
		for key, rstItem := range rst {
			optCfg.Opts[fsym].Prices[strings.ToUpper(key)] = rstItem
		}
	}
	for k, v := range optCfg.Opts {
		log.Warnf("%s = %s", k, xconv.Json(v))
	}
	if len(optCfg.Opts) > 0 {
		optionCurrencyRateOpt.SetOpts(context.TODO(), optCfg)
	}
}

func (s *RechargeTask) doCryptocompareRates(args *CryptoCompareRateReq) (map[string]decimal.Decimal, error) {

	reply := make(map[string]decimal.Decimal)
	err := newCryptoCompareRatesClient().get("/data/price", args, &reply)
	if err != nil {
		return nil, err
	}
	return reply, err
}
