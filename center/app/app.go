package app

import (
	"tron_robot/center/app/internal/tron-task/expenditure"
	"tron_robot/center/app/internal/tron-task/receivable"
	verifytransaction "tron_robot/center/app/internal/tron-task/verify-transaction"
	"tron_robot/internal/dao"
	"xbase/cluster/node"
)

func Init(proxy *node.Proxy) {
	//创建区块拉取需要的表
	dao.InitPaymentWork()
	//tron 区块拉取服务，以及数据验证
	receivable.Instance(proxy)
	expenditure.Instance(proxy)
	verifytransaction.Instance(proxy)
}
