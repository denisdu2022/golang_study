package main

import (
	. "ginRouterZu/route"
	"github.com/gin-gonic/gin"
)

func main() {

	//获取路由对象
	r := gin.Default()

	//路由分组
	InitRoute(r)

	//启动
	r.Run()
}
