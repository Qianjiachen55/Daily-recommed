package model

type News struct {
	UniqueKey string `json:"uniquekey" gorm:"primaryKey"`
	Title string `json:"title"`
	Date string `json:"date"`
	AuthorName string `json:"authorName"`
	Url string `json:"url"`
	ThumbnailPicS string `json:"thumbnail_pic_s"`
}