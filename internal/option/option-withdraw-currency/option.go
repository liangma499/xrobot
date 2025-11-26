package optionWithdrawCurrency

import (
	"context"
	"strings"
	"sync"
	"sync/atomic"
	"xbase/config"
	"xbase/config/etcd"
	"xbase/errors"
	"xbase/log"
	"xrobot/internal/code"
	optionCurrencyDao "xrobot/internal/dao/option-currency"
	optionWithdrawCurrencyDao "xrobot/internal/dao/option-withdraw-currency"
	"xrobot/internal/model"
	"xrobot/internal/xtypes"
)

const (
	Name = xtypes.OptionPrefix + "option-withdraw-currency-cfg"
	file = xtypes.OptionPrefix + "option-withdraw-currency-cfg.json"
)

type Options struct {
	Opts map[string]map[xtypes.Currency]map[string]*model.OptionWithdrawCurrency `json:"option-withdraw-currency-cfg"` // 归集费用映射
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
	opts := &Options{
		Opts: make(map[string]map[xtypes.Currency]map[string]*model.OptionWithdrawCurrency),
	}
	if err := config.Get(Name).Scan(opts); err != nil {
		log.Warnf("baseConfig:%v", err)
	}
	//log.Warnf("baseConfig:%#v", o)
	return opts, nil
}

func GetOptionsByCurrency(whithdrawType string) map[xtypes.Currency]map[string]*model.OptionWithdrawCurrency {
	opts := GetOpts()
	if opts == nil {
		return nil
	}
	if opts.Opts == nil {
		return nil
	}

	if data, ok := opts.Opts[whithdrawType]; ok {
		return data

	}
	return nil
}
func GetOptionByCardCurrencyChannel(whithdrawType string, currency xtypes.Currency, channel string) *model.OptionWithdrawCurrency {
	opts := GetOpts()
	if opts == nil {
		return nil
	}
	if opts.Opts == nil {
		return nil
	}

	if whithdrawData, ok := opts.Opts[whithdrawType]; ok {
		if whithdrawData == nil {
			return nil
		}
		currency = currency.ToUpper()

		if data, ok := whithdrawData[currency]; ok {
			if data == nil {
				return nil
			}
			channel := strings.ToLower(channel)
			if rst, ok := data[channel]; ok {
				return rst
			}
		}
	}

	return nil
}
func SetOpts(ctx context.Context) error {
	currencyCfgs, err := optionCurrencyDao.Instance().FindMany(ctx, func(cols *optionCurrencyDao.Columns) any {
		return map[string]any{
			cols.Status: xtypes.OptionStatus_Normal,
		}
	}, nil, nil)
	if err != nil {
		return errors.NewError(code.LoadOptionErr, err)
	}
	currency := make([]xtypes.Currency, 0)
	for _, item := range currencyCfgs {
		if item.Status != xtypes.OptionStatus_Normal {
			continue
		}
		currency = append(currency, item.Currency.ToUpper())
	}
	if len(currency) == 0 {
		return nil
	}
	opts := GetOpts()
	if opts == nil {
		opts = &Options{
			Opts: make(map[string]map[xtypes.Currency]map[string]*model.OptionWithdrawCurrency),
		}
	}

	opts.Opts = make(map[string]map[xtypes.Currency]map[string]*model.OptionWithdrawCurrency)

	cfgs, err := optionWithdrawCurrencyDao.Instance().FindMany(ctx, func(cols *optionWithdrawCurrencyDao.Columns) any {
		return map[string]any{
			cols.Status:   xtypes.OptionStatus_Normal,
			cols.Currency: currency,
		}
	}, nil, nil)
	if err != nil {
		return errors.NewError(code.LoadOptionErr, err)
	}

	for _, item := range cfgs {
		if item.Status != xtypes.OptionStatus_Normal {
			continue
		}
		if _, ok := opts.Opts[item.WithdrawType]; !ok {
			opts.Opts[item.WithdrawType] = make(map[xtypes.Currency]map[string]*model.OptionWithdrawCurrency)
		}
		//检查货币是否关了
		currency := item.Currency.ToUpper()
		//检查货币渠道是否关了
		if _, ok := opts.Opts[item.WithdrawType][currency]; !ok {
			opts.Opts[item.WithdrawType][currency] = make(map[string]*model.OptionWithdrawCurrency)
		}
		channel := strings.ToLower(item.Channel)
		opts.Opts[item.WithdrawType][currency][channel] = item
	}

	return config.Store(ctx, etcd.Name, file, opts, true)
}
