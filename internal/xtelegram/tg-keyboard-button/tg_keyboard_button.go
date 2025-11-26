package tgkeyboardbutton

import (
	"context"
	"sort"
	optionTelegramKeyboardMarkupButtonDao "tron_robot/internal/dao/option-telegram-keyboard-markup-button"
	"tron_robot/internal/model"
	"tron_robot/internal/xtelegram/telegram/types"
	tgtypes "tron_robot/internal/xtelegram/tg-types"
	"tron_robot/internal/xtypes"
	"xbase/log"

	"gorm.io/gorm/clause"
)

type XTelegramKeyboardButton map[int][]*model.OptionTelegramKeyboardMarkupButton

func (xb XTelegramKeyboardButton) Keyboard() [][]types.KeyboardButton {

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
	keyboard := make([][]types.KeyboardButton, 0)

	for i := 0; i < length; i++ {
		index := sortLine[i]
		if item, ok := xb[index]; ok {
			lenLine := len(item)
			lineButton := make([]types.KeyboardButton, 0)
			for j := 0; j < lenLine; j++ {
				buttonInfo := item[j]
				if buttonInfo == nil {
					continue
				}
				button := types.KeyboardButton{
					Text:            buttonInfo.Text,                 // 按钮的文本。如果没有使用任何可选字段，则在按钮被按下时以消息的形式发送
					RequestUsers:    buttonInfo.RequestUsers.Clone(), // 可选。如果指定，按下按钮时将打开合适用户的列表。点击任何用户将以“users_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
					RequestChat:     buttonInfo.RequestChat.Clone(),  // 可选。如果指定，按下按钮时将打开合适聊天的列表。点击一个聊天将以“chat_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
					RequestContact:  buttonInfo.RequestContact,       // 可选。如果为 true，用户的电话号码将在按钮被按下时作为联系人发送。仅在私聊中可用。
					RequestLocation: buttonInfo.RequestLocation,      // 可选。如果为 true，用户的当前位置将在按钮被按下时发送。仅在私聊中可用。
					RequestPoll:     buttonInfo.RequestPoll.Clone(),  // 可选。如果指定，用户将被要求创建一个投票并在按钮被按下时发送给机器人。仅在私聊中可用。
					WebApp:          buttonInfo.WebApp.Clone(),       // 可选。如果指定，描述的 Web 应用将在按钮被按下时启动。Web 应用能够发送“web_app_data”服务消息。仅在私聊中可用。
				}

				lineButton = append(lineButton, button)
			}
			if len(lineButton) > 0 {
				keyboard = append(keyboard, lineButton)
			}

		}

	}

	return keyboard
}

func TelegramKeyBoardbutton(ctx context.Context, cmd *model.OptionTelegramCmd) *types.ReplyKeyboardMarkup {
	if cmd == nil {
		return nil
	}
	if cmd.CmdKind != tgtypes.CmdKind_KeyboardMarkup {
		return nil
	}
	if buttons, err := optionTelegramKeyboardMarkupButtonDao.Instance().FindMany(ctx, func(cols *optionTelegramKeyboardMarkupButtonDao.Columns) any {
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
			buttosMap := make(XTelegramKeyboardButton)
			for _, itemButton := range buttons {
				if _, ok := buttosMap[itemButton.LineID]; !ok {
					buttosMap[itemButton.LineID] = make([]*model.OptionTelegramKeyboardMarkupButton, 0)
				}
				buttosMap[itemButton.LineID] = append(buttosMap[itemButton.LineID], itemButton)
			}
			rst := &types.ReplyKeyboardMarkup{
				Keyboard: buttosMap.Keyboard(),
			}
			if !cmd.ReplyKeyboardMarkup.IsNull() {
				rst.IsPersistent = cmd.ReplyKeyboardMarkup.IsPersistent
				rst.ResizeKeyboard = cmd.ReplyKeyboardMarkup.ResizeKeyboard
				rst.OneTimeKeyboard = cmd.ReplyKeyboardMarkup.OneTimeKeyboard
				rst.InputFieldPlaceholder = cmd.ReplyKeyboardMarkup.InputFieldPlaceholder
				rst.Selective = cmd.ReplyKeyboardMarkup.Selective
			}
			return rst
		}
	} else {
		log.Errorf("%v", err)
	}
	return nil
}
