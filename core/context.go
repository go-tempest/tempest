package core

import (
    "github.com/go-tempest/tempest/config"
    "github.com/go-tempest/tempest/log"
)

type TempestContext struct {
    Logger             log.Logger
    BootstrapConfig    *config.Bootstrap
    AppConfig          *config.Application
    RegistrationConfig *config.Registration
}
