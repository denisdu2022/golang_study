package config

import (
	"errors"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

//日志初始化函数  返回类型*zap.SugaredLogger

func InitLogger() *zap.SugaredLogger {

	//日志级别 DebugLevel
	logMode := zapcore.DebugLevel

	//创建core对象  zapcore.NewCore(参数1:定义日志格式,参数2:日志写,参数3:日志的级别)
	core := zapcore.NewCore(getEncoder(), getWriteSyncer(), logMode)

	//测试
	test := zap.New(core).Sugar()
	test.Infof("Error: %s", errors.New("test......输出日志....."))

	//返回core
	return zap.New(core).Sugar()
}

// 日志格式函数 返回zapcore.Encoder 日志格式
func getEncoder() zapcore.Encoder {
	//获取encoderConfig对象
	encoderConfig := zap.NewProductionEncoderConfig()
	//TimeKey = time
	encoderConfig.TimeKey = "time"
	//日志级别大写 使用zapcore.CapitalLevelEncoder 内置方法把日志级别转为大写
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//日志记录时间格式化
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		//使用go语言提供的时间格式化方式  "2006-01-02 15:04:05"
		encoder.AppendString(t.Local().Format("2006-01-02 15:04:05"))
	}

	//返回 JSON 的 encoderConfig 对象
	return zapcore.NewJSONEncoder(encoderConfig)
}

//日志写函数getWriteSyncer() 返回类型  zapcore.WriteSyncer

func getWriteSyncer() zapcore.WriteSyncer {

	//-------------日志文件存储------------------
	//路径分隔符
	stSeparator := string(filepath.Separator)
	fmt.Println("路径分隔符: ", stSeparator)
	//当前工作目录
	stRootDir, _ := os.Getwd()
	fmt.Println("当前工作目录: ", stRootDir)
	//日志文件存储路径 当前目录+路径分隔符+目录名+当前年-月-日+文件名
	stLogFilePath := stRootDir + stSeparator + "log" + stSeparator + time.Now().Format("2006-01-02") + ".log"
	fmt.Println("日志文件存储路径: ", stLogFilePath)

	//-------------日志切割------------------
	lumberjackSyncer := &lumberjack.Logger{
		//文件名称
		Filename: stLogFilePath,
		//文件保留大小 单位(MB)
		MaxSize: 1,
		//文件保留个数
		MaxBackups: 10,
		//文件保留天数 单位(day)
		MaxAge: 7,
		//是否启用gzip压缩
		Compress: false,
	}

	//返回日志输出的对象
	return zapcore.AddSync(lumberjackSyncer)
}
