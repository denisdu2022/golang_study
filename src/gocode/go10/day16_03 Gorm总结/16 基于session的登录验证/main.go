package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
	ID   int
	Name string
}

//登录页面

func LoginHtml(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

//登录

func Login(ctx *gin.Context) {
	//获取登录信息
	//roleID := ctx.PostForm("role_id")

	account := ctx.Param("user")
	pwd := ctx.Param("pwd")

	//数据库查询
	var user User
	DB.Where("account = ? and pwd = ?", account, pwd).Take(&user)

	//判断校验
	if user.ID != 0 {
		userID := strconv.Itoa(user.ID)
		//登录成功
		//初始化session对象
		session := sessions.Default(ctx)
		//设置session数据
		session.Set("user_id", userID)
		//保存session数据
		session.Save()
		//登录成功后重定向到首页
		ctx.Redirect(http.StatusMovedPermanently, "/")
	} else {
		//登录失败
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"err": "用户名或者密码错误",
		})
	}

}

//登出

func Logout(ctx *gin.Context) {
	//初始化session对象
	session := sessions.Default(ctx)
	//删除session数据
	session.Delete("user_id")
	//重定向到登录页面
	ctx.Redirect(http.StatusMovedPermanently, "/login")
}

func main() {
	//获取引擎对象
	r := gin.Default()

	//创建基于cookie的存储引擎,css参数是用于加密的密钥,可以随便填写
	store := cookie.NewStore([]byte("css"))
	//设置session中间件,参数mySession,指的是session的名字,也是cookie的名字
	//store是前面创建的引擎
	//r.Use()加载中间件
	r.Use(sessions.Sessions("mySession", store))

	//路由
	r.GET("/test", func(ctx *gin.Context) {
		//初始化session对象
		//从上下文对象中获取session对象
		session := sessions.Default(ctx)

		//通过session.Get读取session值
		//session是键值对格式数据,因此需要通过key查询数据
		if session.Get("hello") != "world" {
			//设置session数据
			session.Set("hello", "world")
			//删除session数据
			session.Delete("delSession")
			//删除整个session
			session.Clear()
			//保存session数据
			session.Save()

		}
		ctx.JSON(http.StatusOK, gin.H{
			"hello": session.Get("hello"),
		})
	})

	//启动并监听8080端口
	r.Run()
}
