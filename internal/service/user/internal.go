package user

import (
	"context"
	"xrobot/internal/code"
	redisdefault "xrobot/internal/component/redis/redis-default"
	optionChannelDao "xrobot/internal/dao/option-channel"
	userDao "xrobot/internal/dao/user"
	userBaseDao "xrobot/internal/dao/user-base"
	userLoginDayStatDao "xrobot/internal/dao/user-login-day-stat"
	userLoginLogDao "xrobot/internal/dao/user-login-log"
	userParentDao "xrobot/internal/dao/user-parent"
	userevt "xrobot/internal/event/user"

	"fmt"
	"xbase/cache"
	"xbase/errors"
	"xbase/log"
	"xbase/utils/xconv"
	"xbase/utils/xrand"
	"xbase/utils/xtime"
	"xrobot/internal/model"
	"xrobot/internal/service/user/pb"
	"xrobot/internal/utils/xcrypt"
	"xrobot/internal/xtypes"

	"gorm.io/gorm"
)

// 执行注册操作
func (s *Server) doRegister(ctx context.Context, args *registerArgs, deviceInfo *pb.DeviceInfo) (*model.User, *model.UserBase, error) {
	var (
		err             error
		salt            string
		password        string
		avatar          = args.avatar
		birthday        string
		supervisorInfos = make(xtypes.SupervisorInfo)
		supervisorUser  *model.UserBase
	)

	if args.inviteCode != "" {
		supervisorUser, err = userBaseDao.Instance().DoGetUserBaseByCode(ctx, args.inviteCode)
		if err != nil {
			return nil, nil, errors.NewError(err, code.InviteCodeNotExist)
		}

		if supervisorUser == nil {
			return nil, nil, errors.NewError(code.InvalidArgument)
		}
		supervisorInfos = supervisorUser.SupervisorInfo.Supervisor(supervisorUser.UID)

	}

	if args.password != "" {
		salt, password, err = xtypes.EncryptPassword(args.password)
		if err != nil {
			return nil, nil, err
		}
	}

	ucode, err := s.doGenUserCode(ctx)
	if err != nil {
		return nil, nil, err
	}

	if avatar == "" {
		avatar = s.doGenAvatar()
	}
	channelCode := args.channelCode
	channelName := args.channelName

	if channelCfg, err := optionChannelDao.Instance().GetChannel(ctx, channelCode); err == nil && channelCfg != nil {

		channelName = channelCfg.Name
		channelCode = channelCfg.ChannelCode

	}

	if args.birthday != "" {
		birthday = args.birthday
	} else {
		birthday = s.doGenBirthday()
	}
	now := xtime.Now()
	zeroTime := xtime.DayHead(0).Unix()
	user := &model.User{
		Account:        args.account,        // 账号
		Email:          args.email,          // 邮箱
		TelegramUserID: args.telegramUserID, //tgUserID
		TelegramPwd:    args.telegramPwd,    // tg密码
		Salt:           salt,                // 盐
		Password:       password,            // 密码
	}

	if _, err = userDao.Instance().Insert(ctx, user); err != nil {
		log.Errorf("insert user failed, user = %s err = %v", xconv.Json(user), err)
		return nil, nil, errors.NewError(err, code.InternalError)
	}
	isSetPasswd := 0
	if args.email != "" && password != "" {
		isSetPasswd = 1
	}
	baseUser := &model.UserBase{
		UID:                   user.ID,                           // 主键
		Code:                  xconv.String(user.TelegramUserID), // 编号
		Account:               user.Account,                      // 账号
		Email:                 user.Email,                        //  邮箱
		Nickname:              args.telegramuserName,             // 昵称
		Avatar:                avatar,                            // 头像
		LastLoginType:         args.registerLoginType,            // 最后登录类型
		RegisterType:          args.registerLoginType,
		IsVerifiedEmail:       args.isVerifiedEmail,    // 用户是否验证了邮箱
		Birthday:              birthday,                // 生日
		UserType:              args.userType,           // 类型
		Status:                xtypes.UserStatusNormal, // 状态
		ChannelName:           channelName,             // 注册渠道
		ChannelCode:           channelCode,             // 渠道编码
		SupervisorInfo:        supervisorInfos.Clone(), // 上级ID
		RegisterIP:            args.clientIP,           // 注册IP
		RegisterAt:            now,                     // 注册时间
		RegisterZero:          zeroTime,                // 上次登录时间
		TotalLoginTimes:       1,                       // 总共登录次数
		TotalLoginDays:        1,                       // 总共登录天数
		ContinuouslyLoginDays: 1,                       // 连续登录天数
		Language:              args.language,           // 用户语言
		IsSetPasswd:           isSetPasswd,             // 是设置过密码
		LastLoginZero:         zeroTime,                // 上次登录时间
		LastLoginAt:           now,                     // 上次登录时间
		LastLoginIP:           args.clientIP,           // 上次登录IP
	}
	if deviceInfo != nil {
		baseUser.Country = deviceInfo.Country
		baseUser.City = deviceInfo.City
		baseUser.DeviceType = xtypes.DeviceTypeEnum(deviceInfo.DeviceType)
		baseUser.DeviceID = deviceInfo.DeviceID
	}
	if _, err = userBaseDao.Instance().Insert(ctx, baseUser); err != nil {
		log.Errorf("insert baseUser failed, user = %s err = %v", xconv.Json(user), err)
		userDao.Instance().Delete(ctx, func(cols *userDao.Columns) any {
			return map[string]any{
				cols.ID: user.ID,
			}
		})
		return nil, nil, errors.NewError(err, code.InternalError)
	}

	s.doUnlockUserCode(ctx, ucode)

	s.doInitUserParent(ctx, baseUser.UID, supervisorUser)

	if args.userType != xtypes.UserTypeRobot {
		userevt.PublishRegister(&userevt.RegisterPayload{
			UID:            baseUser.UID,
			Time:           baseUser.RegisterAt.Unix(),
			SupervisorInfo: baseUser.SupervisorInfo,
			DeviceID:       baseUser.DeviceID,
			RegisterIP:     baseUser.RegisterIP,
		})
	} else {
		redisdefault.Instance().SAdd(ctx, robotSetUserId, baseUser.UID)
	}

	return user, baseUser, nil
}

