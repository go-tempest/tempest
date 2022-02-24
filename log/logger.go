package log

import (
	"github.com/go-tempest/tempest/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

//日志配置
//var filename = "demos/log/test.log" //日志文件的位置
//var maxSize = 1                     //在进行切割之前，日志文件的最大大小（以MB为单位）
//var maxBackups = 5                  //保留旧文件的最大个数
//var maxAge = 30                     //保留旧文件的最大天数
//var compress = true                 //是否压缩/归档旧文件

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger = zap.New(core).Sugar()
}

func GetLogger() *zap.SugaredLogger {
	return logger
}

func Flush() {
	logger.Sync()
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.TempestCfg.Logger.Filename,   //日志文件的位置
		MaxSize:    config.TempestCfg.Logger.MaxSize,    //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: config.TempestCfg.Logger.MaxBackups, //保留旧文件的最大个数
		MaxAge:     config.TempestCfg.Logger.MaxAge,     //保留旧文件的最大天数
		Compress:   config.TempestCfg.Logger.Compress,   //是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}
