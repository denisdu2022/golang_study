package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
