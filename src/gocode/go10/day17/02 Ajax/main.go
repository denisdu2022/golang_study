package main

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func index(ctx *gin.Context) {

	//响应
	ctx.HTML(http.StatusOK, "cal.html", nil)
}

//cal函数

func getCal(ctx *gin.Context) {
	//获取前端数据
	strNum1 := ctx.Query("num1")
	fmt.Println(strNum1)
	strNum2 := ctx.Query("num2")
	fmt.Println(strNum2)
	//ctx.Query()获取到的是string数据 计算前需要转为int类型
	num1, _ := strconv.Atoi(strNum1)
	num2, _ := strconv.Atoi(strNum2)

	sum := num1 + num2
	fmt.Println(sum)
	//响应
	ctx.JSON(http.StatusOK, gin.H{
		"ret": sum,
	})
}

//multitemplate 函数

func CreateMyRender() multitemplate.Renderer {
	//定义render
	render := multitemplate.NewRenderer()
	//页面路径
	render.AddFromFiles("cal.html", "templates/cal.html")

	//返回render
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
	//首页
	r.GET("/", index)
	//cal
	r.Any("/cal", getCal)

	//启动 端口:8060
	r.Run(":8060")
}
