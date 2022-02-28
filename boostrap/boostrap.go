package boostrap

import (
    "github.com/go-tempest/tempest/comp"
    "github.com/go-tempest/tempest/core"
    "sync"
)

type ServerBootstrap struct {
    sync.Once
    comps []comp.Starter
    ctx   *core.Context
}

func (b *ServerBootstrap) initialize() {
    b.ctx = new(core.Context)
    b.comps = []comp.Starter{
        &comp.LoggerStarter{},
        &comp.ConfigStarter{},
        &comp.RegistrationStarter{},
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
