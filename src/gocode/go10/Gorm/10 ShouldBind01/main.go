package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

//定义接収shouldBind的结构体

type User struct {
	User string `json:"name" form:"name"`
	Pwd  string `json:"pwd" form:"pwd"`
}

func main() {

	//1.获取路由对象
	//默认使用了两个中间件Logger(),Recovery()
	r := gin.Default()

	//路由
	r.POST("/shouldBind/data", func(ctx *gin.Context) {
		//"{"user":"luobo","pwd":"123"}"--->反序列化---->映射到结构体对象
		//初始化user结构体对象
		user := User{}
		//通过shouldBind函数,将请求参数绑定到struct对象,处理JSON请求代码是一样的. &user传地址
		//如果是post请求则根据规则Content-Type判断,接收的是JSON数据,还是普通的http请求参数
		if ctx.ShouldBind(&user) == nil {
			//绑定成功,打印请求参数
			log.Println(">>> ", user.User)
			log.Println(">>> ", user.Pwd)
		}

		//响应
		ctx.JSON(200, gin.H{
			"user": user,
		})

	})

	//启动
	r.Run()
}
