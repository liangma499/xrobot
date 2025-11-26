package optiontelegraminlinekeyboardbutton

import (
	"context"
	"sync"
	"xbase/utils/xtime"
	mysqlimp "xrobot/internal/component/mysql/mysql-default"
	"xrobot/internal/dao/option-telegram-inline-keyboard-button/internal"
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

type OptionTelegramInlineKeyboardButton struct {
	*internal.OptionTelegramInlineKeyboardButton
}

func NewOptionTelegramInlineKeyboardButton(db *gorm.DB) *OptionTelegramInlineKeyboardButton {
	return &OptionTelegramInlineKeyboardButton{OptionTelegramInlineKeyboardButton: internal.NewOptionTelegramInlineKeyboardButton(db)}
}

var (
	once     sync.Once
	instance *OptionTelegramInlineKeyboardButton
)

func Instance() *OptionTelegramInlineKeyboardButton {
	once.Do(func() {
		instance = NewOptionTelegramInlineKeyboardButton(mysqlimp.Instance())
	})
	return instance
}
func (dao *OptionTelegramInlineKeyboardButton) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.OptionTelegramInlineKeyboardButton{})
		if err != nil {
			panic(err)
		}
	}
	dao.initInlineKeyboardButton()
	return nil
}
func (dao *OptionTelegramInlineKeyboardButton) initInlineKeyboardButton() error {
	now := xtime.Now()
	ctx := context.Background()

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                              //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_NumberOfTransactionPackages, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_NTP_20Bi,                        // Label text on the button
		LineID:                       1,                                                       //按钮所在行
		Sort:                         1,                                                       //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                              // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_NTP_20Bi.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                              // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                              // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},             // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                           // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})
	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                              //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_NumberOfTransactionPackages, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_NTP_30Bi,                        // Label text on the button
		LineID:                       1,                                                       //按钮所在行
		Sort:                         2,                                                       //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                              // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_NTP_30Bi.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                              // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                              // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},             // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                           // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})
	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                              //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_NumberOfTransactionPackages, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_NTP_50Bi,                        // Label text on the button
		LineID:                       1,                                                       //按钮所在行
		Sort:                         3,                                                       //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                              // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_NTP_50Bi.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                              // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                              // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},             // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                           // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                              //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_NumberOfTransactionPackages, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_NTP_100Bi,                       // Label text on the button
		LineID:                       2,                                                       //按钮所在行
		Sort:                         1,                                                       //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                               // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_NTP_100Bi.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                 // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                               // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                               // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},              // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                            // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                              //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_NumberOfTransactionPackages, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_NTP_200Bi,                       // Label text on the button
		LineID:                       2,                                                       //按钮所在行
		Sort:                         2,                                                       //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                               // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_NTP_200Bi.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                 // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                               // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                               // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},              // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                            // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                              //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_NumberOfTransactionPackages, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_NTP_300Bi,                       // Label text on the button
		LineID:                       2,                                                       //按钮所在行
		Sort:                         3,                                                       //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                               // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_NTP_300Bi.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                 // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                               // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                               // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},              // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                            // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                              //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_NumberOfTransactionPackages, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_NTP_500Bi,                       // Label text on the button
		LineID:                       3,                                                       //按钮所在行
		Sort:                         1,                                                       //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                               // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_NTP_500Bi.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                 // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                               // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                               // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},              // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                            // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                              //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_NumberOfTransactionPackages, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_NTP_1000Bi,                      // Label text on the button
		LineID:                       3,                                                       //按钮所在行
		Sort:                         2,                                                       //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_NTP_1000Bi.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                  // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},               // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                             // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                              //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_NumberOfTransactionPackages, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_NTP_2000Bi,                      // Label text on the button
		LineID:                       3,                                                       //按钮所在行
		Sort:                         3,                                                       //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_NTP_2000Bi.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                  // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},               // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                             // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	//XTelegramCmd_Button_EnergyFlashRental

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                    //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_EnergyFlashRental, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_ExpirationDate,    // Label text on the button
		LineID:                       1,                                             //按钮所在行
		Sort:                         1,                                             //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                        // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_ExpirationDate.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                          // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                        // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                        // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},                       // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                                     // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})
	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                    //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_EnergyFlashRental, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_1Bi,               // Label text on the button
		LineID:                       2,                                             //按钮所在行
		Sort:                         1,                                             //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                             // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_1Bi.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                               // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                             // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                             // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},            // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                          // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                    //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_EnergyFlashRental, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_2Bi,               // Label text on the button
		LineID:                       2,                                             //按钮所在行
		Sort:                         2,                                             //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                             // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_2Bi.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                               // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                             // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                             // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},            // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                          // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                    //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_EnergyFlashRental, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_3Bi,               // Label text on the button
		LineID:                       2,                                             //按钮所在行
		Sort:                         3,                                             //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                             // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_3Bi.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                               // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                             // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                             // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},            // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                          // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                    //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_EnergyFlashRental, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_5Bi,               // Label text on the button
		LineID:                       2,                                             //按钮所在行
		Sort:                         4,                                             //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                             // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_5Bi.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                               // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                             // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                             // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},            // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                          // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                    //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_EnergyFlashRental, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_10Bi,              // Label text on the button
		LineID:                       2,                                             //按钮所在行
		Sort:                         5,                                             //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                              // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_10Bi.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                              // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                              // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},             // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                           // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                    //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_EnergyFlashRental, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_NTP,                   // Label text on the button
		LineID:                       3,                                             //按钮所在行
		Sort:                         1,                                             //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                         // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_NTP.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                           // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                         // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                         // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},        // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                      // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})
	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                    //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_EnergyFlashRental, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_TelegramMember,        // Label text on the button
		LineID:                       3,                                             //按钮所在行
		Sort:                         2,                                             //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                    // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_TelegramMember.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                      // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                    // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                    // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},                   // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                                 // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})
	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                    //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_EnergyFlashRental, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_CustomizeTheSameRobot, // Label text on the button
		LineID:                       4,                                             //按钮所在行
		Sort:                         1,                                             //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                           // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_CustomizeTheSameRobot.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                             // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                           // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                           // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},                          // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                                        // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                         //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_EnergyFlashRental_BiSu, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_RechargeOtherAddresses, // Label text on the button
		LineID:                       1,                                                  //按钮所在行
		Sort:                         1,                                                  //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                                // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_RechargeOtherAddresses.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                                  // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                                // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                                // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},                               // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                                             // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	//XTelegramCmd_Button_RechargeOtherAddresses
	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                                       //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_RechargeOtherAddresses,               //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_RechargeOtherAddressesBalancePayment, // Label text on the button
		LineID:                       1,                                                                //按钮所在行
		Sort:                         1,                                                                //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                                              // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_RechargeOtherAddressesBalancePayment.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                                                // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                                              // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                                              // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},                                             // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                                                           // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                                    //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_RechargeOtherAddresses,            //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_RechargeOtherAddressesCancelOrder, // Label text on the button
		LineID:                       1,                                                             //按钮所在行
		Sort:                         2,                                                             //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                                           // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_RechargeOtherAddressesCancelOrder.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                                             // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                                           // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                                           // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},                                          // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                                                        // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_Recharge,      //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_Recharge50TRX, // Label text on the button
		LineID:                       1,                                         //按钮所在行
		Sort:                         1,                                         //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                       // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_Recharge50TRX.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                         // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                       // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                       // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},                      // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                                    // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                 //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_Recharge,       //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_Recharge100TRX, // Label text on the button
		LineID:                       1,                                          //按钮所在行
		Sort:                         2,                                          //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                        // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_Recharge100TRX.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                          // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                        // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                        // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},                       // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                                     // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                 //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_Recharge,       //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_Recharge300TRX, // Label text on the button
		LineID:                       1,                                          //按钮所在行
		Sort:                         3,                                          //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                        // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_Recharge300TRX.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                          // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                        // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                        // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},                       // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                                     // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                 //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_Recharge,       //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_Recharge500TRX, // Label text on the button
		LineID:                       1,                                          //按钮所在行
		Sort:                         4,                                          //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                        // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_Recharge500TRX.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                          // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                        // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                        // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},                       // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                                     // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                  //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_Recharge,        //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_Recharge1000TRX, // Label text on the button
		LineID:                       2,                                           //按钮所在行
		Sort:                         1,                                           //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                         // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_Recharge1000TRX.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                           // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                         // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                         // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},                        // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                                      // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                  //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_Recharge,        //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_Recharge2000TRX, // Label text on the button
		LineID:                       2,                                           //按钮所在行
		Sort:                         2,                                           //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                         // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_Recharge2000TRX.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                           // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                         // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                         // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},                        // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                                      // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                  //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_Recharge,        //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_Recharge3000TRX, // Label text on the button
		LineID:                       2,                                           //按钮所在行
		Sort:                         3,                                           //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                         // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_Recharge3000TRX.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                           // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                         // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                         // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},                        // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                                      // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                  //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_Button_Recharge,        //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_Recharge5000TRX, // Label text on the button
		LineID:                       2,                                           //按钮所在行
		Sort:                         4,                                           //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                         // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_Recharge5000TRX.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                           // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                         // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                         // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},                        // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                                      // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                 //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_CustomizeTheSameRobot, //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_BuySameRobot,       // Label text on the button
		LineID:                       1,                                          //按钮所在行
		Sort:                         1,                                          //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                  // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_BuySameRobot.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                    // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                  // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                  // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},                 // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                               // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	dao.Table.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramInlineKeyboardButton{
		ChannelCode:                  xtypes.OfficialChannelCode,                                    //            //渠道code
		Cmd:                          tgtypes.XTelegramCmd_BuySameRobot,                             //推送方式                                                                        // 主键                                                                 //                                                                     // 主键                                                                 // 主键
		Name:                         tgtypes.XTelegramButton_EFR_RechargeOtherAddressesCancelOrder, // Label text on the button
		LineID:                       1,                                                             //按钮所在行
		Sort:                         1,                                                             //按钮所在列
		Kind:                         tgtypes.ButtonKind_InlineKeyboardCallbackData,
		URL:                          "",                                                                           // 自选。按下按钮时要打开的 HTTP 或 tg:// URL。如果其隐私设置允许，则链接 tg：//user？id=<user_id> 可用于通过用户的 ID 提及用户，而无需使用用户名。
		CallbackData:                 tgtypes.XTelegramButton_EFR_RechargeOtherAddressesCancelOrder.CallbackData(), // 自选。用户按下按钮时将启动的 Web 应用程序的描述。Web 应用程序将能够使用 answerWebAppQuery 方法代表用户发送任意消息。仅在用户与 bot 之间的私人聊天中可用。
		LoginURL:                     types.LoginURL{},                                                             // 自选。按下按钮时要在回调查询中发送到机器人的数据，1-64 字节
		SwitchInlineQuery:            "",                                                                           // 自选。如果设置，按下该按钮将提示用户选择他们的一个聊天，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。注意：这为用户提供了一种简单的方法，当他们当前正在与您的机器人进行私人聊天时，他们可以在内联模式下开始使用您的机器人。与 switch_pm 结合使用时特别有用...操作 - 在这种情况下，用户将自动返回到他们切换的聊天，跳过聊天选择屏幕。
		SwitchInlineQueryCurrentChat: "",                                                                           // 自选。如果设置，按下该按钮将在当前聊天的输入字段中插入机器人的用户名和指定的内联查询。可能为空，在这种情况下，只会插入机器人的用户名。这为用户提供了一种在同一聊天中以内联模式打开机器人的快速方法 - 非常适合从多个选项中选择内容。
		SwitchInlineQueryChosenChat:  types.SwitchInlineQueryChosenChat{},                                          // 自选。如果设置，按下该按钮将提示用户选择指定类型的聊天之一，打开该聊天并在输入字段中插入机器人的用户名和指定的内联查询。
		Pay:                          false,                                                                        // 自选。指定 True，以发送 Pay 按钮。注意： 这种类型的按钮必须始终是第一行的第一个按钮，并且只能在发票消息中使用。
		IsXtelegramToken:             false,
		Status:                       xtypes.OptionStatus_Normal,
		OperateUid:                   -1,    //
		OperateUser:                  "初始化", //
		CreateAt:                     now,   //
		UpdateAt:                     now,   //
	})

	return nil
}
