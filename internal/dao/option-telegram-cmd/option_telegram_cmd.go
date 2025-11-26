package optiontelegramcmd

import (
	"context"
	"sync"
	mysqlimp "tron_robot/internal/component/mysql/mysql-default"
	"tron_robot/internal/dao/option-telegram-cmd/internal"
	modelpkg "tron_robot/internal/model"
	"tron_robot/internal/xtelegram/telegram/types"
	tgtemplate "tron_robot/internal/xtelegram/tg-template"
	tgtypes "tron_robot/internal/xtelegram/tg-types"
	"tron_robot/internal/xtypes"
	"xbase/utils/xtime"

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
