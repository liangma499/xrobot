package receivable

import (
	"context"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
	"xbase/cluster/node"
	"xbase/log"
	"xbase/task"
	"xbase/utils/xconv"
	"xbase/utils/xrand"
	"xbase/utils/xtime"
	"xrobot/internal/cryptocurrencies/tron/transfer"
	optionListenerAddressDao "xrobot/internal/dao/option-listener-address"
	paymentcryptotransaction "xrobot/internal/dao/payment-crypto-transaction"
	paymentnetwork "xrobot/internal/dao/payment-network"
	"xrobot/internal/model"
	optionCurrencyNetworkCfg "xrobot/internal/option/option-currency-network"
	"xrobot/internal/xtypes"

	"gorm.io/gorm/clause"
)

type TronTask struct {
	bInitTimer      bool
	cycleTime       time.Duration
	cycleTimeVerify time.Duration
	ctx             context.Context
	netWorkInfo     *model.NetworkInfo
	proxy           *node.Proxy
	runRunBolckTask task.Pool
}

var (
	trc20Task *TronTask
	once      sync.Once
)

const (
	cycleTime                  = 2    //订单任务间隔时间
	OrderTaskRunCount          = 30   // 一个订单最多查询次数
	MaxTaskUserInfo            = 1024 //一次处理多少个订单
	netWorkChannelType         = xtypes.NetWorkChannelType_TRON
	apiKind                    = xtypes.APITrongrid
	runRunBolckTaskSize        = 24 //最多少线程
	runRunBolckTaskNonblocking = false
)

