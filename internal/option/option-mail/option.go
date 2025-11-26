package optionmailcfg

import (
	"context"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"xbase/codes"
	"xbase/config"
	"xbase/errors"
	"xbase/log"
	"xbase/utils/xpath"

	"xbase/config/etcd"

	mailcomp "xrobot/internal/component/mail"
	"xrobot/internal/xtypes"

	"github.com/chanyipiaomiao/hltool"
	"gopkg.in/gomail.v2"
)

const (
	Name = xtypes.OptionPrefix + "option-mail"
	file = xtypes.OptionPrefix + "option-mail.toml"
)

type columns struct {
	VerificationCode string
}

type Options struct {
	VerificationCode map[string]*VerificationCodeOptions `json:"verificationCode"` // 验证码配置
}

type VerificationCodeOptions struct {
	From        string `json:"from"`        // 邮件来源
	Subject     string `json:"subject"`     // 邮件主题
	ContentType string `json:"contentType"` // 内容格式
	Body        string `json:"body"`        // 邮件内容
}

var (
	opts    atomic.Value
	once    sync.Once
	Columns = &columns{
		VerificationCode: "verificationCode", // 验证码邮件
	}
)

// GetOpts 读取配置项
func GetOpts(language string) *VerificationCodeOptions {
	once.Do(func() {
		o, err := doLoadOpts()
		if err != nil {
			log.Fatalf("get server's options failed: %v", err)
		}

		config.Watch(func(names ...string) {
			if o, err := doLoadOpts(); err == nil {
				opts.Store(o)
			}
		}, Name)

		opts.Store(o)
	})
	options := opts.Load().(*Options)
	if lanOpt, ok := options.VerificationCode[language]; ok {
		return lanOpt
	}
	lanOpt, ok := options.VerificationCode["en"]
	if !ok {
		log.Warnf("en language is nil")
		return nil
	}
	return lanOpt

}

// SetOpts 设置配置项
func SetOpts(ctx context.Context, optsData any) error {
	return config.Store(ctx, etcd.Name, file, optsData, true)
}

// HasOpts 是否有配置项
func HasOpts() bool {
	return config.Has(Name)
}

// 加载配置项
func doLoadOpts() (*Options, error) {
	o := &Options{
		VerificationCode: make(map[string]*VerificationCodeOptions, 0),
	}

	err := config.Get(Name).Scan(o)
	if err != nil {
		return nil, err
	}
	//log.Warnf("doLoadOpts-option-mail:%#v", o)
	return o, nil
}
func (vc *VerificationCodeOptions) SendMailByHltool(code string, to []string, attaches ...string) error {

	if len(to) == 0 {
		return nil
	}
	accConf := mailcomp.InstanceConf()
	username := accConf.Username //"cryptogirlnancy@gmail.com"
	host := accConf.Host         //"smtp.gmail.com"
	password := accConf.Password //"Everythingisgood..."
	port := accConf.Port         //465

	// 替换验证码
	replaces := make(map[string]string)
	replaces["code"] = code
	body := os.Expand(vc.Body, func(s string) string {
		return replaces[s]
	})
	contentType := vc.ContentType //"text/html"
	attach := ""
	cc := []string{}

	message := hltool.NewEmailMessage("U2Bet Team<U2Bet_Team@gmail.com>", vc.Subject, contentType, body, attach, to, cc)

	sendEmail := hltool.NewEmailClient(host, username, password, port, message)
	isSuccess, err := hltool.SendMessage(sendEmail)
	if !isSuccess {
		log.Warnf("email :%v,accConf:%#v", to, accConf)
		return err
	}
	return nil
}

// 发送邮件
func (vc *VerificationCodeOptions) DoSendEmail(code string, to []string, attaches ...string) error {
	if len(to) == 0 {
		return nil
	}
	accConf := mailcomp.InstanceConf()
	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s<%s>", accConf.SendUsername, accConf.Username))
	//m.SetHeader("From", accConf.SendUsername)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", vc.Subject)

	// 替换验证码
	replaces := make(map[string]string)
	replaces["code"] = code
	body := os.Expand(vc.Body, func(s string) string {
		return replaces[s]
	})
	m.SetBody(vc.ContentType, body)

	for _, attach := range attaches {
		if xpath.IsFile(attach) {
			m.Attach(attach)
		}
	}

	err := mailcomp.Instance().DialAndSend(m)
	if err != nil {
		log.Errorf("send email failed: err = %v", err)
		return errors.NewError(err, codes.InternalError)
	}

	return nil
}
