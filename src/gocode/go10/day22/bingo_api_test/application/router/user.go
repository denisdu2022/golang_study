package router

import (
	"bingotest01/application/api"
	"bingotest01/application/utils"
	"github.com/gin-gonic/gin"
)

//用户路由

func InitUserRouter(Router *gin.RouterGroup) {
	//分路由小组
	UserRouter := Router.Group("user")
	{
		////用户认证登录
		//UserRouter.POST("authenticate", api.UserAuthenticate)

		//路由注册 (路由组对象,[]string{"请求方式"},"路径","路由函数")
		utils.Register(UserRouter, []string{"GET", "POST"}, "authenticate", api.UserAuthenticate)
		//用户创建
		utils.Register(UserRouter, []string{"POST"}, "create", api.UserCreate)
	}
}
