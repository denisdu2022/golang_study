package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

//请求日志中间件

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//开始执行时间
		stTime := time.Now()
		//获取请求路径
		rePath := c.Request.URL.Path

		//继续执行
		c.Next()

		//执行结束时间 time.Since 里边传入开始时间会使用结束时间减去开始时间
		endTime := time.Since(stTime)

		//记录日志
		zap.S().Info(
			//请求路径
			rePath,
			//请求状态
			zap.Int("status", c.Writer.Status()),
			//请求方式
			zap.String("method", c.Request.Method),
			//客户端IP
			zap.String("clientIP", c.ClientIP()),
			//ua头
			zap.String("user-agent", c.Request.UserAgent()),
			//执行时间
			zap.Duration("cost timer", endTime),
		)
	}
}
