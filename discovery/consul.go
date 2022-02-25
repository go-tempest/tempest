package discovery

import (
    "fmt"
    "github.com/go-kit/kit/sd/consul"
    tempestconfig "github.com/go-tempest/tempest/config"
    "github.com/go-tempest/tempest/consts"
    tempesterr "github.com/go-tempest/tempest/error"
    "github.com/go-tempest/tempest/log"
    "github.com/hashicorp/consul/api"
    "github.com/hashicorp/consul/api/watch"
    "strconv"
    "sync"
)

const (
    httpProtocol  string = "http://"
    portSeparator string = ":"

    defaultHealthCheckInterval    string = "15s"
    defaultDeregisterServiceAfter string = "15m"
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

func (wrapper *ConsulDiscoveryClientWrapper) Register(serviceName, instanceId, instanceHost string,
    instancePort int, healthCheckUrl string,
    meta map[string]string, tags ...string) bool {

    registration := &api.AgentServiceRegistration{
        ID:      instanceId,
        Name:    serviceName,
        Address: instanceHost,
        Port:    instancePort,
        Meta:    meta,
        Tags:    tags,
        Check: &api.AgentServiceCheck{
            Interval: getStringVal(
                tempestconfig.TempestConfig.Registration.Service.Health.CheckInerval, defaultHealthCheckInterval),
            DeregisterCriticalServiceAfter: getStringVal(
                tempestconfig.TempestConfig.Registration.Service.DeregisterAfter, defaultDeregisterServiceAfter),
            HTTP: fmt.Sprintf("%s%s%s%s%s", httpProtocol, instanceHost,
                portSeparator, strconv.Itoa(instancePort), healthCheckUrl),
        },
    }

    err := wrapper.client.Register(registration)
    if err != nil {
        log.Logger.Error("Register Service Error", err)
        return false
    }

    log.Logger.Info("Register Service Success!")
    return true
}

func (wrapper *ConsulDiscoveryClientWrapper) Deregister(instanceId string) bool {

    registration := &api.AgentServiceRegistration{
        ID: instanceId,
    }

    err := wrapper.client.Deregister(registration)
    if err != nil {
        log.Logger.Error("Deregister Service Error!", err)
        return false
    }

    log.Logger.Info("Deregister Service Success!")
    return true
}

func (wrapper *ConsulDiscoveryClientWrapper) DiscoverServices(serviceName, tag string) []interface{} {

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
        params, err := getWatchParams("type", "service", "service", serviceName)
        if err != nil {
            log.Logger.Error("Discover Service Error!", err)
            return
        }

        plan, _ := watch.Parse(params)
        plan.Handler = func(_ uint64, i interface{}) {

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

        e := plan.Run(wrapper.config.Address)
        if e != nil {
            log.Logger.Error("Discover Service Error!", e)
        }
    }()

    entries, _, err := wrapper.client.Service(serviceName, tag, true, nil)
    if err != nil {
        wrapper.instanceCache.Store(serviceName, []interface{}{})
        log.Logger.Error("Discover Service Error!", err)
        return nil
    }

    instances := make([]interface{}, len(entries))
    for i := 0; i < len(instances); i++ {
        instances[i] = entries[i].Service
    }
    wrapper.instanceCache.Store(serviceName, instances)

    return instances
}

func getWatchParams(args ...string) (map[string]interface{}, tempesterr.UnifiedErr) {

    l := len(args)
    if l%2 != 0 {
        return nil, tempesterr.SystemErr{
            C: consts.IllegalArgument,
            CustomMessage: fmt.Sprintf(
                "The parameter length of the method is expected to be an even number, but is actually %d", l),
        }
    }

    step := 2
    params := make(map[string]interface{})
    for i := 0; i < l; i = i + step {
        params[args[i]] = args[i+1]
    }

    return params, nil
}

func getStringVal(val, defaultVal string) string {
    if val == "" {
        return defaultVal
    }
    return val
}
