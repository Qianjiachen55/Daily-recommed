package cron

import (
	"github.com/Qianjiachen55/Daily-recommed/global"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"time"
)

func SendMail()  {

	sendMailCron := cron.New(
		cron.WithParser(
			cron.NewParser(
				cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))

	sendMailCron.AddFunc("1-59/10 * * * * *",sendmMail)

	sendMailCron.Start()

}

func sendmMail()  {
	mailSlice := viper.GetStringSlice("email.to")
	for _,mail := range mailSlice{
		global.DrLogger.Info("sendTo:"+cast.ToString(mail))
		global.DrMail.SendMail(cast.ToString(mail),time.Now().Format("2006-01-02 15:04:05"),"hello ccccc!!!!")
	}

}