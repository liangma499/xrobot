package user

import (
	"tron_robot/internal/xtypes"
)

type RegisterPayload struct {
	UID            int64                 `json:"uid"`            // 用户用户
	Time           int64                 `json:"time"`           // 注册时间
	SupervisorInfo xtypes.SupervisorInfo `json:"supervisorInfo"` // 上级ID
	DeviceID       string                `json:"device_id"`      // 设备id
	RegisterIP     string                `json:"register_ip"`    // 注册ip
}

type LoginPayload struct {
	UID  int64 `json:"uid"`  // 登录用户
	Time int64 `json:"time"` // 登录时间
}
