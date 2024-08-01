package main

import (
	"Blog/core"
	"Blog/flag"
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

	// 绑定命令行参数
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	addr := global.Config.System.Addr()
	global.Log.Info(addr)
	router := routers.InitRouter()
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
