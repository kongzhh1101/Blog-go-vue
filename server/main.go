package main

import (
	"Blog/core"
	"Blog/global"
	"fmt"
)

func main() {
	core.InitConf()
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
