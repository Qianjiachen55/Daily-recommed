package initalize

import (
	"fmt"
	"github.com/Qianjiachen55/Daily-recommed/global"
	"github.com/Qianjiachen55/Daily-recommed/model"
	"github.com/Qianjiachen55/Daily-recommed/utils"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mysqlConInstance *gorm.DB
)

func MysqlInit() *gorm.DB {
	mysqlConInstance = mysqlInit()

	return mysqlConInstance
}

func mysqlInit() *gorm.DB {

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
	msqlDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.PaincFun(global.DrLogger)
	}

	//初始化数据库
	if err := msqlDb.AutoMigrate(&model.News{}); err != nil {
		global.DrLogger.Error(cast.ToString(err))
	}
	if err := msqlDb.AutoMigrate(&model.User{}); err != nil {
		global.DrLogger.Error(cast.ToString(err))
	}
	if err := msqlDb.AutoMigrate(&model.UserAction{}); err != nil {
		global.DrLogger.Error(cast.ToString(err))
	}
	global.DrLogger.Info("init mysql")

	return msqlDb
}
