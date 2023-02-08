package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

//创建选课系统的模型类

//基本模型

type BaseModel struct {
	ID         int        `gorm:"primaryKey"`
	Name       string     `gorm:"type:varchar(32);unique;not null"`
	CreateTime *time.Time `gorm:"autoCreateTime"`
	UpdateTime *time.Time `gorm:"autoCreateTime"`
	DeleteTime *time.Time `gorm:"autoCreateTime"`
}

//教师表

type Teacher struct {
	BaseModel
	Tno    int
	Pwd    string `gorm:"type:varchar(100);not null"`
	Tel    string `gorm:"type:char(11);"`
	Birth  *time.Time
	Remark string `gorm:"type:varchar(255);"`
}

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

	//迁移模型类:将模型类转为sql表
	db.AutoMigrate(&Teacher{})

}
