package model

//go:generate xgorm-dao-generator -model-dir=. -model-names=UserLoginDayStat -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type UserLoginDayStat struct {
	ID        int64 `gorm:"column:id;size:64;primarykey;autoIncrement;"`               // 主键
	ZeroTime  int64 `gorm:"column:zero_time;size:64;uniqueIndex:uinque_zerotime_uid;"` // 登录时间
	UID       int64 `gorm:"column:uid;size:64;uniqueIndex:uinque_zerotime_uid;"`       // 用户ID
	LoginTime int64 `gorm:"column:login_time;size:64;"`                                // 登录次数
}

func (u *UserLoginDayStat) TableName() string {
	return "user_login_day_stat"
}
