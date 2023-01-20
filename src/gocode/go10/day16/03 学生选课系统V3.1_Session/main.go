package main

import (
	. "css/databse"
	. "css/render"
	. "css/route"
	"github.com/gin-gonic/gin"
	// 导入session包
	"github.com/gin-contrib/sessions"
	// 导入session存储引擎
	"github.com/gin-contrib/sessions/cookie"
)

func main() {
	//获取路由对象
	r := gin.Default()

	//初始化Session
	//初始化cookie的存储引擎,"mySession"参数是用于加密的秘钥,可以随便写
	store := cookie.NewStore([]byte("mySession"))
	//设置session中间件,参数mySession,指的是session的名字,可以随便写
	r.Use(sessions.Sessions("mySession", store))

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
