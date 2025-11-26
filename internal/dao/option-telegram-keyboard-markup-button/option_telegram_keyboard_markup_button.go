package optiontelegramkeyboardmarkupbutton

import (
	"context"
	"sync"
	"xbase/utils/xtime"
	mysqlimp "xrobot/internal/component/mysql/mysql-default"
	"xrobot/internal/dao/option-telegram-keyboard-markup-button/internal"
	modelpkg "xrobot/internal/model"
	"xrobot/internal/xtelegram/telegram/types"
	tgtypes "xrobot/internal/xtelegram/tg-types"
	"xrobot/internal/xtypes"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	Columns    = internal.Columns
	OrderBy    = internal.OrderBy
	FilterFunc = internal.FilterFunc
	UpdateFunc = internal.UpdateFunc
	ColumnFunc = internal.ColumnFunc
	OrderFunc  = internal.OrderFunc
)

type OptionTelegramKeyboardMarkupButton struct {
	*internal.OptionTelegramKeyboardMarkupButton
}

func NewOptionTelegramKeyboardMarkupButton(db *gorm.DB) *OptionTelegramKeyboardMarkupButton {
	return &OptionTelegramKeyboardMarkupButton{OptionTelegramKeyboardMarkupButton: internal.NewOptionTelegramKeyboardMarkupButton(db)}
}

var (
	once     sync.Once
	instance *OptionTelegramKeyboardMarkupButton
)

func Instance() *OptionTelegramKeyboardMarkupButton {
	once.Do(func() {
		instance = NewOptionTelegramKeyboardMarkupButton(mysqlimp.Instance())
	})
	return instance
}
func (dao *OptionTelegramKeyboardMarkupButton) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.OptionTelegramKeyboardMarkupButton{})
		if err != nil {
			panic(err)
		}
	}
	dao.initKeyboardMarkupButton()
	return nil
}

