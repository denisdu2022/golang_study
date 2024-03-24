package router

import (
	"bingotest01/application/api"
	"bingotest01/application/utils"
	"github.com/gin-gonic/gin"
)

//主机路由

func InitHostRouter(Router *gin.RouterGroup) {

	//主机路由组
	HostRouter := Router.Group("host")
	//HostRouter.Use(middleware.JWTAuthorization())
	{
		//添加主机类别路由
		utils.Register(HostRouter, []string{"POST"}, "category", api.HostCategoryCreate)

		//查看所有主机类别
		utils.Register(HostRouter, []string{"GET"}, "category", api.HostCategoryList)

		//添加主机
		utils.Register(HostRouter, []string{"POST"}, "", api.HostCreate)

		//查看主机类别
		utils.Register(HostRouter, []string{"GET"}, "", api.HostList)

		//删除主机
		utils.Register(HostRouter, []string{"DELETE"}, "", api.HostDelete)

		//Console
		utils.Register(HostRouter, []string{"GET"}, ":id/console", api.Console)

		//主机 批量任务 执行指令
		utils.Register(HostRouter, []string{"POST"}, "cmd", api.HostRunCmdList)

	}

}
