package main

//01 导入gin包
import (
	"github.com/gin-gonic/gin"
)

// *.gin.Context上下文对象(收发) ctx是形参
func getuser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"user_id":     1001,
		"user_name":   "denis",
		"user_passwd": 123,
	})
}

func main() {

	//02.通过gin获取引擎对象,可以理解为路由对象
	r := gin.Default()

	//03.路由映射
	r.GET("/index", getuser)

	//05.加载模板文件
	r.LoadHTMLGlob("templates/*")

	//06.返回页面 使用匿名函数 context是实参
	r.GET("/login", func(context *gin.Context) {
		context.HTML(200, "login.html", nil)
	})

	//07.post请求
	r.POST("/success", func(context *gin.Context) {
		//获取用户提交的数据  上取数据
		user := context.PostForm("user")
		pwd := context.PostForm("pwd")
		//数据校验
		if user == "denis" && pwd == "123" {
			//下做影响
			//context.JSON(200, gin.H{
			//	"login_status": "login success!!!",
			//)
			context.String(200, "登录成功,用户名密码正确!!!")
		} else {
			context.String(200, "登录失败,用户名密码错误...")
		}

	})

	//04.启动,相当于net包里的阻塞
	//默认本机端口8080
	r.Run()

	//
}