func (dao *OptionTelegramKeyboardMarkupButton) initKeyboardMarkupButton() error {
	now := xtime.Now()
	ctx := context.Background()
	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramKeyboardMarkupButton{
		ChannelCode:     xtypes.OfficialChannelCode,                             //            //渠道code
		Cmd:             tgtypes.XTelegramCmd_Start,                             //推送方式                                                                        // 主键
		Text:            tgtypes.XTelegramCmd_Button_EnergyFlashRental.String(), // 按钮的文本。如果没有使用任何可选字段，则在按钮被按下时以消息的形式发送
		LineID:          1,                                                      //按钮所在行
		Sort:            1,                                                      //按钮所在列
		RequestUsers:    types.KeyboardButtonRequestUsers{},                     // 可选。如果指定，按下按钮时将打开合适用户的列表。点击任何用户将以“users_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestChat:     types.KeyboardButtonRequestChat{},                      // 可选。如果指定，按下按钮时将打开合适聊天的列表。点击一个聊天将以“chat_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestContact:  false,                                                  // 可选。如果为 true，用户的电话号码将在按钮被按下时作为联系人发送。仅在私聊中可用。
		RequestLocation: false,                                                  // 可选。如果为 true，用户的当前位置将在按钮被按下时发送。仅在私聊中可用。
		RequestPoll:     types.KeyboardButtonPollType{},                         // 可选。如果指定，用户将被要求创建一个投票并在按钮被按下时发送给机器人。仅在私聊中可用。
		WebApp:          types.WebAppInfo{},                                     // 可选。如果指定，描述的 Web 应用将在按钮被按下时启动。Web 应用能够发送“web_app_data”服务消息。仅在私聊中可用。
		Status:          xtypes.OptionStatus_Normal,
		OperateUid:      -1,     //
		OperateUser:     "init", //
		CreateAt:        now,    //
		UpdateAt:        now,    //
	})
	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramKeyboardMarkupButton{
		ChannelCode:     xtypes.OfficialChannelCode,                      //            //渠道code
		Cmd:             tgtypes.XTelegramCmd_Start,                      //推送方式                                                                        // 主键
		Text:            tgtypes.XTelegramCmd_Button_TRXConvert.String(), // 按钮的文本。如果没有使用任何可选字段，则在按钮被按下时以消息的形式发送
		LineID:          1,                                               //按钮所在行
		Sort:            2,                                               //按钮所在列
		RequestUsers:    types.KeyboardButtonRequestUsers{},              // 可选。如果指定，按下按钮时将打开合适用户的列表。点击任何用户将以“users_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestChat:     types.KeyboardButtonRequestChat{},               // 可选。如果指定，按下按钮时将打开合适聊天的列表。点击一个聊天将以“chat_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestContact:  false,                                           // 可选。如果为 true，用户的电话号码将在按钮被按下时作为联系人发送。仅在私聊中可用。
		RequestLocation: false,                                           // 可选。如果为 true，用户的当前位置将在按钮被按下时发送。仅在私聊中可用。
		RequestPoll:     types.KeyboardButtonPollType{},                  // 可选。如果指定，用户将被要求创建一个投票并在按钮被按下时发送给机器人。仅在私聊中可用。
		WebApp:          types.WebAppInfo{},                              // 可选。如果指定，描述的 Web 应用将在按钮被按下时启动。Web 应用能够发送“web_app_data”服务消息。仅在私聊中可用。
		Status:          xtypes.OptionStatus_Normal,
		OperateUid:      -1,     //
		OperateUser:     "init", //
		CreateAt:        now,    //
		UpdateAt:        now,    //
	})
	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramKeyboardMarkupButton{
		ChannelCode:     xtypes.OfficialChannelCode,                          //            //渠道code
		Cmd:             tgtypes.XTelegramCmd_Start,                          //推送方式                                                                        // 主键
		Text:            tgtypes.XTelegramCmd_Button_EnergyAdvances.String(), // 按钮的文本。如果没有使用任何可选字段，则在按钮被按下时以消息的形式发送
		LineID:          1,                                                   //按钮所在行
		Sort:            2,                                                   //按钮所在列
		RequestUsers:    types.KeyboardButtonRequestUsers{},                  // 可选。如果指定，按下按钮时将打开合适用户的列表。点击任何用户将以“users_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestChat:     types.KeyboardButtonRequestChat{},                   // 可选。如果指定，按下按钮时将打开合适聊天的列表。点击一个聊天将以“chat_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestContact:  false,                                               // 可选。如果为 true，用户的电话号码将在按钮被按下时作为联系人发送。仅在私聊中可用。
		RequestLocation: false,                                               // 可选。如果为 true，用户的当前位置将在按钮被按下时发送。仅在私聊中可用。
		RequestPoll:     types.KeyboardButtonPollType{},                      // 可选。如果指定，用户将被要求创建一个投票并在按钮被按下时发送给机器人。仅在私聊中可用。
		WebApp:          types.WebAppInfo{},                                  // 可选。如果指定，描述的 Web 应用将在按钮被按下时启动。Web 应用能够发送“web_app_data”服务消息。仅在私聊中可用。
		Status:          xtypes.OptionStatus_Normal,
		OperateUid:      -1,     //
		OperateUser:     "init", //
		CreateAt:        now,    //
		UpdateAt:        now,    //
	})
	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramKeyboardMarkupButton{
		ChannelCode:     xtypes.OfficialChannelCode,                                       //            //渠道code
		Cmd:             tgtypes.XTelegramCmd_Start,                                       //推送方式                                                                        // 主键
		Text:            tgtypes.XTelegramCmd_Button_NumberOfTransactionPackages.String(), // 按钮的文本。如果没有使用任何可选字段，则在按钮被按下时以消息的形式发送
		LineID:          2,                                                                //按钮所在行
		Sort:            1,                                                                //按钮所在列
		RequestUsers:    types.KeyboardButtonRequestUsers{},                               // 可选。如果指定，按下按钮时将打开合适用户的列表。点击任何用户将以“users_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestChat:     types.KeyboardButtonRequestChat{},                                // 可选。如果指定，按下按钮时将打开合适聊天的列表。点击一个聊天将以“chat_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestContact:  false,                                                            // 可选。如果为 true，用户的电话号码将在按钮被按下时作为联系人发送。仅在私聊中可用。
		RequestLocation: false,                                                            // 可选。如果为 true，用户的当前位置将在按钮被按下时发送。仅在私聊中可用。
		RequestPoll:     types.KeyboardButtonPollType{},                                   // 可选。如果指定，用户将被要求创建一个投票并在按钮被按下时发送给机器人。仅在私聊中可用。
		WebApp:          types.WebAppInfo{},                                               // 可选。如果指定，描述的 Web 应用将在按钮被按下时启动。Web 应用能够发送“web_app_data”服务消息。仅在私聊中可用。
		Status:          xtypes.OptionStatus_Normal,
		OperateUid:      -1,     //
		OperateUser:     "init", //
		CreateAt:        now,    //
		UpdateAt:        now,    //
	})
	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramKeyboardMarkupButton{
		ChannelCode:     xtypes.OfficialChannelCode,                          //            //渠道code
		Cmd:             tgtypes.XTelegramCmd_Start,                          //推送方式                                                                        // 主键
		Text:            tgtypes.XTelegramCmd_Button_TelegramMember.String(), // 按钮的文本。如果没有使用任何可选字段，则在按钮被按下时以消息的形式发送
		LineID:          2,                                                   //按钮所在行
		Sort:            2,                                                   //按钮所在列
		RequestUsers:    types.KeyboardButtonRequestUsers{},                  // 可选。如果指定，按下按钮时将打开合适用户的列表。点击任何用户将以“users_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestChat:     types.KeyboardButtonRequestChat{},                   // 可选。如果指定，按下按钮时将打开合适聊天的列表。点击一个聊天将以“chat_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestContact:  false,                                               // 可选。如果为 true，用户的电话号码将在按钮被按下时作为联系人发送。仅在私聊中可用。
		RequestLocation: false,                                               // 可选。如果为 true，用户的当前位置将在按钮被按下时发送。仅在私聊中可用。
		RequestPoll:     types.KeyboardButtonPollType{},                      // 可选。如果指定，用户将被要求创建一个投票并在按钮被按下时发送给机器人。仅在私聊中可用。
		WebApp:          types.WebAppInfo{},                                  // 可选。如果指定，描述的 Web 应用将在按钮被按下时启动。Web 应用能够发送“web_app_data”服务消息。仅在私聊中可用。
		Status:          xtypes.OptionStatus_Normal,
		OperateUid:      -1,     //
		OperateUser:     "init", //
		CreateAt:        now,    //
		UpdateAt:        now,    //
	})
	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramKeyboardMarkupButton{
		ChannelCode:     xtypes.OfficialChannelCode,                            //            //渠道code
		Cmd:             tgtypes.XTelegramCmd_Start,                            //推送方式                                                                        // 主键
		Text:            tgtypes.XTelegramCmd_Button_PromoteMakeMoney.String(), // 按钮的文本。如果没有使用任何可选字段，则在按钮被按下时以消息的形式发送
		LineID:          2,                                                     //按钮所在行
		Sort:            3,                                                     //按钮所在列
		RequestUsers:    types.KeyboardButtonRequestUsers{},                    // 可选。如果指定，按下按钮时将打开合适用户的列表。点击任何用户将以“users_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestChat:     types.KeyboardButtonRequestChat{},                     // 可选。如果指定，按下按钮时将打开合适聊天的列表。点击一个聊天将以“chat_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestContact:  false,                                                 // 可选。如果为 true，用户的电话号码将在按钮被按下时作为联系人发送。仅在私聊中可用。
		RequestLocation: false,                                                 // 可选。如果为 true，用户的当前位置将在按钮被按下时发送。仅在私聊中可用。
		RequestPoll:     types.KeyboardButtonPollType{},                        // 可选。如果指定，用户将被要求创建一个投票并在按钮被按下时发送给机器人。仅在私聊中可用。
		WebApp:          types.WebAppInfo{},                                    // 可选。如果指定，描述的 Web 应用将在按钮被按下时启动。Web 应用能够发送“web_app_data”服务消息。仅在私聊中可用。
		Status:          xtypes.OptionStatus_Normal,
		OperateUid:      -1,     //
		OperateUser:     "init", //
		CreateAt:        now,    //
		UpdateAt:        now,    //
	})
	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramKeyboardMarkupButton{
		ChannelCode:     xtypes.OfficialChannelCode,                       //            //渠道code
		Cmd:             tgtypes.XTelegramCmd_Start,                       //推送方式                                                                        // 主键
		Text:            tgtypes.XTelegramCmd_Button_GoodAddress.String(), // 按钮的文本。如果没有使用任何可选字段，则在按钮被按下时以消息的形式发送
		LineID:          3,                                                //按钮所在行
		Sort:            1,                                                //按钮所在列
		RequestUsers:    types.KeyboardButtonRequestUsers{},               // 可选。如果指定，按下按钮时将打开合适用户的列表。点击任何用户将以“users_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestChat:     types.KeyboardButtonRequestChat{},                // 可选。如果指定，按下按钮时将打开合适聊天的列表。点击一个聊天将以“chat_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestContact:  false,                                            // 可选。如果为 true，用户的电话号码将在按钮被按下时作为联系人发送。仅在私聊中可用。
		RequestLocation: false,                                            // 可选。如果为 true，用户的当前位置将在按钮被按下时发送。仅在私聊中可用。
		RequestPoll:     types.KeyboardButtonPollType{},                   // 可选。如果指定，用户将被要求创建一个投票并在按钮被按下时发送给机器人。仅在私聊中可用。
		WebApp:          types.WebAppInfo{},                               // 可选。如果指定，描述的 Web 应用将在按钮被按下时启动。Web 应用能够发送“web_app_data”服务消息。仅在私聊中可用。
		Status:          xtypes.OptionStatus_Normal,
		OperateUid:      -1,     //
		OperateUser:     "init", //
		CreateAt:        now,    //
		UpdateAt:        now,    //
	})
	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramKeyboardMarkupButton{
		ChannelCode:     xtypes.OfficialChannelCode,                            //            //渠道code
		Cmd:             tgtypes.XTelegramCmd_Start,                            //推送方式                                                                        // 主键
		Text:            tgtypes.XTelegramCmd_Button_ListeningAddress.String(), // 按钮的文本。如果没有使用任何可选字段，则在按钮被按下时以消息的形式发送
		LineID:          3,                                                     //按钮所在行
		Sort:            2,                                                     //按钮所在列
		RequestUsers:    types.KeyboardButtonRequestUsers{},                    // 可选。如果指定，按下按钮时将打开合适用户的列表。点击任何用户将以“users_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestChat:     types.KeyboardButtonRequestChat{},                     // 可选。如果指定，按下按钮时将打开合适聊天的列表。点击一个聊天将以“chat_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestContact:  false,                                                 // 可选。如果为 true，用户的电话号码将在按钮被按下时作为联系人发送。仅在私聊中可用。
		RequestLocation: false,                                                 // 可选。如果为 true，用户的当前位置将在按钮被按下时发送。仅在私聊中可用。
		RequestPoll:     types.KeyboardButtonPollType{},                        // 可选。如果指定，用户将被要求创建一个投票并在按钮被按下时发送给机器人。仅在私聊中可用。
		WebApp:          types.WebAppInfo{},                                    // 可选。如果指定，描述的 Web 应用将在按钮被按下时启动。Web 应用能够发送“web_app_data”服务消息。仅在私聊中可用。
		Status:          xtypes.OptionStatus_Normal,
		OperateUid:      -1,     //
		OperateUser:     "init", //
		CreateAt:        now,    //
		UpdateAt:        now,    //
	})
	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramKeyboardMarkupButton{
		ChannelCode:     xtypes.OfficialChannelCode,                          //            //渠道code
		Cmd:             tgtypes.XTelegramCmd_Start,                          //推送方式                                                                        // 主键
		Text:            tgtypes.XTelegramCmd_Button_PersonalCenter.String(), // 按钮的文本。如果没有使用任何可选字段，则在按钮被按下时以消息的形式发送
		LineID:          3,                                                   //按钮所在行
		Sort:            3,                                                   //按钮所在列
		RequestUsers:    types.KeyboardButtonRequestUsers{},                  // 可选。如果指定，按下按钮时将打开合适用户的列表。点击任何用户将以“users_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestChat:     types.KeyboardButtonRequestChat{},                   // 可选。如果指定，按下按钮时将打开合适聊天的列表。点击一个聊天将以“chat_shared”服务消息的形式将其标识符发送给机器人。仅在私聊中可用。
		RequestContact:  false,                                               // 可选。如果为 true，用户的电话号码将在按钮被按下时作为联系人发送。仅在私聊中可用。
		RequestLocation: false,                                               // 可选。如果为 true，用户的当前位置将在按钮被按下时发送。仅在私聊中可用。
		RequestPoll:     types.KeyboardButtonPollType{},                      // 可选。如果指定，用户将被要求创建一个投票并在按钮被按下时发送给机器人。仅在私聊中可用。
		WebApp:          types.WebAppInfo{},                                  // 可选。如果指定，描述的 Web 应用将在按钮被按下时启动。Web 应用能够发送“web_app_data”服务消息。仅在私聊中可用。
		Status:          xtypes.OptionStatus_Normal,
		OperateUid:      -1,     //
		OperateUser:     "init", //
		CreateAt:        now,    //
		UpdateAt:        now,    //
	})
	return nil
}
