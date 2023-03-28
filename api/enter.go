package api

import "ginblog/api/settings_api"

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
}

var ApiGroupApp = new(ApiGroup) //必须实例化结构体才能使用方法
