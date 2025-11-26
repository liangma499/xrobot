package wallet

import (
	"context"
	"tron_robot/internal/code"
	userBaseDao "tron_robot/internal/dao/user-base"
	userTradeDao "tron_robot/internal/dao/user-trade"
	userWalletDao "tron_robot/internal/dao/user-wallet"
	"tron_robot/internal/model"
	optionCurrencyCfg "tron_robot/internal/option/option-currency"
	"tron_robot/internal/service/wallet/pb"
	"tron_robot/internal/service/wallet/rate"
	"tron_robot/internal/xtypes"

	"fmt"
	"strings"
	"xbase/cluster/mesh"
	"xbase/codes"
	"xbase/errors"
	"xbase/log"

	"github.com/shopspring/decimal"

	"gorm.io/gorm/clause"
)

const (
	serviceName = "wallet" // 服务名称
	servicePath = "Wallet" // 服务路径要与pb中的服务路径保持一致
)

var _ pb.WalletAble = &Server{}

type Server struct {
	proxy *mesh.Proxy
}

func NewServer(proxy *mesh.Proxy) *Server {
	return &Server{
		proxy: proxy,
	}
}

func (s *Server) Init() {
	s.proxy.AddServiceProvider(serviceName, servicePath, s)
	rate.Instance().InitTimer()
}

// InitWallet 初始化钱包
func (s *Server) InitWallet(ctx context.Context, args *pb.InitWalletArgs, reply *pb.InitWalletReply) error {
	defaultCurrencyName := xtypes.Currency(args.DefaultCurrency).ToUpper()
	opts := optionCurrencyCfg.GetOpts()
	if opts == nil {
		return errors.NewError(code.OptionNotFound)
	}
	if !defaultCurrencyName.IsValid() {
		var defaultCurrency *model.OptionCurrency
		for _, item := range opts.Opts {
			if item == nil {
				continue
			}
			if item.Status != xtypes.OptionStatus_Normal {
				continue
			}
			if defaultCurrency == nil {
				defaultCurrency = item
			} else {
				if item.Sort < defaultCurrency.Sort {
					defaultCurrency = item
				}
			}
		}
		if defaultCurrency != nil {
			defaultCurrencyName = defaultCurrency.Currency
		}
	}

	for _, item := range opts.Opts {
		if item == nil {
			continue
		}
		if item.Status != xtypes.OptionStatus_Normal {
			continue
		}
		def := xtypes.WalletDefaultNo
		if !defaultCurrencyName.IsValid() {
			defaultCurrencyName = item.Currency
		}
		if item.Currency == defaultCurrencyName {
			def = xtypes.WalletDefaultYes
		}

		initData := &userWalletDao.InitBalanceArgs{
			UID:      args.UID,
			Currency: item.Currency,
			Def:      def,
			Cash:     decimal.NewFromFloat(args.InitCash),
			Used:     decimal.NewFromFloat(0),
		}
		for _, item := range xtypes.WalletAmountKindUsed {
			_, err := userWalletDao.Instance().InitBalance(ctx, initData, item)
			if err != nil {
				log.Warnf("InitWallet:%#v", initData)
			}
		}

	}
	return nil
}

// SetDefaultCurrency 设置默认货币
func (s *Server) SetDefaultCurrency(ctx context.Context, args *pb.SetDefaultCurrencyArgs, reply *pb.SetDefaultCurrencyReply) error {
	currencyAll, currencyInfo := optionCurrencyCfg.GetOptionCurrencyByCurrency(xtypes.Currency(args.Currency))
	if currencyInfo == nil {
		return errors.NewError(code.OptionNotFound)
	}
	if currencyInfo.Status != xtypes.OptionStatus_Normal {
		return errors.NewError(code.CurrencyNotUse)
	}
	if currencyAll == nil {
		return errors.NewError(code.OptionNotFound)
	}

	needSet := true
	for _, item := range currencyAll {
		if item == nil {
			continue
		}
		if item.Status != xtypes.OptionStatus_Normal {
			continue
		}
		balance, err := userWalletDao.Instance().GetBalance(ctx, args.UID, item.Currency, xtypes.WalletAmountKindCash)
		if err != nil {
			needSet = false
			log.Error("SetDefaultCurrency:%v,%v", err, item.Currency)
			continue
		}

		if balance.IsDefault {
			if balance.Currency.String() == args.Currency {
				needSet = false
			} else {
				if err = userWalletDao.Instance().SetBalanceDefaultStatus(ctx, args.UID, item.Currency, xtypes.WalletDefaultNo, xtypes.WalletAmountKindCash); err != nil {
					needSet = false
					log.Error("SetDefaultCurrency:%v,%v", err, item.Currency)
				}
			}
			break
		}
	}

	if needSet {
		return userWalletDao.Instance().SetBalanceDefaultStatus(ctx, args.UID, currencyInfo.Currency, xtypes.WalletDefaultYes, xtypes.WalletAmountKindCash)
	}
	return nil
}

