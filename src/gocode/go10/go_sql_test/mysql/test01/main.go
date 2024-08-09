package main

import (
	"database/sql"
	"fmt"
	"log"

	//数据库驱动
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type USER struct {
	UserName   string
	UserPasswd string
}

//mysql 信息

const (
	USERNAME = "xxxxx"
	PASSWORD = "xxxx"
	IP       = "xxxxx.com"
	PORT     = "3306"
	dbName   = "xxxx"
)

func main() {
	//初始化mysql连接
	InitDB()

	user := USER{
		UserName:   "李白",
		UserPasswd: "1234567",
	}
	//向表中插入数据
	InsertUser(user)
}

//初始化数据库连接方法

func InitDB() *sql.DB {
	//数据库连接字符串
	dsn := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", IP, ":", PORT, ")/", dbName, "?charset=utf8"}, "")
	//连接数据库
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		println("mysql open fail!")
		log.Fatal(err)
		return nil
	}
	//数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//数据库最大空闲连接数
	DB.SetMaxIdleConns(10)
	//判断数据库是否连接成功
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail!")
		//记录日志
		log.Fatal(err)
		return nil
	}
	fmt.Println("数据库连接成功! ")
	return DB
}

//往user表插入数据(注意user表已建立)

func InsertUser(u USER) bool {

	//开启事务
	tx, err := InitDB().Begin()
	if err != nil {
		fmt.Println("tx fail!")
		log.Fatal(err)
		return false
	}

	//准备SQL语句
	stmt, err := InitDB().Prepare("INSERT INTO user (UserName,UserPasswd) values (?,?)")
	if err != nil {
		fmt.Println("Prepare fail!")
		log.Fatal(err)
		return false
	}

	defer stmt.Close()

	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(u.UserName, u.UserPasswd)
	if err != nil {
		fmt.Println("exec fail!")
		log.Fatal(err)
		return false
	}

	//将事务提交
	tx.Commit()

	//获得上一个插入自增的ID
	fmt.Println(res.RowsAffected())
	return true

}
