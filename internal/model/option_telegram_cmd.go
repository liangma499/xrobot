package model

import (
	"time"
	"tron_robot/internal/xtelegram/telegram/types"
	tgtypes "tron_robot/internal/xtelegram/tg-types"
	"tron_robot/internal/xtypes"
)

//go:generate xgorm-dao-generator -model-dir=. -model-names=OptionTelegramCmd -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type OptionTelegramCmd struct {
	ChannelCode         string                      `gorm:"column:channel_code;size:32;primarykey;"` //渠道code
	Cmd                 tgtypes.XTelegramCmd        `gorm:"column:cmd;size:64;primarykey"`           //推送方式
	CmdKind             tgtypes.CmdKind             `gorm:"column:cmd_kind;size:64"`                 //推送方式
	ReplyKeyboardMarkup types.ReplyKeyboardMarkupDb `gorm:"column:reply_keyboard_markup;type:json"`
	PictureUrl          string                      `gorm:"column:picture_url;size:255;"`                               //图片
	Type                tgtypes.RobotMsgType        `gorm:"column:type;size:32;"`                                       //文件类型(1=图片,2=视频)
	Text                string                      `gorm:"column:text;type:text"`                                      //文本内容
	ParseMode           tgtypes.ParseMode           `gorm:"column:parse_mode;size:64;comment:Markdown,MarkdownV2,HTML"` //编码方式
	Status              xtypes.OptionStatus         `gorm:"column:status;size:32;comment:状态( 1启用,2禁用)"`
	OperateUid          int64                       `gorm:"column:operate_uid;size:64;comment:操作用户ID"`      //
	OperateUser         string                      `gorm:"column:operate_user;size:64;comment:操作用户名"`      //
	CreateAt            time.Time                   `gorm:"column:created_at;type:timestamp;comment:创建时间戳"` //
	UpdateAt            time.Time                   `gorm:"column:updated_at;type:timestamp;comment:修改时间戳"` //
}

func (t *OptionTelegramCmd) TableName() string {
	return "option_telegram_cmd"
}