// IncrBalance 增加账户余额
func (s *Server) IncrBalance(ctx context.Context, args *pb.IncrBalanceArgs, reply *pb.IncrBalanceReply) error {
	amountKind := xtypes.WalletAmountKind(args.AmountKind)
	if !amountKind.IsValid() {
		return errors.NewError(code.InvalidArgument)
	}
	currency := xtypes.Currency(args.Currency)
	currencyInfo, err := s.doGetCurrencyInfo(currency)
	if err != nil {
		return err
	}

	user, err := s.doGetUser(ctx, args.UID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.NewError(code.NotFound)
	}
	channelCode := user.ChannelCode
	nickName := user.Nickname

	tradeNO, balance, change, err := s.doIncrBalance(ctx, &incrBalanceArgs{
		UID:             args.UID,
		Currency:        currency,
		Cash:            args.Cash,
		TradeType:       xtypes.TradeType(args.Type),
		TradeStatus:     xtypes.TradeStatusSuccess,
		RelatedID:       args.RelatedID,
		Taxation:        args.Taxation,
		BetAmount:       args.BetAmount,
		Rebate:          args.Rebate,
		UserType:        user.UserType,
		ChannelCode:     channelCode,
		NickName:        nickName,
		GameName:        args.GameName,
		AmountKind:      amountKind,
		PlayCurrencyID:  args.PlayCurrencyID, // 币种
		PlayCurrency:    args.PlayCurrency,   // 币种
		PlayCash:        args.PlayCash,       // 变动的现金
		TransactionID:   args.TransactionID,
		UserControlKind: xtypes.UserControlKind(args.UserControlKind),
	}, currencyInfo, user)
	if err != nil {
		return err
	}

	reply.TradeNO = tradeNO
	reply.Balance = balance
	reply.Change = change
	return nil
}

// DecrBalance 扣减账户余额
func (s *Server) DecrBalance(ctx context.Context, args *pb.DecrBalanceArgs, reply *pb.DecrBalanceReply) error {
	amountKind := xtypes.WalletAmountKind(args.AmountKind)
	if !amountKind.IsValid() {
		return errors.NewError(code.InvalidArgument)
	}
	currency := xtypes.Currency(args.Currency)

	currencyInfo, err := s.doGetCurrencyInfo(currency)
	if err != nil {
		return err
	}
	if currencyInfo == nil {
		return errors.NewError(code.OptionNotFound)
	}

	user, err := s.doGetUser(ctx, args.UID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.NewError(code.InvalidArgument)
	}
	tradeNO, balance, change, err := s.doDecrBalance(ctx, &decrBalanceArgs{
		UID:             args.UID,
		Currency:        currencyInfo.Currency,
		Cash:            args.Cash,
		TradeType:       xtypes.TradeType(args.Type),
		TradeStatus:     xtypes.TradeStatusSuccess,
		RelatedID:       args.RelatedID,
		Taxation:        args.Taxation,
		BetAmount:       args.BetAmount,
		Rebate:          args.Rebate,
		UserType:        user.UserType,
		ChannelCode:     user.ChannelCode,
		NickName:        user.Nickname,
		GameName:        args.GameName,
		AmountKind:      amountKind,
		PlayCurrencyID:  args.PlayCurrencyID, // 币种
		PlayCurrency:    args.PlayCurrency,   // 币种
		PlayCash:        args.PlayCash,       // 变动的现金
		TransactionID:   args.TransactionID,
		UserControlKind: xtypes.UserControlKind(args.UserControlKind),
	}, currencyInfo, user)
	if err != nil {
		return err
	}
	reply.TradeNO = tradeNO
	reply.Balance = balance
	reply.Change = change

	return nil
}

