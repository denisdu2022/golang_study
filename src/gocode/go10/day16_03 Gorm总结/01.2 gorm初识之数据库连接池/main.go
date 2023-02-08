package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

//定义全局数据库对象 DB

var DB *gorm.DB

// db包初始化函数,可以用来初始化gorm
func dbInit() {
	//配置dsn
	dsn := "root:12345678@tcp(127.0.0.1:3306)/css?charset=utf8mb4&parseTime=True&loc=Local"

	//创建日志对象
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //io write
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

	//获取DB对象
	db := GetDB()

	//下边执行数据库查询操作
	fmt.Println(db)

}
