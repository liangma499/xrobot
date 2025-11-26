package user

import (
	"context"
	"tron_robot/internal/code"
	"tron_robot/internal/event/message"

	"fmt"
	"time"
	identitycomp "tron_robot/internal/component/google/identity"
	jwtcomp "tron_robot/internal/component/jwt"
	"tron_robot/internal/xtypes"
	"xbase/cache"
	"xbase/cluster/mesh"
	"xbase/errors"
	"xbase/log"
	"xbase/utils/jwt"
	"xbase/utils/xconv"

	"tron_robot/internal/model"
	"tron_robot/internal/service/user/pb"

	"tron_robot/internal/utils/xresource"
	"xbase/utils/xtime"

	optionChannelDao "tron_robot/internal/dao/option-channel"
	userDao "tron_robot/internal/dao/user"
	userBaseDao "tron_robot/internal/dao/user-base"

	"github.com/go-redis/redis/v8"
)

const (
	serviceName = "user" // 服务名称
	servicePath = "User" // 服务路径要与pb中的服务路径保持一致
)

const (
	defaultUserCodeLength                = 7                   // 默认的用户编号长度
	cacheUserMailToUIDKey                = "user:mail:%s:uid"  // 用户邮箱对就的用户ID
	cacheUserCodeLockKey                 = "user:code:%s:lock" // 用户编号锁
	cacheUserCodeLockValue               = "1"                 // 用户编号锁值
	cacheUserCodeLockExpiration          = 10 * time.Second    // 用户编号锁过期时间
	cacheUserBridgeTokenKey              = "bridgeToken:%s"    // %s=token
	cacheUserBridgeTmpTokenKeyExpiration = 5 * time.Minute
	cacheUserBridgeTokenKeyExpiration    = 1 * time.Hour
	cacheTwitterOAuthKey                 = "twitter:oauth:%s" // 推特授权信息
	cacheTwitterOAuthExpiration          = 20 * time.Minute   // 推特授权信息过期时间
	maxLevel                             = 2                  // 等级
	robotSetUserId                       = "robotUserIDKey"   // 推特授权信息
)

var _ pb.UserAble = &Server{}

type Server struct {
	proxy    *mesh.Proxy
	jwt      *jwt.JWT
	identity *identitycomp.Identity
	redis    redis.UniversalClient
}

func NewServer(proxy *mesh.Proxy) *Server {
	return &Server{
		proxy:    proxy,
		jwt:      jwtcomp.Instance(),
		identity: identitycomp.Instance(),
		redis:    cache.Client().(redis.UniversalClient),
	}
}

func (s *Server) Init() {
	s.proxy.AddServiceProvider(serviceName, servicePath, s)
	message.SubscribeMessageStart(s.doSubscribeMessageStart)
}

// Register 注册
func (s *Server) Register(ctx context.Context, args *pb.RegisterArgs, reply *pb.RegisterReply) error {
	if true {
		return nil
	}
	user, err := userDao.Instance().FindOneByEmail(ctx, args.Email)
	if err != nil {
		return errors.NewError(err, code.InternalError)
	}

	if user != nil {
		return errors.NewError(code.EmailExists)
	}

	_, userBase, err := s.doRegister(ctx, &registerArgs{
		email:             args.Email,
		registerLoginType: xtypes.RegisterLoginType_Email,
		password:          args.Password,
		channelName:       args.Channel,
		clientIP:          args.ClientIP,
		inviteCode:        args.InviteCode,
		userType:          xtypes.UserTypeGeneral,
		isVerifiedEmail:   xtypes.UserVerifiedEmailStatusNo,
		channelCode:       args.Channel,
		avatar:            xresource.RandAvatarUrl(12),
	}, args.DeviceInfo)
	if err != nil {
		return err
	}
	if userBase == nil {
		return errors.NewError(err, code.InternalError)
	}
	reply.UID = userBase.UID
	return nil
}

// Login 登录
func (s *Server) Login(ctx context.Context, args *pb.LoginArgs, reply *pb.LoginReply) error {
	if args.Email == "" {
		return errors.NewError(code.InvalidArgument)
	}

	user, err := userDao.Instance().FindOne(ctx, func(cols *userDao.Columns) any {
		return map[string]any{
			cols.Email: args.Email,
		}
	})
	if err != nil {
		log.Errorf("find user failed, email = %s err = %v", args.Email, err)
		return errors.NewError(err, code.InternalError)
	}

	if user == nil {
		return errors.NewError(code.IncorrectAccountOrPassword)
	}

	ok, err := s.doComparePassword(user.Password, args.Password, user.Salt)
	if err != nil {
		return err
	}

	if !ok {
		return errors.NewError(code.IncorrectAccountOrPassword)
	}

	uid, err := s.doLogin(ctx, user, args.ClientIP, args.DeviceInfo, xtypes.RegisterLoginType_Email)
	if err != nil {
		return err
	}

	reply.UID = uid

	return nil
}

