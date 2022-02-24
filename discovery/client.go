package discovery

import "log"

// Client 服务发现客户端接口
type Client interface {

    // Register 服务注册
    Register(serviceName, instanceId, healthCheckUrl string,
        logger *log.Logger,
        instanceHost string, instancePort int,
        meta map[string]string, tags ...string) bool

    // Deregister 服务注销
    Deregister(instanceId string, logger *log.Logger) bool

    // DiscoverServices 服务发现
    DiscoverServices(serviceName, tag string, logger *log.Logger) []interface{}
}
