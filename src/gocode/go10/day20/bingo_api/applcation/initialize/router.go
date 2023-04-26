package initialize

import (
	"bingo_api/applcation/middleware"
	"bingo_api/applcation/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//1.创建路由
	Router := gin.Default()

	//注册全局中间件
	//Router.Use(middleware.JWTAuthorization(), middleware.RequestLogger(), middleware.GinRecovery(true), middleware.ExceptionMiddleware)
	Router.Use(middleware.RequestLogger(), middleware.GinRecovery(true), middleware.ExceptionMiddleware)
	//加载跨域(CORS)中间件
	Router.Use(middleware.Cors)

	//3.路由分组
	//ApiGroup相当于一个大的路由组
	ApiGroup := Router.Group("/api")

	//4.初始化用户相关的路由
	//router是router包 InitUserRouter相当于是大的路由组下的下组,路由组的嵌套
	router.InitUserRouter(ApiGroup)

	//5.初始化主机相关的路由
	router.InitHostRouter(ApiGroup)

	//2.返回路由对象
	return Router
}
