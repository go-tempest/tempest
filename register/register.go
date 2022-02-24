package register

import (
    "github.com/go-tempest/tempest/config"
    "github.com/go-tempest/tempest/discovery"
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

            if !client.Register(serviceName, instanceId, "/health", nil, "127.0.0.1", 8080, nil) {
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
