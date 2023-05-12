package initialize

import "go.uber.org/zap"

//日志初始化

func InitLogger() {

	//用于项目生产阶段,格式json:适合用于集成到第三方日志分析系统中
	//logger,_ := zap.NewProduction()

	//用于项目开发阶段,格式:普通文本格式:适合在中端查看
	logger, _ := zap.NewDevelopment()

	//把logger替换成zap的全局日志配置
	zap.ReplaceGlobals(logger)
}
