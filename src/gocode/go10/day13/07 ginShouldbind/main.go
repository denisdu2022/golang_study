package main

import "github.com/gin-gonic/gin"

type User struct {
	User string `json:"user"`
	Pwd  int    `json:"pwd"`
}

func main() {

	r := gin.Default()

	//ShouldBind
	r.POST("/shouldBind/data", func(context *gin.Context) {
		//{
		//    "user":"luobo",
		//    "pwd":"123",
		//}  --->反序列化--->映射到结构体对象

		//初始化结构体对象(实例化结构体对象)
		user := User{}

		err := context.ShouldBind(&user)
		if err != nil {
			return
		}
		//响应
		context.JSON(200, gin.H{
			//"ShouldBind": "ShouldBind data",
			"User": user,
		})
	})

	//

	r.Run()

}
