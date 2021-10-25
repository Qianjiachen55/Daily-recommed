package cron

import (
	"github.com/Qianjiachen55/Daily-recommed/global"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func SendMail()  {

	sendMailCron := cron.New(
		cron.WithParser(
			cron.NewParser(
				cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))

	sendMailCron.AddFunc("* * 8 * * * *",sendmMail)
}

func sendmMail()  {
	mailSlice := viper.GetStringSlice("email.to")
	for mail := range(mailSlice){
		global.DrMail.SendMail(cast.ToStringSlice(mail),"asdf","sdfasd")
	}

}