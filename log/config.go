package log

import (
    "github.com/go-tempest/tempest/config"
    "github.com/go-tempest/tempest/log/zap"
)

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
    switch t {
    case Zap:
        zap.Create()
    default:
        zap.Create()
    }
}
