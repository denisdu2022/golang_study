package main

import (
	"github.com/gin-gonic/gin"
	//引入路由包
	. "routeGroup001/route"
)

func main() {
	//1. 获取路由对象
	r := gin.Default()

	//3.加载templates文件
	r.LoadHTMLGlob("templates/*")

	//4.调用路由包
	InitRoute(r)

	//2.启动
	r.Run()
}
