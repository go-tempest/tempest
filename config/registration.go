package config

type RegistrationConfig struct {
    Registration
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
    Port            int
    Tags            []string
    Health
}

type Health struct {
    CheckInerval string `mapstructure:"check-interval"`
    CheckUrl     string `mapstructure:"check-url"`
}
