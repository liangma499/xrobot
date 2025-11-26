package model

import (
	"time"
	"xrobot/internal/xtypes"
)

//go:generate xgorm-dao-generator -model-dir=. -model-names=UserBase -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type UserBase struct {
	UID                   int64                          `gorm:"column:uid;primaryKey;type:bigint;" json:"uid" redis:"uid"`                                         // 主键
	Code                  string                         `gorm:"column:code;size:16;uinque;primaryKey;" json:"code" redis:"code"`                                   // 编号
	Account               string                         `gorm:"column:account;size:128;index"`                                                                     // 账号
	Email                 string                         `gorm:"column:email;size:128;index"`                                                                       // 邮箱
	Nickname              string                         `gorm:"column:nickname;size:64" json:"nickname" redis:"nickname"`                                          // 昵称                                                                                                        // tg用户名
	Avatar                string                         `gorm:"column:avatar;size:255" json:"avatar" redis:"avatar"`                                               // 头像
	LastLoginType         xtypes.RegisterLoginType       `gorm:"column:last_login_type;size:32" json:"last_login_type" redis:"last_login_type"`                     // 最后登录类型
	RegisterType          xtypes.RegisterLoginType       `gorm:"column:register_type;size:32" json:"register_type" redis:"register_type"`                           // 注册类型
	IsVerifiedEmail       xtypes.UserVerifiedEmailStatus `gorm:"column:is_verified_email;size:32" json:"isVerifiedEmail" redis:"isVerifiedEmail"`                   // 用户是否验证了邮箱
	Birthday              string                         `gorm:"column:birthday;size:128" json:"birthday" redis:"birthday"`                                         // 生日
	UserType              xtypes.UserType                `gorm:"column:user_type;size:32" json:"user_type" redis:"user_type"`                                       // 类型
	Status                xtypes.UserStatus              `gorm:"column:status;size:32" json:"status" redis:"status"`                                                // 状态
	ChannelName           string                         `gorm:"column:channel_name;size:64" json:"channel_name" redis:"channel_name"`                              // 渠道名
	ChannelCode           string                         `gorm:"column:channel_code;size:32" json:"channel_code" redis:"channel_code"`                              // 渠道编码
	SupervisorInfo        xtypes.SupervisorInfo          `gorm:"column:supervisor_info;type:json" json:"supervisorInfo" redis:"supervisorInfo"`                     // 上级ID
	DeviceType            xtypes.DeviceTypeEnum          `gorm:"column:device_type;size:32" json:"deviceType" redis:"deviceType"`                                   // 设备类型(1=android、2=ios、3=web)
	DeviceID              string                         `gorm:"column:device_id;size:128" json:"deviceID" redis:"deviceID"`                                        // 设备ID
	Language              string                         `gorm:"column:language;size:32" json:"language" redis:"language"`                                          // 用户语言
	Country               string                         `gorm:"column:country;size:255" json:"country" redis:"country"`                                            // 国家
	City                  string                         `gorm:"column:city;size:255" json:"city" redis:"city"`                                                     // 城市
	IsSetPasswd           int                            `gorm:"column:is_set_passwd;size:32" json:"is_set_passwd" redis:"is_set_passwd"`                           // 是设置过密码
	TotalLoginTimes       int                            `gorm:"column:total_login_times;size:32" json:"totalLoginTimes" redis:"totalLoginTimes"`                   // 总共登录次数
	TotalLoginDays        int                            `gorm:"column:total_login_days;size:32" json:"totalLoginDays" redis:"totalLoginDays"`                      // 总共登录天数
	ContinuouslyLoginDays int                            `gorm:"column:continuously_login_days;size:32" json:"continuouslyLoginDays" redis:"continuouslyLoginDays"` // 连续登录天数
	LastLoginZero         int64                          `gorm:"column:last_login_zero;size:64;" json:"last_login_zero" redis:"last_login_zero"`                    // 上次登录时间
	LastLoginAt           time.Time                      `gorm:"column:last_login_at;type:timestamp;" json:"lastLoginAt" redis:"lastLoginAt"`                       // 上次登录时间
	LastLoginIP           string                         `gorm:"column:last_login_ip" json:"lastLoginIP" redis:"lastLoginIP"`
	RegisterZero          int64                          `gorm:"column:register_zero;size:64;" json:"register_zero" redis:"register_zero"`
	RegisterIP            string                         `gorm:"column:register_ip;size:64" json:"registerIP" redis:"registerIP"`         // 注册IP
	RegisterAt            time.Time                      `gorm:"column:register_at;type:timestamp;" json:"registerAt" redis:"registerAt"` // 注册时间

}

func (u *UserBase) Clone() *UserBase {
	if u == nil {
		return nil
	}
	return &UserBase{
		UID:                   u.UID, // 主键
		Code:                  u.Code,
		Account:               u.Account,                // 账号
		Email:                 u.Email,                  // 邮箱
		Nickname:              u.Nickname,               // 昵称                                                                                                        // tg用户名
		Avatar:                u.Avatar,                 // 头像
		LastLoginType:         u.LastLoginType,          // 最后登录类型
		RegisterType:          u.RegisterType,           // 注册类型
		IsVerifiedEmail:       u.IsVerifiedEmail,        // 用户是否验证了邮箱
		Birthday:              u.Birthday,               // 生日
		UserType:              u.UserType,               // 类型
		Status:                u.Status,                 // 状态
		ChannelName:           u.ChannelName,            // 渠道名
		ChannelCode:           u.ChannelCode,            // 渠道编码
		SupervisorInfo:        u.SupervisorInfo.Clone(), // 上级ID
		DeviceType:            u.DeviceType,             // 设备类型(1=android、2=ios、3=web)
		DeviceID:              u.DeviceID,               // 设备ID
		Language:              u.Language,               // 用户语言
		Country:               u.Country,                // 国家
		City:                  u.City,                   // 城市
		IsSetPasswd:           u.IsSetPasswd,            // 是设置过密码
		TotalLoginTimes:       u.TotalLoginTimes,        // 总共登录次数
		TotalLoginDays:        u.TotalLoginDays,         // 总共登录天数
		ContinuouslyLoginDays: u.ContinuouslyLoginDays,  // 连续登录天数
		LastLoginZero:         u.LastLoginZero,          // 上次登录时间
		LastLoginAt:           u.LastLoginAt,            // 上次登录时间
		LastLoginIP:           u.LastLoginIP,
		RegisterZero:          u.RegisterZero,
		RegisterIP:            u.RegisterIP, // 注册IP
		RegisterAt:            u.RegisterAt, // 注册时间

	}
}
