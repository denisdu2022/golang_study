package api

import (
	"fmt"
	"mysite/db"
)

func RestfulMySql() {
	db.HandleMySql1()
	fmt.Println("Mysql接口数据...")
}
