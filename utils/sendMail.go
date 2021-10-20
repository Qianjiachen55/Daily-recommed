package utils

import "gopkg.in/gomail.v2"

type MailConfig struct {
	username string
	pass     string
	host     string
	port     int
	name     string
}

//UIVDUGPUDVDIKSYV
func SendMail(config MailConfig, user []string, body string, subject string) error {
	m := gomail.NewMessage()

	//m.SetAddressHeader("From", config.username,config.name)
	m.SetAddressHeader("From", config.username, config.name)
	//m.SetHeader("From",config.username)
	m.SetHeader("Subject", subject)
	m.SetHeader("To", user...)
	m.SetHeader("recommend")

	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(config.host, config.port, config.username, config.pass)

	err := dialer.DialAndSend(m)
	return err
}
