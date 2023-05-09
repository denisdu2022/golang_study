package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 自定义中间件
//函数返回一个函数,函数返回函数,使用时需要调用  MyLogger()

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取当前时间戳
		t := time.Now()

		//设置connect
		c.Set("example", "123456")

		foo := time.Now().Add(time.Second * 3)
		fmt.Println("foo 的执行时间: ", foo)

		//让原本改执行的逻辑继续执行 c.Next()
		c.Next()

		//获取执行时间,让结束时间减去开始时间 ,在go语言中可以使用Since
		end := time.Since(t)
		//耗时时间
		fmt.Printf("耗时: %#v\n", end)
		//获取状态信息
		status := c.Writer.Status()
		fmt.Println("状态信息: ", status)
	}
}

func main() {

	//获取路由对象
	router := gin.Default()
	//使用自定义中间件
	router.Use(MyLogger())

	//逻辑路由
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	//启动:端口8083
	router.Run(":8083")
}