// Logout 退出登录
func (s *Server) Logout(ctx context.Context, args *pb.LogoutArgs, reply *pb.LogoutReply) error {
	err := s.jwt.DestroyIdentity(args.UID)
	if err != nil {
		log.Errorf("destroy user's token failed: %v", err)
		return errors.NewError(err, code.InternalError)
	}

	return nil
}

// ModifyAvatar 修改头像
func (s *Server) ModifyAvatar(ctx context.Context, args *pb.ModifyAvatarArgs, reply *pb.ModifyAvatarReply) error {
	user, err := userBaseDao.Instance().GetUserBase(ctx, args.UID)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.NewError(code.NotFound)
	}

	return userBaseDao.Instance().DoUpdateUserBase(ctx, args.UID, map[string]any{
		userBaseDao.Instance().Columns.Avatar: args.Avatar,
	})
}

// ModifyNickname 修改昵称
func (s *Server) ModifyNickname(ctx context.Context, args *pb.ModifyNicknameArgs, reply *pb.ModifyNicknameReply) error {
	user, err := userBaseDao.Instance().GetUserBase(ctx, args.UID)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.NewError(code.NotFound)
	}

	return userBaseDao.Instance().DoUpdateUserBase(ctx, args.UID, map[string]any{
		userBaseDao.Instance().Columns.Nickname: args.Nickname,
	})
}

// ModifyPassword 修改密码
func (s *Server) ModifyPassword(ctx context.Context, args *pb.ModifyPasswordArgs, reply *pb.ModifyPasswordReply) error {
	user, err := userDao.Instance().FindOne(ctx, func(cols *userDao.Columns) any {
		return map[string]any{
			cols.ID: args.UID,
		}

	})
	if err != nil {
		return err
	}

	if user == nil {
		return errors.NewError(code.NotFound)
	}

	ok, err := s.doComparePassword(user.Password, args.OldPassword, user.Salt)
	if err != nil {
		return err
	}

	if !ok {
		return errors.NewError(code.IncorrectPassword)
	}

	salt, password, err := xtypes.EncryptPassword(args.NewPassword)
	if err != nil {
		return err
	}
	_, err = userDao.Instance().Update(ctx, func(cols *userDao.Columns) any {
		return map[string]any{
			cols.ID: args.UID,
		}

	}, func(cols *userDao.Columns) any {
		return map[string]any{
			cols.Salt:     salt,
			cols.Password: password,
		}
	})
	if err != nil {
		return err
	}

	// 修改密码后清除Token
	_ = s.jwt.DestroyIdentity(args.UID)

	return nil
}

// 设置密码
func (s *Server) SetPassword(ctx context.Context, args *pb.SetPasswordArgs, reply *pb.SetPasswordReply) error {
	user, err := userDao.Instance().FindOne(ctx, func(cols *userDao.Columns) any {
		return map[string]any{
			cols.ID: args.UID,
		}

	})
	if err != nil {
		return errors.NewError(code.InternalError)
	}
	if user == nil {
		return errors.NewError(code.NotFound)
	}
	if user.Email == "" {
		return errors.NewError(code.NotFound)
	}
	salt, password, err := xtypes.EncryptPassword(args.Password)
	if err != nil {
		return errors.NewError(code.EmailNotExists)
	}
	_, err = userDao.Instance().Update(ctx, func(cols *userDao.Columns) any {
		return map[string]any{
			cols.ID: args.UID,
		}

	}, func(cols *userDao.Columns) any {
		return map[string]any{
			cols.Salt:     salt,
			cols.Password: password,
		}
	})
	if err != nil {
		return errors.NewError(code.InternalError)
	}
	// 修改密码后清除Token
	//_ = s.jwt.DestroyIdentity(user.ID)
	return nil
}

