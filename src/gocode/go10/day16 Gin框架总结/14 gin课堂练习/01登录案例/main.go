package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 首页路由
func getIndex(ctx *gin.Context) {
	//返回首页
	ctx.HTML(http.StatusOK, "index.html", nil)

}

// 登录get路由函数
func getLogin(ctx *gin.Context) {
	//获取get数据
	user := ctx.Query("user")
	pwd := ctx.Query("pwd")
	//查不到就使用第二个值(第二个参数)
	user = ctx.DefaultQuery("user", "guest")
	//打印获取到的数据
	fmt.Println("GET user: ", user)
	fmt.Println("GET pwd: ", pwd)
	//获取请求方式
	met := ctx.Request.Method
	//获取请求路径
	path := ctx.Request.URL.Path

	//响应
	//http.StatusOK就是状态码200
	ctx.JSON(http.StatusOK, gin.H{
		"用户名: ":  user,
		"密码:":    pwd,
		"请求方式: ": met,
		"请求路径":   path,
	})
}

// 登录post路由函数
func postLogin(ctx *gin.Context) {
	//获取请求数据
	user := ctx.PostForm("user")
	pwd := ctx.PostForm("pwd")
	meth := ctx.Request.Method
	path := ctx.Request.URL.Path
	//打印请求数据
	fmt.Println("POST user: ", user)
	fmt.Println("POST pwd: ", pwd)

	//数据校验
	if user == "luobo" && pwd == "123" {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "请求成功!!!",
			"用户名":  user,
			"密码":   pwd,
			"请求方式": meth,
			"请求路径": path,
		})
	} else {
		ctx.String(http.StatusOK, "用户名或密码错误!!!")
	}

}

// 没有路由匹配
func notFund(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"404": "Not Found 404!!!",
	})
}

//获取动态路径参数函数

func articleYearMonth(ctx *gin.Context) {
	//获取动态路径参数
	year := ctx.Param("year")
	month := ctx.Param("month")

	//响应
	ctx.JSON(http.StatusOK, gin.H{
		"year":  year,
		"month": month,
	})
}

func articleDelete(ctx *gin.Context) {
	//获取动态参数
	deleteID := ctx.Param("delete_id")
	fmt.Println("删除书籍: ", deleteID)
	//响应重定向
	ctx.Redirect(http.StatusMovedPermanently, "/index")
}

func main() {

	//获取路由对象
	r := gin.Default()
	//加载模板文件
	r.LoadHTMLGlob("templates/*")
	//开放静态资源窗口
	r.Static("/static", "./static")

	//路由
	//首页
	r.GET("/", getIndex)
	//登录
	r.GET("/login", getLogin)
	r.POST("/login", postLogin)

	//设置博客路由组
	blogRoute := r.Group("/blog")
	{
		blogRoute.GET("/articles/:year/:month", articleYearMonth)
		blogRoute.GET("/articles/delete/:delete_id", articleDelete)
	}

	//没有路由匹配
	r.NoRoute(notFund)

	//启动
	r.Run()

}
