package main

import (
	"Blog/core"
	"Blog/global"
	"Blog/routers"
)

func main() {
	// 读取配置文件
	core.InitConf()

	// 初始化日志
	global.Log = core.InitLogger()

	// 连接数据库
	global.DB = core.InitGorm()

	addr := global.Config.System.Addr()
	global.Log.Info(addr)
	router := routers.InitRouter()
	router.Run(addr)
}
