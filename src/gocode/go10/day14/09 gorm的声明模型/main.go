package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type t1 struct {
	Name string `gorm:"type:varchar(32);unique;not null"`
	Age  int
}

//创建模型类

func main() {
	//mysql连接信息
	//dsn := "user:pwd@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "denis:CIAy2k0308@tcp(mysql-internet-cn-north-1-72d891cb573c48c8.rds.jdcloud.com:3306)/css?charset=utf8mb4"

	// 创建日志对象
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			//SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel: logger.Info, // Log level
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, //日志配置
	})
	fmt.Println(db, err)
}
