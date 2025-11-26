package app

import (
	"tron_robot/tools/client/app/logic/lobby"
	"xbase/cluster/client"
)

func Init(proxy *client.Proxy) {
	// 初始化大厅逻辑
	lobby.NewLogic(proxy).Init()

}
