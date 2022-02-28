package log

import (
    "github.com/go-tempest/tempest/env"
    "github.com/natefinch/lumberjack"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
)

type ZapLogger struct {
}

func (ZapLogger) createZap(env env.Env, level LoggerLevel, filename string,
    maxSize, maxBackups, maxAge int, compress, logInConsole bool) *zap.SugaredLogger {

    enc := getEncoder(env)
    ws := getLogWriter(filename, maxSize, maxBackups, maxAge, compress, logInConsole)
    core := zapcore.NewCore(enc, ws, getLevel(level))

    return zap.New(core).Sugar()
}

func getLevel(ll LoggerLevel) (zapLevel zapcore.Level) {
    switch ll {
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

func getLogWriter(filename string, maxSize, maxBackups, maxAge int,
    compress, logInConsole bool) zapcore.WriteSyncer {

    lumberJackLogger := &lumberjack.Logger{
        Filename:   filename,
        MaxSize:    maxSize,
        MaxBackups: maxBackups,
        MaxAge:     maxAge,
        Compress:   compress,
    }

    var syncer zapcore.WriteSyncer
    if logInConsole {
        syncer = zapcore.NewMultiWriteSyncer(
            zapcore.AddSync(os.Stdout),
            zapcore.AddSync(lumberJackLogger),
        )
    } else {
        syncer = zapcore.AddSync(lumberJackLogger)
    }

    return syncer
}
