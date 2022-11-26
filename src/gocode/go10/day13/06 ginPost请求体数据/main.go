package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	r.POST("/data", func(context *gin.Context) {
		//获取请求体数据 PostFrom是接收form表单格式的数据
		user := context.PostForm("user")
		pwd := context.PostForm("pwd")
		//多个键值
		//hobby := context.PostForm("hobby")
		hobbys := context.PostFormArray("hobbys")

		//响应
		context.JSON(200, gin.H{
			"data":  "data",
			"User":  user,
			"Pwd":   pwd,
			"Hobby": hobbys,
		})
	})

	r.Run()
}
