package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//获取路由对象
	r := gin.Default()

	//POST

	r.POST("/data", func(ctx *gin.Context) {
		//获取请求体数据  PostForm是接受Form表单格式数据的
		//获取user和pwd 通过PostForm获取到的参数值是string类型
		//user := ctx.PostForm("user")
		//pwd := ctx.PostForm("pwd")
		////PostForm只能接收一个值,需要使用PostFormArray 使用于复选框
		////hobby := ctx.PostForm("hobby")
		//hobby := ctx.PostFormArray("hobby")
		//
		////响应
		//ctx.JSON(200, gin.H{
		//	"user":  user,
		//	"pwd":   pwd,
		//	"hobby": hobby,
		//})

		////DefaultPostForm
		////跟PostForm的区别是可以设置默认值,第二个参数是默认值
		//user := ctx.DefaultPostForm("name", "guest")
		//
		////响应
		//ctx.JSON(200, gin.H{
		//	"user": user,
		//})

		//GetPostForm
		//获取id参数,跟PostForm一样,通过GetPostForm获取的参数值也是string类型
		//区别是GetPostForm返回两个参数,第一个是参数值,第二个参数是参数是否存在的bool值,可以用来判断参数是否存在
		userID, ok := ctx.GetPostForm("id")
		//参数是否存在
		if !ok {
			//参数不存在
			fmt.Println("id参数不存在!")
		}

		//响应
		ctx.JSON(200, gin.H{
			"userID": userID,
		})

	})

	//启动
	r.Run()
}
