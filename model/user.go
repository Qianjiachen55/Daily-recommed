package model

import "github.com/Qianjiachen55/Daily-recommed/global"

type User struct {
	Email string `json:"email" gorm:"primaryKey"`
}

func QueryUserByEmail(email string) ([]User,error) {
	var users []User
	result := global.DrMysql.Where("email = ?",email).Find(&users)
	return users,result.Error
}

func InsertUser(user User) error {
	result := global.DrMysql.Create(&user)
	return result.Error
}