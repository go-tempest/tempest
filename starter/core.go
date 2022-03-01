package starter

import (
    "fmt"
    "github.com/go-tempest/tempest/config"
    "github.com/go-tempest/tempest/core"
    "github.com/spf13/viper"
    "os"
)

type Starter interface {
    Start(ctx *core.BootstrapContext)
}

func parseBootstrapYAML() *config.Bootstrap {

    var b config.Bootstrap

    viper.SetConfigType(config.DefaultConfigType)
    viper.AddConfigPath(*config.GetDefaultConfigPath())
    viper.AddConfigPath(*config.GetFlagConfigPath())
    viper.SetConfigName(config.DefaultBootstrapConfigName)

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
