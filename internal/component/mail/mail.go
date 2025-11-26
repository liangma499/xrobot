package mail

import (
	"crypto/tls"
	"sync"
	"xbase/etc"
	"xbase/log"

	"gopkg.in/gomail.v2"
)

var (
	once           sync.Once
	instance       *Mail
	configMail     *Config
	onceConfigMail sync.Once
)

type Mail = gomail.Dialer

type Config struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	SendUsername string `json:"send-username"`
}

func Instance() *Mail {
	once.Do(func() {
		instance = NewInstance("etc.mail")
	})

	return instance
}
func InstanceConf() *Config {
	onceConfigMail.Do(func() {
		configMail = &Config{}
		if err := etc.Get("etc.mail").Scan(configMail); err != nil {
			log.Fatalf("load mail config failed: %v", err)
		}
	})

	return configMail
}

// NewInstance 新建实例
func NewInstance[T string | Config | *Config](config T) *Mail {
	var (
		conf *Config
		v    any = config
	)

	switch c := v.(type) {
	case string:
		conf = &Config{}
		if err := etc.Get(c).Scan(conf); err != nil {
			log.Fatalf("load mail config failed: %v", err)
		}
	case Config:
		conf = &c
	case *Config:
		conf = c
	}

	dialer := gomail.NewDialer(conf.Host, conf.Port, conf.Username, conf.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return dialer
}
