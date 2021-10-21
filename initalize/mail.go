package initalize

import (
	"github.com/Qianjiachen55/Daily-recommed/utils"
	"github.com/spf13/viper"
)

var mailInstance *utils.Mail

func MailInit() *utils.Mail {
	once.Do(func() {
		mailInstance = mailInit()
	})
	return mailInstance
}

func mailInit() *utils.Mail {
	defer utils.PanicFun()
	username :=viper.GetString("email.username")
	pass := viper.GetString("email.pass")
	host := viper.GetString("email.host")
	port := viper.GetString("email.port")
	name := viper.GetString("email.name")

	return utils.InitMail(username,pass,host,port,name)
}
