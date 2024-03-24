package router

import (
	"api_01/application/api"
	"github.com/gin-gonic/gin"
)

//用户路由组

func InitUser(Router *gin.RouterGroup) {
	Router.GET("user", api.UserCreateDemo)
	Router.POST("/signup", api.Signup)
}
