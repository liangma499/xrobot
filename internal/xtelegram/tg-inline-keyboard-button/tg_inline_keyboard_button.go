package tginlinekeyboardbutton

import (
	"context"
	"net/url"
	"sort"
	optionTelegramInlineKeyboardButtonDao "tron_robot/internal/dao/option-telegram-inline-keyboard-button"
	"tron_robot/internal/model"
	"tron_robot/internal/xtelegram/telegram/telegram"
	"tron_robot/internal/xtelegram/telegram/types"
	tgtypes "tron_robot/internal/xtelegram/tg-types"
	"tron_robot/internal/xtypes"
	"xbase/log"

	"gorm.io/gorm/clause"
)

type XTelegramButton map[int][]*model.OptionTelegramInlineKeyboardButton

func (xb XTelegramButton) Keyboard() *types.InlineKeyboardMarkup {

	botApi, err := telegram.New("test")
	if err != nil {
		log.Warnf("%v", err)
		return nil
	}
	if len(xb) == 0 {
		return nil
	}
	sortLine := make([]int, 0)
	for k := range xb {
		sortLine = append(sortLine, k)
	}
	//排序
	if len(sortLine) > 1 {
		sort.Ints(sortLine)
	}
	//同排按钮排序
	for _, item := range xb {
		if len(item) > 1 {
			sort.SliceStable(item, func(i, j int) bool {
				return item[i].Sort < item[j].Sort
			})
		}

	}
	length := len(sortLine)
	inlineKeyboard := &types.InlineKeyboardMarkup{
		InlineKeyboard: make([][]types.InlineKeyboardButton, 0),
	}

	for i := 0; i < length; i++ {
		index := sortLine[i]
		if item, ok := xb[index]; ok {
			lenLine := len(item)
			lineButton := make([]types.InlineKeyboardButton, 0)
			for j := 0; j < lenLine; j++ {
				buttonInfo := item[j]
				if buttonInfo == nil {
					continue
				}
				switch buttonInfo.Kind {
				case tgtypes.ButtonType_InlineKeyboardButtonURL: //ButtonKind = "InlineKeyboardButtonURL"      //跳转网页
					{
						lineButton = append(lineButton, botApi.NewInlineKeyboardButtonURL(buttonInfo.Name.String(), buttonInfo.URL))
					}
				case tgtypes.ButtonKind_InlineKeyboardWebApp: //ButtonKind = "InlineKeyboardWebApp"         //按钮
					{
						lineButton = append(lineButton, botApi.NewInlineKeyboardWebApp(buttonInfo.Name.String(), buttonInfo.URL))
					}
				case tgtypes.ButtonKind_InlineKeyboardButtonLoginURL: //ButtonKind = "InlineKeyboardButtonLoginURL" //按钮
					{
						lineButton = append(lineButton, botApi.NewInlineKeyboardButtonLoginURL(buttonInfo.Name.String(), types.LoginURL{
							URL:                buttonInfo.LoginURL.URL,                // 按下按钮时打开的 HTTP URL，用户授权数据将添加到查询字符串中。如果用户拒绝提供授权数据，则将打开不带用户信息的原始 URL。添加的数据与接收授权数据中描述的数据相同。注意：您必须始终检查接收数据的哈希，以验证身份验证和数据的完整性，如检查授权中所述。
							ForwardText:        buttonInfo.LoginURL.ForwardText,        // 可选。转发消息中按钮的新文本。
							BotUsername:        buttonInfo.LoginURL.BotUsername,        // 可选。将用于用户授权的机器人的用户名。有关详细信息，请参见设置机器人。如果未指定，将假定当前机器人的用户名。hurl 的域名必须与链接到机器人的域名相同。有关详细信息，请参见将您的域名链接到机器人。
							RequestWriteAccess: buttonInfo.LoginURL.RequestWriteAccess, // 可选。传递 true 以请求机器人向用户发送消息的权限。
						}))
					}
				case tgtypes.ButtonKind_InlineKeyboardButtonSwitch: //ButtonKind = "InlineKeyboardButtonSwitch"
					{
						lineButton = append(lineButton, botApi.NewInlineKeyboardButtonSwitch(buttonInfo.Name.String(), buttonInfo.SwitchInlineQuery))
					}
				case tgtypes.ButtonKind_InlineKeyboardButtonSwitchCurrentChat: //ButtonKind = "InlineKeyboardButtonSwitchCurrentChat"
					{
						lineButton = append(lineButton, botApi.NewInlineKeyboardButtonSwitchCurrentChat(buttonInfo.Name.String(), buttonInfo.SwitchInlineQueryCurrentChat))
					}
				case tgtypes.ButtonKind_InlineKeyboardButtonSwitchChosenChat: // ButtonKind = "InlineKeyboardButtonSwitchChosenChat"
					{
						lineButton = append(lineButton, botApi.NewInlineKeyboardButtonSwitchChosenChat(buttonInfo.Name.String(), &types.SwitchInlineQueryChosenChat{
							Query:             buttonInfo.SwitchInlineQueryChosenChat.Query,             // 可选。要插入输入字段中的默认内联查询。如果留空，则仅插入机器人的用户名
							AllowUserChats:    buttonInfo.SwitchInlineQueryChosenChat.AllowUserChats,    // 可选。如果可以选择与用户的私聊，则为 true
							AllowBotChats:     buttonInfo.SwitchInlineQueryChosenChat.AllowBotChats,     // 可选。如果可以选择与机器人的私聊，则为 true
							AllowGroupChats:   buttonInfo.SwitchInlineQueryChosenChat.AllowGroupChats,   // 可选。如果可以选择群组和超级群组聊天，则为 true
							AllowChannelChats: buttonInfo.SwitchInlineQueryChosenChat.AllowChannelChats, // 可选。如果可以选择频道聊天，则为 true
						}))
					}

				default:
					{
						lineButton = append(lineButton, botApi.NewInlineKeyboardCallbackData(buttonInfo.Name.String(), buttonInfo.CallbackData))
					}
				}

			}
			if len(lineButton) > 0 {
				inlineKeyboard.InlineKeyboard = append(inlineKeyboard.InlineKeyboard, lineButton)
			}

		}

	}

	return inlineKeyboard
}
func InlineKeyboardMarkupWithToken(xb *types.InlineKeyboardMarkup, token string) *types.InlineKeyboardMarkup {

	if xb == nil {
		return nil
	}
	if token == "" {
		return xb
	}
	for _, itemTop := range xb.InlineKeyboard {
		for _, item := range itemTop {
			if item.IsXtelegramToken {
				buttonUrl := ""
				if item.WebApp != nil {
					buttonUrl = item.WebApp.URL
				} else if item.URL != "" {
					buttonUrl = item.URL
				}
				if buttonUrl == "" {
					continue
				}

				if buttonUrl != "" {
					u, err := url.Parse(buttonUrl)
					if err == nil {
						if u.RawQuery == "" {
							buttonUrl += "?token=" + token
						} else {
							buttonUrl += "&token=" + token
						}
					}

					if item.WebApp != nil {
						item.WebApp.URL = buttonUrl
					} else if item.URL != "" {
						item.URL = buttonUrl
					}
					log.Warnf("keyboard:%v", buttonUrl)
				}

			}
		}
	}
	return xb
}

