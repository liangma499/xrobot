package optionCurrencyCfg

import (
	"context"
	"sync"
	"sync/atomic"
	"tron_robot/internal/code"
	optionCurrencyDao "tron_robot/internal/dao/option-currency"
	"tron_robot/internal/model"
	"tron_robot/internal/xtypes"
	"xbase/config"
	"xbase/config/etcd"
	"xbase/errors"
	"xbase/log"
)

const (
	Name = xtypes.OptionPrefix + "option-currency-cfg"
	file = xtypes.OptionPrefix + "option-currency-cfg.json"
)

type columns struct {
	OptionCurrencyKey string
}

type Options struct {
	Opts map[xtypes.Currency]*model.OptionCurrency `json:"option-currency-cfg"` // 归集费用映射
}

var (
	opts    atomic.Value
	once    sync.Once
	Columns = &columns{
		OptionCurrencyKey: "option-currency-cfg", // 验证码邮件
	}
)

// GetOpts 读取配置项
func GetOpts() *Options {
	once.Do(func() {
		o, err := doLoadOpts()
		if err != nil {
			log.Fatalf("option-currency-cfg:%v", err)
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
		Opts: make(map[xtypes.Currency]*model.OptionCurrency),
	}
	if err := config.Get(Name).Scan(o); err != nil {
		log.Warnf("option-currency-cfg:%v", err)
	}
	//log.Warnf("option-currency-cfg:%#v", o)
	return o, nil

}

func GetOptionCurrencyByCurrency(currency xtypes.Currency) (map[xtypes.Currency]*model.OptionCurrency, *model.OptionCurrency) {
	getOpts := GetOpts()
	if getOpts == nil {
		return nil, nil
	}
	if getOpts.Opts == nil {
		return nil, nil
	}

	currency = currency.ToUpper()
	if cfg, ok := getOpts.Opts[currency]; ok && cfg != nil {
		if cfg.Status != xtypes.OptionStatus_Normal {
			return getOpts.Opts, nil
		}
		return getOpts.Opts, cfg
	}
	return getOpts.Opts, nil
}
func GetOptionCurrencyAll() []*model.OptionCurrency {
	getOpts := GetOpts()
	if getOpts == nil {
		return nil
	}
	if getOpts.Opts == nil {
		return nil
	}
	list := make([]*model.OptionCurrency, 0)
	for _, item := range getOpts.Opts {
		if item == nil {
			continue
		}
		if item.Status != xtypes.OptionStatus_Normal {
			continue
		}
		list = append(list, &model.OptionCurrency{
			Currency: item.Currency,
			Url:      item.Url,
			Sort:     item.Sort, //
			Memo:     item.Memo, //
			Status:   item.Status,
		})
	}
	/*
		if len(list) > 0 {
			sort.SliceStable(list, func(i, j int) bool {
				if list[i].Sort != list[j].Sort {
					return list[i].Sort > list[j].Sort
				}
				return strings.Compare(list[i].Currency, list[j].Currency) != 1
			})
		}*/
	return list
}
func SetOpts(ctx context.Context, operate xtypes.OptionOperate, keys ...xtypes.Currency) error {

	getOpts := GetOpts()
	if getOpts == nil {
		getOpts = &Options{
			Opts: make(map[xtypes.Currency]*model.OptionCurrency),
		}
	} else {
		if getOpts.Opts == nil {
			getOpts.Opts = make(map[xtypes.Currency]*model.OptionCurrency)
		}
	}
	if operate == xtypes.OptionOperate_LoadAll {
		getOpts.Opts = make(map[xtypes.Currency]*model.OptionCurrency)
		cfgs, err := optionCurrencyDao.Instance().FindMany(ctx, func(cols *optionCurrencyDao.Columns) any {
			return map[string]any{
				cols.Status: xtypes.OptionStatus_Normal,
			}
		}, nil, nil)
		if err != nil {
			return errors.NewError(code.LoadOptionErr, err)
		}

		for _, item := range cfgs {
			currency := item.Currency.ToUpper()
			getOpts.Opts[currency] = &model.OptionCurrency{
				Currency: currency,
				Url:      item.Url,
				Sort:     item.Sort,
				Memo:     item.Memo,
				Status:   item.Status,
			}
		}

	} else if operate == xtypes.OptionOperate_Modify {

		if keys != nil {
			cfgs, err := optionCurrencyDao.Instance().FindMany(ctx, func(cols *optionCurrencyDao.Columns) any {
				return map[string]any{
					cols.Currency: keys,
				}
			}, nil, nil)
			if err != nil {
				return errors.NewError(code.ModifyOptionErr, err)
			}

			for _, item := range cfgs {
				currency := item.Currency.ToUpper()
				getOpts.Opts[currency] = &model.OptionCurrency{
					Currency: currency,
					Url:      item.Url,
					Sort:     item.Sort,
					Memo:     item.Memo,
					Status:   item.Status,
				}
			}
		}

	} else if operate == xtypes.OptionOperate_Delete {

		for _, item := range keys {
			currency := item.ToUpper()
			if opt, ok := getOpts.Opts[currency]; ok && opt != nil {
				getOpts.Opts[currency].Status = xtypes.OptionStatus_Disable
			}
		}

	}

	return config.Store(ctx, etcd.Name, file, getOpts, true)
}
