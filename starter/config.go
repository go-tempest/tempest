package starter

import (
    "fmt"
    "github.com/go-tempest/tempest/config"
    "github.com/go-tempest/tempest/core"
    "github.com/spf13/viper"
    "os"
)

type ConfigStarter struct {
}

func (cs *ConfigStarter) Start(ctx *core.TempestContext) {

    app := parseAppYAML(ctx.BootstrapConfig)
    r := parseRegistrationYAML(ctx.BootstrapConfig)

    ctx.AppConfig = &app.Application
    ctx.RegistrationConfig = &r.Registration
}

func parseRegistrationYAML(b *config.Bootstrap) *config.RegistrationConfig {

    var registration config.RegistrationConfig
    viper.SetConfigName(fmt.Sprintf(config.DefaultAppConfigName, config.GetEnv(b.Active)))

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

func parseAppYAML(b *config.Bootstrap) *config.AppConfig {

    var app config.AppConfig
    viper.SetConfigName(fmt.Sprintf(config.DefaultAppConfigName, config.GetEnv(b.Active)))

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
