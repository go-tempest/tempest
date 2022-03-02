package discovery

import "github.com/go-tempest/tempest/log"

type Client interface {
    Register(logger log.Logger, serviceName, instanceId, instanceHost string, instancePort int,
        healthCheckUrl, checkInterval, deregisterAfter string, meta map[string]string, tags ...string) bool

    Deregister(logger log.Logger, instanceId string) bool

    DiscoverServices(logger log.Logger, serviceName, tag string) []interface{}
}
