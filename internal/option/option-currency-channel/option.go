package optionCurrencyChannelCfg

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
	optionCurrencyChannelDao "xrobot/internal/dao/option-currency-channel"
	"xrobot/internal/xtypes"

	"github.com/shopspring/decimal"
)

const (
	Name = xtypes.OptionPrefix + "option-currency-channel-cfg"
	file = xtypes.OptionPrefix + "option-currency-channel-cfg.json"
)

type OptionCurrencyChannelCfg struct {
	Currency     xtypes.Currency     `json:"currency"`
	Channel      string              `json:"channel"`
	ThirdKey     string              `json:"third_key"`
	Minutes      string              `json:"minutes"`
	Confirmation string              `json:"confirmation"`
	Second       string              `json:"second"`
	CollectFee   decimal.Decimal     `json:"collect_fee"`
	GasFee       decimal.Decimal     `json:"gasFee"`
	Sort         int                 `json:"sort"` //
	Memo         string              `json:"memo"` //
	Status       xtypes.OptionStatus `json:"status"`
}

type Options struct {
	Opts map[xtypes.Currency]map[string]*OptionCurrencyChannelCfg `json:"option-currency-channel-cfg"` // 归集费用映射
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
		Opts: make(map[xtypes.Currency]map[string]*OptionCurrencyChannelCfg),
	}
	if err := config.Get(Name).Scan(opts); err != nil {
		log.Warnf("baseConfig:%v", err)
	}
	//log.Warnf("baseConfig:%#v", o)
	return opts, nil
}

func GetCollectFeeByCurrency(currency xtypes.Currency) map[string]string {
	opts := GetOpts()
	if opts == nil {
		return nil
	}
	if opts.Opts == nil {
		return nil
	}

	currency = currency.ToUpper()

	if data, ok := opts.Opts[currency]; ok {
		rst := make(map[string]string)
		for _, item := range data {
			if item == nil {
				continue
			}
			key := strings.ToLower(item.ThirdKey)
			rst[key] = item.CollectFee.String()
		}
		return rst
	}
	return nil
}
func GetOptionsByCurrency(currency xtypes.Currency) map[string]*OptionCurrencyChannelCfg {
	opts := GetOpts()
	if opts == nil {
		return nil
	}
	if opts.Opts == nil {
		return nil
	}

	currency = currency.ToUpper()

	if data, ok := opts.Opts[currency]; ok {
		return data
	}
	return nil
}
func GetOptionByCardCurrencyChannel(currency xtypes.Currency, channel string) *OptionCurrencyChannelCfg {
	opts := GetOpts()
	if opts == nil {
		return nil
	}
	if opts.Opts == nil {
		return nil
	}

	currency = currency.ToUpper()

	if data, ok := opts.Opts[currency]; ok {
		if data == nil {
			return nil
		}
		if rst, ok := data[channel]; ok {
			return rst
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
		currency = append(currency, item.Currency)
	}
	if len(currency) == 0 {
		return nil
	}
	opts := GetOpts()
	if opts == nil {
		opts = &Options{
			Opts: make(map[xtypes.Currency]map[string]*OptionCurrencyChannelCfg),
		}
	}

	opts.Opts = make(map[xtypes.Currency]map[string]*OptionCurrencyChannelCfg)

	cfgs, err := optionCurrencyChannelDao.Instance().FindMany(ctx, func(cols *optionCurrencyChannelDao.Columns) any {
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
		//检查货币是否关了
		currency := item.Currency.ToUpper()
		//检查货币渠道是否关了
		channel := strings.ToLower(item.Channel)
		if _, ok := opts.Opts[currency]; !ok {
			opts.Opts[currency] = make(map[string]*OptionCurrencyChannelCfg)
		}
		opts.Opts[currency][channel] = &OptionCurrencyChannelCfg{
			Currency:     item.Currency,
			Channel:      item.Channel,
			ThirdKey:     item.ThirdKey,
			Minutes:      item.Minutes,
			Confirmation: item.Confirmation,
			Second:       item.Second,
			CollectFee:   item.CollectFee,
			GasFee:       item.GasFee,
			Sort:         item.Sort, //
			Memo:         item.Memo, //
			Status:       item.Status,
		}
	}

	return config.Store(ctx, etcd.Name, file, opts, true)
}
