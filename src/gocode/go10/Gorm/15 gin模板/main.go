package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	//获取路由对象
	r := gin.Default()

	//加载模板文件
	r.LoadHTMLGlob("templates/*")

	//timer案例
	//动态请求路由
	r.GET("/timer", func(ctx *gin.Context) {
		now := time.Now()
		//打印当前时间
		fmt.Println("当前时间是: ", now)
		//响应返回一个页面
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"now": now,
		})
	})

	//启动
	r.Run()
}
