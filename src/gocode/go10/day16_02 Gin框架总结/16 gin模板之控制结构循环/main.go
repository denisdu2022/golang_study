package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Student struct {
	Name string
	Age  int
}

// 首页路由函数
func getIndex(ctx *gin.Context) {
	//控制结构之分支结构
	var name = "zhangSan"
	var age = 17
	var book18 = []string{"金瓶梅", "剪灯新话", " 国色天香", "聊斋", "红楼梦"}
	var book17 = []string{"三国演义", "红楼梦", " 哪吒闹海", "王小虎捉妖"}
	name17 := "liu"
	age17 := 17

	var students = []Student{{Name: "liu", Age: 22}, {Name: "zhang", Age: 21}, {"li", 24}}

	//响应返回HTML页面
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"name":     name,
		"age":      age,
		"books":    book18,
		"book17":   book17,
		"name17":   name17,
		"age17":    age17,
		"students": students,
	})

}

func main() {
	//获取路由对象
	r := gin.Default()
	//加载模板文件
	r.LoadHTMLGlob("templates/*")
	//开放静态文件窗口
	r.Static("/static", "./static")

	//路由
	r.GET("/", getIndex)

	//启动
	r.Run()
}
