package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type stuStruct struct {
	Name string
	Age  int
}

func main() {
	//获取路由对象
	r := gin.Default()

	//加载模板文件
	r.LoadHTMLGlob("templates/*")

	//get路由
	r.GET("/index", func(ctx *gin.Context) {

		//响应
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"user":      "haha",
			"state":     "在线",
			"bookSlice": []string{"三国演义", "红楼梦", "水浒传", "西游记"},
			"stuMap": map[string]interface{}{
				"name":  "bin",
				"age":   22,
				"hobby": []string{"篮球", "足球", "大乐透"},
			},
			"stuStruct": stuStruct{Name: "han", Age: 21},
		})
	})

	//启动
	r.Run()
}
