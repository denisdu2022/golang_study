package main

import "github.com/gin-gonic/gin"

func main() {
	//获取路由对象
	r := gin.Default()

	//启动
	r.Run()
}
