package middleware

import (
	"bingo_api/applcation/constants"
	"bingo_api/applcation/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

//JWTAuthorization 验证token

func JWTAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从请求头里获取token
		token := c.Request.Header.Get("Authorization")

		//fmt.Println("token>>> ", token)
		//判断token 如果获取到的token为空
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.CodeAuthticateFail,
				"message": constants.TokenIsInvalid,
			})
			//则直接退出不执行后边的逻辑
			c.Abort()
		}

		//日志记录获取到的token
		zap.S().Debugf("get token: %#v\n", token)

		//新建jwt对象
		j := utils.NewJWT()
		//parseToken 解析token包含的信息
		claims, err := j.GetPayloadByToken(token)
		//判断错误
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.CodeAuthticateFail,
				"message": err.Error(),
				"msg":     "--------------",
			})
			c.Abort()
		}

		//存储认证的载荷信息和token,保留至后面使用
		c.Set("claims", claims)
		c.Set("access_token", token)

		c.Next()

	}
}
