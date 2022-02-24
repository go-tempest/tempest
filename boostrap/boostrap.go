package boostrap

import (
    "flag"
    "fmt"
    "github.com/go-tempest/tempest/config"
    "github.com/go-tempest/tempest/log"
    "github.com/go-tempest/tempest/register"
    "github.com/spf13/viper"
    "os"
    "path"
    "sync"
)

type bootstrap struct {
    sync.Once
}

func (b *bootstrap) start() {
    b.Do(func() {
        new(register.Registration).StartIfNecessary()
        // TODO 启动
    })
}

func getFlagConfigPath() *string {
    var configPath = flag.String("config", "", "Specifies the path used to search the configuration file")
    flag.Parse()
    return configPath
}

func getDefaultConfigPath() *string {
    workDir, err := os.Getwd()
    if err != nil {
        return nil
    }
    p := path.Join(workDir, "resources")
    return &p
}

func parseYaml(v *viper.Viper) {
    if err := v.Unmarshal(&config.TempestConfig); err != nil {
        fmt.Printf("Deserialization configuration failed, error is [%v]\n", err)
        os.Exit(-1)
    }
}

func init() {
    initViper()
    log.Initialize()

    new(bootstrap).start()
}

func initViper() {
    viper.AutomaticEnv()
    viper.SetConfigType(config.DefaultConfigType)
    viper.SetConfigName(config.DefaultConfigName)
    viper.AddConfigPath(*getDefaultConfigPath())
    viper.AddConfigPath(*getFlagConfigPath())

    err := viper.ReadInConfig()
    if err != nil {
        fmt.Printf("Viper initialization failed, error is [%v]\n", err)
        os.Exit(-1)
    }

    parseYaml(viper.GetViper())
}