// FreezeBalance 冻结账户
func (s *Server) FreezeBalance(ctx context.Context, args *pb.FreezeBalanceArgs, reply *pb.FreezeBalanceReply) error {
	amountKind := xtypes.WalletAmountKind(args.AmountKind)
	if !amountKind.IsValid() {
		return errors.NewError(code.InvalidArgument)
	}
	currency := xtypes.Currency(args.Currency)
	currencyInfo, err := s.doGetCurrencyInfo(currency)
	if err != nil {
		return err
	}
	if currencyInfo == nil {
		return errors.NewError(code.OptionNotFound)
	}

	user, err := s.doGetUser(ctx, args.UID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.NewError(code.InvalidArgument)
	}

	tradeNO, balance, change, err := s.doDecrBalance(ctx, &decrBalanceArgs{
		UID:            args.UID,
		Currency:       currencyInfo.Currency,
		Cash:           args.Cash,
		TradeType:      xtypes.TradeType(args.Type),
		TradeStatus:    xtypes.TradeStatusTrading,
		RelatedID:      args.RelatedID,
		Taxation:       0,
		UserType:       user.UserType,
		ChannelCode:    user.ChannelCode,
		NickName:       user.Nickname,
		AmountKind:     amountKind,
		PlayCurrencyID: 0,  // 币种
		PlayCurrency:   "", // 币种
		PlayCash:       0,  // 变动的现金
	}, currencyInfo, user)
	if err != nil {
		return err
	}

	reply.TradeNO = tradeNO
	reply.Balance = balance
	reply.Change = change

	return nil
}

// CancelTrade 取消交易
func (s *Server) CancelTrade(ctx context.Context, args *pb.CancelTradeArgs, reply *pb.CancelTradeReply) error {
	balance, change, err := s.doCancelTrade(ctx, args)
	if err != nil {
		return err
	}
	reply.Balance = balance
	reply.Change = change

	return nil
}

// CompleteTrade 完成交易
func (s *Server) CompleteTrade(ctx context.Context, args *pb.CompleteTradeArgs, reply *pb.CompleteTradeReply) error {
	if err := s.doCompleteTrade(ctx, args); err != nil {
		return err
	}

	return nil
}

// FetchBalance 拉取账户余额
func (s *Server) FetchBalance(ctx context.Context, args *pb.FetchBalanceArgs, reply *pb.FetchBalanceReply) error {

	currency := xtypes.Currency(args.Currency)
	currencyInfo, err := s.doGetCurrencyInfo(currency)
	if err != nil {
		return err
	}
	if currencyInfo == nil {
		return errors.NewError(code.OptionNotFound)
	}
	reply.Balance = make(map[int32]*pb.BalanceInfo)
	for _, item := range xtypes.WalletAmountKindUsed {
		balance, err := userWalletDao.Instance().GetBalance(ctx, args.UID, xtypes.Currency(args.Currency), item)

		if err != nil {
			return err
		}

		reply.Balance[int32(balance.AmountKind)] = &pb.BalanceInfo{
			Currency:   balance.Currency.String(),
			Cash:       balance.Cash,
			Used:       balance.Used,
			IsDefault:  balance.IsDefault,
			AmountKind: int32(balance.AmountKind),
		}

	}

	return nil
}

// FetchBalances 拉取用户余额
func (s *Server) FetchBalances(ctx context.Context, args *pb.FetchBalancesArgs, reply *pb.FetchBalancesReply) error {
	opts := optionCurrencyCfg.GetOpts()
	if opts == nil {
		return errors.NewError(code.OptionNotFound)
	}
	bHaveDefault := false
	reply.List = make(map[string]*pb.FetchBalanceMap)
	for _, item := range opts.Opts {
		if item == nil {
			continue
		}
		if item.Status != xtypes.OptionStatus_Normal {
			continue
		}
		balancePb := make(map[int32]*pb.BalanceInfo)
		for _, itemKind := range xtypes.WalletAmountKindUsed {
			balance, err := userWalletDao.Instance().GetBalance(ctx, args.UID, item.Currency, itemKind)
			if err != nil {
				log.Warnf("err:%v,currency:%v", err, item.Currency)
				continue

			}
			if !bHaveDefault {
				if balance.IsDefault {
					bHaveDefault = true
				}
			}
			balancePb[int32(balance.AmountKind)] = &pb.BalanceInfo{
				Currency:   balance.Currency.String(),
				Cash:       balance.Cash,
				Used:       balance.Used,
				IsDefault:  balance.IsDefault,
				AmountKind: int32(balance.AmountKind),
			}
		}
		reply.List[item.Currency.String()] = &pb.FetchBalanceMap{
			Balance: balancePb,
		}
	}
	if !bHaveDefault {

		for k, item := range reply.List {
			for kv := range item.Balance {
				if kv == xtypes.WalletAmountKindCash.Int32() {
					reply.List[k].Balance[kv].IsDefault = true
					return nil
				}
			}
		}
	}

	return nil
}

