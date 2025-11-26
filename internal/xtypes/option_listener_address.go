package xtypes

// 监听地址类型
type AddressKind int32

const (
	AddressKind_None             AddressKind = 0   // 无
	AddressKind_Official         AddressKind = 1   // 官方
	AddressKind_Customize        AddressKind = 2   // 自定义
	AddressKind_Pay              AddressKind = 3   // 付款方是自己的
	AddressKind_OutVerify        AddressKind = 99  // 给人付款验证
	AddressKind_OtherTransaction AddressKind = 999 // 别人付款取消代理能量监听
)
