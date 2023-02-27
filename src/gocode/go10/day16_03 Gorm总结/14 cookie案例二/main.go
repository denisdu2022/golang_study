package main

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
			//设置登录时间
			now := time.Now().String()[:19]
			fmt.Println(now)

			//给客户端设置cookie
			ctx.SetCookie("lastTime", now, 2000000, "/", "127.0.0.1", false, true)
			//跳转首页
			ctx.Redirect(http.StatusMovedPermanently, "/index")
		} else {
			ctx.Redirect(http.StatusMovedPermanently, "/login")
		}
	}
}

//index

func GetIndex(ctx *gin.Context) {
	//获取cookie
	lastTime, _ := ctx.Cookie("lastTime")
	fmt.Println(lastTime)
	//判断
	if lastTime != "" {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"lastTime": lastTime,
		})
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
