package conf

type AppConfig struct {
    Application
}

type Application struct {
    Name string
    Port int
}
