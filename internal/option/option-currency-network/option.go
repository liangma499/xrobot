package optionCurrencyNetworkCfg

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"tron_robot/internal/code"
	optionCurrencyNetworkDao "tron_robot/internal/dao/option-currency-network"
	"tron_robot/internal/model"
	"tron_robot/internal/xtypes"
	"xbase/config"
	"xbase/errors"
	"xbase/log"

	"xbase/config/etcd"

	"gorm.io/gorm/clause"
)

const (
	Name = "optionCurrencyNetworkCfg"
	file = "optionCurrencyNetworkCfg.json"
)

type Options struct {
	Opts map[int64]*model.OptionCurrencyNetwork `json:"currencyPlatformOpt"` // 用户币种配置
}
type columns struct {
	Opts string
}

var (
	opts    atomic.Value
	once    sync.Once
	Columns = &columns{
		Opts: "currencyPlatformOpt", // 最小提现金额
	}
)

// GetOpts 读取配置项
func GetOpts() *Options {
	once.Do(func() {
		o, err := doLoadOpts()
		if err != nil {
			log.Fatalf("currencyPlatformOpt:%v", err)
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
	opts := GetOpts()
	if opts.Opts == nil {
		opts.Opts = make(map[int64]*model.OptionCurrencyNetwork)
	}

	list, err := optionCurrencyNetworkDao.Instance().FindMany(ctx, func(cols *optionCurrencyNetworkDao.Columns) any {
		condition := make([]clause.Expression, 0)
		if len(keys) > 0 {
			values := make([]any, 0)
			for _, item := range keys {
				values = append(values, item)
			}
			condition = append(condition, clause.IN{
				Column: cols.ID,
				Values: values,
			})
		}
		condition = append(condition, clause.Eq{
			Column: cols.Status,
			Value:  xtypes.OptionStatus_Normal,
		})
		return clause.And(condition...)
	}, nil, nil)

	if err != nil {
		return errors.NewError(code.LoadOptionErr, err)
	}

	if len(list) > 0 {
		for _, item := range list {
			opts.Opts[item.ID] = item.Clone()
		}
	}

	return config.Store(ctx, etcd.Name, file, opts, true)
}

// HasOpts 是否有配置项
func HasOpts() bool {
	return config.Has(Name)
}

// 加载配置项
// 加载配置项
func doLoadOpts() (*Options, error) {
	o := &Options{
		Opts: make(map[int64]*model.OptionCurrencyNetwork),
	}
	err := config.Get(Name).Scan(o)
	if err != nil {
		return nil, err
	}
	//log.Warnf("doLoadOpts-currencyPlatformOpt:%#v", o)
	return o, nil
}
func GetOptByChannelType(channelType xtypes.NetWorkChannelType) *model.OptionCurrencyNetwork {
	opts := GetOpts()
	if opts == nil {
		return nil
	}
	opt := opts.Opts
	if opt == nil {
		return nil
	}
	if len(opt) == 0 {
		return nil
	}
	for _, item := range opt {
		if item == nil {
			return nil
		}
		if item.Status != xtypes.OptionStatus_Normal {
			return nil
		}
		if item.Type == channelType {
			return item.Clone()
		}
	}
	return nil
}

func GetPrivateKeyCfg(channelType xtypes.NetWorkChannelType, currency xtypes.Currency) (*xtypes.PrivateKeyCfg, error) {

	cfg := GetOptByChannelType(channelType)
	if cfg == nil {
		return nil, fmt.Errorf("channeltype is not found")
	}
	config := cfg.PrivateKey.DesConfig()
	if config == nil {
		return nil, fmt.Errorf("config is not found")
	}
	if config.PrivateKeyCfg == nil {
		return nil, fmt.Errorf("privateKeyCfg is not found")
	}
	if rst, ok := config.PrivateKeyCfg[currency]; ok && rst != nil {
		return rst, nil
	}
	return nil, fmt.Errorf("privateKeyCfg is not found")
}
