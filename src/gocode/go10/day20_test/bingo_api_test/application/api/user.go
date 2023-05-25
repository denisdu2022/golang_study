package api

import (
	"bingotest01/application/constants"
	"bingotest01/application/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

//用户认证登录

func UserAuthenticate(ctx *gin.Context) {
	//响应测试
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "用户登录成功",
	})
}

//创建用户

func UserCreate(ctx *gin.Context) {
	id, err := services.CreateUser(ctx)
	if err != nil || id < 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    constants.CodeCreateUserFail,
			"message": constants.CreateUserFail,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.CreateUserSuccess,
		"data":    id,
	})
}
