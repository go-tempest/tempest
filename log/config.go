package log

import (
    "github.com/go-tempest/tempest/config"
)

var Logger Log

type (
    Type  string
    Level string
)

const (
    Zap Type = "zap"
)

const (
    Debug Level = "debug"
    Info        = "info"
    Error       = "error"
    Panic       = "panic"
    Fatal       = "fatal"
)

func Initialize() {
    t := config.TempestConfig.Logger.Type
    switch Type(t) {
    case Zap:
        Logger = Create()
    default:
        Logger = Create()
    }
}
