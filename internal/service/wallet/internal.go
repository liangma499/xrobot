package wallet

import (
	"context"
	"xrobot/internal/code"

	"xbase/errors"
	"xbase/log"
	"xbase/utils/xconv"
	"xbase/utils/xtime"
	userBaseDao "xrobot/internal/dao/user-base"
	userTradeDao "xrobot/internal/dao/user-trade"
	userWalletDao "xrobot/internal/dao/user-wallet"
	walletevt "xrobot/internal/event/wallet"
	"xrobot/internal/model"
	optionCurrencyCfg "xrobot/internal/option/option-currency"
	"xrobot/internal/service/wallet/pb"
	"xrobot/internal/utils/xstr"
	"xrobot/internal/xtypes"

	"github.com/shopspring/decimal"
	"gorm.io/gorm/clause"
)

func (s *Server) doCancelTrade(ctx context.Context, args *pb.CancelTradeArgs) (*pb.BalanceInfo, *pb.ChangeInfo, error) {
	trade, err := s.doQueryTrade(ctx, &queryTradeArgs{
		TradeNO:   args.TradeNO,
		RelatedID: args.RelatedID,
	})
	if err != nil {
		return nil, nil, err
	}

	if trade.Status != xtypes.TradeStatusTrading {
		return nil, nil, errors.NewError(code.IllegalOperation)
	}

	rows, err := userTradeDao.Instance().Update(ctx, func(cols *userTradeDao.Columns) any {
		return clause.And(
			clause.Eq{
				Column: cols.ID,
				Value:  trade.ID,
			},
			clause.Eq{
				Column: cols.Status,
				Value:  xtypes.TradeStatusTrading,
			})
	}, func(cols *userTradeDao.Columns) any {
		return map[string]any{
			cols.Status:    xtypes.TradeStatusFailed,
			cols.UpdatedAt: xtime.Now(),
		}
	})
	if err != nil {
		log.Errorf("update trade status failed, args = %s err = %v", xconv.Json(args), err)
		return nil, nil, errors.NewError(err, code.InternalError)
	}

	if rows == 0 {
		return nil, nil, errors.NewError(code.IllegalOperation)
	}
	currencyInfo, err := s.doGetCurrencyInfo(trade.Currency)
	if err != nil {
		log.Warnf("currencyInfo is nil, trade = %s err = %v", xconv.Json(trade), err)
		return nil, nil, err
	}
	if currencyInfo == nil {
		return nil, nil, errors.NewError(code.OptionNotFound)
	}
	balance, change, err := s.doReturnBalance(ctx, trade, currencyInfo)
	if err != nil {
		return nil, nil, err
	}

	walletevt.PublishBalanceChange(&walletevt.BalanceChangePayload{
		UID:         trade.UID,
		TradeType:   trade.Type,
		TradeStatus: xtypes.TradeStatusFailed,
		RelatedID:   trade.RelatedID,
		Currency:    currencyInfo.Currency,
		BeforeCash:  trade.BeforeCash,
		AfterCash:   trade.AfterCash,
		ChannelCode: trade.ChannelCode,
		Change: &walletevt.ChangeInfo{
			Cash:         decimal.NewFromFloat(change.Cash),
			Currency:     currencyInfo.Currency,
			AmountKind:   trade.AmountKind,
			PlayCurrency: "",           // 币种
			PlayCash:     decimal.Zero, // 变动的现金
		},
	})

	//reply.Balance = balance
	//reply.Change = change

	return balance, change, nil
}