// 初始化用户代上下级关系表
func (s *Server) doInitUserParent(ctx context.Context, uid int64, superiorUser *model.UserBase) {
	if superiorUser == nil {
		return
	}

	if superiorUser.UID > 0 {
		columns := fmt.Sprintf("%d AS %s, %s+1 AS %s,%s",
			uid,
			userParentDao.Instance().Columns.UID,
			userParentDao.Instance().Columns.Level,
			userParentDao.Instance().Columns.Level,
			userParentDao.Instance().Columns.PID,
		)

		query := fmt.Sprintf(
			"%s = ?",
			userParentDao.Instance().Columns.UID,
		)
		resData := make([]*model.UserParent, 0)

		err := userParentDao.Instance().Table.WithContext(ctx).Select(columns).Where(query, superiorUser.UID).Find(&resData).Error

		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				log.Errorf("init , uid = %d err = %#v", uid, superiorUser)
				return
			}

		}
		resData = append(resData, &model.UserParent{
			UID:   uid,
			PID:   superiorUser.UID,
			Level: 1,
		})
		userParentDao.Instance().Insert(ctx, resData...)
	}

}

// 执行登录逻辑
func (s *Server) doLogin(ctx context.Context, user *model.User, clientIP string, deviceInfo *pb.DeviceInfo, lastLoginType xtypes.RegisterLoginType) (int64, error) {
	userbase, err := userBaseDao.Instance().FindOne(ctx, func(cols *userBaseDao.Columns) any {
		return map[string]any{
			cols.UID: user.ID,
		}
	})
	if err != nil {
		return 0, errors.NewError(code.InternalError)
	}
	if userbase == nil {
		return 0, errors.NewError(code.UserForbidden)
	}
	if userbase.Status != xtypes.UserStatusNormal {
		return 0, errors.NewError(code.UserForbidden)
	}

	var (
		zeroTime      = xtime.DayHead(0).Unix()
		now           = xtime.Now()
		yesterdayZero = zeroTime - 86400
		lastLoginDate = userbase.LastLoginZero
	)

	updates := map[string]any{
		userBaseDao.Instance().Columns.TotalLoginTimes: gorm.Expr(fmt.Sprintf("%s + ?", userBaseDao.Instance().Columns.TotalLoginTimes), 1),
		userBaseDao.Instance().Columns.LastLoginAt:     now,
		userBaseDao.Instance().Columns.LastLoginZero:   zeroTime,
		userBaseDao.Instance().Columns.LastLoginIP:     clientIP,
		userBaseDao.Instance().Columns.LastLoginType:   lastLoginType,
	}

	if lastLoginDate != zeroTime {
		updates[userBaseDao.Instance().Columns.TotalLoginDays] = gorm.Expr(fmt.Sprintf("%s + ?", userBaseDao.Instance().Columns.TotalLoginDays), 1)
	}
	if deviceInfo != nil {
		updates[userBaseDao.Instance().Columns.City] = deviceInfo.City
		updates[userBaseDao.Instance().Columns.Country] = deviceInfo.Country
		updates[userBaseDao.Instance().Columns.DeviceID] = deviceInfo.DeviceID
		updates[userBaseDao.Instance().Columns.DeviceType] = deviceInfo.DeviceType
	}

	if lastLoginDate == yesterdayZero {
		updates[userBaseDao.Instance().Columns.ContinuouslyLoginDays] = gorm.Expr(fmt.Sprintf("%s + ?", userBaseDao.Instance().Columns.ContinuouslyLoginDays), 1)
	} else {
		updates[userBaseDao.Instance().Columns.ContinuouslyLoginDays] = 1
	}

	err = userBaseDao.Instance().DoUpdateUserBase(ctx, userbase.UID, updates)
	if err != nil {
		return 0, err
	}

	_, err = userLoginLogDao.Instance().Insert(ctx, &model.UserLoginLog{
		UID:     user.ID,
		LoginAt: now,
		LoginIP: clientIP,
	})
	if err != nil {
		return 0, errors.NewError(err, code.InternalError)
	}
	err = userLoginDayStatDao.Instance().InsertOrUpdate(ctx, userbase.UID, zeroTime)
	if err != nil {
		return 0, errors.NewError(err, code.InternalError)
	}

	// 发布登录事件
	userevt.PublishLogin(&userevt.LoginPayload{
		UID:  userbase.UID,
		Time: now.Unix(),
	})

	return userbase.UID, nil
}

