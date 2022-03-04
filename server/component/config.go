package component

import (
    "fmt"
    "github.com/go-tempest/tempest/boostrap/context"
    "github.com/go-tempest/tempest/config"
    "github.com/spf13/viper"
    "os"
)

type ConfigServerComponent struct {
    bc BootstrapComponent
}

func (cs *ConfigServerComponent) Execute(ctx *context.BootstrapContext) {

    if ctx.BootstrapConfig == nil {
        ctx.BootstrapConfig = cs.bc.parseYAML()
    }

    app, r := cs.parseYAML(ctx.BootstrapConfig)

    ctx.AppConfig = &app.Application
    ctx.RegistrationConfig = &r.Registration
}

func (cs *ConfigServerComponent) parseYAML(b *config.Bootstrap) (*config.AppConfig, *config.RegistrationConfig) {

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

    var registration config.RegistrationConfig
    viper.SetConfigName(fmt.Sprintf(config.DefaultAppConfigName, config.GetEnv(b.Active)))

    err = viper.ReadInConfig()
    if err != nil {
        fmt.Printf("Viper initialization failed, error is [%v]\n", err)
        os.Exit(1)
    }

    if err := v.Unmarshal(&registration); err != nil {
        fmt.Printf("Deserialization configuration failed, error is [%v]\n", err)
        os.Exit(1)
    }

    return &app, &registration
}
