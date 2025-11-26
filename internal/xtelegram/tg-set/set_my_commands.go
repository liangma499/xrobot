package tgSet

import (
	"tron_robot/internal/code"
	"tron_robot/internal/xtelegram/telegram/telegram"
	"tron_robot/internal/xtelegram/telegram/types"
	tgtypes "tron_robot/internal/xtelegram/tg-types"
	"xbase/errors"
	"xbase/log"
)

func SetMyCommands(botToken string) error {

	if botToken == "" {
		return nil
	}

	botApi, err := telegram.New(botToken)
	if err != nil {
		log.Errorf("%v", err)
		return nil
	}

	rst, err := botApi.SetMyCommands(&types.SetMyCommands{
		Scope: &types.BotCommandScope{
			Type: "all_private_chats",
		},

		Commands: []types.BotCommand{
			{
				Command:     tgtypes.XTelegramCmd_Start.String(),
				Description: tgtypes.XTelegramCmd_Start.Description(),
			},
		},
	})

	//6867997452:AAFYZXHAC_TDvcfBiYto2ShutRSiUcboa04
	if err != nil {
		log.Errorf("%v", err)
		return errors.NewError(code.OptionNotFound, err)
	}
	log.Warnf("botToken,botToken:%v %#v", botToken, rst)

	return nil

}
