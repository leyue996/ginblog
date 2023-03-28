package global

import (
	"ginblog/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config  *config.Config
	MysqlDB *gorm.DB
	Log     *logrus.Logger
)
