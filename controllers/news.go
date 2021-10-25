package controllers

import (
	"fmt"
	"github.com/Qianjiachen55/Daily-recommed/global"
	"github.com/Qianjiachen55/Daily-recommed/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func News(ctx *gin.Context) {
	uk := ctx.Param("uk")
	if uk == ""{
		global.DrLogger.Error("uk is nil")
		ctx.IndentedJSON(
			http.StatusBadRequest,
			Response{
				Code:    400,
				Message: fmt.Sprintf("uk is nil"),
			},
		)
	}
	global.DrLogger.Info(fmt.Sprintf("uk :%s",uk))

	news, err := service.News(uk)
	if err != nil {
		global.DrLogger.Error(fmt.Sprintf("uk :%s no exits!!", uk))
		ctx.IndentedJSON(
			http.StatusOK,
			Response{
				Code:    400,
				Message: fmt.Sprintf("uk :%s no exits!!", uk),
			},
		)
	} else {
		global.DrLogger.Info("redirect:" + news.Url)
		ctx.Redirect(http.StatusMovedPermanently, news.Url)
	}

}
