package middleware

import (
	"xbase/cluster/node"
	"xbase/log"

	//optionbaseconfig "tron_robot/internal/option/option-base-config"
	//"tron_robot/internal/xprotocol/common"
	"tron_robot/internal/xtypes"
)

func State(middleware *node.Middleware, ctx node.Context) {
	//state := optionbaseconfig.GetValue(xtypes.ServerStatusKey)
	state := "1"
	switch state {
	case xtypes.Opened.String():
		middleware.Next(ctx)
	case xtypes.Closed.String():
		/*err := ctx.Response(&common.CommonRes{Code: code.ServerIsClosed.Code()})
		if err != nil {
			log.Errorf("response message failed, err: %v", err)
		}*/
		if err := ctx.Disconnect(true); err != nil {
			log.Errorf("disconnect message failed, err: %v", err)
		}

	}
}
