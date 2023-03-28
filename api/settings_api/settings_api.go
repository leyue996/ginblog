package settings_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (SettingsApi) SettingInfoView(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "hahaha"})
}
