package component

import (
    "fmt"
    "github.com/go-tempest/tempest/boostrap/context"
    "github.com/go-tempest/tempest/config"
    "github.com/go-tempest/tempest/log"
    "github.com/spf13/viper"
    "os"
)

type LoggerServerComponent struct {
    bc BootstrapComponent
}

func (ls *LoggerServerComponent) Execute(ctx *context.BootstrapContext) {

    b := ls.bc.parseYAML()
    logger := ls.parseYAML(b)

    lt := log.LoggerType(logger.Type)
    ll := config.GetLoggerLevel(logger.Level)
    e := config.GetEnv(b.Active)

    filename := logger.File.Filename
    maxSize := logger.File.MaxSize
    maxBackups := logger.File.MaxBackups
    maxAge := logger.File.MaxAge
    compress := logger.File.Compress
    lic := logger.Console.LogInConsole

    ctx.BootstrapConfig = b
    ctx.Logger = log.Create(lt, e, ll, filename, maxSize, maxBackups, maxAge, compress, lic)
}

func (ls *LoggerServerComponent) parseYAML(b *config.Bootstrap) *config.LoggerConfig {

    var logger config.LoggerConfig
    viper.SetConfigName(fmt.Sprintf(config.DefaultLoggerConfigName, config.GetEnv(b.Active)))

    err := viper.ReadInConfig()
    if err != nil {
        fmt.Printf("Viper initialization failed, error is [%v]\n", err)
        os.Exit(1)
    }

    v := viper.GetViper()
    if err := v.Unmarshal(&logger); err != nil {
        fmt.Printf("Deserialization configuration failed, error is [%v]\n", err)
        os.Exit(1)
    }

    return &logger
}
