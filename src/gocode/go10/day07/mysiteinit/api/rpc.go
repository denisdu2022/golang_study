package api

import (
	"fmt"
	"mysite/db"
)

func RpcAPI() {
	db.HandleRedis()
	fmt.Println("RpcAPI:reids数据库接口")
}
