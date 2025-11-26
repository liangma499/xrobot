package tgbutton

import (
	"xrobot/internal/xtelegram/telegram/types"
	tginlinekeyboardbutton "xrobot/internal/xtelegram/tg-inline-keyboard-button"
	tgtypes "xrobot/internal/xtelegram/tg-types"
)

type TelegramButton struct {
	CmdKind        tgtypes.CmdKind             `json:"cmd_kind,omitempty"` //推送方式
	InlineKeyboard *types.InlineKeyboardMarkup `json:"inline_keyboard,omitempty"`
	Keyboard       *types.ReplyKeyboardMarkup  `json:"keyboard,omitempty"`
}

func NewTelegramButton(cmdKind tgtypes.CmdKind, keyboard any) *TelegramButton {

	if keyboard == nil {
		return nil
	}
	switch cmdKind {
	case tgtypes.CmdKind_InlineKeyboard:
		{
			inlineKeyboard, ok := keyboard.(*types.InlineKeyboardMarkup)
			if !ok {
				return nil
			}
			return &TelegramButton{
				CmdKind:        cmdKind,
				InlineKeyboard: inlineKeyboard,
			}
		}
	case tgtypes.CmdKind_KeyboardMarkup:
		{
			inlineKeyboard, ok := keyboard.(*types.ReplyKeyboardMarkup)
			if !ok {
				return nil
			}
			return &TelegramButton{
				CmdKind:  cmdKind,
				Keyboard: inlineKeyboard,
			}
		}
	}
	return nil
}

func (otc *TelegramButton) Clone() *TelegramButton {
	if otc == nil {
		return nil
	}
	switch otc.CmdKind {
	case tgtypes.CmdKind_InlineKeyboard:
		{

			rst := &TelegramButton{
				CmdKind:        otc.CmdKind,
				InlineKeyboard: nil,
				Keyboard:       nil,
			}
			if otc.InlineKeyboard != nil {
				rst.InlineKeyboard = &types.InlineKeyboardMarkup{
					InlineKeyboard: make([][]types.InlineKeyboardButton, 0, len(otc.InlineKeyboard.InlineKeyboard)),
				}
				if len(otc.InlineKeyboard.InlineKeyboard) > 0 {

					//copy(rst.InlineKeyboard.InlineKeyboard, otc.InlineKeyboard.InlineKeyboard)

					for _, row := range otc.InlineKeyboard.InlineKeyboard {
						newRow := make([]types.InlineKeyboardButton, len(row))
						copy(newRow, row)
						rst.InlineKeyboard.InlineKeyboard = append(rst.InlineKeyboard.InlineKeyboard, newRow)
					}

				}
			}

			return rst
		}
	case tgtypes.CmdKind_KeyboardMarkup:
		{
			rst := &TelegramButton{
				CmdKind:        otc.CmdKind,
				InlineKeyboard: nil,
				Keyboard:       nil,
			}
			if otc.Keyboard != nil {

				rst.Keyboard = &types.ReplyKeyboardMarkup{
					Keyboard:              make([][]types.KeyboardButton, 0, len(otc.Keyboard.Keyboard)), // 按钮行的数组，每行由 KeyboardButton 对象数组表示,
					IsPersistent:          otc.Keyboard.IsPersistent,                                     // 可选。请求客户端在常规键盘隐藏时始终显示此键盘。默认值为 false，此时自定义键盘可以被隐藏，并通过键盘图标打开。
					ResizeKeyboard:        otc.Keyboard.ResizeKeyboard,                                   // 可选。请求客户端根据最佳适配垂直调整键盘大小（例如，如果只有两行按钮，则使键盘变小）。默认值为 false，此时自定义键盘的高度始终与应用程序的标准键盘相同。
					OneTimeKeyboard:       otc.Keyboard.OneTimeKeyboard,                                  // 可选。请求客户端在使用此键盘后隐藏键盘。键盘仍然可用，但客户端将自动在聊天中显示常规字母键盘 - 用户可以在输入字段中按特殊按钮再次查看自定义键盘。默认值为 false。
					InputFieldPlaceholder: otc.Keyboard.InputFieldPlaceholder,                            // 可选。当键盘处于活动状态时，在输入字段中显示的占位符；1-64 个字符
					Selective:             otc.Keyboard.Selective,                                        // 可选。如果只想对特定用户显示此键盘，请使用此参数。目标：1）在消息对象的文本中提到的用户；2）如果机器人的消息是回复（具有 reply_to_message_id），则是原始消息的发送者。示例：用户请求更改机器人的语言，机器人回复请求并显示选择新语言的键盘。其他用户在群组中看不到此键盘。
				}

				for _, row := range rst.Keyboard.Keyboard {
					newRow := make([]types.KeyboardButton, len(row))
					copy(newRow, row)
					rst.Keyboard.Keyboard = append(rst.Keyboard.Keyboard, newRow)
				}

				/*if len(otc.Keyboard.Keyboard) > 0 {
					copy(rst.Keyboard.Keyboard, otc.Keyboard.Keyboard)
				}*/

			}
			return rst
		}
	}
	return nil
}
func (otc *TelegramButton) Button(token string) any {
	if otc == nil {
		return nil
	}

	switch otc.CmdKind {
	case tgtypes.CmdKind_InlineKeyboard:
		{
			return tginlinekeyboardbutton.InlineKeyboardMarkupWithToken(otc.InlineKeyboard, token)
		}
	case tgtypes.CmdKind_KeyboardMarkup:
		{
			return otc.Keyboard
		}

	}
	return nil
}
