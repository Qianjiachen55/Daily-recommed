package model

import "github.com/Qianjiachen55/Daily-recommed/global"

type News struct {
	UniqueKey     string `json:"unique_key" gorm:"primaryKey"`
	Title         string `json:"title"`
	Date          string `json:"date"`
	AuthorName    string `json:"authorName"`
	Url           string `json:"url"`
	ThumbnailPicS string `json:"thumbnail_pic_s"`
}

func (News) TableName() string{
	return "News"
}

func QueryNewsByUniqueKey(uniqueKey string) ([]News, error) {
	var newsS []News

	result := global.DrMysql.Where("unique_key = ?", uniqueKey).Find(&newsS)
	return newsS,result.Error
}

func InsertNews(news News) (error) {

	result := global.DrMysql.Create(&news)

	return result.Error
}
