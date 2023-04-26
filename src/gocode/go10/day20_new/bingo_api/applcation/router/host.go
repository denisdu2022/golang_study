package router

import (
	"bingo_api/applcation/api"
	"bingo_api/applcation/utils"
	"github.com/gin-gonic/gin"
)

func InitHostRouter(Router *gin.RouterGroup) {
	//创建主机路由组
	CreateHost := Router.Group("host")
	{
		//创建主机
		//CreateHost.POST("", api.HostCreate)
		//Get请求做跨域测试
		utils.Register(CreateHost, []string{"GET", "POST"}, "", api.HostCreate)

	}
}
