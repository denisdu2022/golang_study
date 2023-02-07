package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 结构体
type Student struct {
	Name string
	Age  int
}

// 首页路由函数
func getIndex(ctx *gin.Context) {
	//构建Go语言下的各类数据类型
	//变量
	var name string
	name = "haha"
	var age = 23
	//切片
	var books = []string{"三国演义", "红楼梦", " 哪吒闹海", "王小虎捉妖"}
	//map
	var stuMap = map[string]interface{}{
		"name":   "bin",
		"age":    21,
		"gender": "男",
	}
	//在go语言中取map对象的值是
	fmt.Println("姓名: ", stuMap["name"])

	//实例化结构体对象
	stuStruct := Student{Name: "han", Age: 23}

	//练习
	var students = []Student{{Name: "liu", Age: 22}, {Name: "zhang", Age: 21}, {"li", 24}}

	//响应返回HTML页面
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"name":      name,
		"age":       age,
		"books":     books,
		"stuMap":    stuMap,
		"stuStruct": stuStruct,
		"students":  students,
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
