package option

import (
	"context"
	"xbase/log"
	optionChannelDao "xrobot/internal/dao/option-channel"
	tgWebhook "xrobot/internal/xtelegram/tg-webhook"
	"xrobot/internal/xtypes"
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
