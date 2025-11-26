package xstr

import "tron_robot/internal/component/snowflake"

// SerialNO 序列号
func SerialNO() string {
	return snowflake.Instance().Generate().String()
}

// Version 版本号
func Version() int64 {
	return snowflake.Instance().Generate().Int64()
}
