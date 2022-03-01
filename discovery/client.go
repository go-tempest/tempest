package discovery

import "github.com/go-tempest/tempest/log"

type Client interface {
    Register(logger log.FlagLogger, serviceName, instanceId, instanceHost string, instancePort int,
        healthCheckUrl, checkInterval, deregisterAfter string, meta map[string]string, tags ...string) bool

    Deregister(logger log.FlagLogger, instanceId string) bool

    DiscoverServices(logger log.FlagLogger, serviceName, tag string) []interface{}
}
