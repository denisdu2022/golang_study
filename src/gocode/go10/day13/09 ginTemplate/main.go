package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
)

func main() {
	//获取路由对象
	r := gin.Default()

	//加载HTML文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/index", func(context *gin.Context) {
		context.HTML(200, "index.html", nil)
	})

	//动态请求
	r.GET("/timer", func(context *gin.Context) {
		now := now.EndOfMinute()
		fmt.Println(now)
		context.HTML(200, "timer.html", gin.H{
			"now": now,
		})
	})

	//启动
	r.Run()

}
