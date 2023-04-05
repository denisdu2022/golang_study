package router

import (
	"bingo_api/applcation/api"
	. "bingo_api/applcation/utils"
	"github.com/gin-gonic/gin"
)

//Router *gin.RouterGroup 路由组,接收initialize里边传来的大的路由组对象

func InitUserRouter(Router *gin.RouterGroup) {
	//在使用Router分一个user的小组
	UserRouter := Router.Group("user")
	{
		//用户认证登录 Post请求(    路径,  认证函数 )
		//UserRouter.POST("authenticate", api.UserAuthenticate)
		//注册( 路由,[]string{"请求方式"},路径,handles方法 )
		Register(UserRouter, []string{"GET", "POST"}, "authenticate", api.UserAuthenticate)
	}
}
