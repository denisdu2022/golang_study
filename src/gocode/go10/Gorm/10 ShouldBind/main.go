package main

import (
	"github.com/gin-gonic/gin"
)

//定义接収shouldBind的结构体

type User struct {
	User string `json:"name" form:"name"`
	Pwd  string `json:"pwd" form:"pwd"`
}

func main() {

	//获取路由对象
	r := gin.Default()

	//路由
	r.POST("/shouldBind/data", func(ctx *gin.Context) {
		//"{"user":"luobo","pwd":"123"}"--->反序列化---->映射到结构体对象
		//初始化user结构体对象
		user := User{}
		//通过shouldBind将请求数据放在结构体中,&user传地址
		err := ctx.ShouldBind(&user)
		if err != nil {
			return
		}

		//数据校验判断
		if user.User == "test" && user.Pwd == "123" {
			//响应
			ctx.JSON(200, gin.H{
				"msg":  "登录成功",
				"user": user,
			})
		} else {
			ctx.JSON(200, gin.H{
				"err": "用户名或密码错误",
			})
		}

	})

	//启动
	r.Run()
}
