package boostrap

import (
    "github.com/go-tempest/tempest/boostrap/context"
    "github.com/go-tempest/tempest/starter"
    "os"
    "sync"
)

type ServerBootstrapHook func(*context.BootstrapContext)
type ServerBootstrapLifecycle int

const (
    Pre ServerBootstrapLifecycle = iota
    Post
)

type ServerBootstrap struct {
    sync.Once
    ctx      *context.BootstrapContext
    starters []starter.Starter
    hooks    map[ServerBootstrapLifecycle]ServerBootstrapHook
}

func (b *ServerBootstrap) initialize() {
    b.ctx = new(context.BootstrapContext)
    b.hooks = make(map[ServerBootstrapLifecycle]ServerBootstrapHook)
    b.starters = []starter.Starter{
        &starter.LoggerStarter{},
        &starter.ConfigStarter{},
        &starter.RegistrationStarter{},
    }
}

func (b *ServerBootstrap) ResigerHook(lifecycle ServerBootstrapLifecycle, hook ServerBootstrapHook) *ServerBootstrap {
    return ResigerHook(lifecycle, hook)
}

func (b *ServerBootstrap) Start() {
    if b == nil {
        os.Exit(1)
    }

    b.Do(func() {
        defer func() {
            _ = b.ctx.Logger.Flush()
        }()

        triggerHook(b, Pre)
        for _, c := range b.starters {
            c.Start(b.ctx)
        }
        triggerHook(b, Post)
    })
}

func ResigerHook(lifecycle ServerBootstrapLifecycle, hook ServerBootstrapHook) *ServerBootstrap {

    b := new(ServerBootstrap)
    b.initialize()

    if hook != nil {
        b.hooks[lifecycle] = hook
    }

    return b
}

func triggerHook(b *ServerBootstrap, lifecycle ServerBootstrapLifecycle) {

    if b.hooks == nil || len(b.hooks) == 0 {
        return
    }

    h := b.hooks[lifecycle]
    if h != nil {
        h(b.ctx)
    }
}
