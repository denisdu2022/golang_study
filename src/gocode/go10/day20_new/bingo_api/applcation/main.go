package main

import (
	"bingo_api/applcation/config"
	"bingo_api/applcation/database"
	"bingo_api/applcation/initialize"
	"bingo_api/applcation/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	//initialize.InitLogger()
	if err := initialize.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("初始化日志失败,错误信息是: %v\n", err)
		return
	}

	//调试信息
	zap.S().Debugf("调试信息:%d", config.Conf.Port)

	//4.初始化路由
	Router := initialize.InitRouter()

	//zap提供了一个S函数和一个L函数给开发者使用,调用S函数或者L函数,可以得到一个全局的线程安全的logger对象
	zap.S().Info("服务端启动,端口: ", config.Conf.Port)

	//数据库初始化
	Orm := database.InitDB(config.Conf.DatabaseConfig)

	//数据库迁移创建表
	Orm.AutoMigrate(&model.User{})

	//数据库禁用复数
	Orm.SingularTable(true)
	//延迟注册关闭
	defer Orm.Close()

	//5.启动,监听8080端口
	if err := Router.Run(fmt.Sprintf(":%d", config.Conf.Port)); err != nil {
		zap.S().Panic("服务端启动失败: ", err.Error())
	}

}
