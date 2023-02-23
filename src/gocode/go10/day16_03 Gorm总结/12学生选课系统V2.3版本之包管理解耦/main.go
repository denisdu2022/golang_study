package main

import (
	. "cssv2.3/database"
	. "cssv2.3/render"
	. "cssv2.3/route"
	"github.com/gin-gonic/gin"
)

func main() {

	//获取路由对象
	r := gin.Default()

	//使用HTMLRender调用createMyRender函数   加载模板文件
	r.HTMLRender = CreateMyRender()

	//开放静态文件窗口
	r.Static("/static", "./static")

	// 数据库初始化
	MySqlInit()

	//路由系统
	InitRoute(r)

	//服务器启动 默认端口8080
	r.Run()

}
