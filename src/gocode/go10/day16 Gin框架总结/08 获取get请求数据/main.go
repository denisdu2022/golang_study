package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//获取路由对象
	r := gin.Default()

	//获取get请求数据
	//get路由

	////获取路径参数
	//r.GET("/user:id", func(ctx *gin.Context) {
	//	//获取URL参数id
	//	id := ctx.Param("id")
	//
	//	//响应
	//	ctx.JSON(200, gin.H{
	//		"user_id": id,
	//	})
	//})

	//r.GET("/data", func(ctx *gin.Context) {
	//	//1.获取请求数据参数
	//	kd := ctx.Query("kd")
	//	//打印请求参数
	//	fmt.Println("kd: ", kd)
	//
	//	//响应
	//	ctx.JSON(200, gin.H{
	//		"kd: ": kd,
	//	})
	//
	//})

	//Query
	//r.GET("/data", func(ctx *gin.Context) {
	//	//1.获取请求数据参数
	//  //获取name参数,通过Query获取的参数值是String类型
	//	name := ctx.Query("name")
	//	pwd := ctx.Query("pwd")
	//	//打印请求参数
	//	fmt.Println("name: ", name)
	//	fmt.Println("pwd: ", pwd)
	//
	//	//响应
	//	ctx.JSON(200, gin.H{
	//		"name: ": name,
	//		"pwd: ":  pwd,
	//	})
	//
	//})

	//DefaultQuery 默认值
	//r.GET("/data", func(ctx *gin.Context) {
	//	//1.获取请求数据参数
	//	name := ctx.Query("name")
	//	//当没有传pwd时的默认值
	//  //获取age参数,跟Query函数的区别是,可以通过第二个参数设置默认值
	//	age := ctx.DefaultQuery("age", "18")
	//	//打印请求参数
	//	fmt.Println("name: ", name)
	//	fmt.Println("age: ", age)
	//
	//	//响应
	//	ctx.JSON(200, gin.H{
	//		"name: ": name,
	//		"age: ":  age,
	//	})
	//
	//})

	//GetQuery
	r.GET("/data", func(ctx *gin.Context) {
		//获取请求数据参数
		//获取id参数,通过GetQuery获取的参数值也是String类型
		//GetQuery返回两个参数,第一个是参数值,第二个是参数是否存在的bool值,可以用来判断参数是否存在
		//GetQuery函数,判断参数是否存在的逻辑是,参数值为空,参数也算存在,只是没有提交参数,才算参数不存在.
		userID, ok := ctx.GetQuery("id")
		//打印请求参数
		fmt.Println("userID: ", userID)
		fmt.Println("ok: ", ok)

		//响应
		ctx.JSON(200, gin.H{
			"userID: ": userID,
			"ok: ":     ok,
		})

	})

	//启动
	r.Run()
}
