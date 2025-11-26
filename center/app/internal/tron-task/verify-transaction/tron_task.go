package verifytransaction

import (
	"context"
	"sync"
	"time"
	tronscanapi "tron_robot/internal/cryptocurrencies/tron/tronscan-api"
	paymentCryptoTransactionDao "tron_robot/internal/dao/payment-crypto-transaction"
	"tron_robot/internal/event/cryptocurrencyevent"
	"tron_robot/internal/model"
	optionCurrencyNetworkCfg "tron_robot/internal/option/option-currency-network"
	"tron_robot/internal/xtypes"
	"xbase/cluster/node"
	"xbase/log"
	"xbase/task"
	"xbase/utils/xtime"

	"gorm.io/gorm/clause"
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
	cycleTime           = 2    //订单任务间隔时间
	OrderTaskRunCount   = 30   // 一个订单最多查询次数
	MaxTaskUserInfo     = 1024 //一次处理多少个订单
	netWorkChannelType  = xtypes.NetWorkChannelType_TRON
	apiKind             = xtypes.APITrongrid
	runRunBolckTaskSize = 24 //最多少线程
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
		s.initVerifyTransactionTask()
	})
}

func (s *TronTask) initVerifyTransactionTask() {
	time.Sleep(10 * time.Second)

	for {
		s.timeAfterVerify()
		s.verifyTransactionTask()
		log.Warnf("验证订单")
	}

}

func (s *TronTask) timeAfterVerify() {
	tmr := time.After(s.cycleTimeVerify * time.Second)
	<-tmr
}
func (s *TronTask) verifyTransactionTask() {
	s.mu.Lock()
	defer s.mu.Unlock()
	ctx := context.Background()
	trxs, err := paymentCryptoTransactionDao.Instance().FindMany(ctx, func(cols *paymentCryptoTransactionDao.Columns) any {
		return clause.And(clause.Eq{
			Column: cols.NetWork,
			Value:  xtypes.TRON,
		}, clause.Eq{
			Column: cols.Stauts,
			Value:  xtypes.Transaction_Verified,
		})
	}, nil, nil, xtypes.TransactionLimit, 0)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	if trxs == nil {
		return
	}
	updates := make([]*model.PaymentCryptoTransaction, 0)

	now := xtime.Now()

	for _, item := range trxs {
		if item == nil {
			continue
		}
		apiCfg := optionCurrencyNetworkCfg.Instance().GetNeedApiByChannelType(netWorkChannelType, xtypes.APITronscan)
		if apiCfg == nil {
			continue
		}
		resp, err := tronscanapi.GetTransactionInfo(apiCfg.Url, apiCfg.AppID, &tronscanapi.TransactionInfoReq{
			Hash: item.TransactionHash,
		})
		if err != nil {
			log.Errorf("%v", err)
			continue
		}
		if resp == nil {
			continue
		}
		//有可以是空结构
		if resp.Hash != item.TransactionHash {
			continue
		}
		if resp.ContractRet == "SUCCESS" && resp.Confirmed {
			item.Stauts = xtypes.Transaction_Confirmed
		} else {
			if item.VerifyCount >= 6 {
				item.Stauts = xtypes.Transaction_Fail
			}

		}
		item.VerifyCount += 1

		if resp.Cost != nil {
			item.EnergyFee = resp.Cost.EnergyFee
			item.NetUsage = resp.Cost.NetUsage
			item.EnergyUsage = resp.Cost.EnergyUsage
		}

		item.UpdateAt = now
		updates = append(updates, item)
	}
	if len(updates) > 0 {
		if err = paymentCryptoTransactionDao.Instance().Table.WithContext(ctx).Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: paymentCryptoTransactionDao.Instance().Columns.ID}},
			DoUpdates: clause.AssignmentColumns([]string{
				paymentCryptoTransactionDao.Instance().Columns.Stauts,
				paymentCryptoTransactionDao.Instance().Columns.VerifyCount,
				paymentCryptoTransactionDao.Instance().Columns.EnergyFee,
				paymentCryptoTransactionDao.Instance().Columns.NetUsage,
				paymentCryptoTransactionDao.Instance().Columns.EnergyUsage,
				paymentCryptoTransactionDao.Instance().Columns.UpdateAt}),
		}).Create(updates).Error; err == nil {
			//这里在别的平台就调用回调
			log.Warnf("更新成功")
			cryptocurrencyevent.PublishCryptoCurrency(&cryptocurrencyevent.CryptoCurrencyMsg{
				TimeUninx: xtime.Now().Unix(),
			})
		} else {
			log.Errorf("%v", err)
		}
	}
}
