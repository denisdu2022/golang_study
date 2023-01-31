package main

import "github.com/gin-gonic/gin"

func main() {
	//获取路由对象
	r := gin.Default()

	//POST

	r.POST("/data", func(ctx *gin.Context) {
		//获取请求体数据  PostForm是接受Form表单格式数据的
		user := ctx.PostForm("user")
		pwd := ctx.PostForm("pwd")
		//PostForm只能接收一个值,需要使用PostFormArray 使用于复选框
		//hobby := ctx.PostForm("hobby")
		hobby := ctx.PostFormArray("hobby")

		//响应
		ctx.JSON(200, gin.H{
			"user":  user,
			"pwd":   pwd,
			"hobby": hobby,
		})

	})

	//启动
	r.Run()
}
