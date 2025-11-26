package optionBaseConfigCfg

import (
	"context"
	"sync"
	"sync/atomic"
	"xbase/config"
	"xbase/config/etcd"
	"xbase/errors"
	"xbase/log"
	"xrobot/internal/code"
	optionBaseConfigDao "xrobot/internal/dao/option-base-config"
	"xrobot/internal/xtypes"
)

const (
	Name = xtypes.OptionPrefix + "option-base-config"
	file = xtypes.OptionPrefix + "tron-robot-option-base-config.json"
)

type Options struct {
	Opts map[string]string `json:"option-base-config"` // 税率映射
}

var (
	opts atomic.Value
	once sync.Once
)

// GetOpts 读取配置项
func GetOpts() *Options {
	once.Do(func() {
		o, err := doLoadOpts()
		if err != nil {
			log.Fatalf("currencyRateOpt:%v", err)
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

// HasOpts 是否有配置项
func HasOpts() bool {
	return config.Has(Name)
}

// 加载配置项
func doLoadOpts() (*Options, error) {
	o := &Options{
		Opts: make(map[string]string),
	}
	if err := config.Get(Name).Scan(o); err != nil {
		log.Warnf("baseConfig:%v", err)
	}
	//log.Warnf("baseConfig:%#v", o)
	return o, nil
}

func GetValue(key string) string {
	opts := GetOpts()
	if opts == nil {
		return ""
	}
	if opts.Opts == nil {
		return ""
	}
	if val, ok := opts.Opts[key]; ok {
		return val
	}
	return ""
}
func SetOpts(ctx context.Context, operate xtypes.OptionOperate, keys ...string) error {
	opts := GetOpts()
	if opts == nil {
		opts = &Options{
			Opts: make(map[string]string),
		}
	} else {
		if opts.Opts == nil {
			opts.Opts = make(map[string]string)
		}
	}
	if operate == xtypes.OptionOperate_LoadAll {
		opts.Opts = make(map[string]string)
		cfgs, err := optionBaseConfigDao.Instance().FindMany(ctx, func(cols *optionBaseConfigDao.Columns) any {
			return map[string]any{
				cols.Status: xtypes.OptionStatus_Normal,
			}
		}, nil, nil)
		if err != nil {
			return errors.NewError(code.LoadOptionErr, err)
		}

		for _, item := range cfgs {
			opts.Opts[item.Key] = item.Value
		}

	} else if operate == xtypes.OptionOperate_Modify {

		if keys != nil {
			cfgs, err := optionBaseConfigDao.Instance().FindMany(ctx, func(cols *optionBaseConfigDao.Columns) any {
				return map[string]any{
					cols.Key: keys,
				}
			}, nil, nil)
			if err != nil {
				return errors.NewError(code.ModifyOptionErr, err)
			}

			for _, item := range cfgs {
				if item.Status == xtypes.OptionStatus_Normal {
					delete(opts.Opts, item.Key)
				} else {
					opts.Opts[item.Key] = item.Value
				}

			}
		}

	} else if operate == xtypes.OptionOperate_Delete {

		for _, item := range keys {
			delete(opts.Opts, item)
		}

	}

	//opts.Store(optsData)
	return config.Store(ctx, etcd.Name, file, opts)
}
