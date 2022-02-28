package discovery

import "github.com/go-tempest/tempest/log"

// Client 服务发现客户端接口
type Client interface {

    // Register 服务注册
    Register(logger log.FlagLogger, serviceName, instanceId, instanceHost string, instancePort int,
        healthCheckUrl, checkInterval, deregisterAfter string, meta map[string]string, tags ...string) bool

    // Deregister 服务注销
    Deregister(logger log.FlagLogger, instanceId string) bool

    // DiscoverServices 服务发现
    DiscoverServices(logger log.FlagLogger, serviceName, tag string) []interface{}
}
