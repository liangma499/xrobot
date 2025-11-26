package tgSet

import (
	"xbase/errors"
	"xbase/log"
	"xrobot/internal/code"
	"xrobot/internal/xtelegram/telegram/telegram"
	"xrobot/internal/xtelegram/telegram/types"
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
