package component

import (
    "github.com/go-tempest/tempest/boostrap/context"
    "github.com/go-tempest/tempest/register"
)

type RegisterComponent struct {
}

func (cs *RegisterComponent) Execute(ctx *context.BootstrapContext) {
    new(register.Register).StartIfNecessary(ctx)
}
