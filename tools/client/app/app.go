package app

import (
	"xbase/cluster/client"
	"xrobot/tools/client/app/logic/lobby"
)

func Init(proxy *client.Proxy) {
	// 初始化大厅逻辑
	lobby.NewLogic(proxy).Init()

}
