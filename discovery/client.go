package discovery

// Client 服务发现客户端接口
type Client interface {

    // Register 服务注册
    Register(serviceName, instanceId, instanceHost string,
        instancePort int, healthCheckUrl string,
        meta map[string]string, tags ...string) bool

    // Deregister 服务注销
    Deregister(instanceId string) bool

    // DiscoverServices 服务发现
    DiscoverServices(serviceName, tag string) []interface{}
}
