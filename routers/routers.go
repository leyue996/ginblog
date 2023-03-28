package routers

import (
	"ginblog/global"
	"github.com/gin-gonic/gin"
)

//type ApiRouter struct{}
//
//var apiRouterGroup = new(ApiRouter)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	apiRouterGroup := router.Group("api")
	{
		SettingsRouter(apiRouterGroup) //系统配置api
	}

	return router
}