func (s *Server) doCompleteTrade(ctx context.Context, args *pb.CompleteTradeArgs) error {
	trade, err := s.doQueryTrade(ctx, &queryTradeArgs{
		TradeNO:   args.TradeNO,
		RelatedID: args.RelatedID,
	})
	if err != nil {
		return err
	}

	if trade.Status != xtypes.TradeStatusTrading {
		return errors.NewError(code.IllegalOperation)
	}

	rows, err := userTradeDao.Instance().Update(ctx, func(cols *userTradeDao.Columns) any {
		return clause.And(
			clause.Eq{
				Column: cols.ID,
				Value:  trade.ID,
			},
			clause.Eq{
				Column: cols.Status,
				Value:  xtypes.TradeStatusTrading,
			})
	}, func(cols *userTradeDao.Columns) any {
		return map[string]any{
			cols.Status:    xtypes.TradeStatusSuccess,
			cols.UpdatedAt: xtime.Now(),
		}
	})
	if err != nil {
		log.Errorf("update trade status failed, args = %s err = %v", xconv.Json(args), err)
		return errors.NewError(err, code.InternalError)
	}

	if rows == 0 {
		return errors.NewError(code.IllegalOperation)
	}
	currencyInfo, err := s.doGetCurrencyInfo(trade.Currency)
	if err != nil {
		return err
	}
	if currencyInfo == nil {
		return errors.NewError(code.OptionNotFound)
	}

	walletevt.PublishBalanceChange(&walletevt.BalanceChangePayload{
		UID:         trade.UID,
		TradeType:   trade.Type,
		TradeStatus: xtypes.TradeStatusSuccess,
		RelatedID:   trade.RelatedID,
		Currency:    currencyInfo.Currency,
		BeforeCash:  trade.BeforeCash,
		AfterCash:   trade.AfterCash,
		ChannelCode: trade.ChannelCode,
		Change: &walletevt.ChangeInfo{
			Cash:         trade.ChangeCash,
			Currency:     currencyInfo.Currency,
			AmountKind:   trade.AmountKind,
			PlayCurrency: "",           // 币种
			PlayCash:     decimal.Zero, // 变动的现金
		},
	})

	return nil
}

// 增加账户余额
func (s *Server) doIncrBalance(ctx context.Context, args *incrBalanceArgs, currencyInfo *model.OptionCurrency, user *model.UserBase) (string, *pb.BalanceInfo, *pb.ChangeInfo, error) {

	//icount := baseConfigOpt.GetMaxRegisterIp()
	//log.Warnf("GetMaxRegisterIp:%v", icount)
	chageInfo, err := userWalletDao.Instance().IncrBalance(ctx, user.UID, decimal.NewFromFloat(args.Cash), args.Currency, args.AmountKind)
	if err != nil {
		return "", nil, nil, err
	}

	//beforeCash, afterCash, changeCash, def, otherCash := xconv.Float64(rst[1]), xconv.Float64(rst[2]), xconv.Float64(rst[3]), xconv.Int(rst[4]), xconv.Float64(rst[5])
	changeCashDecimal := decimal.NewFromFloat(chageInfo.ChangeCash)
	afterCashDecimal := decimal.NewFromFloat(chageInfo.AfterCash)
	beforeCashDecimal := decimal.NewFromFloat(chageInfo.BeforeCash)
	no := xstr.SerialNO()
	now := xtime.Now()
	trade := &model.UserTrade{
		NO:          no,
		UID:         args.UID,
		Currency:    currencyInfo.Currency,
		BeforeCash:  beforeCashDecimal,
		AfterCash:   afterCashDecimal,
		ChangeCash:  changeCashDecimal,
		RelatedID:   args.RelatedID,
		Type:        args.TradeType,
		Channel:     args.TradeType.TradeChannel(),
		ChannelCode: args.ChannelCode,
		Status:      args.TradeStatus,
		BetCash:     decimal.NewFromFloat(args.BetAmount),
		UserType:    args.UserType,
		CreatedAt:   now,
		UpdatedAt:   now,
		Taxation:    decimal.NewFromFloat(args.Taxation),
		AmountKind:  args.AmountKind,
	}

	if _, err = userTradeDao.Instance().Insert(ctx, trade); err != nil {
		log.Errorf("incr user's wallet balance failed, args = %s err = %v", xconv.Json(args), err)

		_, _, _ = s.doReturnBalance(ctx, trade, currencyInfo)

		return "", nil, nil, errors.NewError(err, code.InternalError)
	}

	var (
		isDefault = xtypes.WalletDefaultStatus(chageInfo.Def) == xtypes.WalletDefaultYes
	)

	walletevt.PublishBalanceChange(&walletevt.BalanceChangePayload{
		UID:          args.UID,
		TradeType:    args.TradeType,
		TradeStatus:  args.TradeStatus,
		RelatedID:    args.RelatedID,
		Currency:     currencyInfo.Currency,
		BetAmount:    decimal.NewFromFloat(args.BetAmount),
		Taxation:     decimal.NewFromFloat(args.Taxation),
		Rebate:       args.Rebate,
		UserType:     args.UserType,
		BeforeCash:   trade.BeforeCash,
		AfterCash:    trade.AfterCash,
		RegisterTime: user.RegisterAt.Unix(),
		ChannelCode:  user.ChannelCode,
		Change: &walletevt.ChangeInfo{
			Currency:        currencyInfo.Currency,
			Cash:            changeCashDecimal,
			AmountKind:      args.AmountKind,                              // 1真金 2奖金
			PlayCurrency:    xtypes.Currency(args.PlayCurrency).ToUpper(), // 币种
			PlayCash:        decimal.NewFromFloat(args.PlayCash),          // 变动的现金
			UserControlKind: args.UserControlKind,
		},
	})

	return trade.NO, &pb.BalanceInfo{
			Currency:   currencyInfo.Currency.String(),
			Cash:       chageInfo.AfterCash,
			Used:       chageInfo.Used,
			IsDefault:  isDefault,
			AmountKind: int32(args.AmountKind),
		}, &pb.ChangeInfo{
			Currency:   currencyInfo.Currency.String(),
			Cash:       changeCashDecimal.InexactFloat64(),
			AmountKind: int32(args.AmountKind),
		}, nil
}

