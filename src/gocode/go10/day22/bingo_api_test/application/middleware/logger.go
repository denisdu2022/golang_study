package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

//RequestLogger 接收gin框架的请求日志

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//执行开始时间
		start := time.Now()
		//获取请求路径
		path := c.Request.URL.Path

		//继续执行
		c.Next()

		//执行结束的时间 time.Since输入开始时间就会使用执行结束时间减去开始时间获得执行时间
		cost := time.Since(start)

		//记录日志
		zap.S().Info(
			path,                                            //路径
			zap.Int("status", c.Writer.Status()),            //状态
			zap.String("method", c.Request.Method),          //请求方式
			zap.String("ip", c.ClientIP()),                  //客户端ip
			zap.String("user-agent", c.Request.UserAgent()), //ua头
			zap.Duration("cost timer", cost),                //执行时间
		)
	}
}
