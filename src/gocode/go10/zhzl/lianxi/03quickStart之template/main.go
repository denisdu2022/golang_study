package main

//导入gin
import "github.com/gin-gonic/gin"

// 首页函数
func getIndex(ctx *gin.Context) {

	//响应json数据
	ctx.JSON(200, gin.H{
		"user": "haha",
		"pwd":  "123",
	})
}

// login函数
func getLogin(ctx *gin.Context) {
	//响应一个页面
	ctx.HTML(200, "login.html", nil)
}

func main() {
	//1.获取路由对象
	r := gin.Default()

	//加载模板文件(需要返回页面时)
	//加载templates/*下的所有文件
	r.LoadHTMLGlob("templates/*")

	//3.路由
	r.GET("/index", getIndex)

	//login路由
	r.GET("login", getLogin)
	//可以使用匿名函数
	r.POST("/login", func(ctx *gin.Context) {
		//1.获取用户数据
		user := ctx.PostForm("user")
		pwd := ctx.PostForm("pwd")

		//2.模拟数据校验
		if user == "text" && pwd == "123" {
			//响应登录成功
			ctx.JSON(200, gin.H{
				"msg": "登录成功!",
			})
		} else {
			//响应登录失败
			ctx.JSON(200, gin.H{
				"msg": "登录失败,用户名或密码错误!",
			})
		}

	})

	//2.启动
	r.Run()

}
