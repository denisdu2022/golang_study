package main

import "github.com/gin-gonic/gin"

// 首页函数
// *gin.Context是上下文对象(收发),ctx是形参
func index(ctx *gin.Context) {
	//4.响应
	//状态码200 响应一个字符串
	ctx.String(200, "Hello WEB!")
}

func main() {
	//1.通过Gin获取引擎对象,可以理解为路由对象
	r := gin.Default()

	//3.路由映射
	r.GET("/", index)

	//2.启动相当于net阻塞
	//如果不写端口,默认是本机8080端口
	r.Run()
}
