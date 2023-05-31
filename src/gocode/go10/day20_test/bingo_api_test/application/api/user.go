package api

import (
	"bingotest01/application/constants"
	"bingotest01/application/services"
	"bingotest01/application/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//用户认证登录

func UserAuthenticate(ctx *gin.Context) {
	////响应测试
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"msg":  "用户登录成功",
	//})

	//用户登录认证业务
	user, err := services.UserLogin(ctx)
	//判断
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    constants.CodeFail,
			"message": err.Error(),
		})
		return
	}

	//实例化token对象
	newJwt := utils.JWT{}
	//Claims
	publicClaims := utils.PublicClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}
	//生成token
	token, err := newJwt.AccessToken(publicClaims)
	if err != nil {
		panic(err.Error())
	}
	//响应返回token
	ctx.JSON(http.StatusCreated, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data": map[string]interface{}{
			"token": token,
		},
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
