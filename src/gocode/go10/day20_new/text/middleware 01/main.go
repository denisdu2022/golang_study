package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 自定义中间件
//函数返回一个函数,函数返回函数,使用时需要调用  MyLogger()

//func MyLogger() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		//获取当前时间戳
//		t := time.Now()
//
//		//设置connect
//		c.Set("example", "123456")
//
//		foo := time.Now().Add(time.Second * 3)
//		fmt.Println("foo 的执行时间: ", foo)
//
//		//return是否有作用?
//		return
//
//		//让原本改执行的逻辑继续执行 c.Next()
//		c.Next()
//
//		//获取执行时间,让结束时间减去开始时间 ,在go语言中可以使用Since
//		end := time.Since(t)
//		//耗时时间
//		fmt.Printf("耗时: %#v\n", end)
//		//获取状态信息
//		status := c.Writer.Status()
//		fmt.Println("状态信息: ", status)
//	}
//}

//假设需要token验证后才可以进行路由响应

func TokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		//定义token类型
		var token string
		//token一般都是放在header中的
		//获取header
		for k, v := range c.Request.Header {
			//判断token是否存在
			if k == "X-Token" {
				//value会解析称为[]string
				token = v[0]
			} else {
				fmt.Println("打印K和V的值: ", k, v)
			}
			//fmt.Println("k,v,token ", k, v, token)
		}

		//判断token的值
		if token != "bobby" {
			//http.StatusUnauthorized 401没有登录
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登录",
			})
			//如果token的值不对,则直接返回,虽然return是不起作用的,这里只是示例
			//return //不生效
			//使用 c.Abort()
			c.Abort()
		}
		//token的值正确继续执行下边的逻辑
		c.Next()
	}
}

func main() {

	//获取路由对象
	router := gin.Default()
	//使用自定义中间件
	//router.Use(MyLogger())
	router.Use(TokenRequired())

	//逻辑路由
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	//启动:端口8083
	router.Run(":8083")
}
