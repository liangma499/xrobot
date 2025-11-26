package paymentamountuser

import (
	"context"
	"fmt"
	"sync"
	"xbase/errors"
	"xbase/log"
	"xbase/utils/xconv"
	"xrobot/internal/code"
	mysqlimp "xrobot/internal/component/mysql/mysql-default"
	redisCryptoCurrencies "xrobot/internal/component/redis/redis-crypto-currencies"
	"xrobot/internal/dao/payment-amount-user/internal"
	modelpkg "xrobot/internal/model"
	"xrobot/internal/xtypes"

	"github.com/shopspring/decimal"
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

type PaymentAmountUser struct {
	*internal.PaymentAmountUser
	mu sync.Mutex
}

func NewPaymentAmountUser(db *gorm.DB) *PaymentAmountUser {
	return &PaymentAmountUser{PaymentAmountUser: internal.NewPaymentAmountUser(db)}
}

var (
	once     sync.Once
	instance *PaymentAmountUser
)

func Instance() *PaymentAmountUser {
	once.Do(func() {
		instance = NewPaymentAmountUser(mysqlimp.Instance())
	})
	return instance
}
func (dao *PaymentAmountUser) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.PaymentAmountUser{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (dao *PaymentAmountUser) GetUniqueAmount(currency xtypes.Currency, amount decimal.Decimal) (decimal.Decimal, error) {
	dao.mu.Lock()
	defer dao.mu.Unlock()
	key := fmt.Sprintf(xtypes.UniqueAmountKey, currency.String())
	for i := 1; i < 1000; i++ {
		amountNew := amount.Add(decimal.NewFromInt32(xconv.Int32(i)).Div(decimal.NewFromFloat(1000)))
		if count, err := redisCryptoCurrencies.Instance().SAdd(context.Background(), key, amountNew.String()).Result(); err != nil {
			log.Errorf("%v", err)
			return decimal.Zero, errors.NewError(code.NotUniqueAmount, err)
		} else if count > 0 {
			return amountNew, nil
		}
	}

	return decimal.Zero, errors.NewError(code.NotUniqueAmount)
}
func (dao *PaymentAmountUser) ClearAmount(currency xtypes.Currency, amount string, orderID string) error {
	dao.mu.Lock()
	defer dao.mu.Unlock()
	err := dao.doClearAmount(currency, amount)
	if err != nil {
		return err
	}
	if orderID != "" {
		list, _ := dao.FindMany(context.Background(), func(cols *internal.Columns) any {
			return clause.Eq{
				Column: cols.OrderID,
				Value:  orderID,
			}
		}, nil, nil)
		if list != nil {
			if len(list) > 0 {
				for _, item := range list {
					dao.doClearAmount(item.Currency, item.Amount)
				}
			}
		}

	}
	return nil
}
func (dao *PaymentAmountUser) ClearAmountByOrderID(orderID string) error {
	dao.mu.Lock()
	defer dao.mu.Unlock()

	list, _ := dao.FindMany(context.Background(), func(cols *internal.Columns) any {
		return clause.Eq{
			Column: cols.OrderID,
			Value:  orderID,
		}
	}, nil, nil)
	if list != nil {
		if len(list) > 0 {
			for _, item := range list {
				dao.doClearAmount(item.Currency, item.Amount)
			}
		}
	}

	return nil
}
func (dao *PaymentAmountUser) doClearAmount(currency xtypes.Currency, amount string) error {

	ctx := context.Background()

	_, err := dao.Delete(ctx, func(cols *internal.Columns) any {
		return clause.And(clause.Eq{
			Column: cols.Amount,
			Value:  amount,
		}, clause.Eq{
			Column: cols.Currency,
			Value:  currency.String(),
		})
	})
	if err != nil {
		log.Errorf("%v", err)
		return err
	}
	key := fmt.Sprintf(xtypes.UniqueAmountKey, currency.String())
	_, err = redisCryptoCurrencies.Instance().SRem(ctx, key, amount).Result()
	if err != nil {
		log.Errorf("%v", err)
		return err
	}

	return nil
}

func (dao *PaymentAmountUser) AmountToUser(currency xtypes.Currency, amount string) (*modelpkg.PaymentAmountUser, error) {
	dao.mu.Lock()
	defer dao.mu.Unlock()
	key := fmt.Sprintf(xtypes.UniqueAmountKey, currency.String())
	ctx := context.Background()

	have, err := redisCryptoCurrencies.Instance().SIsMember(ctx, key, amount).Result()
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	if !have {
		return nil, nil
	}
	return dao.FindOne(ctx, func(cols *internal.Columns) any {
		return clause.And(clause.Eq{
			Column: cols.Amount,
			Value:  amount,
		}, clause.Eq{
			Column: cols.Currency,
			Value:  currency.String(),
		})
	})
}
