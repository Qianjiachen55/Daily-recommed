package utils

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetNewsFromJuhe(key string, newsType string, page int, pageSize int, isFilter int) string{
	if key==""{
		key = viper.GetString("juhe.key")
	}
	if newsType=="" {
		newsType = viper.GetString("juhe.type")
	}
	if page==0 {
		page = viper.GetInt("juhe.page")
	}
	if pageSize==0{
		pageSize = viper.GetInt("juhe.pageSize")
	}
	if isFilter==0 {
		isFilter = viper.GetInt("juhe.isFilter")
	}
	//建立http请求
	params := url.Values{}
	Url, err := url.Parse("http://v.juhe.cn/toutiao/index")
	if err != nil {
		return ""
	}

	params.Set("key", key)
	params.Set("type",newsType)
	params.Set("page", cast.ToString(page))
	params.Set("page_size",cast.ToString( pageSize))
	params.Set("is_filter",cast.ToString(isFilter))

	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	resp,err := http.Get(urlPath)
	defer resp.Body.Close()

	body,_ := ioutil.ReadAll(resp.Body)
	return cast.ToString(body)
}

//func init()  {
//	viper.AddConfigPath("../conf")
//	viper.SetConfigType("toml")
//	viper.SetConfigName("config")
//	viper.AutomaticEnv()
//
//	if err := viper.ReadInConfig(); err != nil {
//		fmt.Println("error:", err)
//		panic(err)
//	}
//}