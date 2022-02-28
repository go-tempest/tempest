package core

import (
    "github.com/go-tempest/tempest/conf"
    "github.com/go-tempest/tempest/log"
)

type Context struct {
    Logger log.FlagLogger

    BootstrapConfig    *conf.Bootstrap
    AppConfig          *conf.Application
    RegistrationConfig *conf.Registration
}
