package conf

import "github.com/go-tempest/tempest/log"

type LoggerConfig struct {
    Logger
}

type Logger struct {
    Type    string
    Level   log.LoggerLevel
    File    LoggerFile
    Console LoggerConsole
}

type LoggerConsole struct {
    LogInConsole bool `mapstructure:"log-in-console"`
}

type LoggerFile struct {
    LogInFile  bool   `mapstructure:"log-in-file"`
    Filename   string `mapstructure:"filename"`
    MaxSize    int    `mapstructure:"max-size"`
    MaxBackups int    `mapstructure:"max-backups"`
    MaxAge     int    `mapstructure:"max-age"`
    Compress   bool
}
