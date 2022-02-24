package discovery

import (
    "fmt"
    "github.com/go-kit/kit/sd/consul"
    "github.com/hashicorp/consul/api"
    "github.com/hashicorp/consul/api/watch"
    "log"
    "strconv"
    "sync"
)

const (
    httpProtocol  string = "http://"
    portSeparator string = ":"
)

type ConsulDiscoveryClientWrapper struct {
    Host          string
    Port          int
    client        consul.Client
    config        *api.Config
    lock          sync.Mutex
    instanceCache sync.Map
}

func New(registerHost string, registerPort int) (Client, error) {

    config := api.DefaultConfig()
    config.Address = fmt.Sprintf("%s%s%s", registerHost, portSeparator, strconv.Itoa(registerPort))

    cli, err := api.NewClient(config)
    if err != nil {
        return nil, err
    }

    return &ConsulDiscoveryClientWrapper{
        Host:   registerHost,
        Port:   registerPort,
        config: config,
        client: consul.NewClient(cli),
    }, err
}

func (wrapper *ConsulDiscoveryClientWrapper) Register(serviceName, instanceId, healthCheckUrl string,
    logger *log.Logger,
    instanceHost string, instancePort int,
    meta map[string]string, tags ...string) bool {

    registration := &api.AgentServiceRegistration{
        ID:      instanceId,
        Name:    serviceName,
        Address: instanceHost,
        Port:    instancePort,
        Meta:    meta,
        Tags:    tags,
        Check: &api.AgentServiceCheck{
            DeregisterCriticalServiceAfter: "30s",
            HTTP: fmt.Sprintf("%s%s%s%s%s", httpProtocol, instanceHost,
                portSeparator, strconv.Itoa(instancePort), healthCheckUrl),
            Interval: "15s",
        },
    }

    err := wrapper.client.Register(registration)
    if err != nil {
        logger.Println("Register Service Error!", err)
        return false
    }

    logger.Println("Register Service Success!")
    return true
}

func (wrapper *ConsulDiscoveryClientWrapper) Deregister(instanceId string, logger *log.Logger) bool {

    registration := &api.AgentServiceRegistration{
        ID: instanceId,
    }

    err := wrapper.client.Deregister(registration)
    if err != nil {
        logger.Println("Deregister Service Error!", err)
        return false
    }

    logger.Println("Deregister Service Success!")
    return true
}

func (wrapper *ConsulDiscoveryClientWrapper) DiscoverServices(
    serviceName, tag string,
    logger *log.Logger) []interface{} {

    instanceList, ok := wrapper.instanceCache.Load(serviceName)
    if ok {
        return instanceList.([]interface{})
    }

    wrapper.lock.Lock()
    defer wrapper.lock.Unlock()

    instanceList, ok = wrapper.instanceCache.Load(serviceName)
    if ok {
        return instanceList.([]interface{})
    }

    go func() {
        params := make(map[string]interface{})
        params["type"] = "service"
        params["service"] = serviceName

        plan, _ := watch.Parse(params)
        plan.Handler = func(u uint64, i interface{}) {

            if i == nil {
                return
            }

            v, ok := i.([]*api.ServiceEntry)
            if !ok {
                return
            }

            if len(v) == 0 {
                wrapper.instanceCache.Store(serviceName, []interface{}{})
            } else {
                var healthServices []interface{}
                for _, service := range v {
                    if service.Checks.AggregatedStatus() == api.HealthPassing {
                        healthServices = append(healthServices, service.Service)
                    }
                }
                wrapper.instanceCache.Store(serviceName, healthServices)
            }
        }
        defer plan.Stop()

        err := plan.Run(wrapper.config.Address)
        if err != nil {
            logger.Println("Deregister Service Error!", err)
        }
    }()

    entries, _, err := wrapper.client.Service(serviceName, tag, false, nil)
    if err != nil {
        wrapper.instanceCache.Store(serviceName, []interface{}{})
        logger.Println("Discover Service Error!", err)
        return nil
    }

    instances := make([]interface{}, len(entries))
    for i := 0; i < len(instances); i++ {
        instances[i] = entries[i].Service
    }
    wrapper.instanceCache.Store(serviceName, instances)

    return instances
}
