package model

import (
	"time"
	"xrobot/internal/xtelegram/telegram/types"
	tgtypes "xrobot/internal/xtelegram/tg-types"
	"xrobot/internal/xtypes"
)

//go:generate xgorm-dao-generator -model-dir=. -model-names=OptionTelegramKeyboardMarkupButton -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type OptionTelegramKeyboardMarkupButton struct {
	ChannelCode     string                           `gorm:"column:channel_code;size:32;primarykey;"`             //渠道code
	Cmd             tgtypes.XTelegramCmd             `gorm:"column:cmd;size:64;primarykey"`                       //推送方式                                                                        // 主键
	Text            string                           `gorm:"column:text;size:32;not null;primarykey" json:"text"` // 按钮的文本。如果没有使用任何可选字段，则在按钮被按下时以消息的形式发送
	LineID          int                              `gorm:"column:line_id;size:32;default:0" json:"line_id"`     //按钮所在行
	Sort            int                              `gorm:"column:sort;size:32;default:0" json:"sort"`           //按钮所在列
	RequestUsers    types.KeyboardButtonRequestUsers `gorm:"request_users;type:json"`                             // 可选。如果指定，按下按钮时将打开合适用户的列表。点击任何用户将以“users_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
	RequestChat     types.KeyboardButtonRequestChat  `gorm:"request_chat;type:json"`                              // 可选。如果指定，按下按钮时将打开合适聊天的列表。点击一个聊天将以“chat_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
	RequestContact  bool                             `gorm:"request_contact;type:bool"`                           // 可选。如果为 true，用户的电话号码将在按钮被按下时作为联系人发送。仅在私聊中可用。
	RequestLocation bool                             `gorm:"request_location;type:bool"`                          // 可选。如果为 true，用户的当前位置将在按钮被按下时发送。仅在私聊中可用。
	RequestPoll     types.KeyboardButtonPollType     `gorm:"request_poll;type:json"`                              // 可选。如果指定，用户将被要求创建一个投票并在按钮被按下时发送给机器人。仅在私聊中可用。
	WebApp          types.WebAppInfo                 `gorm:"web_app;type:json"`                                   // 可选。如果指定，描述的 Web 应用将在按钮被按下时启动。Web 应用能够发送“web_app_data”服务消息。仅在私聊中可用。
	Status          xtypes.OptionStatus              `gorm:"column:status;size:32;comment:状态( 1启用,2禁用)"`
	OperateUid      int64                            `gorm:"column:operate_uid;size:64;comment:操作用户ID"`      //
	OperateUser     string                           `gorm:"column:operate_user;size:64;comment:操作用户名"`      //
	CreateAt        time.Time                        `gorm:"column:created_at;type:timestamp;comment:创建时间戳"` //
	UpdateAt        time.Time                        `gorm:"column:updated_at;type:timestamp;comment:修改时间戳"` //
}

func (t *OptionTelegramKeyboardMarkupButton) TableName() string {
	return "option_telegram_keyboard_markup_button"
}
