package main

import (
	"Blog/core"
	"Blog/global"
	"fmt"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("aaa")
	global.Log.Error("aaa")
	global.Log.Info("aaa")
	// 连接数据库
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
