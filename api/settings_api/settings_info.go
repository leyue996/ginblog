package settings_api

import (
	"ginblog/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingInfoView(ctx *gin.Context) {
	res.Ok(map[string]any{"MSG": "111"}, "chenggongxinxi", ctx)
	// res.FailWithCode(res.SettingsError, ctx)
}
