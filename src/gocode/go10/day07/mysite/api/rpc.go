package api

import (
	"fmt"
	"mysite/db"
)

func Rpc() {
	db.HandleRedis1()
	fmt.Println("Redis接口数据....")
}