// 扣减账户余额
func (s *Server) doDecrBalance(ctx context.Context, args *decrBalanceArgs, currencyInfo *model.OptionCurrency, user *model.UserBase) (string, *pb.BalanceInfo, *pb.ChangeInfo, error) {

	if currencyInfo == nil {
		return "", nil, nil, errors.NewError(code.OptionNotFound)
	}
	if args.UID <= 0 {
		return "", nil, nil, errors.NewError(code.InvalidArgument)
	}
	chageInfo, err := userWalletDao.Instance().DecrBalance(ctx, user.UID, decimal.NewFromFloat(args.Cash), args.Currency, args.AmountKind)
	if err != nil {
		return "", nil, nil, err
	}

	//beforeCash, afterCash, changeCash, def, otherCash := xconv.Float64(rst[1]), xconv.Float64(rst[2]), xconv.Float64(rst[3]), xconv.Int(rst[4]), xconv.Float64(rst[5])
	no := xstr.SerialNO()
	changeCashDecimal := decimal.NewFromFloat(chageInfo.ChangeCash)
	afterCashDecimal := decimal.NewFromFloat(chageInfo.AfterCash)
	beforeCashDecimal := decimal.NewFromFloat(chageInfo.BeforeCash)
	now := xtime.Now()
	trade := &model.UserTrade{
		NO:          no,
		UID:         args.UID,
		Currency:    currencyInfo.Currency,
		BeforeCash:  beforeCashDecimal,
		AfterCash:   afterCashDecimal,
		ChangeCash:  changeCashDecimal,
		RelatedID:   args.RelatedID,
		Type:        args.TradeType,
		Channel:     args.TradeType.TradeChannel(),
		ChannelCode: args.ChannelCode,
		Status:      args.TradeStatus,
		BetCash:     decimal.NewFromFloat(args.BetAmount),
		UserType:    args.UserType,
		CreatedAt:   now,
		UpdatedAt:   now,
		Taxation:    decimal.NewFromFloat(args.Taxation),
		AmountKind:  args.AmountKind,
	}

	if _, err = userTradeDao.Instance().Insert(ctx, trade); err != nil {
		log.Errorf("decr trade user's wallet balance failed, args = %s err = %v", xconv.Json(args), err)

		_, _, _ = s.doReturnBalance(ctx, trade, currencyInfo)

		return "", nil, nil, errors.NewError(err, code.InternalError)
	}

	walletevt.PublishBalanceChange(&walletevt.BalanceChangePayload{
		UID:          args.UID,
		Currency:     currencyInfo.Currency,
		TradeType:    args.TradeType,
		TradeStatus:  args.TradeStatus,
		RelatedID:    args.RelatedID,
		BetAmount:    decimal.NewFromFloat(args.BetAmount),
		Rebate:       args.Rebate,
		UserType:     args.UserType,
		BeforeCash:   trade.BeforeCash,
		AfterCash:    trade.AfterCash,
		RegisterTime: user.RegisterAt.Unix(),
		ChannelCode:  user.ChannelCode,
		Change: &walletevt.ChangeInfo{
			Cash:            changeCashDecimal,
			Currency:        currencyInfo.Currency,
			AmountKind:      args.AmountKind,
			PlayCurrency:    xtypes.Currency(args.PlayCurrency).ToUpper(), // 币种
			PlayCash:        decimal.NewFromFloat(args.PlayCash),          // 变动的现金
			UserControlKind: args.UserControlKind,
		},
	})

	return trade.NO, &pb.BalanceInfo{
			Currency:   currencyInfo.Currency.String(),
			Cash:       afterCashDecimal.InexactFloat64(),
			Used:       chageInfo.Used,
			IsDefault:  xtypes.WalletDefaultStatus(chageInfo.Def) == xtypes.WalletDefaultYes,
			AmountKind: int32(args.AmountKind),
		}, &pb.ChangeInfo{
			Currency:   currencyInfo.Currency.String(),
			Cash:       changeCashDecimal.InexactFloat64(),
			AmountKind: int32(args.AmountKind),
		}, nil

}

