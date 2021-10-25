package utils

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
)

type NewsResp struct {
	Resason    string `json:"reason"`
	Result     Result `json:"result"`
	Error_code int    `json:"error_code"`
}
type Result struct {
	Stat     string `json:"stat"`
	Data     []Data `json:"data"`
	Page     string `json:"page"`
	PageSize string `json:"pageSize"`
}
type Data struct {
	Uniquekey       string `json:"uniquekey"`
	Title           string `json:"title"`
	Date            string `json:"date"`
	Category        string `json:"category"`
	AuthorName      string `json:"author_name"`
	Url             string `json:"url"`
	ThumbnailPicS   string `json:"thumbnail_pic_s"`
	ThumbnailPicS02 string `json:"thumbnail_pics02"`
	ThumbnailPicS03 string `json:"thumbnail_pic_s03"`
	IsContent       string `json:"is_content"`
}

func GetNewsFromJuhe(key string, newsType string, page int, pageSize int, isFilter int) string {
	if key == "" {
		key = viper.GetString("juhe.key")
	}
	if newsType == "" {
		newsType = viper.GetString("juhe.type")
	}
	if page == 0 {
		page = viper.GetInt("juhe.page")
	}
	if pageSize == 0 {
		pageSize = viper.GetInt("juhe.pageSize")
	}
	if isFilter == 0 {
		isFilter = viper.GetInt("juhe.isFilter")
	}
	//建立http请求
	params := url.Values{}
	Url, err := url.Parse("http://v.juhe.cn/toutiao/index")
	if err != nil {
		return ""
	}

	params.Set("key", key)
	params.Set("type", newsType)
	params.Set("page", cast.ToString(page))
	params.Set("page_size", cast.ToString(pageSize))
	params.Set("is_filter", cast.ToString(isFilter))

	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
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