func TelegramInlineKeyBoardbutton(ctx context.Context, cmd *model.OptionTelegramCmd) *types.InlineKeyboardMarkup {
	if cmd == nil {
		return nil
	}
	if cmd.CmdKind != tgtypes.CmdKind_InlineKeyboard {
		return nil
	}
	if buttons, err := optionTelegramInlineKeyboardButtonDao.Instance().FindMany(ctx, func(cols *optionTelegramInlineKeyboardButtonDao.Columns) any {
		return clause.And(
			clause.Eq{Column: cols.ChannelCode, Value: cmd.ChannelCode},
			clause.Eq{Column: cols.Cmd, Value: cmd.Cmd.String()},
			clause.Eq{Column: cols.Status, Value: xtypes.OptionStatus_Normal},
		)
	}, nil, nil); err == nil {
		if buttons != nil {
			if len(buttons) == 0 {
				return nil
			}
			buttosMap := make(XTelegramButton)
			for _, itemButton := range buttons {
				if _, ok := buttosMap[itemButton.LineID]; !ok {
					buttosMap[itemButton.LineID] = make([]*model.OptionTelegramInlineKeyboardButton, 0)
				}
				buttosMap[itemButton.LineID] = append(buttosMap[itemButton.LineID], itemButton)
			}

			return buttosMap.Keyboard()
		}
	} else {
		log.Errorf("%v", err)
	}
	return nil
}

/*
func doReplyKeyboardMarkupWithToken(xb *types.InlineKeyboardMarkup, token string) *types.ReplyKeyboardMarkup {

	if xb == nil {
		return nil
	}
	replyKeyboardMarkup := &types.ReplyKeyboardMarkup{
		Keyboard: xb,
	}
	return replyKeyboardMarkup
	if token == "" {
		return xb
	}
	for _, itemTop := range xb.InlineKeyboard {
		for _, item := range itemTop {
			if item.IsXtelegramToken {
				buttonUrl := ""
				if item.WebApp != nil {
					buttonUrl = item.WebApp.URL
				} else if item.URL != "" {
					buttonUrl = item.URL
				}
				if buttonUrl == "" {
					continue
				}

				if buttonUrl != "" {
					u, err := url.Parse(buttonUrl)
					if err == nil {
						if u.RawQuery == "" {
							buttonUrl += "?token=" + token
						} else {
							buttonUrl += "&token=" + token
						}
					}

					if item.WebApp != nil {
						item.WebApp.URL = buttonUrl
					} else if item.URL != "" {
						item.URL = buttonUrl
					}
					log.Warnf("keyboard:%v", buttonUrl)
				}

			}
		}
	}
	return xb
}
*/