// 退还到余额
func (s *Server) doReturnBalance(ctx context.Context, trade *model.UserTrade, currencyInfo *model.OptionCurrency) (*pb.BalanceInfo, *pb.ChangeInfo, error) {

	if currencyInfo == nil {
		return nil, nil, errors.NewError(code.OptionNotFound)
	}

	changeCash := trade.ChangeCash
	if changeCash.GreaterThan(decimal.Zero) {
		changeCash = trade.ChangeCash.Neg()
	}

	chageInfo, err := userWalletDao.Instance().IncrBalance(ctx, trade.UID, changeCash, trade.Currency, trade.AmountKind)
	if err != nil {
		return nil, nil, err
	}

	return &pb.BalanceInfo{
			Currency:   trade.Currency.String(),
			Cash:       chageInfo.AfterCash,
			Used:       chageInfo.Used,
			IsDefault:  xtypes.WalletDefaultStatus(chageInfo.Def) == xtypes.WalletDefaultYes,
			AmountKind: int32(trade.AmountKind),
		}, &pb.ChangeInfo{
			Currency:   trade.Currency.String(),
			Cash:       trade.ChangeCash.InexactFloat64(),
			AmountKind: int32(trade.AmountKind),
		}, nil
}

// 查询交易
func (s *Server) doQueryTrade(ctx context.Context, args *queryTradeArgs) (*model.UserTrade, error) {
	var filter userTradeDao.FilterFunc

	switch {
	case args.TradeNO != "":
		filter = func(cols *userTradeDao.Columns) any {
			return map[string]any{cols.NO: args.TradeNO}
		}
	case args.RelatedID != "":
		filter = func(cols *userTradeDao.Columns) any {
			return map[string]any{cols.RelatedID: args.RelatedID}
		}

	default:
		return nil, errors.NewError(code.InvalidArgument)
	}

	trade, err := userTradeDao.Instance().FindOne(ctx, filter)
	if err != nil {
		log.Errorf("find trade failed, args = %s err = %v", xconv.Json(args), err)
		return nil, errors.NewError(err, code.InternalError)
	}

	if trade == nil {
		return nil, errors.NewError(code.NotFound)
	}

	return trade, nil
}
func (s *Server) doGetCurrencyInfo(currency xtypes.Currency) (*model.OptionCurrency, error) {

	_, opt := optionCurrencyCfg.GetOptionCurrencyByCurrency(currency)

	if opt == nil {
		return nil, errors.NewError(code.OptionNotFound)
	}
	// 检测用户信息
	return opt, nil
}

// 获取用户
func (s *Server) doGetUser(ctx context.Context, uid int64) (*model.UserBase, error) {
	return userBaseDao.Instance().GetUserBase(ctx, uid)
}
