package api

import (
	"api_01/application/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserCreateDemo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "用户创建成功!",
	})
}

func Signup(ctx *gin.Context) {
	var sign model.SignUpParam
	if err := ctx.ShouldBindJSON(&sign); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	//成功保存入库
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "用户登录成功!",
	})
}
