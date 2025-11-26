package xfloat2str

import (
	"github.com/shopspring/decimal"
)

// 向下取整
func Float64ToSrtFloor(f float64, places int) decimal.Decimal {

	if places < 0 {
		places = 0
	}
	mut := 1
	for i := 0; i < places; i++ {
		mut *= 10
	}
	deMut := decimal.NewFromFloat(float64(mut))
	return decimal.NewFromFloat(f).Mul(deMut).Floor().Div(deMut)

}

// 向下取整
func Float32ToSrtFloor(f float64, places int) decimal.Decimal {
	if places < 0 {
		places = 0
	}
	mut := 1
	for i := 0; i < places; i++ {
		mut *= 10
	}
	deMut := decimal.NewFromFloat(float64(mut))
	return decimal.NewFromFloat(f).Mul(deMut).Floor().Div(deMut)

}
