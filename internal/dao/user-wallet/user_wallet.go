package userwallet

import (
	"context"
	"fmt"
	"sync"
	"tron_robot/internal/code"
	mysqlimp "tron_robot/internal/component/mysql/mysql-default"
	redisUserWallet "tron_robot/internal/component/redis/redis-userwallet"
	"tron_robot/internal/dao/user-wallet/internal"

	modelpkg "tron_robot/internal/model"
	"tron_robot/internal/xtypes"
	"xbase/errors"
	"xbase/log"
	"xbase/task"
	"xbase/utils/xconv"

	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type (
	Columns    = internal.Columns
	OrderBy    = internal.OrderBy
	FilterFunc = internal.FilterFunc
	UpdateFunc = internal.UpdateFunc
	ColumnFunc = internal.ColumnFunc
	OrderFunc  = internal.OrderFunc
)

type UserWallet struct {
	*internal.UserWallet
	incrBalanceScript *redis.Script
}

func NewUserWallet(db *gorm.DB) *UserWallet {
	return &UserWallet{
		UserWallet:        internal.NewUserWallet(db),
		incrBalanceScript: redis.NewScript(incrBalanceScript),
	}
}

var (
	once     sync.Once
	instance *UserWallet
)

func Instance() *UserWallet {
	once.Do(func() {
		instance = NewUserWallet(mysqlimp.Instance())
	})
	return instance
}
func (dao *UserWallet) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.UserWallet{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}
func (dao *UserWallet) updateUserCash(uid int64, currency xtypes.Currency, afterCash, usedCash decimal.Decimal, walletAmountKind xtypes.WalletAmountKind) (int64, error) {

	cash := afterCash.InexactFloat64()
	i, err := dao.Update(context.TODO(), func(cols *internal.Columns) any {
		return map[string]any{
			cols.UID:              uid,
			cols.Currency:         currency,
			cols.WalletAmountKind: walletAmountKind,
		}

	}, func(cols *internal.Columns) any {
		return map[string]any{
			cols.Cash: cash,
			cols.Used: usedCash.InexactFloat64(),
		}
	})
	if err != nil {
		log.Warnf("%d:%s:%f err:%v", uid, currency, cash, err)
	}
	return i, err
}

// 初始化资产
func (dao *UserWallet) InitBalance(ctx context.Context, args *InitBalanceArgs, walletAmountKind xtypes.WalletAmountKind) (*BalanceInfo, error) {
	//货币基础信息

	wallet, err := dao.FindOne(ctx, func(cols *internal.Columns) any {
		return map[string]any{
			cols.UID:              args.UID,
			cols.Currency:         args.Currency.String(),
			cols.WalletAmountKind: walletAmountKind,
		}
	})
	if err != nil {
		log.Errorf("find user's wallet balance failed, args = %s err = %v", xconv.Json(args), err)
		return nil, errors.NewError(err, code.InternalError)
	}

	if wallet == nil {

		wallet = &modelpkg.UserWallet{
			UID:              args.UID,
			Currency:         args.Currency,
			Def:              args.Def,
			Cash:             xtypes.KeepDecimal(args.Cash, args.Currency),
			Used:             xtypes.KeepDecimal(args.Used, args.Currency),
			WalletAmountKind: walletAmountKind,
		}

		if _, err = dao.Insert(ctx, wallet); err != nil {
			log.Errorf("insert user's wallet balance failed, args = %s err = %v", xconv.Json(args), err)
			return nil, errors.NewError(err, code.InternalError)
		}
	}

	key := fmt.Sprintf(xtypes.UserWalletBalanceKey, args.UID, args.Currency, walletAmountKind)

	if err = redisUserWallet.Instance().HMSet(ctx, key, map[string]any{
		xtypes.UserWalletFieldid:       wallet.ID,
		xtypes.UserWalletFieldUID:      wallet.UID,
		xtypes.UserWalletFieldCurrency: wallet.Currency.String(),
		xtypes.UserWalletFieldCash:     xtypes.DecimalInt64(wallet.Cash, args.Currency),
		xtypes.UserWalletFieldUsed:     xtypes.DecimalInt64(wallet.Used, args.Currency),
		xtypes.UserWalletFieldDef:      int(wallet.Def),
		xtypes.UserWalletAmountKindKey: int32(walletAmountKind),
	}).Err(); err != nil {
		log.Errorf("hmset user's wallet balance failed, args = %s err = %v", xconv.Json(args), err)
		return nil, errors.NewError(err, code.InternalError)
	}

	return &BalanceInfo{
		Currency:   wallet.Currency,
		Cash:       wallet.Cash.InexactFloat64(),
		Used:       wallet.Used.InexactFloat64(),
		IsDefault:  wallet.Def == xtypes.WalletDefaultYes,
		AmountKind: walletAmountKind,
	}, nil
}

func (dao *UserWallet) DecrBalance(ctx context.Context, uid int64, changCage decimal.Decimal, currency xtypes.Currency, walletAmountKind xtypes.WalletAmountKind) (*ChangeInfo, error) {

	if uid <= 0 {
		return nil, errors.NewError(code.InvalidArgument)
	}

	key := fmt.Sprintf(xtypes.UserWalletBalanceKey, uid, currency.String(), walletAmountKind)

	keyDecimal, _ := xtypes.KeepDecimalDigits(currency)

	if changCage.GreaterThan(decimal.Zero) {
		changCage = changCage.Neg()
	}
	rst, err := dao.incrBalanceScript.Run(ctx, redisUserWallet.Instance(), []string{key, keyDecimal.String()}, changCage.InexactFloat64()).StringSlice()
	if err != nil {
		log.Errorf("decr user's wallet balance failed, uid = %d changCage:%s currency:%s rst = %v err = %v",
			uid, changCage.String(), currency.String(), rst, err)
		return nil, errors.NewError(err, code.InternalError)
	}

	//baseInfo
	switch rst[0] {
	case "1": // 未初始化
		_, err = dao.InitBalance(ctx, &InitBalanceArgs{
			UID:      uid,
			Currency: currency,
		}, walletAmountKind)
		if err != nil {
			return nil, err
		}

		return dao.DecrBalance(ctx, uid, changCage, currency, walletAmountKind)
	case "2": // 余额不足
		return nil, errors.NewError(code.BalanceInsufficient)
	case "-1", "-2", "-3":
		return nil, errors.NewError(code.InvalidArgument, rst[0])
	}
	afterCash := decimal.NewFromFloat(xconv.Float64(rst[2]))
	used := decimal.NewFromFloat(xconv.Float64(rst[5]))
	task.AddTask(func() {
		dao.updateUserCash(uid, currency, afterCash, used, walletAmountKind)
	})
	return &ChangeInfo{
		BeforeCash: xconv.Float64(rst[1]),
		AfterCash:  afterCash.InexactFloat64(),
		ChangeCash: xconv.Float64(rst[3]),
		Def:        xconv.Int(rst[4]),
		Used:       used.InexactFloat64(),
	}, nil
}

func (dao *UserWallet) IncrBalance(ctx context.Context, uid int64, changCage decimal.Decimal, currency xtypes.Currency, walletAmountKind xtypes.WalletAmountKind) (*ChangeInfo, error) {

	if uid <= 0 {
		return nil, errors.NewError(code.InvalidArgument)
	}
	if changCage.LessThan(decimal.Zero) {
		return nil, errors.NewError(code.InvalidArgument)
	}
	key := fmt.Sprintf(xtypes.UserWalletBalanceKey, uid, currency.String(), walletAmountKind)

	keyDecimal, _ := xtypes.KeepDecimalDigits(currency)

	rst, err := dao.incrBalanceScript.Run(ctx, redisUserWallet.Instance(), []string{key, keyDecimal.String()}, changCage.InexactFloat64()).StringSlice()
	if err != nil {
		log.Errorf("incr user's wallet balance failed, uid = %d changCage:%s currency:%s rst = %v err = %v",
			uid, changCage.String(), currency.String(), rst, err)
		return nil, errors.NewError(err, code.InternalError)
	}

	//baseInfo
	switch rst[0] {
	case "1": // 未初始化
		_, err = dao.InitBalance(ctx, &InitBalanceArgs{
			UID:      uid,
			Currency: currency,
		}, walletAmountKind)
		if err != nil {
			return nil, err
		}

		return dao.IncrBalance(ctx, uid, changCage, currency, walletAmountKind)
	case "2": // 余额不足
		return nil, errors.NewError(code.BalanceInsufficient)
	case "-1", "-2", "-3":
		return nil, errors.NewError(code.InvalidArgument, rst[0])
	}
	afterCash := decimal.NewFromFloat(xconv.Float64(rst[2]))
	used := decimal.NewFromFloat(xconv.Float64(rst[5]))
	task.AddTask(func() {
		dao.updateUserCash(uid, currency, afterCash, used, walletAmountKind)
	})
	return &ChangeInfo{
		BeforeCash: xconv.Float64(rst[1]),
		AfterCash:  afterCash.InexactFloat64(),
		ChangeCash: xconv.Float64(rst[3]),
		Def:        xconv.Int(rst[4]),
		Used:       used.InexactFloat64(),
	}, nil
}

// 获取账户余额
func (dao *UserWallet) GetBalance(ctx context.Context, uid int64, currency xtypes.Currency, walletAmountKind xtypes.WalletAmountKind) (*BalanceInfo, error) {

	wallet := dao.doGetBalanceFormRedis(ctx, uid, currency, walletAmountKind)

	if wallet.UID != 0 {
		return &BalanceInfo{
			Currency:   currency,
			Cash:       xtypes.DecimalFloat64(wallet.Cash, currency).InexactFloat64(),
			Used:       xtypes.DecimalFloat64(wallet.Used, currency).InexactFloat64(),
			IsDefault:  wallet.Def == xtypes.WalletDefaultYes,
			AmountKind: walletAmountKind,
		}, nil
	}

	return dao.InitBalance(ctx, &InitBalanceArgs{UID: uid, Currency: currency}, walletAmountKind)
}

// 获取账户余额
func (dao *UserWallet) doGetBalanceFormRedis(ctx context.Context, uid int64, currency xtypes.Currency, walletAmountKind xtypes.WalletAmountKind) *modelpkg.UserWallet {
	var (
		key    = fmt.Sprintf(xtypes.UserWalletBalanceKey, uid, currency, walletAmountKind)
		wallet = &modelpkg.UserWallet{}
	)

	walletMap, err := redisUserWallet.Instance().HGetAll(ctx, key).Result()
	if err != nil {

		log.Errorf("get user's wallet balance failed, uid = %d currency = %s walletAmountKind:%d  err = %v",
			uid, currency, walletAmountKind, err)
		return wallet
	}
	if len(walletMap) == 0 {
		return wallet
	}
	if ret, ok := walletMap[xtypes.UserWalletFieldid]; ok {
		wallet.ID = xconv.Int64(ret)
	} else {
		log.Errorf("get user's wallet balance cash failed, uid = %d currency = %s walletAmountKind:%d  not ID",
			uid, currency, walletAmountKind)
		return &modelpkg.UserWallet{}
	}

	if ret, ok := walletMap[xtypes.UserWalletFieldUID]; ok {
		wallet.UID = xconv.Int64(ret)
	} else {
		log.Errorf("get user's wallet balance cash failed, uid = %d currency = %s walletAmountKind:%d  not UID",
			uid, currency, walletAmountKind)
		return &modelpkg.UserWallet{}
	}
	if ret, ok := walletMap[xtypes.UserWalletFieldCurrency]; ok {
		wallet.Currency = xtypes.Currency(ret)
	} else {
		log.Errorf("get user's wallet balance cash failed, uid = %d currency = %s walletAmountKind:%d  not Currency",
			uid, currency, walletAmountKind)
		return &modelpkg.UserWallet{}
	}
	if ret, ok := walletMap[xtypes.UserWalletFieldCash]; ok {
		wallet.Cash, err = decimal.NewFromString(ret)
		if err != nil {
			log.Errorf("get user's wallet balance cash failed, uid = %d currency = %s walletAmountKind:%d  err = %v",
				uid, currency, walletAmountKind, err)
			return &modelpkg.UserWallet{}
		}
	} else {
		log.Errorf("get user's wallet balance cash failed, uid = %d currency = %s walletAmountKind:%d  not Cash",
			uid, currency, walletAmountKind)
		return &modelpkg.UserWallet{}
	}
	if ret, ok := walletMap[xtypes.UserWalletFieldUsed]; ok {
		wallet.Used, err = decimal.NewFromString(ret)
		if err != nil {
			log.Errorf("get user's wallet balance cash failed, uid = %d currency = %s walletAmountKind:%d  err = %v",
				uid, currency, walletAmountKind, err)
			return &modelpkg.UserWallet{}
		}
	} else {
		log.Errorf("get user's wallet balance cash failed, uid = %d currency = %s walletAmountKind:%d  not Used",
			uid, currency, walletAmountKind)
		return &modelpkg.UserWallet{}
	}
	if ret, ok := walletMap[xtypes.UserWalletFieldDef]; ok {
		wallet.Def = xtypes.WalletDefaultStatus(xconv.Int(ret))
	}
	if ret, ok := walletMap[xtypes.UserWalletAmountKindKey]; ok {
		wallet.WalletAmountKind = xtypes.WalletAmountKind(xconv.Int(ret))
	} else {
		log.Errorf("get user's wallet balance cash failed, uid = %d currency = %s walletAmountKind:%d  not WalletAmountKind",
			uid, currency, walletAmountKind)
		return &modelpkg.UserWallet{}
	}
	return wallet
}

// 设置钱包余额默认状态
func (dao *UserWallet) SetBalanceDefaultStatus(ctx context.Context, uid int64, currency xtypes.Currency, def xtypes.WalletDefaultStatus, walletAmountKind xtypes.WalletAmountKind) error {
	key := fmt.Sprintf(xtypes.UserWalletBalanceKey, uid, currency.String(), walletAmountKind)

	err := redisUserWallet.Instance().HSet(ctx, key, xtypes.UserWalletFieldDef, int(def)).Err()
	if err != nil {
		log.Errorf("set default currency of user's wallet failed, uid = %d currency = %s err = %v", uid, currency.String(), err)
		return errors.NewError(err, code.InternalError)
	}

	_, err = dao.Update(ctx, func(cols *internal.Columns) any {
		return map[string]any{
			cols.UID:              uid,
			cols.Currency:         currency.String(),
			cols.WalletAmountKind: walletAmountKind,
		}
	}, func(cols *internal.Columns) any {
		return map[string]any{
			cols.Def: def,
		}
	})
	if err != nil {
		log.Errorf("update default currency of user's wallet failed, uid = %d currency = %s err = %v", uid, currency.String(), err)
		return errors.NewError(err, code.InternalError)
	}

	return nil
}
