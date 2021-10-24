package routers

import (
	"github.com/Qianjiachen55/Daily-recommed/controllers"
	"github.com/gin-gonic/gin"
)

func LoadNewsRouters(engine *gin.Engine)  {
	engine.GET("/api/news/:uk",controllers.News)
}