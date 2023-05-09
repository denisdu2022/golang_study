package router

import (
	"bingo_api/applcation/api"
	"bingo_api/applcation/middleware"
	"bingo_api/applcation/utils"
	"github.com/gin-gonic/gin"
)

func InitHostRouter(Router *gin.RouterGroup) {
	//创建主机路由组
	HostRouter := Router.Group("host")
	//jwt认证中间件
	HostRouter.Use(middleware.JWTAuthorization())
	{

		//Get请求做跨域测试
		utils.Register(HostRouter, []string{"GET"}, "", api.HostList)
		//创建主机
		//CreateHost.POST("", api.HostCreate)
		utils.Register(HostRouter, []string{"POST"}, "", api.HostCreate)
	}
}
