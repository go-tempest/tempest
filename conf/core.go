package conf

import (
    "flag"
    "github.com/go-tempest/tempest/env"
    "github.com/go-tempest/tempest/log"
    "os"
    "path"
)

const (
    DefaultConfigType          string = "yaml"
    DefaultBootstrapConfigName string = "boostrap.yaml"
    DefaultLoggerConfigName    string = "logger-%s.yaml"
    DefaultAppConfigName       string = "application-%s.yaml"
    DefaultConfigFileDirName   string = "resources"
)

func GetEnv(e env.Env) env.Env {
    if e == "" {
        return env.Dev
    }
    return e
}

func GetLoggerLevel(ll log.LoggerLevel) log.LoggerLevel {
    if ll == "" {
        return log.Debug
    }
    return ll
}

func GetFlagConfigPath() *string {
    configPath := flag.String(
        "conf",
        "",
        "Specifies the path used to search the configuration file")
    flag.Parse()
    return configPath
}

func GetDefaultConfigPath() *string {
    workDir, err := os.Getwd()
    if err != nil {
        return nil
    }
    p := path.Join(workDir, DefaultConfigFileDirName)
    return &p
}
