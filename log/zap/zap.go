package zap

import (
    "github.com/go-tempest/tempest/config"
    "github.com/go-tempest/tempest/env"
    "github.com/go-tempest/tempest/log"
    "github.com/natefinch/lumberjack"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func Create() {
    enc := getEncoder(config.TempestConfig.Application.Profiles.Active)
    ws := getLogWriter()
    core := zapcore.NewCore(enc, ws, getLevel(config.TempestConfig.Logger.Level))
    Logger = zap.New(core).Sugar()
}

func getLevel(l log.Level) (zapLevel zapcore.Level) {
    switch l {
    case log.Debug:
        zapLevel = zapcore.DebugLevel
    case log.Info:
        zapLevel = zapcore.InfoLevel
    case log.Error:
        zapLevel = zapcore.ErrorLevel
    case log.Panic:
        zapLevel = zapcore.PanicLevel
    case log.Fatal:
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
    return zapcore.AddSync(lumberJackLogger)
}
