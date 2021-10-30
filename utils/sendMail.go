package utils

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type Mail struct {
	username string
	pass     string
	host     string
	port     int
	name     string
}

func InitMail(username string,pass string,host string,port string,name string) *Mail {
	if username==""{
		username = viper.GetString("email.username")
	}
	if pass==""{
		pass=viper.GetString("email.pass")
	}
	if host=="" {
		host = viper.GetString("email.host")
	}
	if port=="" {
		port = viper.GetString("email.port")
	}
	if name==""{
		name = viper.GetString("email.name")
	}
	return &Mail{
		username: username,
		pass:     pass,
		host:     host,
		port:     cast.ToInt(port),
		name:     name,
	}
}


func (mail *Mail)SendMail(user string, body string, subject string) error {
	m := gomail.NewMessage()

	//m.SetAddressHeader("From", config.username,config.name)
	m.SetAddressHeader("From", mail.username, mail.name)
	//m.SetHeader("From",config.username)
	//主题
	m.SetHeader("Subject", subject)
	m.SetHeader("To", user)
	m.SetHeader("recommend","recommend")

	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(mail.host, mail.port, mail.username, mail.pass)

	err := dialer.DialAndSend(m)
	return err
}