// FetchTradeList 拉取交易记录列表
func (s *Server) FetchTradeList(ctx context.Context, args *pb.FetchTradeListArgs, reply *pb.FetchTradeListReply) error {
	db := userTradeDao.Instance().Table.WithContext(ctx)

	if args.UID > 0 {
		db = db.Where(clause.Eq{Column: userTradeDao.Instance().Columns.UID, Value: args.UID})
	} else {
		db = db.Joins(fmt.Sprintf(
			"LEFT JOIN %s ON %s.%s = %s.%s ",
			userBaseDao.Instance().TableName,
			userBaseDao.Instance().TableName,
			userBaseDao.Instance().Columns.UID,
			userTradeDao.Instance().TableName,
			userTradeDao.Instance().Columns.UID,
		))
	}
	db = db.Where("1 = ?", 1)
	if len(args.Types) > 0 {
		values := make([]any, 0)
		for _, item := range args.Types {
			values = append(values, item)
		}
		db = db.Where(clause.IN{Column: userTradeDao.Instance().TableName + "." + userTradeDao.Instance().Columns.Type, Values: values})
		/*db = db.Where(map[string]any{
			userTradeDao.Instance().Columns.Type: args.Types,
		})*/
	}

	if args.StartTime != "" {
		db = db.Where(clause.Gte{
			Column: userTradeDao.Instance().TableName + "." + userTradeDao.Instance().Columns.CreatedAt,
			Value:  args.StartTime,
		})
	}

	if args.EndTime != "" {
		db = db.Where(clause.Gte{
			Column: userTradeDao.Instance().TableName + "." + userTradeDao.Instance().Columns.CreatedAt,
			Value:  args.EndTime,
		})
	}

	total := int64(0)
	if !args.HasMore {
		err := db.Count(&total).Error
		if err != nil {
			log.Errorf("count trade failed: %v", err)
			return errors.NewError(err, code.InternalError)
		}

	}

	reply.Page = args.Page
	reply.Limit = args.Limit
	reply.Total = int32(total)

	offset := int((args.Page - 1) * args.Limit)
	limit := int(args.Limit)
	if !args.HasMore && int64(offset) >= total {
		return nil
	}
	//数据倒序
	db.Order(fmt.Sprintf("%s.%s DESC", userTradeDao.Instance().TableName, userTradeDao.Instance().Columns.ID))

	list := make([]*TradeList, 0, limit)
	queryLimit := limit
	if args.HasMore {
		queryLimit += 1
	}

	if args.UID > 0 {
		err := db.Offset(offset).Limit(queryLimit).Scan(&list).Error
		if err != nil {
			log.Errorf("query trade list failed: %v", err)
			return errors.NewError(err, codes.InternalError)
		}
	} else {
		err := db.Select(strings.Join([]string{
			userTradeDao.Instance().TableName + ".*",
			userBaseDao.Instance().TableName + "." + userBaseDao.Instance().Columns.Code,
			userBaseDao.Instance().TableName + "." + userBaseDao.Instance().Columns.Nickname,
		}, ",")).Offset(offset).Limit(queryLimit).Scan(&list).Error
		if err != nil {
			log.Errorf("query trade list failed: %v", err)
			return errors.NewError(err, codes.InternalError)
		}
	}
	lenght := len(list)
	reply.HasMore = lenght > limit

	reply.List = make([]*pb.TradeInfo, 0, len(list))
	for i := 0; i < lenght && i < limit; i++ {
		item := list[i]
		reply.List = append(reply.List, &pb.TradeInfo{
			ID:          item.ID,
			NO:          item.NO,
			UID:         item.UID,
			Currency:    item.Currency.String(),
			BeforeCash:  item.BeforeCash,
			AfterCash:   item.AfterCash,
			ChangeCash:  item.ChangeCash,
			BetCash:     item.BetCash,
			UserCode:    item.Code,
			NickName:    item.Nickname,
			ChannelCode: item.ChannelCode,
			Type:        int32(item.Type),
			Status:      int32(item.Status),
			Channel:     string(item.Channel),
			CreatedAt:   item.CreatedAt.Unix(),
			UpdatedAt:   item.UpdatedAt.Unix(),
			AmountKind:  int32(item.AmountKind),
		})
	}

	return nil
}
