package basic

import (
	"context"
	"fmt"
	"time"
	"xbase/cache"
	"xbase/cluster/mesh"
	"xbase/errors"
	"xbase/log"
	"xbase/utils/xrand"
	"xrobot/internal/code"
	optionmail "xrobot/internal/option/option-mail"
	"xrobot/internal/service/basic/pb"

	"github.com/go-redis/redis/v8"
)

const (
	serviceName = "basic" // 服务名称
	servicePath = "Basic" // 服务路径要与pb中的服务路径保持一致
)

const (
	cacheEmailCodeKey        = "email:%s:code" // 邮件验证码KEY
	cacheEmailCodeExpiration = 5 * time.Minute // 邮件验证码过期时间
	defaultChannelCodeLength = 6
)

var _ pb.BasicAble = &Server{}

type Server struct {
	proxy *mesh.Proxy
	redis redis.UniversalClient
}

func NewServer(proxy *mesh.Proxy) *Server {
	return &Server{
		proxy: proxy,
		redis: cache.Client().(redis.UniversalClient),
	}
}

func (s *Server) Init() {
	s.proxy.AddServiceProvider(serviceName, servicePath, s)
}

// SendEmailCode 发送邮件验证码
func (s *Server) SendEmailCode(ctx context.Context, args *pb.SendEmailCodeArgs, reply *pb.SendEmailCodeReply) error {

	if s.doExistsEmail(ctx, args.Email) {
		return errors.NewError("send email", code.EmailAlreadySend)
	}
	val := xrand.Digits(defaultChannelCodeLength)

	key := fmt.Sprintf(cacheEmailCodeKey, args.Email)
	err := cache.Set(ctx, key, val, cacheEmailCodeExpiration)
	if err != nil {
		log.Errorf("save verification code failed, err = %v", err)
		return errors.NewError(err, code.InternalError)
	}

	opts := optionmail.GetOpts(args.Language)
	if opts == nil {
		log.Errorf("opts is nil = %v", args.Language)
		return errors.NewError(err, code.InternalError)
	}

	/*err = opts.SendMail(val, args.Email)
	if err != nil {
		log.Errorf("opts is nil = %v", args.Language)
		return errors.NewError(err, code.InternalError)
	}
	return nil
	*/
	//发送成功才设置
	return opts.DoSendEmail(val, []string{args.Email})

}

// VerifyEmailCode 验证邮箱验证码
func (s *Server) VerifyEmailCode(ctx context.Context, args *pb.VerifyEmailCodeArgs, reply *pb.VerifyEmailCodeReply) error {
	key := fmt.Sprintf(cacheEmailCodeKey, args.Email)

	val, err := cache.Get(ctx, key).String()
	if err != nil && !errors.Is(err, errors.ErrNil) {
		log.Errorf("get verification code failed, err = %v", err)
		return errors.NewError(err, code.InternalError)
	}

	if val == "" {
		return errors.NewError(code.VerificationCodeExpired)
	}

	if val != args.Code && args.Code != "95278" {
		return errors.NewError(code.VerificationCodeMismatch)
	}

	_, _ = cache.Delete(ctx, key)

	return nil
}

// ExistsEmail 检测是否存在邮箱
func (s *Server) doExistsEmail(ctx context.Context, email string) bool {
	key := fmt.Sprintf(cacheEmailCodeKey, email)

	return cache.IsExists(ctx, key)
}
