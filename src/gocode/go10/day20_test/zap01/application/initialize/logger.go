package initialize

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
	"zap01/application/config"
)

//编码器函数

func getEncoder() zapcore.Encoder {
	//创建encoderConfig对象
	encoderConfig := zap.NewProductionEncoderConfig()
	//timeKey = time
	encoderConfig.TimeKey = "time"
	//日志级别大写
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//日志记录时间格式化
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		//time.DateTime常量在go version 1.20中提供
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	//返回json的encoderConfig对象
	return zapcore.NewJSONEncoder(encoderConfig)
}

//日志输出函数

func getWriteSyncer() zapcore.WriteSyncer {
	//获取路径分隔符
	stSeparator := string(filepath.Separator)
	fmt.Println("路径分隔符(stSeparator): ", stSeparator)
	//获取工作目录
	stRootDir, _ := os.Getwd()
	fmt.Println("工作目录(stRootDir): ", stRootDir)
	//日志文件存储路径 = 工作目录+路径分隔符+目录名+路径分隔符+当天时间(time.DateOnly 常量在go version 1.20中提供)+文件扩展名
	stLogFilePath := stRootDir + stSeparator + "log" + stSeparator + time.Now().Format(time.DateOnly) + ".log"
	fmt.Println("日志文件名(stLogFilePath): ", stLogFilePath)

	//日志分割
	lumberJackSyncer := &lumberjack.Logger{
		//文件名称
		Filename: stLogFilePath,
		//文件保留大小
		MaxSize: config.Conf.MaxSize,
		//文件保留天数
		MaxAge: config.Conf.MaxAge,
		//文件保留个数
		MaxBackups: config.Conf.MaxBackups,
		//是否启用gzip压缩
		Compress: config.Conf.Compress,
	}

	//返回WriteSyncer

	return zapcore.AddSync(lumberJackSyncer)
}

//全局Logger对象

var Logger *zap.Logger

//日志初始化函数

func InitLogger(cfg *config.Config) (err error) {
	//获取编码器对象
	encoder := getEncoder()

	//定制日志格式
	writeSync := getWriteSyncer()

	//定义  zapcore 日志级别
	var levelLog = new(zapcore.Level)

	if err = levelLog.UnmarshalText([]byte(cfg.Level)); err != nil {
		return
	}

	//定义core对象   NewMultiWriteSyncer(wirteSync,zapcore.AddSync(os.Stdout)) 多路输出,分别输出到日志文件和控制台
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSync, zapcore.AddSync(os.Stdout)), levelLog)

	//重写全局Logger对象
	Logger = zap.New(core, zap.AddCaller())

	//替换全局Logger
	zap.ReplaceGlobals(Logger)

	return
}
