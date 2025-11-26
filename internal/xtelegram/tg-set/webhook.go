package tgSet

import (
	"xbase/errors"
	"xbase/log"
	"xrobot/internal/code"
	optionBaseConfigCfg "xrobot/internal/option/option-base-config"
	"xrobot/internal/xtypes"
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
	//SetChatMenuButton(botToken)
	SetMyCommands(botToken)
	SetStickerEmojiList(botToken)
	return nil

}
