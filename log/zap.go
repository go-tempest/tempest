package log

import (
	"github.com/go-tempest/tempest/config"
	"github.com/go-tempest/tempest/env"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type ZapLog struct {
	Logger *zap.SugaredLogger
}

func (zl *ZapLog) With(args ...interface{}) *zap.SugaredLogger {
	return zl.Logger.With(args)
}

func (zl *ZapLog) Debug(args ...interface{}) {
	zl.Logger.Debug(args)
}

func (zl *ZapLog) Info(args ...interface{}) {
	zl.Logger.Info(args)
}

func (zl *ZapLog) Warn(args ...interface{}) {
	zl.Logger.Warn(args)
}

func (zl *ZapLog) Error(args ...interface{}) {
	zl.Logger.Error(args)
}

func Create() Log {
	enc := getEncoder(config.TempestConfig.Application.Profiles.Active)
	ws := getLogWriter()
	core := zapcore.NewCore(enc, ws, getLevel(config.TempestConfig.Logger.Level))
	logger := zap.New(core).Sugar()
	return &ZapLog{
		Logger: logger,
	}
}

func getLevel(l string) (zapLevel zapcore.Level) {
	switch Level(l) {
	case Debug:
		zapLevel = zapcore.DebugLevel
	case Info:
		zapLevel = zapcore.InfoLevel
	case Error:
		zapLevel = zapcore.ErrorLevel
	case Panic:
		zapLevel = zapcore.PanicLevel
	case Fatal:
		zapLevel = zapcore.FatalLevel
	default:
		zapLevel = zapcore.InfoLevel
	}
	return
}

func getEncoder(e env.Env) zapcore.Encoder {

	var encConfig zapcore.EncoderConfig

	if env.Prod == e {
		encConfig = zap.NewProductionEncoderConfig()
	} else {
		encConfig = zap.NewDevelopmentEncoderConfig()
	}

	return zapcore.NewJSONEncoder(encConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.TempestConfig.Logger.Filename,
		MaxSize:    config.TempestConfig.Logger.MaxSize,
		MaxBackups: config.TempestConfig.Logger.MaxBackups,
		MaxAge:     config.TempestConfig.Logger.MaxAge,
		Compress:   config.TempestConfig.Logger.Compress,
	}
	var syncer zapcore.WriteSyncer
	if config.TempestConfig.Logger.LogInConsole {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	} else {
		syncer = zapcore.AddSync(lumberJackLogger)
	}
	return syncer
}
