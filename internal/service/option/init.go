package option

import (
	"context"
	optionChannelDao "tron_robot/internal/dao/option-channel"
	tgWebhook "tron_robot/internal/xtelegram/tg-webhook"
	"tron_robot/internal/xtypes"
	"xbase/log"
)

// 开卡
func (s *Server) initWebhook() {
	opts, err := optionChannelDao.Instance().GetChannel(context.Background(), xtypes.OfficialChannelCode)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	if opts == nil {
		return
	}
	if opts.TelegramCfg == nil {
		return
	}

	if err := tgWebhook.SetWebHook(opts.TelegramCfg.MainRobotToken, xtypes.OfficialChannelCode); err != nil {
		log.Errorf("%v", err)
	}

}
