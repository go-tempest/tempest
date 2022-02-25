package register

import (
    "fmt"
    "github.com/go-tempest/tempest/config"
    "github.com/go-tempest/tempest/consts"
    "github.com/go-tempest/tempest/discovery"
    tempesterr "github.com/go-tempest/tempest/error"
    "github.com/go-tempest/tempest/log"
    "github.com/go-tempest/tempest/utils"
    "os"
    "strconv"
)

const (
    instanceIdSeparator   string = "_"
    defaultHealthCheckUrl string = "/health"
)

type Registration struct {
}

func (r *Registration) StartIfNecessary() {
    enabled := config.TempestConfig.Registration.Enabled
    if enabled {
        client, err := r.connect()
        if err != nil {
            os.Exit(-1)
        }

        registerSelf := config.TempestConfig.Registration.RegisterSelf
        if registerSelf {

            serviceName := config.TempestConfig.Application.Name
            instanceId, err := createInstanceId(serviceName)
            if err != nil {
                log.Logger.Error(fmt.Sprintf("Service [%s] register failed", serviceName), err)
                os.Exit(-1)
            }
            instanceHost := getLocalHost()
            instancePort := config.TempestConfig.Application.Port
            healthCheckUrl := getHealthCheckUrl()
            tags := config.TempestConfig.Registration.Service.Tags

            if !client.Register(serviceName, instanceId, instanceHost,
                instancePort, healthCheckUrl, nil, tags...) {

                log.Logger.Error("Failed to register for service")
                os.Exit(-1)
            }
        }
    }
}

func (r *Registration) connect() (discovery.Client, error) {

    client, err := discovery.New(
        config.TempestConfig.Registration.Address,
        config.TempestConfig.Registration.Port,
    )

    if err != nil {
        return nil, err
    }

    return client, nil
}

func getHealthCheckUrl() string {
    url := config.TempestConfig.Registration.Service.Health.CheckUrl
    if url == "" {
        url = defaultHealthCheckUrl
    }
    return url
}

func getLocalHost() string {
    instanceHost := config.TempestConfig.Registration.Service.Host
    if instanceHost == "" {
        ip, err := utils.GetLocalIP()
        if err != nil {
            fmt.Printf("Failed to get local IP, error is [%v]\n", err)
            os.Exit(-1)
        }
        instanceHost = ip.String()
    }
    return instanceHost
}

func createInstanceId(svcName string) (string, tempesterr.UnifiedErr) {

    if svcName == "" {
        return "", tempesterr.SystemErr{
            C:             consts.IllegalArgument,
            CustomMessage: "argument [svcName] is empty",
        }
    }

    env := config.TempestConfig.Application.Profiles.Active
    return svcName + instanceIdSeparator + string(env) + instanceIdSeparator +
        getLocalHost() + ":" + strconv.Itoa(config.TempestConfig.Application.Port), nil
}
