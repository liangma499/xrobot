package optioncurrencyRateOpt

import (
	"context"
	"strings"
	"sync"
	"sync/atomic"
	"tron_robot/internal/code"
	"xbase/config"
	"xbase/config/etcd"
	"xbase/errors"
	"xbase/log"

	"github.com/shopspring/decimal"
)

const (
	Name = "currencyRateOpt-v3"
	file = "currencyRateOpt-v3.json"
)

type columns struct {
	RateOptions string
}

type OptionRate struct {
	Prices map[string]decimal.Decimal `json:"prices"`
}
type OptionRateRes struct {
	Opts map[string]OptionRate `json:"rates"`
}

var (
	opts    atomic.Value
	once    sync.Once
	Columns = &columns{
		RateOptions: "currencyRateOpt", // 验证码邮件
	}
)

// GetOpts 读取配置项
func GetOpts() *OptionRateRes {
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

	data, ok := opts.Load().(*OptionRateRes)
	if !ok {
		return nil
	}
	return data
}

// SetOpts 设置配置项
func SetOpts(ctx context.Context, optsData any) error {
	//opts.Store(optsData)
	return config.Store(ctx, etcd.Name, file, optsData, false)
}

// HasOpts 是否有配置项
func HasOpts() bool {
	return config.Has(Name)
}

// 加载配置项
func doLoadOpts() (*OptionRateRes, error) {
	o := &OptionRateRes{
		Opts: make(map[string]OptionRate),
	}

	err := config.Get(Name).Scan(o)
	if err != nil {
		return nil, err
	}
	//log.Warnf("doLoadOpts-currencyRateOpt:%#v", o)
	return o, nil
}

func DoRateToDecimalRateErr(name, toCurrency string) (decimal.Decimal, error) {

	decimalRate := decimal.NewFromInt(1)
	if toCurrency == "" {
		return decimalRate, errors.NewError(code.InvalidArgument)
	}

	toCurrency = strings.ToUpper(toCurrency)
	name = strings.ToUpper(name)
	if name == toCurrency {
		return decimalRate, nil
	}
	opts := GetOpts()
	if opts == nil {
		return decimalRate, errors.NewError(code.OptionNotFound)
	}

	opt, ok := opts.Opts[name]
	if !ok {
		return decimalRate, errors.NewError(code.OptionNotFound)
	}

	if price, ok := opt.Prices[toCurrency]; ok {
		return price, nil
	}

	return decimalRate, errors.NewError(code.OptionNotFound)
}
func DoRateToCurrencyList(name string) map[string]decimal.Decimal {
	var Prices = make(map[string]decimal.Decimal)
	opts := GetOpts()
	if opts == nil {
		return nil
	}
	//计算出1单位的name对应opts里面所有币种的汇率
	baseDecimal := decimal.NewFromFloat(1)
	name = strings.ToUpper(name)
	for k, v := range opts.Opts {
		if price, ok := v.Prices[name]; ok {
			if price.GreaterThan(decimal.Zero) {
				Prices[k] = baseDecimal.Div(price)
			}

		}
	}
	return Prices
}
