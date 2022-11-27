package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	User string `form:"name" json:"name"`
	Pwd  int    `form:"pwd" json:"pwd"`
}

func main() {

	r := gin.Default()

	//ShouldBind
	r.POST("/shouldBind/data", func(context *gin.Context) {
		//{
		//    "name":"luobo",
		//    "pwd":"123",
		//}  --->反序列化--->映射到结构体对象

		//初始化结构体对象(实例化结构体对象)
		user := User{}

		err := context.ShouldBind(&user)
		if err != nil {
			return
		}
		fmt.Println(context.ShouldBind(&user))

		if user.User == "luobo" && user.Pwd == 123 {
			//响应
			context.JSON(200, gin.H{
				"ShouldBind": "ShouldBind data",
				"msg":        "登录成功用户名密码正确!!!",
				"User":       user,
			})
		} else {
			context.JSON(200, gin.H{
				"ShouldBind": "ShouldBind data",
				"msg":        "登录失败用户名密码不正确...",
				"User":       user,
			})
		}

	})

	//

	r.Run()

}
