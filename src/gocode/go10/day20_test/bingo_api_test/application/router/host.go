package router

import (
	"bingotest01/application/api"
	"bingotest01/application/middleware"
	"bingotest01/application/utils"
	"github.com/gin-gonic/gin"
)

//主机路由

func InitHostRouter(Router *gin.RouterGroup) {

	//主机路由组
	HostRouter := Router.Group("host")
	HostRouter.Use(middleware.JWTAuthorization())
	{
		//创建主机
		utils.Register(HostRouter, []string{"GET", "POST"}, "", api.HostCreate)
	}

}
