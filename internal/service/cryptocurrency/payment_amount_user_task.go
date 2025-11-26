package cryptocurrency

import (
	"context"
	"time"
	optionChannelDao "tron_robot/internal/dao/option-channel"
	paymentAmountUserDao "tron_robot/internal/dao/payment-amount-user"
	tgmsg "tron_robot/internal/xtelegram/tg-msg"
	"xbase/task"
	"xbase/utils/xconv"
	"xbase/utils/xtime"

	"github.com/smallnest/rpcx/log"
	"gorm.io/gorm/clause"
)

func (s *Server) doBalanceTimer() {
	task.AddTask(func() {
		for {
			s.doPaymentAmountIserTask()
			<-time.After(30 * time.Second)
		}
	})
}

func (s *Server) doPaymentAmountIserTask() {
	s.mu.Lock()
	defer s.mu.Unlock()
	ctx := context.Background()

	userInfo, err := paymentAmountUserDao.Instance().FindMany(ctx, func(cols *paymentAmountUserDao.Columns) any {
		return clause.Lt{
			Column: cols.ExpirationTime,
			Value:  xtime.Now().Unix(),
		}
	}, nil, nil, 100)
	if err != nil {
		return
	}
	if userInfo == nil {
		return
	}
	if len(userInfo) == 0 {
		return
	}
	for _, item := range userInfo {
		err := paymentAmountUserDao.Instance().ClearAmount(item.Currency, item.Amount, "")
		if err != nil {
			log.Errorf("err:%v data:%s", err, xconv.Json(item))
			continue
		}
		if item.TelegramUid > 0 {
			channelCfg, err := optionChannelDao.Instance().GetChannel(ctx, item.ChannelCode)
			if err != nil {
				log.Errorf("%v", err)
				continue
			}
			if channelCfg == nil {
				log.Errorf("channelCfg is nil")
				continue
			}
			xMsg, err := tgmsg.NewXTelegramMessage(channelCfg.TelegramCfg.MainRobotToken)
			if err != nil {
				continue
			}
			xMsg.DeleteMessage(item.TelegramUid, item.Extend.MessageID)
		}
	}
}
