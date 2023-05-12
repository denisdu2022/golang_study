package api

import (
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
