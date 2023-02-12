package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//获取路由对象
	r := gin.Default()

	//获取get的请求参数

	//Get路由
	//	r.GET("/user:id", func(ctx *gin.Context) {
	//
	//		//获取URL路径参数
	//		ID := ctx.Param("id")
	//
	//		//响应
	//		ctx.JSON(http.StatusOK, gin.H{
	//			"userID": ID,
	//		})
	//})

	//Query
	//r.GET("/data", func(ctx *gin.Context) {
	//
	//	//获取请求数据参数
	//	kd := ctx.Query("kd")
	//
	//	//响应
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"kd": kd,
	//	})
	//})

	//r.GET("/data", func(ctx *gin.Context) {
	//	//获取用户名密码
	//	name := ctx.Query("user")
	//	pwd := ctx.Query("pwd")
	//
	//	//数据校验
	//	if name == "test" && pwd == "123" {
	//		fmt.Println("登录成功!!!")
	//	} else {
	//		fmt.Println("登录失败,用户名或者密码错误!!!")
	//	}
	//
	//	//响应
	//	ctx.JSON(200, gin.H{
	//		"msg": "test",
	//	})
	//})

	//DefaultQuery
	//r.GET("/data", func(ctx *gin.Context) {
	//	name := ctx.Query("user")
	//	age := ctx.DefaultQuery("age", "18")
	//
	//	//响应
	//	ctx.JSON(200, gin.H{
	//		"name": name,
	//		"age":  age,
	//	})
	//})

	//GetQuery 返回一个bool值,判断是否存在
	r.GET("/data", func(ctx *gin.Context) {
		//获取用户id
		userID, ok := ctx.GetQuery("id")

		//响应
		ctx.JSON(200, gin.H{
			"user": userID,
			"是否存在": ok,
		})
	})

	//启动
	r.Run()
}
