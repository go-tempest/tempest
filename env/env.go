package env

type Env string

const (
    Prod  Env = "prod"
    Test  Env = "test"
    Dev   Env = "dev"
    Local Env = "local"
)
