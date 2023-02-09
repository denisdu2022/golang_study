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

	//定义dsn数据库连接信息
	//用户名:密码@协议(ip+端口)/库名?设置字符集&解析时间&local
	dsn := "root:12345678@tcp(127.0.0.1:3306)/css?charset=utf8mb4&parseTime=True&loc=Local"

	//创建日志对象
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //io write
		logger.Config{
			//SlowThreshold: time.Second, //慢 SQL阈值
			LogLevel: logger.Info, //Log level(日志级别)
		},
	)

	//db类似文件句柄
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, //日志配置
	})
	if err != nil {
		panic("Connect database to failed!!!")
	}
	fmt.Println("数据库连接成功: ", db)

}
