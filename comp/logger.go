package comp

import (
    "fmt"
    "github.com/go-tempest/tempest/conf"
    "github.com/go-tempest/tempest/core"
    "github.com/go-tempest/tempest/log"
    "github.com/spf13/viper"
    "os"
)

type LoggerStarter struct {
}

func (ls *LoggerStarter) Start(ctx *core.Context) {

    b := parseBootstrapYAML()
    logger := parseLoggerYAML(b)

    lt := log.LoggerType(logger.Type)
    ll := conf.GetLoggerLevel(logger.Level)
    e := conf.GetEnv(b.Active)

    filename := logger.File.Filename
    maxSize := logger.File.MaxSize
    maxBackups := logger.File.MaxBackups
    maxAge := logger.File.MaxAge
    compress := logger.File.Compress
    lic := logger.Console.LogInConsole

    ctx.Logger = log.Wrap(lt, e, ll, filename, maxSize, maxBackups, maxAge, compress, lic)
    ctx.BootstrapConfig = b
}

func parseLoggerYAML(b *conf.Bootstrap) *conf.LoggerConfig {

    var logger conf.LoggerConfig
    viper.SetConfigName(fmt.Sprintf(conf.DefaultLoggerConfigName, conf.GetEnv(b.Active)))

    err := viper.ReadInConfig()
    if err != nil {
        fmt.Printf("Viper initialization failed, error is [%v]\n", err)
        os.Exit(-1)
    }

    v := viper.GetViper()
    if err := v.Unmarshal(&logger); err != nil {
        fmt.Printf("Deserialization configuration failed, error is [%v]\n", err)
        os.Exit(-1)
    }

    return &logger
}