// ModifyBirthday 修改生日
func (s *Server) ModifyBirthday(ctx context.Context, args *pb.ModifyBirthdayArgs, reply *pb.ModifyBirthdayReply) error {
	birthday, err := xtime.Parse(xtime.DateLayout, args.Birthday)
	if err != nil {
		return errors.NewError(err, code.InvalidArgument)
	}

	user, err := userBaseDao.Instance().GetUserBase(ctx, args.UID)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.NewError(code.NotFound)
	}

	return userBaseDao.Instance().DoUpdateUserBase(ctx, args.UID, map[string]any{
		userBaseDao.Instance().Columns.Birthday: birthday,
	})
}

// ModifyEmailVerified 修改邮箱已验证状态
func (s *Server) ModifyEmailVerified(ctx context.Context, args *pb.ModifyEmailVerifiedArgs, reply *pb.ModifyEmailVerifiedReply) error {
	user, err := userBaseDao.Instance().FindOneByEmail(ctx, args.Email)
	if err != nil {
		return errors.NewError(err, code.InternalError)
	}

	if user == nil {
		return errors.NewError(code.NotFound)
	}

	if user.IsVerifiedEmail == xtypes.UserVerifiedEmailStatusYes {
		return nil
	}

	err = userBaseDao.Instance().DoUpdateUserBase(ctx, user.UID, map[string]any{
		userBaseDao.Instance().Columns.IsVerifiedEmail: xtypes.UserVerifiedEmailStatusYes,
	})
	if err != nil {
		return err
	}

	return nil
}

