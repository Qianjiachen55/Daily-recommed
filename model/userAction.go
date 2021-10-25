package model

type UserAction struct {
	Email string `json:"email" gorm:"primaryKey"`
	NewsUK string `json:"news_uk"`
	Click bool `json:"click"`
}


