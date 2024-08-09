package router

import (
	"bingotest01/application/api"
	"bingotest01/application/utils"
	"github.com/gin-gonic/gin"
)

//创建指令路由

func InitCmdRouter(Route *gin.RouterGroup) {

	//指令模板相关的路由组
	cmdRoute := Route.Group("cmd")
	//cmdRoute.Use(middleware.JWTAuthorization())
	{
		//指令模板类别 -添加
		utils.Register(cmdRoute, []string{"POST"}, "category", api.CmdCategoryCreate)
		//指令模板类别 -列表
		utils.Register(cmdRoute, []string{"GET"}, "category", api.CmdCategoryList)
		//指令模板 -添加
		utils.Register(cmdRoute, []string{"POST"}, "", api.CmdCreate)
		//指令模板 -列表
		utils.Register(cmdRoute, []string{"GET"}, "", api.CmdList)
		//指令模板 -信息
		utils.Register(cmdRoute, []string{"GET"}, ":id/info", api.CmdInfo)
	}
}
