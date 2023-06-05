package middleware

import (
	"bingotest01/application/constants"
	"bingotest01/application/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func JWTAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从请求头里获取token
		token := c.Request.Header.Get("Authorization")

		//判断token,如果获取到的token为空
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.CodeAuthticateFail,
				"message": constants.TokenIsInvalid,
			})
			c.Abort()
		}

		//日志记录获取到的token
		zap.S().Debugf("GET token:%#v", token)

		//获取jwt对象
		j := utils.JWT{}
		//parseToken 解析token
		claims, err := j.GetPayloadByToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.CodeAuthticateFail,
				"message": err.Error(),
			})
			c.Abort()
		}
		//存储认证的载荷信息和token,保留后面使用
		c.Set("claims", claims)
		c.Set("access_token", token)

		c.Next()
	}
}
