package app

import (
	"xbase/cluster/node"
	"xrobot/center/app/internal/tron-task/expenditure"
	"xrobot/center/app/internal/tron-task/receivable"
	verifytransaction "xrobot/center/app/internal/tron-task/verify-transaction"
	"xrobot/internal/dao"
)

func Init(proxy *node.Proxy) {
	//创建区块拉取需要的表
	dao.InitPaymentWork()
	//tron 区块拉取服务，以及数据验证
	receivable.Instance(proxy)
	expenditure.Instance(proxy)
	verifytransaction.Instance(proxy)
}
