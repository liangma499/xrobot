package model

//go:generate xgorm-dao-generator -model-dir=. -model-names=User -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type User struct {
	ID             int64  `gorm:"column:id;size:64;primarykey;autoIncrement"` // 主键
	Account        string `gorm:"column:account;size:128;index"`              // 账号
	Email          string `gorm:"column:email;size:128;index"`                // 邮箱
	TelegramUserID int64  `gorm:"column:telegram_user_id;size:64;index"`      //tgUserID
	TelegramPwd    string `gorm:"column:telegram_pwd;size:128"`               // tg密码
	Salt           string `gorm:"column:salt;size:16"`                        // 盐
	Password       string `gorm:"column:password;size:128"`                   // 密码
}
