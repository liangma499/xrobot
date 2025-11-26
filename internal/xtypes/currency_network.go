package xtypes

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"math"
	"xbase/log"
	"xbase/utils/xconv"

	"github.com/shopspring/decimal"
)

const (
	Trc20_Places            = 6 //trc20扩大倍数
	TransactionLimit        = 10
	Tron_Wallet_Must_Blance = float64(0.01)
	Tron_Wallet_Net_Fee     = float64(0.5)
	Tron_MaxGasFeeLimmit    = float64(3)
)

// 扩大倍数
func CoefficientInt64(amount decimal.Decimal, places int8) decimal.Decimal {
	return amount.Mul(decimal.NewFromFloat(math.Pow(10, xconv.Float64(places))))
}

// 缩小倍数
func CoefficientToFloat64(amount decimal.Decimal, places int8) decimal.Decimal {
	if places > 0 {
		places = 0 - places
	}
	return amount.Mul(decimal.NewFromFloat(math.Pow(10, xconv.Float64(places))))
}

// 支付相关的定义在这个定方
// 金额对用户中的作用
type Usage int32

const (
	Usage_Recharge              Usage = 1 //充值
	Usage_Recharge_OtherAddress Usage = 2 //给其他地址充值
	Usage_Recharge_BiShu        Usage = 3 //充值笔数
	Usage_BuySameRobot          Usage = 4 //充值笔数
)

type PaymentStatus int32

const (
	PaymentStatus_Normal     PaymentStatus = 1 //可用
	PaymentStatus_Completion PaymentStatus = 2 //已经完成
	PaymentStatus_Clean      PaymentStatus = 3 //已经清理
)

type NetWork string

const (
	None   NetWork = "None"
	Solana NetWork = "Solana"
	Ton    NetWork = "TON"
	Crypto NetWork = "Crypto"
	TRON   NetWork = "TRON"
	BTC    NetWork = "BTC"
	ETH    NetWork = "ETH"
	BNB    NetWork = "BNB"
)

type NetWorkChannelType int

const (
	NetWorkChannelType_None      NetWorkChannelType = 0
	NetWorkChannelType_Solana    NetWorkChannelType = 1
	NetWorkChannelType_Ton       NetWorkChannelType = 2
	NetWorkChannelType_TonCrypto NetWorkChannelType = 3
	NetWorkChannelType_TRON      NetWorkChannelType = 4
	NetWorkChannelType_BTC       NetWorkChannelType = 5
	NetWorkChannelType_ETH       NetWorkChannelType = 6
	NetWorkChannelType_BNB       NetWorkChannelType = 7
)

func (n NetWork) String() string {
	return string(n)
}
func (n NetWork) NetWorkChannelType() NetWorkChannelType {
	switch n {
	case TRON:
		{
			return NetWorkChannelType_TRON
		}
	case Solana:
		{
			return NetWorkChannelType_Solana
		}
	case Ton:
		{
			return NetWorkChannelType_Ton
		}
	case Crypto:
		{
			return NetWorkChannelType_TonCrypto
		}
	case BTC:
		{
			return NetWorkChannelType_BTC
		}
	case ETH:
		{
			return NetWorkChannelType_ETH
		}
	case BNB:
		{
			return NetWorkChannelType_BNB
		}
	}

	return NetWorkChannelType_None

}

func (r NetWorkChannelType) NetWork() NetWork {
	switch r {
	case NetWorkChannelType_TRON:
		{
			return TRON
		}
	case NetWorkChannelType_Solana:
		{
			return Solana
		}
	case NetWorkChannelType_Ton:
		{
			return Ton
		}
	case NetWorkChannelType_TonCrypto:
		{
			return Crypto
		}
	case NetWorkChannelType_BTC:
		{
			return BTC
		}
	case NetWorkChannelType_ETH:
		{
			return ETH
		}
	case NetWorkChannelType_BNB:
		{
			return BNB
		}

	}

	return None
}

type TransactionStatus int32

const (
	Transaction_Verified  TransactionStatus = 1 //待验证
	Transaction_Confirmed TransactionStatus = 2 //交易成功
	Transaction_Fail      TransactionStatus = 3 //交易失败
	Transaction_Finish    TransactionStatus = 4 //完成
)

type TransactionKind int32

const (
	Transaction_None             TransactionKind = 0   //无
	Transaction_Recharge         TransactionKind = 1   //充值
	Transaction_EnergyIn         TransactionKind = 2   //能量转入
	Transaction_OutVerify        TransactionKind = 99  //转出订单
	Transaction_OtherTransaction TransactionKind = 999 //别人付款取消代理能量
)

