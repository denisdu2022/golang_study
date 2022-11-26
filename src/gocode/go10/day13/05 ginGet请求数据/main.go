package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//GET参数请求数据

	r := gin.Default()

	r.GET("data", func(context *gin.Context) {
		//获取Get参数请求数据
		//fmt.Println("KD: ", context.Query("kd"))
		//kd := context.Query("kd")

		//获取Get参数请求数据 DefaultQuery
		fmt.Println("age : ", context.DefaultQuery("age", "18"))
		age := context.DefaultQuery("age", "18")
		//响应
		context.JSON(200, gin.H{
			//"KD": kd,
			"age": age,
		})
	})

	r.Run(":8070")
}
