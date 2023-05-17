package initialize

import (
	"bingotest01/application/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 编码器函数
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 获取日志写函数
func getLogWrite(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}

	return zapcore.AddSync(lumberJackLogger)
}

//定义全局logger对象

var Logger *zap.Logger

func InitLogger(cfg *config.LogConfig) (err error) {

	//日志初始化

	//定制日志的格式
	writeSyncer := getLogWrite(config.Conf.Filename, config.Conf.MaxSize, config.Conf.MaxBackups, config.Conf.MaxAge)

	//定义编码器
	encoder := getEncoder()

	//定义日志级别
	var levelLog = new(zapcore.Level)

	if err = levelLog.UnmarshalText([]byte(config.Conf.Level)); err != nil {
		return
	}

	//定义core对象
	core := zapcore.NewCore(encoder, writeSyncer, levelLog)

	//对全局Logger对象进行重写赋值
	Logger = zap.New(core, zap.AddCaller())

	//用于项目生产阶段,格式json:适合用于集成到第三方日志分析系统中
	//logger,_ := zap.NewProduction()

	////用于项目开发阶段,格式:普通文本格式:适合在中端查看
	//logger, _ := zap.NewDevelopment()

	//把logger替换成zap的全局日志配置
	zap.ReplaceGlobals(Logger)

	return
}
