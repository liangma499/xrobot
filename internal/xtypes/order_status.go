package xtypes

type OrderStatus int32

const (
	OrderStatus_None                   OrderStatus = 0  // 无
	OrderStatus_ToBeActivated          OrderStatus = 1  // 待激活
	OrderStatus_ActivateAuthentication OrderStatus = 2  // 激活转账完成，待验证激活是否成功
	OrderStatus_Activated              OrderStatus = 3  // 已激活，代理能量
	OrderStatus_Exchange               OrderStatus = 4  // 兑换订单
	OrderStatus_Finished               OrderStatus = 99 // 完成
)
