package initialize

import (
	"bingotest01/application/middleware"
	"bingotest01/application/router"
	"github.com/gin-gonic/gin"
)

//路由初始化函数

func InitRouter() *gin.Engine {
	//创建路由
	Router := gin.Default()

	//使用全局中间件
	Router.Use(middleware.RequestLogger())

	//路由分组
	ApiGroup := Router.Group("/api")

	//初始化用户路由组
	router.InitUserRouter(ApiGroup)

	//初始化主机路由组
	router.InitHostRouter(ApiGroup)

	//返回路由对象
	return Router
}
