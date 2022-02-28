package comp

import (
    "fmt"
    "github.com/go-tempest/tempest/conf"
    "github.com/go-tempest/tempest/core"
    "github.com/spf13/viper"
    "os"
)

type Starter interface {
    Start(ctx *core.Context)
}

func parseBootstrapYAML() *conf.Bootstrap {

    var b conf.Bootstrap

    viper.SetConfigType(conf.DefaultConfigType)
    viper.AddConfigPath(*conf.GetDefaultConfigPath())
    viper.AddConfigPath(*conf.GetFlagConfigPath())
    viper.SetConfigName(conf.DefaultBootstrapConfigName)

    err := viper.ReadInConfig()
    if err != nil {
        fmt.Printf("Viper initialization failed, error is [%v]\n", err)
        os.Exit(-1)
    }

    v := viper.GetViper()
    if err := v.Unmarshal(&b); err != nil {
        fmt.Printf("Deserialization configuration failed, error is [%v]\n", err)
        os.Exit(-1)
    }

    return &b
}
