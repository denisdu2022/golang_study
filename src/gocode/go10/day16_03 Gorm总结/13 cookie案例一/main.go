package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Login

func login(ctx *gin.Context) {
	//因为是任意路由,所以我们需要先获取请求方式
	//如果是Get请求返回登录页面
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "login.html", nil)
	} else {
		//获取参数
		user := ctx.PostForm("user")
		pwd := ctx.PostForm("pwd")
		if user == "admin" && pwd == "123" {
			//给客户端设置cookie
			//name 键
			//value 值
			//maxAge int类型,单位秒(s)
			//path cookie 所在目录
			//domain string 域名
			//secure 是否只能通过https访问,一般测试环境也允许http访问所以这里设置为false
			//httponly bool 是否允许别人通过js获取自己的cookie
			ctx.SetCookie("isLogin", "true", 2000000, "/", "127.0.0.1", false, true)
			//设置第二条cookie
			ctx.SetCookie("username", user, 2000000, "/", "127.0.0.1", false, true)
			//跳转首页
			ctx.Redirect(http.StatusMovedPermanently, "/index")
		} else {
			ctx.Redirect(http.StatusFound, "/login")
		}
	}
}

//index

func GetIndex(ctx *gin.Context) {
	//获取cookie
	isLogin, _ := ctx.Cookie("isLogin")
	//判断
	if isLogin == "true" {
		user, _ := ctx.Cookie("username")
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"username": user,
		})
	} else {
		ctx.Redirect(http.StatusFound, "/login")
	}
}

//multitemplate函数

func CreateMyRender() multitemplate.Renderer {
	//定义 render对象
	render := multitemplate.NewRenderer()

	//页面路径
	render.AddFromFiles("login.html", "templates/login.html")
	render.AddFromFiles("index.html", "templates/index.html")
	//返回render对象
	return render
}

func main() {
	//获取路由对象
	r := gin.Default()

	//开放静态文件窗口
	r.Static("/static", "./static")

	//加载模板文件
	r.HTMLRender = CreateMyRender()

	//路由
	r.Any("/login", login)

	r.GET("/index", GetIndex)

	//启动 默认8080端口
	r.Run()
}
