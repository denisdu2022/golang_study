package initialize

import (
	"api_01/application/router"
	"github.com/gin-gonic/gin"
)

//路由初始化

func InitRouter() *gin.Engine {
	//获取路由对象
	Router := gin.Default()

	//路由分组
	ApiGroup := Router.Group("/api")

	router.InitUser(ApiGroup)

	//返回路由对象
	return Router
}
