package model

import (
	"github.com/Qianjiachen55/Daily-recommed/global"
)

type News struct {
	UniqueKey     string `json:"unique_key" gorm:"primaryKey"`
	Category      string `json:"category"`
	Title         string `json:"title"`
	Date          string `json:"date"`
	AuthorName    string `json:"authorName"`
	Url           string `json:"url"`
	ThumbnailPicS string `json:"thumbnail_pic_s"`

}

func Query(query string, by string, model interface{})  error {
	result := global.DrMysql.Where(query,by).Find(model)

	return result.Error
}

func QueryNewsByUniqueKey(uniqueKey string) ([]News, error) {
	var newsS []News

	result := global.DrMysql.Where("unique_key = ?", uniqueKey).Find(&newsS)
	return newsS, result.Error
}

func InsertNews(news News) error {

	result := global.DrMysql.Create(&news)

	return result.Error
}
