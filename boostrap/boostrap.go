package boostrap

import (
    "github.com/go-tempest/tempest/core"
    "github.com/go-tempest/tempest/starter"
    "sync"
)

type ServerBootstrap struct {
    sync.Once
    starters []starter.Starter
    ctx      *core.TempestContext
}

func (b *ServerBootstrap) initialize() {
    b.ctx = new(core.TempestContext)
    b.starters = []starter.Starter{
        &starter.LoggerStarter{},
        &starter.ConfigStarter{},
        &starter.RegistrationStarter{},
    }
}

func (b *ServerBootstrap) start() {
    b.Do(func() {
        for _, c := range b.starters {
            c.Start(b.ctx)
        }
    })
}

func init() {
    b := new(ServerBootstrap)
    b.initialize()
    b.start()
    ctx = b.ctx
}

var ctx *core.TempestContext

func GetContext() *core.TempestContext {
    return ctx
}
