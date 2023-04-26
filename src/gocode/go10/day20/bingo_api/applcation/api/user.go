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

func UserAuthenticate(ctx *gin.Context) {

	//用户登录认证业务
	user, err := services.UserLogin(ctx)

	//判断
	//如果出现错误
	if err != nil {
		//则响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    constants.CodeAuthticateFail,
			"message": err.Error(),
		})
		return
	}

	//生成token
	//实例化jwt对象
	newJWT := utils.NewJWT()
	//创建载荷Claims
	publicClaims := utils.PublicClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}
	//创建token
	token, err := newJWT.AccessToken(publicClaims)
	if err != nil {
		panic(err.Error())
	}
	//响应
	ctx.JSON(http.StatusOK, gin.H{
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
