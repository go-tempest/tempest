package config

const (
    DefaultConfigType string = "yaml"
    DefaultConfigName string = "boostrap.yaml"
)

type Tempest struct {
    Application  Application
    Registration Registration
}

type Application struct {
    Name string
    Port int
}

type Registration struct {
    Address      string
    Port         int
    RegisterSelf bool `yaml:"register-self"`
    Enabled      bool
}

var TempestCfg Tempest
