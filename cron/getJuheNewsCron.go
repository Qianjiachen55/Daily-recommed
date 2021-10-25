package cron

import (
	"encoding/json"
	"github.com/Qianjiachen55/Daily-recommed/global"
	"github.com/Qianjiachen55/Daily-recommed/model"
	"github.com/Qianjiachen55/Daily-recommed/utils"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func GetNews()  {
	categories := viper.GetStringSlice("juhe.category")
	for _,category := range categories{
		body := utils.GetNewsFromJuhe("", cast.ToString(category),0,0,0)
		newsResp := utils.NewsResp{}
		json.Unmarshal([]byte(body),&newsResp)

		newsS := newsResp.Result.Data
		//var news utils.Data
		for _,news := range newsS {
			go func() {
				err := model.InsertNews(model.News{
					UniqueKey:     news.Uniquekey,
					Category:      news.Category,
					Title:         news.Title,
					Date:          news.Date,
					AuthorName:    news.AuthorName,
					Url:           news.Url,
					ThumbnailPicS: news.ThumbnailPicS,
				})
				if err!=nil {
					global.DrLogger.Error(cast.ToString(err))
				}
			}()

		}
	}

}
