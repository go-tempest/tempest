package starter

import (
    "github.com/go-tempest/tempest/boostrap/context"
    "github.com/go-tempest/tempest/register"
)

type RegistrationStarter struct {
}

func (cs *RegistrationStarter) Start(ctx *context.BootstrapContext) {
    new(register.Register).StartIfNecessary(ctx)
}
