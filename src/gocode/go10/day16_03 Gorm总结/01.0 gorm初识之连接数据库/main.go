package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//定义dsn数据库连接信息
	//用户名:密码@协议(ip+端口)/库名?设置字符集&解析时间&local
	dsn := "root:12345678@tcp(127.0.0.1:3306)/css?charset=utf8mb4&parseTime=True&loc=Local"
	//db类似文件句柄
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)
	fmt.Println("数据库连接成功!!!")

}
