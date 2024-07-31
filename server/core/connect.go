package core

import (
	"Blog/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Path == "" {
		global.Log.Warn("未配置mysql，取消gorm连接")
		return nil
	}
	dns := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		// 开发环境显示所有sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		// 只打印错误的sql
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: mysqlLogger,
	})

	if err != nil {
		global.Log.Fatalf("[%s] mysql连接失败", dns)
	}

	// 连接池
	sdbDB, _ := db.DB()                 // 获取通用数据库对象 sdb.DB ，然后使用其提供的功能
	sdbDB.SetMaxIdleConns(10)           // SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sdbDB.SetMaxOpenConns(100)          // SetMaxOpenConns 设置打开数据库连接的最大数量。
	sdbDB.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime 设置了连接可复用的最大时间。

	return db
}
