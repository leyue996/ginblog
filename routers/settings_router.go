package routers

import (
	"ginblog/api"
	"github.com/gin-gonic/gin"
)

func SettingsRouter(router *gin.RouterGroup) {
	settingApi := api.ApiGroupApp.SettingsApi
	router.GET("settings", settingApi.SettingInfoView)
}
