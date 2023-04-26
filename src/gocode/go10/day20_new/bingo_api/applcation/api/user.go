package api

import (
	"bingo_api/applcation/constants"
	"bingo_api/applcation/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

//用户认证

func UserAuthenticate(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"func": "UserAuthenticate",
	})
}

//创建用户

func UserCreate(ctx *gin.Context) {
	//获取用户ID
	id, err := services.CreateUser(ctx)

	if err != nil || id < 1 {
		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    constants.CodeCreateUserFail,
			"message": constants.CreateUserFai,
		})
		//记录到日志中
		zap.S().Error(constants.CreateUserFai)
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"code":    constants.CodeSuccess,
			"message": constants.CreateUserSuccess,
			"data":    id,
		})
		//记录到日志中
		zap.S().Error(constants.CreateUserSuccess)
	}

}
