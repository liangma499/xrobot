package model

import "time"

//go:generate xgorm-dao-generator -model-dir=. -model-names=UserParent -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type UserParent struct {
	ID        int64     `gorm:"column:id;size:64;primarykey;autoIncrement"`     // 自增ID
	UID       int64     `gorm:"column:uid;size:64;uniqueIndex:uinque_uid_pid;"` // 父级用户ID
	PID       int64     `gorm:"column:pid;size:64;uniqueIndex:uinque_uid_pid;"` // 直属上线
	Level     int32     `gorm:"column:level;size:32"`                           // 直属等级
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;"`              // 创建时间
}

func (u *UserParent) TableName() string {
	return "user_parent"
}
