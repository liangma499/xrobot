package xtypes

import (
	"math"
	"strings"

	"github.com/shopspring/decimal"
)

type Currency string

const (
	USDT   Currency = "USDT"
	TRX    Currency = "TRX"
	ENERGY Currency = "ENERGY"
	BISHU  Currency = "BISHU"
)

func (c Currency) String() string {
	return string(c)
}
func (c Currency) ToUpper() Currency {
	return Currency(strings.ToUpper(string(c)))
}

const (
	Trc20ContractUSTD   = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t" //USDT
	Trc20ContractEnergy = "TU2MJ5Veik1LRAgjeSzEdvmDYx7mefJZvd"
	Trc20ContractTRX    = "-"
)

func (currency Currency) Trc20Contract() string {
	switch currency {
	case USDT:
		{
			return Trc20ContractUSTD
		}
	case ENERGY:
		{
			return Trc20ContractEnergy
		}
	case TRX:
		{
			return Trc20ContractTRX
		}
	}
	return ""
}
func (currency Currency) IsValid() bool {
	switch currency {
	case USDT, TRX:
		{
			return true
		}

	}
	return false
}
func Trc20ContractToCurrency(contract string) string {
	switch contract {
	case Trc20ContractUSTD:
		{
			return USDT.String()
		}
	case Trc20ContractEnergy:
		{
			return ENERGY.String()
		}
	}
	return ""
}
func doGetdecimalPlace(currency Currency) int32 {
	//这个函数不要改，改了一定要清用户金币缓存
	decimalPlace := int32(6)
	switch currency {
	case USDT, TRX:
		{
			decimalPlace = 6
		}
	}
	return decimalPlace
}
func doKeepDecimalDigits(currency Currency) (decimal.Decimal, int32) {

	decimalPlace := doGetdecimalPlace(currency)

	return decimal.NewFromFloat(math.Pow10(int(decimalPlace))), int32(decimalPlace)
}
func KeepDecimalDigits(currency Currency) (decimal.Decimal, int32) {

	return doKeepDecimalDigits(currency)
}

// 将数据放大
func DecimalInt64(f decimal.Decimal, currency Currency) int64 {
	rate, _ := doKeepDecimalDigits(currency)
	return f.Mul(rate).BigInt().Int64()
}

// 将数据减小
func DecimalFloat64(f decimal.Decimal, currency Currency) decimal.Decimal {
	rate, decimalPlace := doKeepDecimalDigits(currency)
	return f.Div(rate).RoundFloor(decimalPlace)
}
func KeepDecimal(source decimal.Decimal, currency Currency) decimal.Decimal {
	decimalPlace := doGetdecimalPlace(currency)
	return source.RoundFloor(decimalPlace)
}
