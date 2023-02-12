package main

import (
	"github.com/gin-gonic/gin"
	. "quicstart1005/route"
)

func main() {
	//获取gin 路由对象
	r := gin.Default()

	//路由系统
	InitRoute(r)

	//启动,默认端口8080
	r.Run()
}
