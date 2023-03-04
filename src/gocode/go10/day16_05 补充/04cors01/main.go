package main

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//定义全局的CORS中间件 里边return 匿名参数

func Cors() gin.HandlerFunc {
	////允许所有来源
	//return func(ctx *gin.Context) {
	//	ctx.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	//	ctx.Next()
	//}

	//高级请求配置
	return func(ctx *gin.Context) {
		//获取请求方式
		method := ctx.Request.Method
		//获取Origin请求头
		origin := ctx.GetHeader("origin")
		//注意这里,来源不能配置为通配符" * "号
		ctx.Header("Access-Control-Allow-Origin", origin)
		//这里必须设定为true
		ctx.Header("Access-Control-Allow-Credentials", "ture")
		//自定义的head字段都需要在这里声明
		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers,Cookie,Origin,X-Requested-With,Content-Type,Accept,Authorization,Token,Timestamp,Userid")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type,cache-control")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			//ctx.AbortWithStatus(http.StatusNoContent)
			ctx.AbortWithStatus(http.StatusOK)
		}
		// 处理请求
		ctx.Next()

	}

}

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

	//解决跨域问题
	//增加Access-Control-Allow-Origin 头,值写范围ip,这里使用*允许所有
	ctx.Writer.Header().Add("Access-Control-Allow-Origin", "*")

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
	//r.Static("/static", "./static")

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
