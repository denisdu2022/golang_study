package api

import (
	"bingo_api/applcation/constants"
	"bingo_api/applcation/services"
	"bingo_api/applcation/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

//用户认证

//func UserAuthenticate(ctx *gin.Context) {
//	ctx.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"func": "UserAuthenticate",
//	})
//}

func UserAuthenticate(ctx *gin.Context) {

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

	//生成token
	//实例化JWT对象
	newJwt := utils.JWT{}
	//claims
	publicClaims := utils.PublicClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}

	//生成token
	token, err := newJwt.AccessToken(publicClaims)
	//判断
	if err != nil {
		panic(err.Error())
	}

	//响应返回
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
	//获取用户ID
	id, err := services.CreateUser(ctx)

	if err != nil || id < 1 {
		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    constants.CodeCreateUserFail,
			"message": constants.CreateUserFail,
		})
		//记录到日志中
		zap.S().Error(constants.CreateUserFail)
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
