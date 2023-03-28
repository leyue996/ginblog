package main

import (
	"ginblog/core"
	"ginblog/global"
	"ginblog/routers"

	_ "github.com/sirupsen/logrus"
)

func main() {
	core.Initconfig() //读取配置文件

	global.Log = core.InitLogger() //初始化log
	// global.Log.Warnln("噶")
	// global.Log.Error("噶")
	// logrus.Infof("噶")

	global.MysqlDB = core.InitGorm() //连接数据库

	router := routers.InitRouter() //初始化路由

	addr := global.Config.System.Addr() //获取地址端口

	global.Log.Infof("server运行在%s", addr) //打印log
	router.Run(addr)

}
