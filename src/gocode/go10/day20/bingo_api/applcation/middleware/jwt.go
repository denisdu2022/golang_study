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
		//获取请求头
		token := c.Request.Header.Get("Authorization")
		//判断token 如果token为空
		if token == "" {
			//响应认证失败状态码和令牌认证无效
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.CodeAuthticateFail,
				"message": constants.TokenIsInvalid,
			})
			c.Abort()
			return
		}

		//日志记录
		zap.S().Debugf("get token: %#v", token)

		//实例化jwt对象
		j := utils.NewJWT()
		//parseToken 解析token包含的信息
		claims, err := j.GetPayloadByToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.CodeAuthticateFail,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		//存储认证的载荷信息和token,保留至后边使用
		c.Set("claims", claims)
		c.Set("access_token", token)
		c.Next()
	}
}
