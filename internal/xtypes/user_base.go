package xtypes

const (
	UserTypeGuest     UserType = 1 // 游客用户
	UserTypeTg        UserType = 2 // TG普通用户
	UserTypeRobot     UserType = 3 // 机器人用户
	UserTypeSystem    UserType = 4 // 系统用户
	UserTypePlayAlong UserType = 5 // 陪玩用户
	UserTypeGeneral   UserType = 6 // 普通用户
)

const (
	UserStatusNormal    UserStatus = 1 // 正常
	UserStatusForbidden UserStatus = 2 // 禁用
)

const (
	UserVerifiedEmailStatusYes UserVerifiedEmailStatus = 1 // 已验证
	UserVerifiedEmailStatusNo  UserVerifiedEmailStatus = 2 // 未验证
)

const (
	Android DeviceTypeEnum = 1 // 安卓
	Ios     DeviceTypeEnum = 2 // ios
	Web     DeviceTypeEnum = 3 // 匿名
)

// UserType 用户类型
type UserType int

// UserStatus 用户状态
type UserStatus int

// UserVerifiedEmailStatus 用户验证邮箱状态
type UserVerifiedEmailStatus int

type DeviceTypeEnum int32

type RegisterLoginType int32

const (
	RegisterLoginType_Guest   RegisterLoginType = 1
	RegisterLoginType_TG      RegisterLoginType = 2
	RegisterLoginType_Email   RegisterLoginType = 3
	RegisterLoginType_Twitter RegisterLoginType = 4
	RegisterLoginType_Wallet  RegisterLoginType = 5
	RegisterLoginType_Google  RegisterLoginType = 6
	RegisterLoginType_Robot   RegisterLoginType = 99
)
