package main

//导入gin

import "github.com/gin-gonic/gin"

// 首页函数
// *gin.Context上下文对象(收发) ctx是形参
func getIndex(ctx *gin.Context) {
	//4.响应
	//响应一个字符串,状态吗是200,参数是string
	//ctx.String(200, "Gin Quick Start!!!")

	//响应JSON
	ctx.JSON(200, gin.H{
		"userID":   "1001",
		"userName": "luobo",
	})
}

func main() {

	//1.通过gin获取引擎对象,可以理解为路由对象
	r := gin.Default()

	//加载模板文件(返回页面时)
	//加载"templates/*"模板文件夹下所有文件(也可以直接写"templates/login.html但是会有多个文件,直接加载所有文件")
	r.LoadHTMLGlob("templates/*")

	//3.路由映射
	r.GET("/", getIndex)

	//登录  后边的路由函数也可以使用匿名函数
	r.GET("/login", func(ctx *gin.Context) {
		//200状态码  HTML页面   先写一个nil对象
		ctx.HTML(200, "login.html", nil)
	})

	//post请求
	r.POST("/login", func(ctx *gin.Context) {
		//获取用户提交的数据
		user := ctx.PostForm("user")
		pwd := ctx.PostForm("pwd")
		//判断
		if user == "luobo" && pwd == "123" {
			//响应一个登录成功的
			ctx.JSON(200, gin.H{
				"msg": "登录成功",
			})
		} else {
			//响应一个登录失败的
			ctx.JSON(200, gin.H{
				"msg": "用户名密码错误",
			})
		}

	})

	//2.启动,相当于net包里的阻塞
	//默认不写端口未本机端口8080
	r.Run()
}
