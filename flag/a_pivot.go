package flag

import (
	sys_flag "flag"
)

type Option struct {
	DB bool
}

// 解析命令行参数
func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	// 解析命令行参数写入注册的flag里
	sys_flag.Parse()
	return Option{
		DB: *db,
	}
}

// 是否停止web项目
func IsWebStop(option Option) bool {
	return option.DB
}

// 根据命令，执行不同函数
func SwitchOption(option Option) {
	if option.DB {
		MakeMigration()
	}
}
