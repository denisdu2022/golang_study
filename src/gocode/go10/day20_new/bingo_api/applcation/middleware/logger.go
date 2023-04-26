package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

//RequestLogger 接收gin框架的请求日志

func RequestLogger() gin.HandlerFunc {
	//需要记录的是每次客户端访问时的上下文信息,所以此处返回了一个匿名函数,在客户端请求时,触发中间件的时候,自动执行这个匿名函数
	return func(c *gin.Context) {
		//执行开始时间
		start := time.Now()
		//获取请求路径
		path := c.Request.URL.Path

		c.Next()

		//执行结束的时间  time.Since(传入开始时间) 会自动减去开始时间,得出用时
		cost := time.Since(start)

		//记录日志
		zap.S().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("cost timer", cost),
		)
	}
}

func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a 检查是否有断开的连接,因为它实际上不是
				// condition that warrants a panic stack trace. 保证进行紧急堆栈跟踪的条件。
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection rest by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.S().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.S().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.S().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
