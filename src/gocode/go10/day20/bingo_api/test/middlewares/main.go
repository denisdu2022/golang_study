package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func M1() gin.HandlerFunc {
	//中间件里写匿名函数
	return func(ctx *gin.Context) {
		//路由处理前执行
		fmt.Println("M1中间件执行,请求")

		ctx.Next()

		//路由处理后执行
		fmt.Println("M1中间件执行,响应")

	}
}

func main() {

	//获取路由对象
	Router := gin.Default()

	//加载全局中间件
	Router.Use(M1())

	//路由
	Router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "访问成功!!!",
		})
	})

	//启动
	Router.Run(":8060")

}
