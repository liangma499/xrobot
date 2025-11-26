package cryptocurrency

import (
	"context"
	"xbase/errors"
	"xbase/log"
	"xbase/utils/xconv"
	"xbase/utils/xtime"
	"xrobot/internal/code"
	paymentAmountUserDao "xrobot/internal/dao/payment-amount-user"
	paymentCryptoTransactionDao "xrobot/internal/dao/payment-crypto-transaction"
	platformpendingorderdao "xrobot/internal/dao/platform-pending-order"
	"xrobot/internal/event/cryptocurrencyevent"
	"xrobot/internal/model"
	walletsvc "xrobot/internal/service/wallet"
	walletpb "xrobot/internal/service/wallet/pb"
	"xrobot/internal/xtypes"

	"github.com/shopspring/decimal"
	"gorm.io/gorm/clause"
)

func (s *Server) doSubscribeCryptoCurrency(uuid string, payload *cryptocurrencyevent.CryptoCurrencyMsg) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if payload == nil {
		return
	}
	ctx := context.Background()
	trx, err := paymentCryptoTransactionDao.Instance().FindMany(ctx, func(cols *paymentCryptoTransactionDao.Columns) any {
		return clause.Or(clause.Eq{
			Column: cols.Stauts,
			Value:  xtypes.Transaction_Confirmed,
		}, clause.Eq{
			Column: cols.Stauts,
			Value:  xtypes.Transaction_Fail,
		})
	}, nil, nil)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	for _, item := range trx {

		userInfo, err := paymentAmountUserDao.Instance().AmountToUser(item.Currency, item.Amount.String())
		if err != nil {
			log.Errorf("err:%v data:%s", err, xconv.Json(item))
			continue
		}
		if item.Stauts == xtypes.Transaction_Fail {
			if item.TransactionKind == xtypes.Transaction_OtherTransaction {
				log.Warnf("trxID:%s TransactionKind:%d fail", item.TransactionHash, item.TransactionKind)
			} else if item.TransactionKind == xtypes.Transaction_OutVerify {
				log.Warnf("trxID:%s TransactionKind:%d fail", item.TransactionHash, item.TransactionKind)
			}

		} else {
			//转出订单处理
			if item.TransactionKind == xtypes.Transaction_OtherTransaction {

			} else if item.TransactionKind == xtypes.Transaction_OutVerify {

			} else {
				if userInfo == nil {
					//处理兑换信息
					s.doExchange(userInfo, item)
				} else {

					if userInfo.Extend != nil && userInfo.Currency == item.Currency {
						if err := s.doUserRecharge(userInfo, item); err != nil {
							log.Errorf("not userInfo:%v data:%s err:%v", xconv.Json(userInfo), xconv.Json(item), err)
						}
					} else {
						log.Warnf("not userInfo:%v data:%s", xconv.Json(userInfo), xconv.Json(item))
					}
					//清除金额
					paymentAmountUserDao.Instance().ClearAmount(item.Currency, userInfo.Amount, userInfo.OrderID)
				}
			}

		}

	}
}
func (s *Server) doUserRecharge(userInfo *model.PaymentAmountUser, trx *model.PaymentCryptoTransaction) error {
	if userInfo == nil || trx == nil {
		return errors.NewError(code.InvalidArgument)
	}
	switch userInfo.Usage {
	case xtypes.Usage_Recharge: //充值
		{
			client, err := walletsvc.NewClient(s.proxy.NewMeshClient)
			if err != nil {
				return errors.NewError(code.InternalError, err)
			}
			_, err = client.IncrBalance(context.Background(), &walletpb.IncrBalanceArgs{
				UID:             userInfo.UID,
				Currency:        trx.Currency.String(),
				Cash:            trx.Amount.Abs().InexactFloat64(),
				Type:            int32(xtypes.TradeTypeRecharge),
				Rebate:          false,
				AmountKind:      xtypes.WalletAmountKindCash.Int32(),
				UserControlKind: int32(xtypes.UserNone),
			})
			if err != nil {
				return err
			}
			//发送到对应的渠道群
			return nil
		}
	case xtypes.Usage_Recharge_OtherAddress: //充值
		{
			addressInfoLen := len(userInfo.Extend.AddressInfo)

			if addressInfoLen == 0 {
				return nil
			}
			platformPendingOrder := make([]*model.PlatformPendingOrder, 0)
			now := xtime.Now()
			for i := 0; i < addressInfoLen; i++ {
				item := userInfo.Extend.AddressInfo[i]
				if item == nil {
					continue
				}
				stauts := xtypes.OrderStatus_Activated
				if !item.Activated {
					stauts = xtypes.OrderStatus_ToBeActivated
				}

				platformPendingOrder = append(platformPendingOrder, &model.PlatformPendingOrder{
					UID:          userInfo.UID,                       // 主键
					Code:         xconv.String(userInfo.TelegramUid), // 编号
					ChannelCode:  userInfo.ChannelCode,               // 渠道编码
					Address:      item.Address,
					ToCurrency:   "", // 币种
					ToAmount:     decimal.Zero,
					FromCurrency: "", // 币种
					FromAmount:   decimal.Zero,
					Energy:       0,
					Stauts:       stauts,
					CreateAt:     now,
					UpdateAt:     now,
				})
			}
			if len(platformPendingOrder) > 0 {
				_, err := platformpendingorderdao.Instance().Insert(context.Background(), platformPendingOrder...)
				return errors.NewError(code.InternalError, err)
			}
			return nil
		}
	case xtypes.Usage_Recharge_BiShu: //充值
		{
			//往用户向上加笔数
			client, err := walletsvc.NewClient(s.proxy.NewMeshClient)
			if err != nil {
				return errors.NewError(code.InternalError, err)
			}
			_, err = client.IncrBalance(context.Background(), &walletpb.IncrBalanceArgs{
				UID:             userInfo.UID,
				Currency:        xtypes.BISHU.String(),
				Cash:            float64(userInfo.Extend.BiShu),
				Type:            int32(xtypes.TradeTypeRechargeBiShu),
				Rebate:          false,
				AmountKind:      xtypes.WalletAmountKindCash.Int32(),
				UserControlKind: int32(xtypes.UserNone),
			})
			if err != nil {
				return err
			}
			//发送到对应的渠道群
			return nil
		}

	}
	return nil
}

func (s *Server) doExchange(userInfo *model.PaymentAmountUser, trx *model.PaymentCryptoTransaction) error {
	if userInfo == nil || trx == nil {
		return errors.NewError(code.InvalidArgument)
	}
	return nil
}

// 佣金返回
func (s *Server) doRebateCommission(userInfo *model.PaymentAmountUser, trx *model.PaymentCryptoTransaction) error {
	if userInfo == nil || trx == nil {
		return errors.NewError(code.InvalidArgument)
	}
	return nil
}
