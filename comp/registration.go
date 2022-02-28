package comp

import (
    "github.com/go-tempest/tempest/core"
    "github.com/go-tempest/tempest/register"
)

type RegistrationStarter struct {
}

func (cs *RegistrationStarter) Start(ctx *core.Context) {
    new(register.Register).StartIfNecessary(ctx)
}