// 生成用户编号
func (s *Server) doGenUserCode(ctx context.Context) (string, error) {
	for {
		ucode := xrand.Digits(defaultUserCodeLength)

		key := cache.AddPrefix(fmt.Sprintf(cacheUserCodeLockKey, ucode))

		val, err := s.redis.SetEX(ctx, key, cacheUserCodeLockValue, cacheUserCodeLockExpiration).Result()
		if err != nil {
			log.Errorf("lock user code failed, user code = %s , err = %v", ucode, err)
			return "", errors.NewError(code.InternalError, err)
		}

		if val == "0" {
			continue
		}

		count, err := userBaseDao.Instance().Count(ctx, func(cols *userBaseDao.Columns) any {
			return map[string]any{
				cols.Code: ucode,
			}
		})
		if err != nil {
			log.Errorf("count user failed, user code = %s err = %v", ucode, err)
			return "", errors.NewError(err, code.InternalError)
		}

		if count == 0 {
			return ucode, nil
		}
	}
}

// 对比密码
func (s *Server) doComparePassword(hashed, password, salt string) (bool, error) {
	ok, err := xcrypt.Compare(hashed, password, salt)
	if err != nil {
		log.Errorf("compare password failed, hashedPassword = %s, password = %s salt = %s err = %s", hashed, password, salt, err)
		return false, errors.NewError(err, code.InternalError)
	}

	return ok, nil
}

// 解锁用户编号
func (s *Server) doUnlockUserCode(ctx context.Context, ucode string) {
	s.redis.Del(ctx, cache.AddPrefix(fmt.Sprintf(cacheUserCodeLockKey, ucode)))
}

// 生成昵称
/*
func (s *Server) doGenNickname() string {
	return "U" + xrand.Digits(8)
}
*/
// 生成头像
func (s *Server) doGenAvatar() string {
	return ""
}

// 生成生日
func (s *Server) doGenBirthday() string {
	return xtime.Now().AddDate(0-xrand.Int(0, 50), 0-xrand.Int(0, 12), 0-xrand.Int(0, 31)).Format(xtime.DatetimeLayout)
}
