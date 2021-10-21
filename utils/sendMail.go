package utils

import "gopkg.in/gomail.v2"

type Mail struct {
	username string
	pass     string
	host     string
	port     int
	name     string
}


func (mail *Mail)SendMail(user []string, body string, subject string) error {
	m := gomail.NewMessage()

	//m.SetAddressHeader("From", config.username,config.name)
	m.SetAddressHeader("From", mail.username, mail.name)
	//m.SetHeader("From",config.username)
	m.SetHeader("Subject", subject)
	m.SetHeader("To", user...)
	m.SetHeader("recommend")

	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(mail.host, mail.port, mail.username, mail.pass)

	err := dialer.DialAndSend(m)
	return err
}
