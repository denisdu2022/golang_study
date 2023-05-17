package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"path/filepath"
	"zap01/application/config"
	"zap01/application/initialize"
	"zap01/application/middleware"
)

func main() {

	//获取工作目录
	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		panic(err.Error())
	}

	//配置文件初始化
	if err := config.Init(fmt.Sprintf("%s/config.json", dir)); err != nil {
		panic(err.Error())
	}

	//设置gin的调试模式
	gin.SetMode(config.Conf.Mode)

	//初始化日志配置
	if err := initialize.InitLogger(config.Conf); err != nil {
		fmt.Printf("初始化日志配置失败,错误信息是:%v\n", err)
		return
	}

	//测试:
	zap.S().Debugf("模式:%s", config.Conf.Mode)
	zap.S().Debugf("服务启动端口:%d", config.Conf.Port)

	//获取路由对象
	router := gin.Default()

	//加载日志中间件
	router.Use(middleware.RequestLogger())

	//路由测试
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "200",
		})
	})

	//启动
	if err := router.Run(fmt.Sprintf(":%d", config.Conf.Port)); err != nil {
		fmt.Println("服务启动失败")
	}
}
