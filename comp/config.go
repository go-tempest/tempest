package comp

import (
    "fmt"
    "github.com/go-tempest/tempest/conf"
    "github.com/go-tempest/tempest/core"
    "github.com/spf13/viper"
    "os"
)

type ConfigStarter struct {
}

func (cs *ConfigStarter) Start(ctx *core.Context) {

    app := parseAppYAML(ctx.BootstrapConfig)
    r := parseRegistrationYAML(ctx.BootstrapConfig)

    ctx.AppConfig = &app.Application
    ctx.RegistrationConfig = &r.Registration
}

func parseRegistrationYAML(b *conf.Bootstrap) *conf.RegistrationConfig {

    var registration conf.RegistrationConfig
    viper.SetConfigName(fmt.Sprintf(conf.DefaultAppConfigName, conf.GetEnv(b.Active)))

    err := viper.ReadInConfig()
    if err != nil {
        fmt.Printf("Viper initialization failed, error is [%v]\n", err)
        os.Exit(-1)
    }

    v := viper.GetViper()
    if err := v.Unmarshal(&registration); err != nil {
        fmt.Printf("Deserialization configuration failed, error is [%v]\n", err)
        os.Exit(-1)
    }

    return &registration
}

func parseAppYAML(b *conf.Bootstrap) *conf.AppConfig {

    var app conf.AppConfig
    viper.SetConfigName(fmt.Sprintf(conf.DefaultAppConfigName, conf.GetEnv(b.Active)))

    err := viper.ReadInConfig()
    if err != nil {
        fmt.Printf("Viper initialization failed, error is [%v]\n", err)
        os.Exit(-1)
    }

    v := viper.GetViper()
    if err := v.Unmarshal(&app); err != nil {
        fmt.Printf("Deserialization configuration failed, error is [%v]\n", err)
        os.Exit(-1)
    }

    return &app
}
