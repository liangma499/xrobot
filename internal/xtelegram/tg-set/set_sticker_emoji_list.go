package tgSet

import (
	"tron_robot/internal/code"
	"tron_robot/internal/xtelegram/telegram/telegram"
	"tron_robot/internal/xtelegram/telegram/types"
	"xbase/errors"
	"xbase/log"
)

func SetStickerEmojiList(botToken string) error {

	if botToken == "" {
		return nil
	}

	botApi, err := telegram.New(botToken)
	if err != nil {
		log.Errorf("%v", err)
		return nil
	}

	/*rst, err := botApi.SetStickerKeywords(&types.SetStickerKeywords{
		Sticker:  "Abc",
		Keywords: []string{"💹TRX闪兑", "💹TRX闪兑01"},
	})
	*/
	rst, err := botApi.SetStickerEmojiList(&types.SetStickerEmojiList{
		Sticker:   "rand55566",
		EmojiList: []string{"🏧bbc", "💹"},
	})

	if err != nil {
		log.Errorf("%v", err)
		return errors.NewError(code.OptionNotFound, err)
	}
	log.Warnf("botToken,botToken:%v %#v", botToken, rst)

	return nil

}
