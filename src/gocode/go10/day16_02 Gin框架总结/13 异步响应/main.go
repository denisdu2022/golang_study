package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {

	//1.创建路由对象
	r := gin.Default()

	//路由

	//3.异步
	r.GET("/async", func(ctx *gin.Context) {

		//创建副本
		copyContext := ctx.Copy()

		//异步处理
		go func() {
			fmt.Println("--------------------start----------------------------------------")
			time.Sleep(3 * time.Second)
			log.Println("异步执行: " + copyContext.Request.URL.Path)
			fmt.Println("异步执行完毕!!!")
			fmt.Println("--------------------end-------------------------------------------")
		}()

		//响应
		ctx.JSON(200, gin.H{
			"msg": "异步执行中!!!",
		})
	})

	//2.同步
	r.GET("sync", func(ctx *gin.Context) {
		fmt.Println("--------------------start----------------------------------------")
		time.Sleep(3 * time.Second)
		log.Println("同步执行: ", ctx.Request.URL.Path)
		fmt.Println("同步执行完毕!!!")
		//响应
		ctx.JSON(200, gin.H{
			"msg": "同步执行完毕中",
		})
		fmt.Println("--------------------end-------------------------------------------")
	})

	//2.启动
	r.Run()
}
