package main

import (
	"github.com/gin-gonic/gin"
	//引入route包
	. "routeGroup1/route"
)

func main() {

	//获取路由对象
	r := gin.Default()

	//加载模板文件
	r.LoadHTMLGlob("templates/*")

	//路由分组
	//调用route包
	InitRoute(r)

	//启动
	r.Run(":8090")
}
