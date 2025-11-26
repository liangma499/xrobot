package optiontelegramcmd

import (
	"context"
	"sync"
	"xbase/utils/xtime"
	mysqlimp "xrobot/internal/component/mysql/mysql-default"
	"xrobot/internal/dao/option-telegram-cmd/internal"
	modelpkg "xrobot/internal/model"
	"xrobot/internal/xtelegram/telegram/types"
	tgtemplate "xrobot/internal/xtelegram/tg-template"
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

type OptionTelegramCmd struct {
	*internal.OptionTelegramCmd
}

func NewOptionTelegramCmd(db *gorm.DB) *OptionTelegramCmd {
	return &OptionTelegramCmd{OptionTelegramCmd: internal.NewOptionTelegramCmd(db)}
}

var (
	once     sync.Once
	instance *OptionTelegramCmd
)

func Instance() *OptionTelegramCmd {
	once.Do(func() {
		instance = NewOptionTelegramCmd(mysqlimp.Instance())
	})
	return instance
}
func (dao *OptionTelegramCmd) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.OptionTelegramCmd{})
		if err != nil {
			panic(err)
		}
	}
	dao.initCmd()
	return nil
}
func (dao *OptionTelegramCmd) initCmd() {
	now := xtime.Now()
	dao.Table.WithContext(context.Background()).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramCmd{
		ChannelCode: xtypes.OfficialChannelCode, //渠道code
		Cmd:         tgtypes.XTelegramCmd_Start, //
		CmdKind:     tgtypes.CmdKind_KeyboardMarkup,
		ReplyKeyboardMarkup: types.ReplyKeyboardMarkupDb{
			ResizeKeyboard:  true,
			OneTimeKeyboard: false,
		},
		PictureUrl:  "",                        //图片
		Type:        tgtypes.RobotMsgTypePhoto, //文件类型(1=图片,2=视频)
		Text:        tgtemplate.Start(),        //文本内容
		ParseMode:   tgtypes.ModeNone,
		Status:      xtypes.OptionStatus_Normal,
		OperateUid:  0,      //
		OperateUser: "init", //
		CreateAt:    now,    //
		UpdateAt:    now,    //
	})

	dao.Table.WithContext(context.Background()).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramCmd{
		ChannelCode: xtypes.OfficialChannelCode,                    //渠道code
		Cmd:         tgtypes.XTelegramCmd_Button_EnergyFlashRental, //
		CmdKind:     tgtypes.CmdKind_InlineKeyboard,
		ReplyKeyboardMarkup: types.ReplyKeyboardMarkupDb{
			ResizeKeyboard:  true,
			OneTimeKeyboard: false,
		},
		PictureUrl:  "",                             //图片
		Type:        tgtypes.RobotMsgTypePhoto,      //文件类型(1=图片,2=视频)
		Text:        tgtemplate.EnergyFlashRental(), //文本内容
		ParseMode:   tgtypes.ModeMarkdown,
		Status:      xtypes.OptionStatus_Normal,
		OperateUid:  0,      //
		OperateUser: "init", //
		CreateAt:    now,    //
		UpdateAt:    now,    //
	})
	dao.Table.WithContext(context.Background()).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramCmd{
		ChannelCode: xtypes.OfficialChannelCode,             //渠道code
		Cmd:         tgtypes.XTelegramCmd_Button_TRXConvert, //
		CmdKind:     tgtypes.CmdKind_InlineKeyboard,
		ReplyKeyboardMarkup: types.ReplyKeyboardMarkupDb{
			ResizeKeyboard:  true,
			OneTimeKeyboard: false,
		},
		PictureUrl:  "",                        //图片
		Type:        tgtypes.RobotMsgTypePhoto, //文件类型(1=图片,2=视频)
		Text:        tgtemplate.Start(),        //文本内容
		ParseMode:   tgtypes.ModeMarkdown,
		Status:      xtypes.OptionStatus_Normal,
		OperateUid:  0,      //
		OperateUser: "init", //
		CreateAt:    now,    //
		UpdateAt:    now,    //
	})

	dao.Table.WithContext(context.Background()).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramCmd{
		ChannelCode: xtypes.OfficialChannelCode,                 //渠道code
		Cmd:         tgtypes.XTelegramCmd_Button_EnergyAdvances, //
		CmdKind:     tgtypes.CmdKind_InlineKeyboard,
		ReplyKeyboardMarkup: types.ReplyKeyboardMarkupDb{
			ResizeKeyboard:  true,
			OneTimeKeyboard: false,
		},
		PictureUrl:  "",                        //图片
		Type:        tgtypes.RobotMsgTypePhoto, //文件类型(1=图片,2=视频)
		Text:        tgtemplate.Start(),        //文本内容
		ParseMode:   tgtypes.ModeMarkdown,
		Status:      xtypes.OptionStatus_Normal,
		OperateUid:  0,      //
		OperateUser: "init", //
		CreateAt:    now,    //
		UpdateAt:    now,    //
	})

	dao.Table.WithContext(context.Background()).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramCmd{
		ChannelCode: xtypes.OfficialChannelCode,                              //渠道code
		Cmd:         tgtypes.XTelegramCmd_Button_NumberOfTransactionPackages, //
		CmdKind:     tgtypes.CmdKind_InlineKeyboard,
		ReplyKeyboardMarkup: types.ReplyKeyboardMarkupDb{
			ResizeKeyboard:  true,
			OneTimeKeyboard: false,
		},
		PictureUrl:  "",                        //图片
		Type:        tgtypes.RobotMsgTypePhoto, //文件类型(1=图片,2=视频)
		Text:        tgtemplate.Start(),        //文本内容
		ParseMode:   tgtypes.ModeMarkdown,
		Status:      xtypes.OptionStatus_Normal,
		OperateUid:  0,      //
		OperateUser: "init", //
		CreateAt:    now,    //
		UpdateAt:    now,    //
	})

	dao.Table.WithContext(context.Background()).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramCmd{
		ChannelCode: xtypes.OfficialChannelCode,                 //渠道code
		Cmd:         tgtypes.XTelegramCmd_Button_TelegramMember, //
		CmdKind:     tgtypes.CmdKind_InlineKeyboard,
		ReplyKeyboardMarkup: types.ReplyKeyboardMarkupDb{
			ResizeKeyboard:  true,
			OneTimeKeyboard: false,
		},
		PictureUrl:  "",                        //图片
		Type:        tgtypes.RobotMsgTypePhoto, //文件类型(1=图片,2=视频)
		Text:        tgtemplate.Start(),        //文本内容
		ParseMode:   tgtypes.ModeMarkdown,
		Status:      xtypes.OptionStatus_Normal,
		OperateUid:  0,      //
		OperateUser: "init", //
		CreateAt:    now,    //
		UpdateAt:    now,    //
	})
	dao.Table.WithContext(context.Background()).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramCmd{
		ChannelCode: xtypes.OfficialChannelCode,                   //渠道code
		Cmd:         tgtypes.XTelegramCmd_Button_PromoteMakeMoney, //
		CmdKind:     tgtypes.CmdKind_InlineKeyboard,
		ReplyKeyboardMarkup: types.ReplyKeyboardMarkupDb{
			ResizeKeyboard:  true,
			OneTimeKeyboard: false,
		},
		PictureUrl:  "",                        //图片
		Type:        tgtypes.RobotMsgTypePhoto, //文件类型(1=图片,2=视频)
		Text:        tgtemplate.Start(),        //文本内容
		ParseMode:   tgtypes.ModeMarkdown,
		Status:      xtypes.OptionStatus_Normal,
		OperateUid:  0,      //
		OperateUser: "init", //
		CreateAt:    now,    //
		UpdateAt:    now,    //
	})
	dao.Table.WithContext(context.Background()).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramCmd{
		ChannelCode: xtypes.OfficialChannelCode,              //渠道code
		Cmd:         tgtypes.XTelegramCmd_Button_GoodAddress, //
		CmdKind:     tgtypes.CmdKind_InlineKeyboard,
		ReplyKeyboardMarkup: types.ReplyKeyboardMarkupDb{
			ResizeKeyboard:  true,
			OneTimeKeyboard: false,
		},
		PictureUrl:  "",                        //图片
		Type:        tgtypes.RobotMsgTypePhoto, //文件类型(1=图片,2=视频)
		Text:        tgtemplate.Start(),        //文本内容
		ParseMode:   tgtypes.ModeMarkdown,
		Status:      xtypes.OptionStatus_Normal,
		OperateUid:  0,      //
		OperateUser: "init", //
		CreateAt:    now,    //
		UpdateAt:    now,    //
	})
	dao.Table.WithContext(context.Background()).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramCmd{
		ChannelCode: xtypes.OfficialChannelCode,                   //渠道code
		Cmd:         tgtypes.XTelegramCmd_Button_ListeningAddress, //
		CmdKind:     tgtypes.CmdKind_InlineKeyboard,
		ReplyKeyboardMarkup: types.ReplyKeyboardMarkupDb{
			ResizeKeyboard:  true,
			OneTimeKeyboard: false,
		},
		PictureUrl:  "",                        //图片
		Type:        tgtypes.RobotMsgTypePhoto, //文件类型(1=图片,2=视频)
		Text:        tgtemplate.Start(),        //文本内容
		ParseMode:   tgtypes.ModeMarkdown,
		Status:      xtypes.OptionStatus_Normal,
		OperateUid:  0,      //
		OperateUser: "init", //
		CreateAt:    now,    //
		UpdateAt:    now,    //
	})
	dao.Table.WithContext(context.Background()).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramCmd{
		ChannelCode: xtypes.OfficialChannelCode,                 //渠道code
		Cmd:         tgtypes.XTelegramCmd_Button_PersonalCenter, //
		CmdKind:     tgtypes.CmdKind_InlineKeyboard,
		ReplyKeyboardMarkup: types.ReplyKeyboardMarkupDb{
			ResizeKeyboard:  true,
			OneTimeKeyboard: false,
		},
		PictureUrl:  "",                        //图片
		Type:        tgtypes.RobotMsgTypePhoto, //文件类型(1=图片,2=视频)
		Text:        tgtemplate.Start(),        //文本内容
		ParseMode:   tgtypes.ModeMarkdown,
		Status:      xtypes.OptionStatus_Normal,
		OperateUid:  0,      //
		OperateUser: "init", //
		CreateAt:    now,    //
		UpdateAt:    now,    //
	})

	dao.Table.WithContext(context.Background()).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramCmd{
		ChannelCode: xtypes.OfficialChannelCode,                         //渠道code
		Cmd:         tgtypes.XTelegramCmd_Button_EnergyFlashRental_BiSu, //
		CmdKind:     tgtypes.CmdKind_InlineKeyboard,
		ReplyKeyboardMarkup: types.ReplyKeyboardMarkupDb{
			ResizeKeyboard:  true,
			OneTimeKeyboard: false,
		},
		PictureUrl:  "",                                  //图片
		Type:        tgtypes.RobotMsgTypePhoto,           //文件类型(1=图片,2=视频)
		Text:        tgtemplate.EnergyFlashRentalBiShu(), //文本内容
		ParseMode:   tgtypes.ModeMarkdown,
		Status:      xtypes.OptionStatus_Normal,
		OperateUid:  0,      //
		OperateUser: "init", //
		CreateAt:    now,    //
		UpdateAt:    now,    //
	})
	dao.Table.WithContext(context.Background()).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramCmd{
		ChannelCode: xtypes.OfficialChannelCode,                         //渠道code
		Cmd:         tgtypes.XTelegramCmd_Button_RechargeOtherAddresses, //
		CmdKind:     tgtypes.CmdKind_InlineKeyboard,
		ReplyKeyboardMarkup: types.ReplyKeyboardMarkupDb{
			ResizeKeyboard:  true,
			OneTimeKeyboard: false,
		},
		PictureUrl:  "",                                     //图片
		Type:        tgtypes.RobotMsgTypeText,               //文件类型(1=图片,2=视频)
		Text:        tgtemplate.RechargeOtherAddressesRet(), //文本内容
		ParseMode:   tgtypes.ModeMarkdown,
		Status:      xtypes.OptionStatus_Normal,
		OperateUid:  0,      //
		OperateUser: "init", //
		CreateAt:    now,    //
		UpdateAt:    now,    //
	})
	dao.Table.WithContext(context.Background()).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramCmd{
		ChannelCode: xtypes.OfficialChannelCode,           //渠道code
		Cmd:         tgtypes.XTelegramCmd_Button_Recharge, //
		CmdKind:     tgtypes.CmdKind_InlineKeyboard,
		ReplyKeyboardMarkup: types.ReplyKeyboardMarkupDb{
			ResizeKeyboard:  false,
			OneTimeKeyboard: false,
		},
		PictureUrl:  "",                       //图片
		Type:        tgtypes.RobotMsgTypeText, //文件类型(1=图片,2=视频)
		Text:        tgtemplate.RechargeRet(), //文本内容
		ParseMode:   tgtypes.ModeMarkdown,
		Status:      xtypes.OptionStatus_Normal,
		OperateUid:  0,      //
		OperateUser: "init", //
		CreateAt:    now,    //
		UpdateAt:    now,    //
	})
	dao.Table.WithContext(context.Background()).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramCmd{
		ChannelCode: xtypes.OfficialChannelCode,                 //渠道code
		Cmd:         tgtypes.XTelegramCmd_CustomizeTheSameRobot, //
		CmdKind:     tgtypes.CmdKind_InlineKeyboard,
		ReplyKeyboardMarkup: types.ReplyKeyboardMarkupDb{
			ResizeKeyboard:  false,
			OneTimeKeyboard: false,
		},
		PictureUrl:  "",                                 //图片
		Type:        tgtypes.RobotMsgTypeText,           //文件类型(1=图片,2=视频)
		Text:        tgtemplate.CustomizeTheSameRobot(), //文本内容
		ParseMode:   tgtypes.ModeMarkdown,
		Status:      xtypes.OptionStatus_Normal,
		OperateUid:  0,      //
		OperateUser: "init", //
		CreateAt:    now,    //
		UpdateAt:    now,    //
	})
	dao.Table.WithContext(context.Background()).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&modelpkg.OptionTelegramCmd{
		ChannelCode: xtypes.OfficialChannelCode,        //渠道code
		Cmd:         tgtypes.XTelegramCmd_BuySameRobot, //
		CmdKind:     tgtypes.CmdKind_InlineKeyboard,
		ReplyKeyboardMarkup: types.ReplyKeyboardMarkupDb{
			ResizeKeyboard:  false,
			OneTimeKeyboard: false,
		},
		PictureUrl:  "",                                 //图片
		Type:        tgtypes.RobotMsgTypeText,           //文件类型(1=图片,2=视频)
		Text:        tgtemplate.CustomizeBuySameRobot(), //文本内容
		ParseMode:   tgtypes.ModeMarkdown,
		Status:      xtypes.OptionStatus_Normal,
		OperateUid:  0,      //
		OperateUser: "init", //
		CreateAt:    now,    //
		UpdateAt:    now,    //
	})
}
