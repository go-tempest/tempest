package log

import (
    "github.com/go-tempest/tempest/env"
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

type Logger interface {
    BaseLogger
    FlagLogger
    FlushLogger
}

type FlushLogger interface {
    Flush() error
}

type FlagLogger interface {
    With(args ...interface{}) BaseLogger
}

type BaseLogger interface {
    Debug(args ...interface{})
    Debugf(template string, args ...interface{})

    Info(args ...interface{})
    Infof(template string, args ...interface{})

    Warn(args ...interface{})
    Warnf(template string, args ...interface{})

    Error(args ...interface{})
    Errorf(template string, args ...interface{})
}

func Create(lt LoggerType, env env.Env, level LoggerLevel,
    filename string, maxSize, maxBackups, maxAge int,
    compress, logInConsole bool) Logger {

    switch lt {
    case Zap:
        fallthrough
    default:
        zl := new(ZapLogger)
        zl.create(env, level, filename, maxSize,
            maxBackups, maxAge, compress, logInConsole)
        return zl
    }
}