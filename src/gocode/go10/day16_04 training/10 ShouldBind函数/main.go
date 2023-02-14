package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//获取路由对象
	r := gin.Default()

	//获取post请求体数据
	//r.POST("/data", func(ctx *gin.Context) {
	//	//获取请求体数据,PostForm是接受Form表单格式数据的
	//	//获取name和pwd,通过PostForm获取到的参数值是string类型
	//	name := ctx.PostForm("user")
	//	pwd := ctx.PostForm("pwd")
	//	//PostForm只能接收一个值,PostFormArray()可以接收多个值,一般用于复选框
	//	hobby := ctx.PostFormArray("hobby")
	//
	//	//响应
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"name":  name,
	//		"pwd":   pwd,
	//		"hobby": hobby,
	//	})
	//})

	//DefaultPostForm 设置一个默认值,第二个参数是默认值
	//r.POST("/book", func(ctx *gin.Context) {
	//	bookName := ctx.DefaultPostForm("bookName", "三国演义")
	//	//响应
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"bookName": bookName,
	//	})
	//})

	//GetPostForm
	r.POST("/book", func(ctx *gin.Context) {
		//获取参数id 返回两个值,第一个是参数值本身,第二个是bool值
		bookID, ok := ctx.GetPostForm("bookID")
		if !ok {
			//参数不存在
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "bookID不存在...",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": bookID,
			})
		}
	})

	//启动
	r.Run()
}
