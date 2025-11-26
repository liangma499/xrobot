package xmath

import (
	"github.com/shopspring/decimal"
	"golang.org/x/exp/constraints"
	"math"
	"math/big"
)

// Floor 舍去取整保留n位小数
func Floor(f float64, n ...int) float64 {
	s := float64(1)

	if len(n) > 0 {
		s = math.Pow10(n[0])
	}

	return math.Floor(f*s+0.00000001) / s
}

// Ceil 进一取整保留n位小数
func Ceil(f float64, n ...int) float64 {
	s := float64(1)

	if len(n) > 0 {
		s = math.Pow10(n[0])
	}

	return math.Ceil(f*s) / s
}

// Round 四舍五入保留n位小数
func Round(f float64, n ...int) float64 {
	s := float64(1)

	if len(n) > 0 {
		s = math.Pow10(n[0])
	}

	return math.Round(f*s) / s
}

func ArrayContainValue(arr []int64, search int64) bool {
	for _, v := range arr {
		if v == search {
			return true
		}
	}

	return false
}

// 两个flota运算
func FloatOperation[T constraints.Integer | constraints.Float](symbol byte, t1, t2 T) (r float64) {
	f1 := big.NewFloat(float64(t1))
	f2 := big.NewFloat(float64(t2))
	switch symbol {
	case '+':
		r, _ = new(big.Float).Add(f1, f2).Float64()
	case '-':
		r, _ = new(big.Float).Sub(f1, f2).Float64()
	case '*':
		r, _ = new(big.Float).Mul(f1, f2).Float64()
	case '/':
		r, _ = new(big.Float).Quo(f1, f2).Float64()
	}
	return
}

func DecimalDip[T constraints.Integer | constraints.Float](symbol byte, t1, t2 T) decimal.Decimal {
	f1 := decimal.NewFromFloat(float64(t1))
	f2 := decimal.NewFromFloat(float64(t2))
	var r decimal.Decimal
	switch symbol {
	case '+':
		r = f1.Add(f2)
	case '-':
		r = f1.Sub(f2)
	case '*':
		r = f1.Mul(f2)
	case '/':
		r = f1.Div(f2)
	}

	return r
}

// 对float四舍五入保留小数
func DecimalRound(f float64, n int32) float64 {
	return decimal.NewFromFloat(f).Round(n).InexactFloat64()
	//pow10_n := math.Pow10(n)
	//return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
