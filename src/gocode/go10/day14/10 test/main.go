package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func main() {
	//连接mysql
	dsn := "root:indionce418@tcp(127.0.0.1:3306)/css?charset=utf8mb4&parseTime=true&loc=local"

	// 创建日志对象
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			//SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel: logger.Info, // Log level
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // 日志配置
	})

	fmt.Println(db,err)
}
