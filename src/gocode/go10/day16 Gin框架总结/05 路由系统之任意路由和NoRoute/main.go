package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	//获取路由对象
	r := gin.Default()

	//加载模板文件
	r.LoadHTMLGlob("templates/*")

	////根据提交请求的方式
	////Get login
	//r.GET("/login", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"msg": "Get Login!!!",
	//	})
	//})
	////Post Login
	//r.POST("/login", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"msg": "Post Login!!!",
	//	})
	//})

	//任意路由
	r.Any("/login", func(context *gin.Context) {
		//请求路径
		path := context.FullPath()
		fmt.Println(path)
		//根据请求方式拿到的请求信息
		/*
			&{POST /login HTTP/1.1 1 1 ....}
			&{GET /login HTTP/1.1 1 ....}
		*/
		req := context.Request
		fmt.Println(req)
		context.JSON(200, gin.H{
			"path": "请求路径是" + path,
		})
	})

	//NoRoute路由匹配不成功,404
	//没有路径,只有handler
	r.NoRoute(func(context *gin.Context) {
		//http.StatusNotFound就是404也可以直接写404
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	//启动
	r.Run(":8090")
}
