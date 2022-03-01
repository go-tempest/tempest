package discovery

type Client interface {
    Register(serviceName, instanceId, instanceHost string, instancePort int,
        healthCheckUrl, checkInterval, deregisterAfter string, meta map[string]string, tags ...string) bool

    Deregister(instanceId string) bool

    DiscoverServices(serviceName, tag string) []interface{}
}
