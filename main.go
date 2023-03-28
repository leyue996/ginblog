package main

import (
	"ginblog/core"
	"ginblog/global"

	_ "github.com/sirupsen/logrus"
)

func main() {
	core.Initconfig() //读取配置文件

	global.Log = core.InitLogger() //初始化log
	// global.Log.Warnln("噶")
	// global.Log.Error("噶")
	// logrus.Infof("噶")

	global.MysqlDB = core.InitGorm() //连接数据库
}
