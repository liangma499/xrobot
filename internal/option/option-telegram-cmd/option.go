package optionTelegramCmdCfg

import (
	"context"
	"sync"
	"sync/atomic"
	"tron_robot/internal/code"
	optionTelegramCmdDao "tron_robot/internal/dao/option-telegram-cmd"
	"tron_robot/internal/model"

	tgbutton "tron_robot/internal/xtelegram/tg-button"
	tginlinekeyboardbutton "tron_robot/internal/xtelegram/tg-inline-keyboard-button"
	tgkeyboardbutton "tron_robot/internal/xtelegram/tg-keyboard-button"

	tgtypes "tron_robot/internal/xtelegram/tg-types"

	"tron_robot/internal/xtypes"
	"xbase/config"
	"xbase/config/etcd"
	"xbase/errors"
	"xbase/log"
)

const (
	Name = xtypes.OptionPrefix + "option-telegram-cmd-cfg_v3"
	file = xtypes.OptionPrefix + "option-telegram-cmd-cfg_v3.json"
)

type OptionTelegramCmdCfg struct {
	ChannelCode string                   `json:"channelCode"` //渠道code
	Cmd         tgtypes.XTelegramCmd     `json:"cmd"`         //推送方式
	Keyboard    *tgbutton.TelegramButton `json:"keyboard"`
	PictureUrl  string                   `json:"picture_url"` //图片
	Type        tgtypes.RobotMsgType     `json:"type"`        //文件类型(1=图片,2=视频)
	Text        string                   `json:"text"`        //文本内容
	ParseMode   tgtypes.ParseMode        `json:"parseMode"`   //解码方式
	Status      xtypes.OptionStatus      `json:"status"`      //状态( 1启用,2禁用)
}

type Options struct {
	Opts map[string]map[tgtypes.XTelegramCmd]*OptionTelegramCmdCfg `json:"option-telegram-cmd-cfg"` // 用户币种配置
}
type columns struct {
	TelegramRobotCmdOpts string
}

var (
	opts    atomic.Value
	once    sync.Once
	Columns = &columns{
		TelegramRobotCmdOpts: "option-telegram-cmd-cfg", // 最小提现金额
	}
)

// GetOpts 读取配置项
func GetOpts() *Options {
	once.Do(func() {
		o, err := doLoadOpts()
		if err != nil {
			log.Fatalf("option-telegram-cmd-cfg:%v", err)
		}
		config.Watch(func(names ...string) {
			if o, err := doLoadOpts(); err == nil {
				opts.Store(o)
			}
		}, Name)

		opts.Store(o)
	})
	data, ok := opts.Load().(*Options)
	if !ok {
		return nil
	}
	return data

}

// SetOpts 设置配置项
func SetOpts(ctx context.Context, operate xtypes.OptionOperate, keys ...int64) error {
	opts := &Options{
		Opts: make(map[string]map[tgtypes.XTelegramCmd]*OptionTelegramCmdCfg),
	}

	cmds, err := optionTelegramCmdDao.Instance().FindMany(ctx, func(cols *optionTelegramCmdDao.Columns) any {
		return map[string]any{
			cols.Status: xtypes.OptionStatus_Normal,
		}
	}, nil, nil)
	if err != nil {
		return errors.NewError(code.LoadOptionErr, err)
	}

	if len(cmds) > 0 {
		for _, item := range cmds {
			if _, ok := opts.Opts[item.ChannelCode]; !ok {
				opts.Opts[item.ChannelCode] = make(map[tgtypes.XTelegramCmd]*OptionTelegramCmdCfg)
			}
			opts.Opts[item.ChannelCode][item.Cmd] = &OptionTelegramCmdCfg{
				ChannelCode: item.ChannelCode, //渠道code
				Cmd:         item.Cmd,         //推送方式
				Keyboard:    doTelegramButton(ctx, item),
				PictureUrl:  item.PictureUrl, //图片
				Type:        item.Type,       //文件类型(1=图片,2=视频)
				Text:        item.Text,       //文本内容
				ParseMode:   item.ParseMode,  //编码方式
				Status:      item.Status,     //状态( 1启用,2禁用)

			}

		}
	}

	return config.Store(ctx, etcd.Name, file, opts, true)
}
func doTelegramButton(ctx context.Context, cmd *model.OptionTelegramCmd) *tgbutton.TelegramButton {
	if cmd == nil {
		return nil
	}
	switch cmd.CmdKind {
	case tgtypes.CmdKind_InlineKeyboard:
		{
			return tgbutton.NewTelegramButton(cmd.CmdKind, tginlinekeyboardbutton.TelegramInlineKeyBoardbutton(ctx, cmd))
		}
	case tgtypes.CmdKind_KeyboardMarkup:
		{
			return tgbutton.NewTelegramButton(cmd.CmdKind, tgkeyboardbutton.TelegramKeyBoardbutton(ctx, cmd))
		}
	}
	return nil
}

// HasOpts 是否有配置项
func HasOpts() bool {
	return config.Has(Name)
}

// 加载配置项
// 加载配置项
func doLoadOpts() (*Options, error) {
	o := &Options{
		Opts: make(map[string]map[tgtypes.XTelegramCmd]*OptionTelegramCmdCfg),
	}

	err := config.Get(Name).Scan(o)
	if err != nil {
		return nil, err
	}
	//log.Warnf("doLoadOpts-telegramrobotcmdopt:%#v", o)
	return o, nil
}
func GetChanCodeCmd(channelCode string, cmd tgtypes.XTelegramCmd) *OptionTelegramCmdCfg {
	opts := GetOpts()
	if opts == nil {
		return nil
	}
	opt := opts.Opts
	if opt == nil {
		return nil
	}
	channcfg, ok := opt[channelCode]
	if !ok {
		return nil
	}
	if channcfg == nil {
		return nil
	}
	cfg, ok := channcfg[cmd]
	if !ok {
		return nil
	}
	if cfg.Status != xtypes.OptionStatus_Normal {
		return nil
	}
	return cfg
}
