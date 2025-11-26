package middleware

import (
	"fmt"
	"strings"
	optionbaseconfig "tron_robot/internal/option/option-base-config"
	"tron_robot/internal/xtypes"
	"xbase/cluster/node"
	"xbase/log"
)

func Auth(middleware *node.Middleware, ctx node.Context) {

	uid := ctx.UID()
	if uid == 0 {

		log.Warnf("Auth Disconnect :%v", ctx.CID())
		if err := ctx.Disconnect(true); err != nil {
			log.Errorf("disconnect message failed, err: %v", err)
		}
	} else {
		state := optionbaseconfig.GetValue(xtypes.ServerStatusKey)

		if state == xtypes.Maintain.String() {
			isAllow := false
			ustr := fmt.Sprintf(";%d;", ctx.UID())
			whitelist := optionbaseconfig.GetValue(xtypes.WhiteListKey)
			isAllow = strings.Contains(whitelist, ustr)

			if isAllow {
				middleware.Next(ctx)
			} else {
				/*
					err := ctx.Response(&common.CommonRes{Code: code.ServerIsMaintain.Code()})
					if err != nil {
						log.Errorf("response message failed, err: %v", err)
					}*/
				if err := ctx.Disconnect(true); err != nil {
					log.Errorf("disconnect message failed, err: %v", err)
				}
			}
		} else {
			middleware.Next(ctx)
		}
	}
}
