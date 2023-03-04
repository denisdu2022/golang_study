package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func app01(ctx *gin.Context) {
	fmt.Println("这是app01...")

	//响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "这是app01...",
	})
}

func app02(ctx *gin.Context) {
	fmt.Println("这是app02...")

	//响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "这是app02...",
	})
}

//1. 构建中间件函数

//func M1() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		fmt.Println("这是M1中间件:::M1")
//		//获取客户端IP
//		fmt.Println("客户端IP>>> ", ctx.ClientIP())
//	}
//}
//
//func M2() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		fmt.Println("这是M2中间件:::M2")
//		//获取请求方式
//		fmt.Println("请求路径是>>> ", ctx.Request.URL)
//	}
//}

////中间件应用场景一  IP黑名单
//
//func M1() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		fmt.Println("这是M1中间件:::M1")
//
//		//IP黑名单
//		var blackMenu = "127.0.0.1"
//
//		//获取客户端IP
//		fmt.Println("客户端IP>>> ", ctx.ClientIP())
//		//判断客户端是否在IP黑名单
//		if blackMenu == ctx.ClientIP() {
//			//拦截语法 不会在进入路由系统,直接返回
//			ctx.String(403, "Forbidden!")
//			ctx.Abort()
//		}
//
//	}
//}
//
//func M2() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		fmt.Println("这是M2中间件:::M2")
//		//获取请求方式
//		fmt.Println("请求路径是>>> ", ctx.Request.URL)
//	}
//}

////中间件应用场景二 权限控制
//
//func M1() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		fmt.Println("这是M1中间件:::M1")
//		//获取客户端IP
//		fmt.Println("客户端IP>>> ", ctx.ClientIP())
//
//	}
//}
//
//func M2() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		fmt.Println("这是M2中间件:::M2")
//		//获取请求方式
//		fmt.Println("请求路径是>>> ", ctx.FullPath())
//
//		//访问路径
//		var urlAcl = "/app01"
//
//		//如果访问是路径是/app01则拒绝,返回权限错误
//		if urlAcl == ctx.FullPath() {
//			ctx.String(403, "访问权限错误!!!")
//			//拒绝
//			ctx.Abort()
//		}
//	}
//}

////next()语法 作用一 直接执行下一个中间件
//
//func M1() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//
//		//IP
//		var blackMenu = "127.0.0.1"
//
//		//获取客户端IP
//		fmt.Println("客户端IP>>> ", ctx.ClientIP())
//		//判断客户端是否在IP黑名单
//		if blackMenu == ctx.ClientIP() {
//			//直接执行下一个中间件
//			ctx.Next()
//		} else {
//			fmt.Println("这是M1中间件:::M1")
//		}
//
//	}
//}
//
//func M2() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		fmt.Println("这是M2中间件:::M2")
//		//获取请求方式
//		fmt.Println("请求路径是>>> ", ctx.Request.URL)
//	}
//}

//next()语法 作用二 直接执行下一个中间件后回来在执行next()后的

func M1() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//IP
		var blackMenu = "127.0.0.1"

		//获取客户端IP
		fmt.Println("客户端IP>>> ", ctx.ClientIP())
		//判断客户端是否在IP黑名单
		if blackMenu == ctx.ClientIP() {
			//直接执行下一个中间件
			ctx.Next()
		} else {
			fmt.Println("这是M1中间件:::M1")
		}

		fmt.Println("next()执行后:::M1")

	}
}

func M2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("这是M2中间件:::M2")
		//获取请求方式
		fmt.Println("请求路径是>>> ", ctx.Request.URL)
	}
}

func main() {

	//获取路由对象
	app := gin.Default()

	//2.注册中间件
	//app.Use() 加载中间件
	app.Use(M1())
	app.Use(M2())

	//中间件是有顺序的 当请求进来时是依次执行的
	//app.Use(M2())
	//app.Use(M1())

	//路由
	app.GET("/app01", app01)

	app.GET("/app02", app02)

	//启动 默认8080端口
	app.Run()

}
