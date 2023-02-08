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

//定义全局数据库对象 DB

var DB *gorm.DB

// db包初始化函数,可以用来初始化gorm
func dbInit() {
	//配置dsn
	dsn := "root:12345678@tcp(127.0.0.1:3306)/css?charset=utf8mb4&parseTime=True&loc=Local"

	//创建日志对象
	newLogger := logger.New(
		log.New(os.Stdin, "\r\n", log.LstdFlags), //io write
		logger.Config{
			//SlowThreshold:             time.Second, //慢SQL 阈值
			//Colorful:                  false,
			//IgnoreRecordNotFoundError: false,
			//ParameterizedQueries:      false,
			LogLevel: logger.Info,
		},
	)

	//定义err
	var err error
	//连接mysql获取db实例
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, //日志配置
	})

	if err != nil {
		panic("连接数据库失败,error=" + err.Error())
	} else {
		panic("连接数据库成功!!!")
	}

	//设置数据库连接池参数
	sqlDB, _ := DB.DB()
	//设置数据库连接池最大连接数
	sqlDB.SetMaxOpenConns(100)
	//连接池最大允许的空闲连接数,如果没有SQL任务,需要执行的连接数大于20,超过的连接会被连接池关闭
	sqlDB.SetMaxIdleConns(20)

}

//获取gorm db,其他包调用此方法即可拿到DB
//无需担心不同携程并发使用这个DB对象,会公用一个连接,因为DB在调用方法时候会从数据库连接池获取新的连接

func GetDB() *gorm.DB {
	return DB
}

func main() {

	//迁移模型类:将模型类转为sql表
	DB.AutoMigrate(&Teacher{})

}
