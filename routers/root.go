package routers

import (
	"github.com/Qianjiachen55/Daily-recommed/controllers"
	"github.com/gin-gonic/gin"
)

func LoadRootRouter(engine *gin.Engine)  {
	engine.GET("/",controllers.Root)
}
