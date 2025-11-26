package tgwebhook

import (
	"tron_robot/internal/code"
	optionBaseConfigCfg "tron_robot/internal/option/option-base-config"
	"tron_robot/internal/xtypes"
	"xbase/errors"
	"xbase/log"
)

func SetWebHook(botToken, channelCode string) error {

	if botToken == "" {
		return nil
	}
	callBack := optionBaseConfigCfg.GetValue(xtypes.TelegramWebhookUrlKey)
	if callBack == "" {
		return errors.NewError(code.OptionNotFound)
	}

	client := NewClient(botToken)

	data := make(map[string]string)

	data["url"] = callBack
	data["secret_token"] = channelCode
	var res any
	err := client.Get("/setWebhook", data, res)
	//6867997452:AAFYZXHAC_TDvcfBiYto2ShutRSiUcboa04
	log.Warnf("channelCode:%v CallBack:%v, %#v:%v", channelCode, callBack, res, err)
	if err != nil {
		return errors.NewError(code.OptionNotFound, err)
	}
	return nil

}