type BlockExtend struct {
	Start int64 `json:"start,omitempty"`
	End   int64 `json:"end,omitempty"`
	Now   int64 `json:"now,omitempty"`
	Key   int64 `json:"key,omitempty"`
}

func (ot *BlockExtend) Clone() *BlockExtend {
	if ot == nil {
		return nil
	}
	return &BlockExtend{
		Start: ot.Start,
		End:   ot.End,
		Now:   ot.Now,
		Key:   ot.Key,
	}
}

type BlockExtendMap map[int64]*BlockExtend

func (ot BlockExtendMap) Clone() BlockExtendMap {
	if ot == nil {
		return nil
	}
	rst := make(BlockExtendMap)
	for k, v := range ot {
		rst[k] = v.Clone()
	}
	return rst
}

type NetworkInfo struct {
	Block       int64          `json:"block,omitempty"` //Start number. Default 0
	BlockExtend BlockExtendMap `json:"block_extend,omitempty"`
}

func (ot *NetworkInfo) Clone() *NetworkInfo {
	if ot == nil {
		return nil
	}
	rst := &NetworkInfo{
		Block:       ot.Block, //Block
		BlockExtend: ot.BlockExtend.Clone(),
	}
	return rst
}
func (ot NetworkInfo) Value() (driver.Value, error) {

	s, err := json.Marshal(ot)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return nil, err
	}
	return s, nil
}

func (ot *NetworkInfo) Scan(value any) error {
	if value == nil {
		return nil
	}
	if ot == nil {
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid Scan Source")
	}
	err := json.Unmarshal(s, ot)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return nil
	}
	return nil
}

type PlatformExtraCfg struct {
	ExchangeAddress string          `json:"exchangeAddress,omitempty"` //兑换能量地址
	PriceU          decimal.Decimal `json:"priceU,omitempty"`          //有U的价格
	PriceNoU        decimal.Decimal `json:"priceNoU,omitempty"`        //无U的价格
	EnergyU         decimal.Decimal `json:"energyU,omitempty"`         //有U的需要能量
	EnergyNoU       decimal.Decimal `json:"energyNoU,omitempty"`       //无U的需要能量
	IsUserEnergy    bool            `json:"isUserEnergy,omitempty"`    //是否需要能量
}

func (ot *PlatformExtraCfg) Clone() *PlatformExtraCfg {
	if ot == nil {
		return nil
	}

	return &PlatformExtraCfg{
		ExchangeAddress: ot.ExchangeAddress,
		PriceU:          ot.PriceU.Copy(),
		PriceNoU:        ot.PriceNoU.Copy(),
		EnergyU:         ot.EnergyU.Copy(),
		EnergyNoU:       ot.EnergyNoU.Copy(),
		IsUserEnergy:    ot.IsUserEnergy,
	}
}

func (ot PlatformExtraCfg) Value() (driver.Value, error) {

	s, err := json.Marshal(ot)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return nil, err
	}
	return s, nil
}

func (ot *PlatformExtraCfg) Scan(value any) error {
	if value == nil {
		return nil
	}
	if ot == nil {
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid Scan Source")
	}
	err := json.Unmarshal(s, ot)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return nil
	}
	return nil
}

type EnergyExtra struct {
	ExchangeAddress string          `json:"exchangeAddress,omitempty"` //兑换能量地址
	ExchangeTXHash  string          `json:"exchangeTXHash,omitempty"`  //交易Hash
	ToBalance       decimal.Decimal `json:"toBalance,omitempty"`       //是否需要能量
	ExchangePrice   decimal.Decimal `json:"exchangePrice,omitempty"`   //有U的价格
}

func (ot *EnergyExtra) Clone() *EnergyExtra {
	if ot == nil {
		return nil
	}

	return &EnergyExtra{
		ExchangeAddress: ot.ExchangeAddress,
		ExchangeTXHash:  ot.ExchangeTXHash,
		ToBalance:       ot.ToBalance.Copy(),
		ExchangePrice:   ot.ExchangePrice.Copy(),
	}
}

func (ot EnergyExtra) Value() (driver.Value, error) {

	s, err := json.Marshal(ot)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return nil, err
	}
	return s, nil
}

func (ot *EnergyExtra) Scan(value any) error {
	if value == nil {
		return nil
	}
	if ot == nil {
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid Scan Source")
	}
	err := json.Unmarshal(s, ot)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return nil
	}
	return nil
}
