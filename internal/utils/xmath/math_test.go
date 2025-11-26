package xmath_test

import (
	"testing"
	"tron_robot/internal/utils/xmath"
)

func TestFloor(t *testing.T) {
	f := 3.1415926

	t.Log(xmath.Floor(f, 2))
}

func TestCeil(t *testing.T) {
	f := 3.1415926

	t.Log(xmath.Ceil(f, 2))
}

func TestRound(t *testing.T) {
	f := 3.1415926

	t.Log(xmath.Round(f, 2))
}
