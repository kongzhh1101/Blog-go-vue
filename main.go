package main

import (
	"Blog/core"
	docs "Blog/docs"
	"Blog/flag"
	"Blog/global"
	"Blog/routers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           博客后端
// @version         1.0
// @BasePath /api

func main() {
	// 读取配置文件
	core.InitConfig()
	docs.SwaggerInfo.BasePath = "/api"

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(addr)
	if err != nil {
		global.Logger.Fatalf(err.Error())
	}
}
