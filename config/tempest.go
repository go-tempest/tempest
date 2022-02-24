package config

const (
	DefaultConfigType string = "yaml"
	DefaultConfigName string = "boostrap.yaml"
)

type Tempest struct {
	Application  Application
	Registration Registration
	Logger       Logger
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

type Logger struct {
	Filename   string //日志文件的位置
	MaxSize    int    //在进行切割之前，日志文件的最大大小（以MB为单位）
	MaxBackups int    //保留旧文件的最大个数
	MaxAge     int    //保留旧文件的最大天数
	Compress   bool   //是否压缩/归档旧文件
}

var TempestCfg Tempest
