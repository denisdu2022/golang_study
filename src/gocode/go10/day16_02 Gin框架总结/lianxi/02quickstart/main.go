package main

import "github.com/gin-gonic/gin"

// 首页函数
// *gin.Context是上下文对象,上取数据,下作响应
func getIndex(ctx *gin.Context) {
	//响应 状态码200 响应一个字符串
	ctx.String(200, "hello web!!!")

}

func main() {
	//1.获取路由对象
	//通过gin获取引擎对象,可以理解为路由对象
	r := gin.Default()

	//3.路由映射
	r.GET("/", getIndex)

	//2.启动,相当于net的阻塞
	//不写默认本机8080端口
	r.Run()
}
