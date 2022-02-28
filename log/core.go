package log

import (
    "github.com/go-tempest/tempest/env"
    "go.uber.org/zap"
)

type LoggerLevel string

const (
    Debug LoggerLevel = "debug"
    Info              = "info"
    Error             = "error"
    Panic             = "panic"
    Fatal             = "fatal"
)

type (
    LoggerType string
)

const (
    Zap LoggerType = "zap"
)

type FlagLogger func(...interface{}) Logger

type LoggerWrapper struct {
    flags  []interface{}
    logger Logger
    lt     LoggerType
}

func (lw *LoggerWrapper) bindFlag() {

    if len(lw.flags) == 0 {
        return
    }

    switch lw.lt {
    case Zap:
        fallthrough
    default:
        l, ok := lw.logger.(*zap.SugaredLogger)
        if ok {
            l.With(lw.flags)
        }
    }
}

func (lw *LoggerWrapper) Log(args ...interface{}) error {
    lw.bindFlag()
    lw.logger.Error(args)
    return nil
}

func (lw *LoggerWrapper) Debug(args ...interface{}) {
    lw.bindFlag()
    lw.logger.Debug(args)
}

func (lw *LoggerWrapper) Debugf(template string, args ...interface{}) {
    lw.bindFlag()
    lw.logger.Debugf(template, args)
}

func (lw *LoggerWrapper) Info(args ...interface{}) {
    lw.bindFlag()
    lw.logger.Info(args)
}

func (lw *LoggerWrapper) Infof(template string, args ...interface{}) {
    lw.bindFlag()
    lw.logger.Infof(template, args)
}

func (lw *LoggerWrapper) Warn(args ...interface{}) {
    lw.bindFlag()
    lw.logger.Warn(args)
}

func (lw *LoggerWrapper) Warnf(template string, args ...interface{}) {
    lw.bindFlag()
    lw.logger.Warn(template, args)
}

func (lw *LoggerWrapper) Error(args ...interface{}) {
    lw.bindFlag()
    lw.logger.Error(args)
}

func (lw *LoggerWrapper) Errorf(template string, args ...interface{}) {
    lw.bindFlag()
    lw.logger.Error(template, args)
}

type Logger interface {
    Debug(args ...interface{})
    Debugf(template string, args ...interface{})

    Info(args ...interface{})
    Infof(template string, args ...interface{})

    Warn(args ...interface{})
    Warnf(template string, args ...interface{})

    Error(args ...interface{})
    Errorf(template string, args ...interface{})
}

func Wrap(lt LoggerType, env env.Env, level LoggerLevel,
    filename string, maxSize, maxBackups, maxAge int,
    compress, logInConsole bool) FlagLogger {

    switch lt {
    case Zap:
        fallthrough
    default:

        var zl ZapLogger
        logger := zl.createZap(env, level, filename, maxSize,
            maxBackups, maxAge, compress, logInConsole)

        EmptyFlagLogger = &LoggerWrapper{
            logger: logger,
            lt:     Zap,
        }

        return func(args ...interface{}) Logger {
            return &LoggerWrapper{
                logger: logger,
                lt:     Zap,
                flags:  args,
            }
        }
    }
}

var EmptyFlagLogger *LoggerWrapper
