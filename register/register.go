package register

import (
    "github.com/go-tempest/tempest/consts"
    "github.com/go-tempest/tempest/core"
    "github.com/go-tempest/tempest/discovery"
    tempesterr "github.com/go-tempest/tempest/error"
    "github.com/go-tempest/tempest/utils"
    "os"
    "strconv"
)

const (
    instanceIdSeparator   string = "_"
    defaultHealthCheckUrl string = "/health"
)

type Register struct {
}

func (r *Register) StartIfNecessary(ctx *core.Context) {

    enabled := ctx.RegistrationConfig.Enabled
    if enabled {
        client, err := r.connect(ctx)
        if err != nil {
            ctx.Logger().Errorf("Failed to connect to registry\n", err)
            os.Exit(-1)
        }

        registerSelf := ctx.RegistrationConfig.RegisterSelf
        if registerSelf {

            serviceName := ctx.AppConfig.Name
            instanceId, err := createInstanceId(ctx, serviceName)
            if err != nil {
                ctx.Logger().Errorf("Service [%s] register failed\n", serviceName, err)
                os.Exit(-1)
            }
            instanceHost := getLocalHost(ctx)
            instancePort := ctx.AppConfig.Port
            healthCheckUrl := getHealthCheckUrl(ctx)
            tags := ctx.RegistrationConfig.Service.Tags

            checkInerval := ctx.RegistrationConfig.Health.CheckInerval
            deregisterAfter := ctx.RegistrationConfig.Service.DeregisterAfter

            if !client.Register(ctx.Logger, serviceName, instanceId, instanceHost,
                instancePort, healthCheckUrl, checkInerval,
                deregisterAfter, nil, tags...) {

                ctx.Logger().Errorf("Service [%s] register failed\n", serviceName)
                os.Exit(-1)
            }
        }
    }
}

func (r *Register) connect(ctx *core.Context) (discovery.Client, error) {

    client, err := discovery.New(
        ctx.RegistrationConfig.Address,
        ctx.RegistrationConfig.Port,
    )

    if err != nil {
        return nil, err
    }

    return client, nil
}

func getHealthCheckUrl(ctx *core.Context) string {
    url := ctx.RegistrationConfig.Service.Health.CheckUrl
    if url == "" {
        url = defaultHealthCheckUrl
    }
    return url
}

func getLocalHost(ctx *core.Context) string {
    instanceHost := ctx.RegistrationConfig.Service.Host
    if instanceHost == "" {
        ip, err := utils.GetLocalIP()
        if err != nil {
            ctx.Logger().Errorf("Failed to get local IP, error is [%v]\n", err)
            os.Exit(-1)
        }
        instanceHost = ip.String()
    }
    return instanceHost
}

func createInstanceId(ctx *core.Context, svcName string) (string, tempesterr.UnifiedErr) {

    if svcName == "" {
        return "", tempesterr.SystemErr{
            C:             consts.IllegalArgument,
            CustomMessage: "argument [svcName] is empty",
        }
    }

    env := ctx.BootstrapConfig.Active
    return svcName + instanceIdSeparator + string(env) + instanceIdSeparator +
        getLocalHost(ctx) + ":" + strconv.Itoa(ctx.AppConfig.Port), nil
}
