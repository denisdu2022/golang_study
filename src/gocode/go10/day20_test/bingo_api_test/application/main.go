package main

import (
	"bingotest01/application/config"
	"bingotest01/application/database"
	"bingotest01/application/initialize"
	"bingotest01/application/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"path/filepath"
)

func main() {

	//获取一个基于整个目录入口所在的路径
	//filepath.Abs获取文件的绝对路径(filepath.Dir获取文件目录"." .是当前路径)
	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		panic(err.Error())
	}

	//配置文件初始化
	if err := config.Init(fmt.Sprintf("%s/config.json", dir)); err != nil {
		panic(err.Error())
	}

	//设置调试模式
	gin.SetMode(config.Conf.Mode)

	//初始化日志配置
	//initialize.InitLogger()
	if err := initialize.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("初始化日志失败,错误信息是:%v\n", err)
		return
	}

	//调试信息
	zap.S().Debugf("调试信息:%d", config.Conf.Port)

	//初始化路由
	Router := initialize.InitRouter()

	//zap提供了一个S函数和一个L函数给开发者使用,调用S函数或者L函数,可以得到一个全局的线程安全的logger对象
	zap.S().Info("服务已启动,调试模式是:", config.Conf.Mode)
	zap.S().Info("服务已启动,端口:", config.Conf.Port)

	//数据库初始化
	Orm := database.InitDB(config.Conf.DatabaseConfig)
	//数据库迁移
	Orm.AutoMigrate(&model.User{})
	//禁用复数
	Orm.SingularTable(true)
	defer Orm.Close()

	//启动
	if err := Router.Run(fmt.Sprintf(":%d", config.Conf.Port)); err != nil {
		zap.S().Panic("服务端启动失败: ", err.Error())
	}

}
