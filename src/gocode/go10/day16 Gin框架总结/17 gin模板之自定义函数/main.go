package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type Student struct {
	Name string
	Age  int
}

// 首页路由函数
func getIndex(ctx *gin.Context) {

	var book18 = []string{"金瓶梅", "剪灯新话", " 国色天香", "聊斋", "红楼梦"}

	//响应返回HTML页面
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"books": book18,
	})

}

func main() {
	//获取路由对象
	r := gin.Default()
	//加载自定义函数 需要放在加载模板文件之前
	//模板自定义函数 需要引入"html/template"包
	r.SetFuncMap(template.FuncMap{
		//add 是自定义函数名称 匿名函数是函数体
		"add": func(x, y int) int {
			return x + y
		},
	})
	//加载模板文件
	r.LoadHTMLGlob("templates/*")
	//开放静态文件窗口
	r.Static("/static", "./static")

	//路由
	r.GET("/", getIndex)

	//启动
	r.Run(":8060")
}
