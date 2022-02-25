package config

import (
    "github.com/go-tempest/tempest/env"
)

const (
    DefaultConfigType string = "yaml"
    DefaultConfigName string = "boostrap.yaml"
)

type Tempest struct {
    Application
    Registration
    Logger
}

type Application struct {
    Name string
    Port int
    Profiles
}

type Profiles struct {
    Active env.Env
}

type Registration struct {
    Enabled bool
    Address string
    Port    int
    Service
}

type Service struct {
    RegisterSelf    bool   `mapstructure:"register-self"`
    DeregisterAfter string `mapstructure:"deregister-after"`
    Host            string
    Tags            []string
    Health
}

type Health struct {
    CheckInerval string `mapstructure:"check-interval"`
    CheckUrl     string `mapstructure:"check-url"`
}

type Logger struct {
    Type       string
    Level      string
    Filename   string `mapstructure:"filename"`
    MaxSize    int    `mapstructure:"max-size"`
    MaxBackups int    `mapstructure:"max-backups"`
    MaxAge     int    `mapstructure:"max-age"`
    Compress   bool
}

var TempestConfig Tempest
