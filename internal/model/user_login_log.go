package model

import "time"

//go:generate xgorm-dao-generator -model-dir=. -model-names=UserLoginLog -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type UserLoginLog struct {
	ID      int64     `gorm:"column:id;size:64;primarykey;autoIncrement"` // 主键
	UID     int64     `gorm:"column:uid;size:64"`                         // 用户ID
	LoginIP string    `gorm:"column:login_ip"`                            // 登录IP
	LoginAt time.Time `gorm:"column:login_at;type:timestamp;"`            // 登录时间
}

func (u *UserLoginLog) TableName() string {
	return "user_login_log"
}
