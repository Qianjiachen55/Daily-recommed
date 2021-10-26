package cron

import (
	"encoding/json"
	"github.com/Qianjiachen55/Daily-recommed/global"
	"github.com/Qianjiachen55/Daily-recommed/model"
	"github.com/Qianjiachen55/Daily-recommed/utils"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func GetNews() {
	categories := viper.GetStringSlice("juhe.category")
	newsChan := make(chan utils.Data)
	//canQuit := make(chan bool)

	for _, category := range categories {
		global.DrLogger.Info(category)
		body := utils.GetNewsFromJuhe("", cast.ToString(category), 0, 0, 0)

		newsResp := utils.NewsResp{}
		json.Unmarshal([]byte(body), &newsResp)
		//fmt.Println("解析完成")
		newsS := newsResp.Result.Data
		//var news utils.Data

		for _,news := range newsS{

			go func() {
				n := <-newsChan
				//fmt.Println(n)
				err := model.InsertNews(model.News{
					UniqueKey:     n.Uniquekey,
					Category:      n.Category,
					Title:         n.Title,
					Date:          n.Date,
					AuthorName:    n.AuthorName,
					Url:           n.Url,
					ThumbnailPicS: n.ThumbnailPicS,
				})
				if err != nil {
					global.DrLogger.Error(cast.ToString(err))
				}
			}()

			newsChan <- news
		}


	}
	close(newsChan)

}
