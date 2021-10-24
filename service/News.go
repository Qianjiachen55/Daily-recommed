package service

import (
	"github.com/Qianjiachen55/Daily-recommed/global"
	"github.com/Qianjiachen55/Daily-recommed/model"
	"github.com/spf13/cast"
)

func News(uk string) (model.News,error) {
	n := model.News{
		UniqueKey:     uk,
		Title:         "",
		Date:          "",
		AuthorName:    "",
		Url:           "",
		ThumbnailPicS: "",
	}
	res := global.DrMysql.First(&n)
	if res.Error != nil{
		global.DrLogger.Error(cast.ToString(res.Error))
	}

	return n,res.Error

}
