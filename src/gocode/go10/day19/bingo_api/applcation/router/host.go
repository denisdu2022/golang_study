package router

import (
	"bingo_api/applcation/api"
	"github.com/gin-gonic/gin"
)

func InitHostRouter(Router *gin.RouterGroup) {
	//创建主机路由组
	CreateHost := Router.Group("host")
	{
		//创建主机
		CreateHost.POST("", api.HostCreate)
	}
}
