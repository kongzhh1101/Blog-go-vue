package main

import (
	"Blog/core"
	"Blog/flag"
	"Blog/global"
	"Blog/routers"
)

func main() {
	// 读取配置文件
	core.InitConfig()

	// 初始化日志
	global.Logger = core.InitLogger()

	// 连接数据库
	global.DB = core.InitGorm()

	// 绑定命令行参数
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	addr := global.Config.System.Addr()
	global.Logger.Info(addr)
	router := routers.InitRouter()
	err := router.Run(addr)
	if err != nil {
		global.Logger.Fatalf(err.Error())
	}
}
