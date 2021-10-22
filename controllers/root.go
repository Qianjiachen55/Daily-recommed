package controllers

import (
	"github.com/Qianjiachen55/Daily-recommed/service"
	"github.com/gin-gonic/gin"
	"net/http"
)



func Root(ctx *gin.Context)  {
	body,err := service.Root()
	if err==nil {
		ctx.IndentedJSON(http.StatusOK,body)
	}

}
