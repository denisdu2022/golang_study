package main

import (
	"github.com/gin-gonic/gin"
	"zap01/config"
)

func main() {

	//获取路由对象
	router := gin.Default()

	//初始化日志
	loger := config.InitLogger()

	loger.Info("服务已启动")
	//启动 端口8081
	router.Run(":8081")

}
