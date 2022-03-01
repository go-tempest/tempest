package boostrap

import (
    "github.com/go-tempest/tempest/core"
    "github.com/go-tempest/tempest/starter"
    "sync"
)

type ServerBootstrap struct {
    sync.Once
    comps []starter.Starter
    ctx   *core.BootstrapContext
}

func (b *ServerBootstrap) initialize() {
    b.ctx = new(core.BootstrapContext)
    b.comps = []starter.Starter{
        &starter.LoggerStarter{},
        &starter.ConfigStarter{},
        &starter.RegistrationStarter{},
    }
}

func (b *ServerBootstrap) start() {
    b.Do(func() {
        for _, c := range b.comps {
            c.Start(b.ctx)
        }
    })
}

func init() {
    b := new(ServerBootstrap)
    b.initialize()
    b.start()
}
