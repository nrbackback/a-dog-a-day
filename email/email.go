package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type EmailConfig struct {
	From     string `yaml:"from"`
	Password string `yaml:"password"`
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	To       string `yaml:"to"`
}

type EmailSender struct {
	Conf EmailConfig
}

var WeiboSender EmailSender

func Send(title, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", WeiboSender.Conf.From)
	m.SetHeader("To", WeiboSender.Conf.To)

	m.SetHeader("Subject", title)
	content = fmt.Sprintf(`<img src="%s"  alt="%s" />`, content, content)

	m.SetBody("text/html", content)

	d := gomail.NewPlainDialer(WeiboSender.Conf.Server, WeiboSender.Conf.Port, WeiboSender.Conf.From, WeiboSender.Conf.Password)
	err := d.DialAndSend(m)
	return err
}
