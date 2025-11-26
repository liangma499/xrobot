package wallet_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"xbase/eventbus"
	"xbase/eventbus/nats"
	"xbase/log"
	"xbase/utils/xconv"
	"xbase/utils/xtime"
	"xrobot/internal/event/message"
	userevt "xrobot/internal/event/user"
	walletsvc "xrobot/internal/service/wallet"
	walletpb "xrobot/internal/service/wallet/pb"
	"xrobot/internal/utils/xstr"
	tgtypes "xrobot/internal/xtelegram/tg-types"
	"xrobot/internal/xtypes"

	"xbase/registry/consul"
	"xbase/transport/rpcx"
)

const (
	uid      int64 = 1
	currency       = xtypes.USDT
)

var transporter = rpcx.NewTransporter(
	rpcx.WithClientDiscovery(consul.NewRegistry()),
)

func init() {
	eventbus.SetEventbus(nats.NewEventbus())
}
func TestClient_InitWallet(t *testing.T) {
	client, err := walletsvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.InitWallet(context.Background(), &walletpb.InitWalletArgs{
		UID:             uid,
		DefaultCurrency: currency.String(),
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_SetDefaultCurrency(t *testing.T) {
	client, err := walletsvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.SetDefaultCurrency(context.Background(), &walletpb.SetDefaultCurrencyArgs{
		UID:      uid,
		Currency: currency.String(),
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_FetchBalance(t *testing.T) {
	client, err := walletsvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.FetchBalance(context.Background(), &walletpb.FetchBalanceArgs{
		UID:      uid,
		Currency: currency.String(),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("balance: %s", xconv.Json(reply))
}

func TestClient_FetchBalances(t *testing.T) {
	client, err := walletsvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.FetchBalances(context.Background(), &walletpb.FetchBalancesArgs{
		UID: uid,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("balances: %s", xconv.Json(reply))
}

const (
	transactionID = "aaaaaaaaaaatttttttttttt03"
	roundId       = "aaaaaaaaaaatttttttttttt03"
	amountKind    = xtypes.WalletAmountKindCash
	tradeType     = xtypes.TradeTypeGameBet
)

func TestClient_IncrBalance01(t *testing.T) {
	client, err := walletsvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.IncrBalance(context.Background(), &walletpb.IncrBalanceArgs{
		UID:             uid,
		Currency:        currency.String(),
		Cash:            10,
		BetAmount:       10,
		Type:            int32(xtypes.TradeTypeAffiliate),
		Rebate:          true,
		AmountKind:      int32(amountKind),
		TransactionID:   xstr.SerialNO(),
		UserControlKind: int32(xtypes.UserLose),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %s", xconv.Json(reply))
	/*reply, err = client.IncrBalance(context.Background(), &walletpb.IncrBalanceArgs{
		UID:             uid,
		Currency:      currency.String(),
		Cash:            10,
		BetAmount:       10,
		Type:            int32(tradeType),
		Rebate:          true,
		AmountKind:      int32(amountKind),
		DeveloperKind:   xtypes.DeveloperKind_TGThirdGame,
		TransactionID:   xstr.SerialNO(),
		UserControlKind: int32(xtypes.UserWin),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %s", xconv.Json(reply))

	reply, err = client.IncrBalance(context.Background(), &walletpb.IncrBalanceArgs{
		UID:             uid,
		Currency:      currency.String(),
		Cash:            10,
		BetAmount:       10,
		Type:            int32(tradeType),
		Rebate:          true,
		AmountKind:      int32(amountKind),
		DeveloperKind:   xtypes.DeveloperKind_TGThirdGame,
		TransactionID:   xstr.SerialNO(),
		UserControlKind: int32(xtypes.UserNone),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %s", xconv.Json(reply))*/
}
func TestClient_IncrBalance02(t *testing.T) {
	client, err := walletsvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.IncrBalance(context.Background(), &walletpb.IncrBalanceArgs{
		UID:             uid,
		Currency:        currency.String(),
		Cash:            10000,
		BetAmount:       10,
		Type:            int32(xtypes.TradeTypeRecharge),
		AmountKind:      xtypes.WalletAmountKindCash.Int32(),
		TransactionID:   xstr.SerialNO(),
		UserControlKind: int32(xtypes.UserLose),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("reply: %s", xconv.Json(reply))
}

func TestClient_DecrBalance1(t *testing.T) {
	client, err := walletsvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.DecrBalance(context.Background(), &walletpb.DecrBalanceArgs{
		UID:           uid,
		Currency:      currency.String(),
		Cash:          11,
		BetAmount:     10,
		Type:          int32(xtypes.TradeTypeRecharge),
		Rebate:        true,
		AmountKind:    xtypes.WalletAmountKindCash.Int32(),
		TransactionID: xstr.SerialNO(),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %s", xconv.Json(reply))
}

func TestClient_DecrBalance2(t *testing.T) {
	client, err := walletsvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.DecrBalance(context.Background(), &walletpb.DecrBalanceArgs{
		UID:        uid,
		Currency:   currency.String(),
		Cash:       95,
		Type:       int32(xtypes.TradeTypeGameWin),
		AmountKind: xtypes.WalletAmountKindCash.Int32(),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %s", xconv.Json(reply))
}

func TestClient_FreezeBalance1(t *testing.T) {
	client, err := walletsvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.FreezeBalance(context.Background(), &walletpb.FreezeBalanceArgs{
		UID:        uid,
		Currency:   currency.String(),
		Cash:       100,
		Type:       int32(xtypes.TradeTypeWithdraw),
		AmountKind: xtypes.WalletAmountKindCash.Int32(),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %s", xconv.Json(reply))
}

func TestClient_CancelTrade(t *testing.T) {
	client, err := walletsvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.CancelTrade(context.Background(), &walletpb.CancelTradeArgs{
		TradeNO: "1854438829780369408",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %s", xconv.Json(reply))
}

func TestClient_CompleteTrade(t *testing.T) {
	client, err := walletsvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.CompleteTrade(context.Background(), &walletpb.CompleteTradeArgs{
		TradeNO: "1854439230957158400",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %s", xconv.Json(reply))
}

func TestClient_FetchTradeList(t *testing.T) {
	client, err := walletsvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.FetchTradeList(context.Background(), &walletpb.FetchTradeListArgs{
		Page:  1,
		Limit: 10,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %s", xconv.Json(reply))
}

func TestClient_xutctime(t *testing.T) {

	log.Warnf("reply: %v", xtime.Now())
}

func TestClient_Expand(t *testing.T) {
	replaces := make(map[string]string)
	replaces["winner_name"] = "tttttttt"
	replaces["win_amount"] = "aaaaaa"
	replaces["multiplier"] = fmt.Sprintf("%.4f", 0.111111111111)

	send := `🎉 *Congratulations, ${winner_name}!* 🎉

			You won with a *${multiplier}x* multiplier!
		

			You're on fire! 🔥 Keep it up, aim for the moon!`

	body := os.Expand(send, func(s string) string {
		return replaces[s]
	})
	log.Warnf("%v", body)

}

func TestClient_PublishMessageStart(t *testing.T) {
	commonMsg := message.MessageCommon{
		Button:      tgtypes.XTelegramButton_Start,
		Type:        message.MessageType_Private,
		ChatID:      1111, //用户Telegram ID
		UserID:      1111,
		ChannelCode: xtypes.OfficialChannelCode, //渠道码
		ClientIP:    "127.0.0.1",
		OrderID:     "",
	}
	message.PublishMessageStart(&message.MessageStart{
		MessageCommon: commonMsg,
		UserName:      "TTTTT",
		InviteCode:    "",      //邀请码
		FirstName:     "TTTTT", //姓
		LastName:      "bbbbb", //名

	})

}

//func TestClient_Debit(t *testing.T) {
//	client, err := walletsvc.NewClient(transporter.NewClient)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	reply, err := client.Debit(context.Background(), &walletpb.DebitArgs{
//		Uid:        uid,
//		OrderId:    xuuid.UUID(),
//		GameId:     gameID,
//		RoundId:    xuuid.UUID(),
//		Amount:     1.9000070001,
//		Currency:   currency,
//		Extend:     "",
//		Channel:    string(model.BillChannelBridge),
//		ActionType: int32(model.BillTypeGameBet),
//	})
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	t.Logf("success: %v,balance: %v", reply.Success, reply.Balance)
//}
//
//func TestClient_CancelDebit(t *testing.T) {
//	client, err := walletsvc.NewClient(transporter.NewClient)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	reply, err := client.CancelDebit(context.Background(), &walletpb.CancelDebitArgs{
//		Uid:       uid,
//		CancelId:  xuuid.UUID(),
//		RelatedId: "37e10a84-b53a-11ee-abc4-d8bbc146a9a7", // 数据库中已存在
//		Channel:   string(model.BillChannelBridge),
//	})
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	t.Logf("balance: %v", reply.Balance)
//}
//
//func TestClient_Credit(t *testing.T) {
//	client, err := walletsvc.NewClient(transporter.NewClient)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	reply, err := client.Credit(context.Background(), &walletpb.CreditArgs{
//		Uid:        uid,
//		OrderId:    xuuid.UUID(),
//		GameId:     gameID,
//		RoundId:    xuuid.UUID(),
//		Amount:     3.3396,
//		Currency:   currency,
//		Extend:     "",
//		Channel:    string(model.BillChannelBridge),
//		ActionType: int32(model.BillTypeGameBetWin),
//	})
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	t.Logf("balance: %v", reply.Balance)
//}
//
//func TestClient_UserBalance(t *testing.T) {
//	client, err := walletsvc.NewClient(transporter.NewClient)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	reply, err := client.UserBalance(context.Background(), &walletpb.UserBalanceArgs{
//		Currency: currency,
//		Uid:      uid,
//	})
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	t.Logf("balance: %v", reply.Balance)
//}
//
//func TestClient_UserAllBalance(t *testing.T) {
//	client, err := walletsvc.NewClient(transporter.NewClient)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	reply, err := client.UserAllBalance(context.Background(), &walletpb.UserAllBalanceArgs{
//		Uid: uid,
//	})
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	t.Logf("balance: %v", reply.Balance)
//}

func TestClient_SubscribeRegister(t *testing.T) {
	userevt.PublishRegister(&userevt.RegisterPayload{UID: 15})
}

func TestClient_MonthTail(t *testing.T) {

	tailTime := xtime.MonthTail().Unix()
	log.Warnf("%d", tailTime)
}
