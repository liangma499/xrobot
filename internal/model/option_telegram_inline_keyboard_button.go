package model

import (
	"fmt"
	"time"
	"tron_robot/internal/xtelegram/telegram/types"
	tgtypes "tron_robot/internal/xtelegram/tg-types"
	"tron_robot/internal/xtypes"
)

//go:generate xgorm-dao-generator -model-dir=. -model-names=OptionTelegramInlineKeyboardButton -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type OptionTelegramInlineKeyboardButton struct {
	ChannelCode                  string                            `gorm:"column:channel_code;size:32;primarykey;"`                                                                        //渠道code
	Cmd                          tgtypes.XTelegramCmd              `gorm:"column:cmd;size:64;primarykey"`                                                                                  //推送方式                                                                        // 主键
	Name                         tgtypes.XTelegramButton           `gorm:"column:name;size:32;not null;primarykey" json:"name"`                                                            // Label text on the button
	LineID                       int                               `gorm:"column:line_id;size:32;default:0" json:"line_id"`                                                                //按钮所在行
	Sort                         int                               `gorm:"column:sort;size:32;default:0" json:"sort"`                                                                      //按钮所在列
	Kind                         tgtypes.ButtonKind                `gorm:"column:kind;size:32;default:''"  json:"kind"`                                                                    //(1=跳转网页，2=跳转应用)
	URL                          string                            `gorm:"column:url;size:512;" json:"url,omitempty"`                                                                      // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
	CallbackData                 string                            `gorm:"column:callback_data;size:512;default:'';" json:"callback_data,omitempty"`                                       // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
	LoginURL                     types.LoginURL                    `gorm:"column:login_url;type:json" json:"login_url,omitempty"`                                                          // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
	SwitchInlineQuery            string                            `gorm:"column:switch_inline_query;size:512;default:'';" json:"switch_inline_query,omitempty"`                           // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
	SwitchInlineQueryCurrentChat string                            `gorm:"column:switch_inline_query_current_chat;size:512;default:'';" json:"switch_inline_query_current_chat,omitempty"` // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
	SwitchInlineQueryChosenChat  types.SwitchInlineQueryChosenChat `gorm:"column:switch_inline_query_chosen_chat;type:json" json:"switch_inline_query_chosen_chat,omitempty"`              // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
	Pay                          bool                              `gorm:"column:pay;type:bool;default:false" json:"pay,omitempty"`                                                        // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
	IsXtelegramToken             bool                              `gorm:"column:is_xtelegram_token;type:bool;default:false" json:"is_xtelegram_token,omitempty"`
	Status                       xtypes.OptionStatus               `gorm:"column:status;size:32;comment:状态( 1启用,2禁用)"`
	OperateUid                   int64                             `gorm:"column:operate_uid;size:64;comment:操作用户ID"`      //
	OperateUser                  string                            `gorm:"column:operate_user;size:64;comment:操作用户名"`      //
	CreateAt                     time.Time                         `gorm:"column:created_at;type:timestamp;comment:创建时间戳"` //
	UpdateAt                     time.Time                         `gorm:"column:updated_at;type:timestamp;comment:修改时间戳"` //
}

func (t *OptionTelegramInlineKeyboardButton) TableName() string {
	return "option_telegram_inline_keyboard_button"
}

func (buttonInfo *OptionTelegramInlineKeyboardButton) CheckFormat() error {
	if buttonInfo == nil {
		return fmt.Errorf("data is err")
	}
	switch buttonInfo.Kind {
	case tgtypes.ButtonType_InlineKeyboardButtonURL: //ButtonKind = "InlineKeyboardButtonURL"      //跳转网页
		{
			if buttonInfo.URL == "" {
				return fmt.Errorf("%v: must be url", tgtypes.ButtonType_InlineKeyboardButtonURL)
			}
		}
	case tgtypes.ButtonKind_InlineKeyboardWebApp: //ButtonKind = "InlineKeyboardWebApp"         //按钮
		{
			if buttonInfo.URL == "" {
				return fmt.Errorf("%v: must be url", tgtypes.ButtonKind_InlineKeyboardWebApp)
			}
		}
	case tgtypes.ButtonKind_InlineKeyboardButtonLoginURL: //ButtonKind = "InlineKeyboardButtonLoginURL" //按钮
		{
			if buttonInfo.LoginURL.IsNull() {
				return fmt.Errorf("%v: must be loginURL", tgtypes.ButtonKind_InlineKeyboardWebApp)
			}
		}
	case tgtypes.ButtonKind_InlineKeyboardButtonSwitch: //ButtonKind = "InlineKeyboardButtonSwitch"
		{
			if buttonInfo.SwitchInlineQuery == "" {
				return fmt.Errorf("%v: must be switchInlineQuery", tgtypes.ButtonKind_InlineKeyboardButtonSwitch)
			}
		}
	case tgtypes.ButtonKind_InlineKeyboardButtonSwitchCurrentChat: //ButtonKind = "InlineKeyboardButtonSwitchCurrentChat"
		{
			if buttonInfo.SwitchInlineQueryCurrentChat == "" {
				return fmt.Errorf("%v: must be switchInlineQueryCurrentChat", tgtypes.ButtonKind_InlineKeyboardButtonSwitchCurrentChat)
			}
		}
	case tgtypes.ButtonKind_InlineKeyboardButtonSwitchChosenChat: // ButtonKind = "InlineKeyboardButtonSwitchChosenChat"
		{
			if buttonInfo.SwitchInlineQueryChosenChat.IsNull() {
				return fmt.Errorf("%v: must be switchInlineQueryChosenChat", tgtypes.ButtonKind_InlineKeyboardButtonSwitchChosenChat)
			}
		}
	case tgtypes.ButtonKind_InlineKeyboardCallbackData: // ButtonKind = "InlineKeyboardButtonSwitchChosenChat"
		{
			if buttonInfo.CallbackData == "" {
				return fmt.Errorf("%v: must be switchInlineQueryChosenChat", tgtypes.ButtonKind_InlineKeyboardButtonSwitchChosenChat)
			}
		}
	}
	return nil
}
