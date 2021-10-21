package initalize

import (
	"fmt"
	"github.com/Qianjiachen55/Daily-recommed/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mysqlConInstance *gorm.DB

)

func MysqlInit() *gorm.DB {
	once.Do(func() {
		mysqlConInstance = mysqlInit()
	})
	return mysqlConInstance
}

func mysqlInit() *gorm.DB {
	defer utils.PanicFun()

	hostname := viper.GetString("mysql.hostname")
	user := viper.GetString("mysql.user")
	port := viper.GetString("mysql.port")
	password := viper.GetString("mysql.password")
	db := viper.GetString("mysql.db")
//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		hostname,
		port,
		db,
		)

	msqlDb,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!= nil {
		utils.ErrFun(err)
		panic(err)
	}

	return msqlDb
}