// ModifyLanguage 修改语言
func (s *Server) ModifyLanguage(ctx context.Context, args *pb.ModifyLanguageArgs, reply *pb.ModifyLanguageReply) error {
	user, err := userBaseDao.Instance().GetUserBase(ctx, args.UID)
	if err != nil {
		return errors.NewError(err, code.InternalError)
	}

	if user == nil {
		return errors.NewError(code.NotFound)
	}

	err = userBaseDao.Instance().DoUpdateUserBase(ctx, user.UID, map[string]any{
		userBaseDao.Instance().Columns.Language: args.Language,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) ModifyUserType(ctx context.Context, args *pb.ModifyUserTypeArgs, reply *pb.ModifyUserTypeReply) error {
	user, err := userBaseDao.Instance().GetUserBase(ctx, args.UID)
	if err != nil {
		return errors.NewError(err, code.InternalError)
	}

	if user == nil {
		return errors.NewError(code.NotFound)
	}

	err = userBaseDao.Instance().DoUpdateUserBase(ctx, user.UID, map[string]any{
		userBaseDao.Instance().Columns.UserType: args.UserType,
	})
	if err != nil {
		return err
	}

	return nil
}

// ResetPassword 重置密码
func (s *Server) ResetPassword(ctx context.Context, args *pb.ResetPasswordArgs, reply *pb.ResetPasswordReply) error {
	user, err := userDao.Instance().FindOneByEmail(ctx, args.Email)
	if err != nil {
		return errors.NewError(err, code.InternalError)
	}

	if user == nil {
		return errors.NewError(code.NotFound)
	}

	salt, password, err := xtypes.EncryptPassword(args.Password)
	if err != nil {
		return err
	}
	_, err = userDao.Instance().Update(ctx, func(cols *userDao.Columns) any {
		return map[string]any{
			cols.ID: user.ID,
		}
	}, func(cols *userDao.Columns) any {
		return map[string]any{
			cols.Salt:     salt,
			cols.Password: password,
		}
	})

	if err != nil {
		return err
	}

	// 修改密码后清除Token
	_ = s.jwt.DestroyIdentity(user.ID)

	return nil
}

// FetchUser 拉取用户
func (s *Server) FetchUser(ctx context.Context, args *pb.FetchUserArgs, reply *pb.FetchUserReply) error {
	var (
		err      error
		userBase *model.UserBase
	)

	switch {
	case args.UID != 0:
		userBase, err = userBaseDao.Instance().GetUserBase(ctx, args.UID)
	case args.Code != "":
		userBase, err = userBaseDao.Instance().DoGetUserBaseByCode(ctx, args.Code)
	default:
		return errors.NewError(code.InvalidArgument)
	}
	if err != nil {
		return err
	}

	if userBase == nil {
		return errors.NewError(code.NotFound)
	}

	isSetPasswd := false
	if userBase.Email != "" && userBase.IsSetPasswd == 0 {
		isSetPasswd = true
	}
	reply.User = &pb.UserInfo{
		ID:              userBase.UID,
		Code:            userBase.Code,
		Nickname:        userBase.Nickname,
		Avatar:          userBase.Avatar,
		AvatarUrl:       xresource.ToAvatarUrl(userBase.Avatar),
		Birthday:        userBase.Birthday,
		Email:           userBase.Email,
		IsVerifiedEmail: userBase.IsVerifiedEmail == xtypes.UserVerifiedEmailStatusYes,
		SupervisorInfo:  userBase.SupervisorInfo,
		Status:          int32(userBase.Status),
		RegisterAt:      userBase.RegisterAt.Unix(),
		UserType:        int32(userBase.UserType),
		DeviceID:        userBase.DeviceID,
		RegisterIP:      userBase.RegisterIP,
		IsSetPasswd:     isSetPasswd,
		RegisterType:    int32(userBase.RegisterType),
		Language:        userBase.Language,
		ChannelName:     userBase.ChannelName,
		ChannelCode:     userBase.ChannelCode,
		DeviceType:      int32(userBase.DeviceType),
		Country:         userBase.Country,
		City:            userBase.City,
		LastLoginIP:     userBase.LastLoginIP,
		LastLoginAt:     userBase.LastLoginAt.Unix(),
	}

	return nil
}

// FetchUserProfile 拉取用户资料
func (s *Server) FetchUserProfile(ctx context.Context, args *pb.FetchUserProfileArgs, reply *pb.FetchUserProfileReply) error {
	var (
		err      error
		userBase *model.UserBase
	)

	switch {
	case args.UID != 0:
		userBase, err = userBaseDao.Instance().GetUserBase(ctx, args.UID)
	case args.Code != "":
		userBase, err = userBaseDao.Instance().DoGetUserBaseByCode(ctx, args.Code)
	default:
		return errors.NewError(code.InvalidArgument)
	}
	if err != nil {
		return err
	}

	if userBase == nil {
		return errors.NewError(code.NotFound)
	}

	if userBase.Status != xtypes.UserStatusNormal {
		return errors.NewError(code.UserForbidden)
	}

	isSetPasswd := false
	if userBase.Email != "" && userBase.IsSetPasswd == 0 {
		isSetPasswd = true
	}
	reply.User = &pb.UserInfo{
		ID:              userBase.UID,
		Code:            userBase.Code,
		Nickname:        userBase.Nickname,
		Avatar:          userBase.Avatar,
		AvatarUrl:       xresource.ToAvatarUrl(userBase.Avatar),
		Birthday:        userBase.Birthday,
		Email:           userBase.Email,
		IsVerifiedEmail: userBase.IsVerifiedEmail == xtypes.UserVerifiedEmailStatusYes,
		SupervisorInfo:  userBase.SupervisorInfo,
		Status:          int32(userBase.Status),
		RegisterAt:      userBase.RegisterAt.Unix(),
		UserType:        int32(userBase.UserType),
		DeviceID:        userBase.DeviceID,
		RegisterIP:      userBase.RegisterIP,
		IsSetPasswd:     isSetPasswd,
		RegisterType:    int32(userBase.RegisterType),
		Language:        userBase.Language,
		ChannelName:     userBase.ChannelName,
		ChannelCode:     userBase.ChannelCode,
		DeviceType:      int32(userBase.DeviceType),
		Country:         userBase.Country,
		City:            userBase.City,
		LastLoginIP:     userBase.LastLoginIP,
		LastLoginAt:     userBase.LastLoginAt.Unix(),
	}

	return nil
}

// ExistsEmail 检测是否存在邮箱
func (s *Server) ExistsEmail(ctx context.Context, args *pb.ExistsEmailArgs, reply *pb.ExistsEmailReply) error {
	count, err := userDao.Instance().CountByEmail(ctx, args.Email)
	if err != nil {
		return errors.NewError(err, code.InternalError)
	}

	reply.OK = count > 0

	return nil
}

// ValidateToken 验证Token
func (s *Server) ValidateToken(ctx context.Context, args *pb.ValidateTokenArgs, reply *pb.ValidateTokenReply) error {
	identity, err := s.jwt.ExtractIdentity(args.Token)
	if err != nil {
		switch {
		case jwt.IsMissingToken(err):
			return errors.NewError(err, code.Unauthorized)
		case jwt.IsInvalidToken(err):
			return errors.NewError(err, code.Unauthorized)
		case jwt.IsIdentityMissing(err):
			return errors.NewError(err, code.Unauthorized)
		case jwt.IsExpiredToken(err):
			return errors.NewError(err, code.AuthorizationExpired)
		case jwt.IsAuthElsewhere(err):
			return errors.NewError(err, code.AuthorizationElsewhere)
		default:
			return errors.NewError(err, code.Unauthorized)
		}
	}

	userBase, err := userBaseDao.Instance().GetUserBase(ctx, xconv.Int64(identity))
	if err != nil {
		return err
	}
	isSetPasswd := false
	if userBase.Email != "" && userBase.IsSetPasswd == 0 {
		isSetPasswd = true
	}
	reply.User = &pb.UserInfo{
		ID:              userBase.UID,
		Code:            userBase.Code,
		Nickname:        userBase.Nickname,
		Avatar:          userBase.Avatar,
		AvatarUrl:       xresource.ToAvatarUrl(userBase.Avatar),
		Birthday:        userBase.Birthday,
		Email:           userBase.Email,
		IsVerifiedEmail: userBase.IsVerifiedEmail == xtypes.UserVerifiedEmailStatusYes,
		SupervisorInfo:  userBase.SupervisorInfo,
		Status:          int32(userBase.Status),
		RegisterAt:      userBase.RegisterAt.Unix(),
		UserType:        int32(userBase.UserType),
		DeviceID:        userBase.DeviceID,
		RegisterIP:      userBase.RegisterIP,
		IsSetPasswd:     isSetPasswd,
		RegisterType:    int32(userBase.RegisterType),
		Language:        userBase.Language,
		ChannelName:     userBase.ChannelName,
		ChannelCode:     userBase.ChannelCode,
		DeviceType:      int32(userBase.DeviceType),
		Country:         userBase.Country,
		City:            userBase.City,
		LastLoginIP:     userBase.LastLoginIP,
		LastLoginAt:     userBase.LastLoginAt.Unix(),
	}

	return nil
}

// 绑定邮箱
func (s *Server) BindEmail(ctx context.Context, args *pb.BindEmailArgs, reply *pb.BindEmailReply) error {
	user, err := userDao.Instance().FindOneByEmail(ctx, args.Email)
	if err != nil {
		log.Errorf("query user failed, uid = %d err = %v", args.Uid, err)
		return errors.NewError(err, code.InternalError)
	}
	if user == nil {
		log.Errorf("query user failed, uid = %d err = %v", args.Uid, err)
		return errors.NewError(err, code.AccountNotExist)
	} else {
		if user.ID > 0 {
			if user.ID != args.Uid {
				log.Errorf("query user failed, uid = %d err = %v", args.Uid, err)
				return errors.NewError(err, code.EmailAlreadyBind)
			} else {
				if args.PassWd != "" {
					_, err = userDao.Instance().Update(ctx, func(cols *userDao.Columns) any {
						return map[string]any{
							cols.ID: user.ID,
						}
					}, func(cols *userDao.Columns) any {
						updateUser := map[string]any{
							cols.Email: args.Email,
						}
						if args.PassWd != "" {
							salt, password, err := xtypes.EncryptPassword(args.PassWd)
							if err != nil {
								return err
							}
							updateUser[cols.Salt] = salt
							updateUser[cols.Password] = password
						}
						return updateUser
					})
					if err != nil {
						return errors.NewError(err, code.InternalError)
					}
				}
				return nil
			}
		}
	}
	_, err = userDao.Instance().Update(ctx, func(cols *userDao.Columns) any {
		return map[string]any{
			cols.ID: user.ID,
		}
	}, func(cols *userDao.Columns) any {
		updateUser := map[string]any{
			cols.Email: args.Email,
		}
		if args.PassWd != "" {
			salt, password, err := xtypes.EncryptPassword(args.PassWd)
			if err != nil {
				return err
			}
			updateUser[cols.Salt] = salt
			updateUser[cols.Password] = password
		}
		return updateUser
	})
	if err != nil {
		return errors.NewError(err, code.InternalError)
	}
	update := map[string]any{
		userBaseDao.Instance().Columns.Email:           args.Email,
		userBaseDao.Instance().Columns.IsVerifiedEmail: xtypes.UserVerifiedEmailStatusYes,
	}

	err = userBaseDao.Instance().DoUpdateUserBase(ctx, args.Uid, update)
	if err != nil {
		userDao.Instance().Update(ctx, func(cols *userDao.Columns) any {
			return map[string]any{
				cols.ID: user.ID,
			}
		}, func(cols *userDao.Columns) any {
			return map[string]any{
				cols.Email:    user.Email,
				cols.Salt:     user.Salt,
				cols.Password: user.Password,
			}
		})
		return errors.NewError(err, code.InternalError)
	}
	return nil
}

// TG登录
func (s *Server) TelegramLogin(ctx context.Context, args *pb.TelegramLoginArgs, reply *pb.TelegramLoginReply) error {
	/*
		if args.UserID == "" {
			return errors.NewError(code.InvalidArgument)
		}
		user, err := userDao.Instance().FindOneByTelegramUserID(ctx, args.UserID)
		if err != nil {
			return errors.NewError(code.InternalError)
		}
		var userBase *model.UserBase
		if user == nil {

			user, userBase, err = s.doRegister(ctx, &registerArgs{
				tguserName:        args.UserName,
				tgPwd:             xrand.Str(xrand.LetterSeed+xrand.DigitSeed, 32),
				registerLoginType: xtypes.RegisterLoginType_TG,
				tguserID:          args.UserID,
				channelName:       args.Channel,
				inviteCode:        args.InviteCode,
				userType:          xtypes.UserTypeTg,
				avatar:            args.Avatar,
				clientIP:          args.ClientIP,
				isVerifiedEmail:   xtypes.UserVerifiedEmailStatusNo,
				channelCode:       args.ChannelCode,
			}, args.DeviceInfo)
			if err != nil {
				return errors.NewError(err, code.InternalError)
			}
			if userBase == nil {
				return errors.NewError(err, code.InternalError)
			}

		} else {
			userBase, err = userBaseDao.Instance().GetUserBase(ctx, user.ID)
			if userBase == nil {
				return errors.NewError(err, code.InternalError)
			}
			if user.TgPwd == "" {
				user.TgPwd = xrand.Str(xrand.LetterSeed+xrand.DigitSeed, 32)
				userDao.Instance().Update(ctx, func(cols *userDao.Columns) any {
					return map[string]any{
						cols.ID: user.ID,
					}
				}, func(cols *userDao.Columns) any {
					return map[string]any{
						cols.TgPwd: user.TgPwd,
					}
				})
			} else {
				if !args.IsRobot && user.TgPwd != args.PassWord {
					return errors.NewError(code.IncorrectPassword)
				}
			}
			if userBase.UserType != xtypes.UserTypeTg && userBase.UserType != xtypes.UserTypePlayAlong && userBase.UserType != xtypes.UserTypeSystem {
				return errors.NewError(err, code.Unauthorized)
			}
			if args.Avatar != "" {
				userBaseDao.Instance().DoUpdateUserBase(ctx, user.ID, map[string]any{
					userBaseDao.Instance().Columns.Avatar: args.Avatar,
				})
			}

		}

		uid, err := s.doLogin(ctx, user, args.ClientIP, args.DeviceInfo, xtypes.RegisterLoginType_TG)
		if err != nil {
			return err
		}
		reply.Uid = uid


			if userBase.ChannelCode != "" {
				redisdefault.Instance().ZAdd(ctx, xtypes.TelegramUserIDKey, &redis.Z{Score: 0, Member: fmt.Sprintf("%s:%s", args.UserID, userBase.ChannelCode)})
				redisdefault.Instance().HSet(ctx, xtypes.TelegramUserIDToChannel, args.UserID, userBase.ChannelCode)
			}
	*/
	//activityplatformevt.PublishTaskTgProgressMember(xconv.Int64(args.UserID), user.ID, userBase.ChannelCode)
	return nil
}

func (s *Server) FetchInviteUrl(ctx context.Context, args *pb.FetchInviteUrlArgs, reply *pb.FetchInviteUrlReply) (err error) {

	user, err := userBaseDao.Instance().GetUserBase(ctx, args.UID)
	if err != nil {
		return errors.NewError(code.InvalidArgument)
	}
	if user == nil {
		return errors.NewError(code.InternalError)
	}
	channelCfg, err := optionChannelDao.Instance().GetChannel(ctx, user.ChannelCode)
	if err != nil {
		return errors.NewError(code.InviteUrlNotExist, err)
	}
	if channelCfg == nil {
		return errors.NewError(code.InviteUrlNotExist)
	}
	reply.InviteUrl = fmt.Sprintf("%s?start=%s", channelCfg.TelegramCfg.MainRobotLink, user.Code)
	/*
		if user.LastLoginType == xtypes.RegisterLoginType_TG {
			reply.InviteUrl = fmt.Sprintf("%s?start=%s", channelCfg.MainRobotLink, user.Code)
		} else {
			reply.InviteUrl = xresource.InviteCodeUrl(user.Code, user.ChannelCode)
		}
	*/
	return nil
}
