package main

import (
	"bingo_api/applcation/config"
	"bingo_api/applcation/initialize"
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func main() {

	//1.获取一个基于整个目录入口所在的路径  filepath.Abs文件的绝对路径(filepath.Dir文件目录(".") .是当前文件)
	dir, err := filepath.Abs(filepath.Dir("."))
	//判断如果遇到错误就抛出
	if err != nil {
		panic(err.Error())
	}

	//2.配置初始化
	//fmt.Sprintf() 把打印的内容作为一个值返回
	if err := config.Init(fmt.Sprintf("%s/config.json", dir)); err != nil {
		panic(err.Error())
	}

	//设置调试模式
	gin.SetMode(config.Conf.Mode)

	//3.日志初始化
	initialize.InitLogger()

	//4.初始化路由
	Router := initialize.InitRouter()

	//5.启动,监听8080端口
	Router.Run(fmt.Sprintf(":%d", config.Conf.Port))

}
