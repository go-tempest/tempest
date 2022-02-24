package register

import (
    "github.com/go-tempest/tempest/config"
    "github.com/go-tempest/tempest/discovery"
    "github.com/go-tempest/tempest/utils"
    uuid "github.com/satori/go.uuid"
    "os"
)

const instanceIdSeparator string = "_"

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
            instanceId := serviceName + instanceIdSeparator + uuid.NewV4().String()
            instanceHost := getLocalHost()
            instancePort := config.TempestConfig.Application.Port
            healthCheckUrl := config.TempestConfig.Registration.Service.Health.CheckUrl

            if !client.Register(serviceName, instanceId, instanceHost, instancePort, healthCheckUrl, nil) {
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

func getLocalHost() string {
    instanceHost := config.TempestConfig.Registration.Service.Host
    if instanceHost == "" {
        ip, err := utils.GetLocalIP()
        if err != nil {
            os.Exit(-1)
        }
        instanceHost = ip.String()
    }
    return instanceHost
}
