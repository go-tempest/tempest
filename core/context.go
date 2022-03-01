package core

import (
    "github.com/go-tempest/tempest/config"
    "github.com/go-tempest/tempest/log"
)

type BootstrapContext struct {
    Logger             log.FlagLogger
    BootstrapConfig    *config.Bootstrap
    AppConfig          *config.Application
    RegistrationConfig *config.Registration
}
