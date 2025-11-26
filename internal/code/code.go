package code

import (
	"xbase/codes"
)

var (
	OK                         = codes.NewCode(1, "ok")                              // 成功
	Canceled                   = codes.NewCode(2, "canceled")                        // 已取消
	Unknown                    = codes.NewCode(3, "unknown")                         // 未知错误
	InvalidArgument            = codes.NewCode(4, "invalid argument")                // 无效参数
	DeadlineExceeded           = codes.NewCode(5, "deadline exceeded")               // 执行超时
	NotFound                   = codes.NewCode(6, "not found")                       // 未找到相关资源
	InternalError              = codes.NewCode(7, "internal error")                  // 服务器内部错误
	Unauthorized               = codes.NewCode(8, "unauthorized")                    // 未授权
	IllegalInvoke              = codes.NewCode(9, "illegal invoke")                  // 非法调用
	IllegalRequest             = codes.NewCode(10, "illegal request")                // 非法请求
	OperationTooFast           = codes.NewCode(11, "operation too fast")             // 操作太快
	ServerIsMaintain           = codes.NewCode(12, "server is maintain")             // 维服中
	ServerIsClosed             = codes.NewCode(13, "server is closed ")              // 服务已关闭
	ServerIsBusy               = codes.NewCode(14, "server is busy")                 // 数据已存在
	LoadOptionErr              = codes.NewCode(15, "load option err")                // 加载配置错误
	OptionNotFound             = codes.NewCode(16, "option not found")               // 配置出错
	ModifyOptionErr            = codes.NewCode(17, "modify option err")              // 修改配置出错
	InValidateAddress          = codes.NewCode(18, "invalidate address")             // 无效的地下
	CurrencyNotUse             = codes.NewCode(19, "currency not use")               // 币种禁用
	BalanceInsufficient        = codes.NewCode(20, "balance insufficient")           // 余额不足
	AuthorizationExpired       = codes.NewCode(100, "authorization expired")         // 授权已过期
	AuthorizationElsewhere     = codes.NewCode(101, "authorization elsewhere")       // 异地登录
	NoPermission               = codes.NewCode(102, "no permission")                 // 暂无权限
	IllegalOperation           = codes.NewCode(103, "illegal operation")             // 非法操作
	IncorrectAccountOrPassword = codes.NewCode(104, "incorrect account or password") // 账号或密码错误
	IncorrectPassword          = codes.NewCode(105, "incorrect password")            // 密码错误
	EmailExists                = codes.NewCode(106, "email exists")                  // 邮箱已存在
	VerificationCodeExpired    = codes.NewCode(107, "verification code expired")     // 验证码已过期
	VerificationCodeMismatch   = codes.NewCode(108, "verification code mismatch")    // 验证码不匹配
	UserForbidden              = codes.NewCode(109, "user is forbidden")             // 用户已禁用
	EmailIsNotVerified         = codes.NewCode(110, "email is not verified")         // 邮箱未验证
	EmailAlreadySend           = codes.NewCode(111, "email already send")            // 邮件已经发送了
	EmailAlreadyBind           = codes.NewCode(112, "email already bind")            // 邮箱已经绑定了
	EmailNotExists             = codes.NewCode(113, "email not exist")               // 邮箱不存在
	AccountNotExist            = codes.NewCode(114, "account not exist")             // 账号不存在
	InviteCodeNotExist         = codes.NewCode(115, "invite code not exist")         // 邀请码不存在
	InviteUrlNotExist          = codes.NewCode(116, "invite url not exist")          // 邀请链接存在

	ChannelNotExist = codes.NewCode(200, "channel not exist") // 渠道不存在

	ButtonInvalid   = codes.NewCode(300, "button invalid")     // 按钮参数无效
	Trc20RpcErr     = codes.NewCode(8000, "trc20 rpc err")     // rpc设置错误
	NotUniqueAmount = codes.NewCode(8001, "not unique amount") // 没有唯一金额
)
