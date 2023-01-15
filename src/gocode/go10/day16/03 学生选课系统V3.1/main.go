package main

import (
	. "css/databse"
	. "css/render"
	. "css/route"
	"github.com/gin-gonic/gin"
)

func main() {
	//获取路由对象
	r := gin.Default()

	//初始化数据库
	MySqlInit()

	//加载模版文件
	r.HTMLRender = CreateMyRender()

	//设置静态文件资源窗口
	r.Static("/static", "./static")

	//路由系统
	InitRoute(r)

	//服务启动
	r.Run(":8090")

}
