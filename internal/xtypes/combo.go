package xtypes

import (
	"xbase/utils/xtime"
)

type ComboKind int32

const (
	ComboKind_None    ComboKind = 0 //无
	ComboKind_Hour    ComboKind = 1 //小时
	ComboKind_Day     ComboKind = 2 //天
	ComboKind_Month   ComboKind = 3 //月
	ComboKind_Year    ComboKind = 4 //年
	ComboKind_Forever ComboKind = 5 //永久
)

func (ck ComboKind) Name() string {
	switch ck {
	case ComboKind_Hour:
		{
			return "小时"
		}
	case ComboKind_Day:
		{
			return "天"
		}
	case ComboKind_Month:
		{
			return "月"
		}
	case ComboKind_Year:
		{
			return "年"
		}
	case ComboKind_Forever:
		{
			return "永久"
		}
	}
	return "none"
}

// 过期时间
func (ck ComboKind) ExpirationTime(num int) int64 {

	switch ck {
	case ComboKind_Hour:
		{
			now := xtime.Now().Unix()
			return now + 3600*int64(num)
		}
	case ComboKind_Day:
		{
			return xtime.Day(num).Unix()
		}
	case ComboKind_Month:
		{
			return xtime.Month(num).Unix()
		}
	case ComboKind_Year:
		{
			return xtime.Month(12 * num).Unix()
		}
	case ComboKind_Forever:
		{
			return -1
		}
	}
	return 0
}

// 存在时长
func (ck ComboKind) Duration(num int) int64 {

	expiration := ck.ExpirationTime(num)
	if expiration < 0 {
		return -1
	}
	if expiration == 0 {
		return 0
	}
	expiration -= xtime.Now().Unix()
	if expiration < 0 {
		expiration = 0
	}
	return expiration
}