func Instance(proxy *node.Proxy) *TronTask {

	once.Do(func() {
		trc20Task = &TronTask{
			cycleTime:       cycleTime,
			cycleTimeVerify: 1,
			ctx:             context.Background(),
			bInitTimer:      false,
			proxy:           proxy,
			runRunBolckTask: task.NewPool(task.WithSize(runRunBolckTaskSize), task.WithNonblocking(runRunBolckTaskNonblocking)),
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
	s.initNetWork()
	for {
		s.timeAfter()
		s.fetchTransaction()
		log.Warnf("拉取交易")
	}
}

func (s *TronTask) initNetWork() {
	netWorkInfo, err := paymentnetwork.Instance().FindOne(s.ctx, func(cols *paymentnetwork.Columns) any {
		return clause.Eq{
			Column: cols.NetWork,
			Value:  xtypes.TRON,
		}
	})
	if netWorkInfo == nil {
		s.netWorkInfo = &model.NetworkInfo{
			Block: -1,
		}
		now := xtime.Now()
		paymentnetwork.Instance().Insert(s.ctx, &model.PaymentNetwork{
			NetWork:     xtypes.TRON, // 主键
			NetworkInfo: s.netWorkInfo.Clone(),
			CreateAt:    now,
			UpdateAt:    now,
		})
	} else {
		if netWorkInfo.NetworkInfo == nil {
			s.netWorkInfo = &model.NetworkInfo{
				Block: -1,
			}
			s.upDateNetworkInfo()
		} else {
			s.netWorkInfo = netWorkInfo.NetworkInfo.Clone()
		}
	}
	if err != nil {
		panic(err)
	}
}
func (s *TronTask) upDateNetworkInfo() {
	_, err := paymentnetwork.Instance().Update(s.ctx, func(cols *paymentnetwork.Columns) any {
		return clause.Eq{
			Column: cols.NetWork,
			Value:  xtypes.TRON,
		}
	}, func(cols *paymentnetwork.Columns) any {
		return map[string]any{
			cols.NetworkInfo: s.netWorkInfo.Clone(),
			cols.UpdateAt:    xtime.Now(),
		}
	})
	if err != nil {
		log.Errorf("%v", err)
		return
	}
}
func (s *TronTask) timeAfter() {
	tmr := time.After(s.cycleTime * time.Second)
	<-tmr
}

// 拉取主线程
func (s *TronTask) fetchTransaction() {
	apiCfg, allApi := optionCurrencyNetworkCfg.Instance().GetApiByChannelType(netWorkChannelType, apiKind)
	if apiCfg == nil {
		return
	}

	//正常流程
	s.fetchTransactionNormal(apiCfg)

	//追的流程
	s.fetchTransactionAdded(allApi)

	s.upDateNetworkInfo()
}
func (s *TronTask) fetchTransactionNormal(args *xtypes.ApiCfg) {
	transferIn := transfer.NewTransfer(args.Url, args.AppID)

	block, err := transferIn.GetNowBlockNum()
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	block -= 1
	if s.netWorkInfo.Block <= 0 {
		s.netWorkInfo.Block = block
	}

	if block-s.netWorkInfo.Block > 10 {
		if s.netWorkInfo.BlockExtend == nil {
			s.netWorkInfo.BlockExtend = make(model.BlockExtendMap)
		}
		key := xtime.Now().Unix()
		s.netWorkInfo.BlockExtend[key] = &model.BlockExtend{
			Start: s.netWorkInfo.Block,
			End:   block, //最后一个不用拉
			Now:   s.netWorkInfo.Block,
			Key:   key,
		}
		s.netWorkInfo.Block = block
	}
	//取上一个块
	if s.netWorkInfo.Block+1 > block {
		return
	}
	runBlock := s.netWorkInfo.Block
	if err := s.addTask(transferIn, runBlock); err != nil {
		log.Warnf("TRC20:%v", err)
		return
	}

	s.netWorkInfo.Block++

}

func (s *TronTask) fetchTransactionAdded(allApi xtypes.KeyToApiCfgInfo) {
	if s.netWorkInfo.BlockExtend == nil {
		return
	}
	if len(s.netWorkInfo.BlockExtend) == 0 {
		return
	}
	if allApi == nil {
		return
	}
	var addBlock *model.BlockExtend
	deleteIds := make([]int64, 0)
	for k, v := range s.netWorkInfo.BlockExtend {
		if v.Now+1 == v.End || v.Start == 0 {
			deleteIds = append(deleteIds, k)
		} else if addBlock == nil {
			addBlock = v.Clone()
		}
	}
	for _, v := range deleteIds {
		delete(s.netWorkInfo.BlockExtend, v)
	}
	if addBlock == nil {
		return
	}
	//这里会利用正常的拿取API 如果 不复用配一个的情况就不会追加
	bDel := false
	for _, item := range allApi {
		if addBlock.Now+1 >= addBlock.End {
			delete(s.netWorkInfo.BlockExtend, addBlock.Key)
			bDel = true
			break
		}
		transferIn := transfer.NewTransfer(item.Url, item.AppID)
		runBlock := addBlock.Now

		if err := s.addTask(transferIn, runBlock); err != nil {
			log.Warnf("TRC20:%v", err)
			break
		}
		addBlock.Now++

	}
	if !bDel {
		s.netWorkInfo.BlockExtend[addBlock.Key].Now = addBlock.Now
	}

}

func (s *TronTask) doGetBlockByNum(transferIn *transfer.TransferInfo, block int64) {
	for i := 0; i < 3; i++ {
		if transferIn == nil {
			return
		}
		blockInfo, err := transferIn.GetBlockByNum(block)
		if err != nil {
			s.randSleep()
			continue
		}
		if blockInfo == nil {
			return
		}
		now := xtime.Now()
		rst := make([]*model.PaymentCryptoTransaction, 0)
		for _, item := range blockInfo.Transactions {
			if item.Transaction == nil {
				continue
			}
			if item.Transaction.RawData == nil {
				continue
			}
			if item.Transaction.RawData.Contract == nil {
				continue
			}
			for _, contractItem := range item.Transaction.RawData.Contract {
				if contractItem == nil {
					continue
				}
				trf := transferIn.UnmarshalTo(contractItem)
				if trf == nil {
					continue
				}
				transactionHash := hex.EncodeToString(item.Txid)
				rstCheck := s.doGetBlockCheckAddress(trf.Currency, trf.ToAddress, trf.FromAddress, transactionHash)
				if rstCheck == nil {
					continue
				}
				//将数据入库
				rst = append(rst, &model.PaymentCryptoTransaction{
					NetWork:         xtypes.TRON,
					ChannelCode:     rstCheck.channelCode,
					AddressKind:     rstCheck.addressKind,
					TransactionHash: transactionHash,
					BlockHash:       hex.EncodeToString(blockInfo.Blockid),
					BlockNum:        blockInfo.BlockHeader.RawData.Number,
					Protocol:        trf.Protocol,
					ToAddress:       trf.ToAddress,
					FromAddress:     trf.FromAddress,
					Amount:          xtypes.CoefficientToFloat64(trf.RealAmount, xtypes.Trc20_Places),
					RealAmount:      trf.RealAmount.Copy(),
					Contract:        trf.Contract,
					Currency:        trf.Currency,
					EnergyFee:       0,
					NetUsage:        0,
					EnergyUsage:     0,
					VerifyCount:     0,
					Stauts:          xtypes.Transaction_Verified,
					TransactionKind: rstCheck.transactionKind,
					CreateAt:        now,
					UpdateAt:        now,
				})
			}

		}

		if err := s.insertIntoTrx(rst); err != nil {
			log.Errorf("err:%v trxer:%s", err, xconv.Json(rst))
			s.randSleep()
			continue
		} else {
			break
		}
	}

}
func (s *TronTask) doGetBlockCheckAddress(currency xtypes.Currency, toAddress, fromAddress string, transactionHash string) *addressCheckRes {

	//收款交易只监听收款
	addressInfo := optionListenerAddressDao.Instance().CheckListenerAddress(xtypes.TRON, toAddress)
	if addressInfo != nil {

		rst := &addressCheckRes{
			transactionKind: xtypes.Transaction_Recharge,
			channelCode:     addressInfo.ChannelCode,
			addressKind:     addressInfo.AddressKind,
		}
		if currency == xtypes.ENERGY {
			rst.transactionKind = xtypes.Transaction_EnergyIn
		}
		return rst
	}
	//监听是否已经转账
	if currency == xtypes.USDT && optionListenerAddressDao.Instance().CheckListenerTransactionAddress(xtypes.TRON, fromAddress) {
		rst := &addressCheckRes{
			transactionKind: xtypes.Transaction_OtherTransaction,
			channelCode:     "",
			addressKind:     xtypes.AddressKind_OtherTransaction,
		}
		if currency == xtypes.ENERGY {
			rst.transactionKind = xtypes.Transaction_EnergyIn
		}
		return rst
	}

	//监听交易hash 验证订单
	if currency == xtypes.USDT && optionListenerAddressDao.Instance().CheckListenerTrxID(xtypes.TRON, transactionHash) {
		rst := &addressCheckRes{
			transactionKind: xtypes.Transaction_OutVerify,
			channelCode:     "",
			addressKind:     xtypes.AddressKind_OutVerify,
		}
		if currency == xtypes.ENERGY {
			rst.transactionKind = xtypes.Transaction_EnergyIn
		}
		return rst
	}
	return nil
}
func (s *TronTask) insertIntoTrx(trx []*model.PaymentCryptoTransaction) error {
	if trx == nil {
		return nil
	}
	lenTrx := len(trx)
	if lenTrx <= 0 {
		return nil
	}
	if lenTrx == 1 {
		return paymentcryptotransaction.Instance().Table.WithContext(s.ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(trx[0]).Error

	} else {
		return paymentcryptotransaction.Instance().Table.WithContext(s.ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(trx).Error
	}

}
func (s *TronTask) randSleep() {
	randTime := xrand.Int64(1, 3)
	tmr := time.After(time.Duration(randTime) * time.Second)
	<-tmr
}

func (s *TronTask) addTask(transferIn *transfer.TransferInfo, runBlock int64) error {
	waitingCount := s.runRunBolckTask.Waiting()
	if !runRunBolckTaskNonblocking && waitingCount >= runRunBolckTaskSize {
		return fmt.Errorf("max run task")
	}
	return s.runRunBolckTask.AddTask(func() {
		s.doGetBlockByNum(transferIn.Clone(), runBlock)
	})
}
