package xtypes

type OptionStatus int32

const (
	OptionStatus_Normal  OptionStatus = 1 //可用
	OptionStatus_Disable OptionStatus = 2 //不可用
)

type OptionOperate int32

const (
	OptionOperate_LoadAll OptionOperate = 1 //加载全部
	OptionOperate_Modify  OptionOperate = 2 //修改
	OptionOperate_Delete  OptionOperate = 3 //删除
)
