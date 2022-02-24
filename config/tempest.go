package config

import (
	"github.com/go-tempest/tempest/env"
	"github.com/go-tempest/tempest/log"
)

const (
	DefaultConfigType string = "yaml"
	DefaultConfigName string = "boostrap.yaml"
)

type Tempest struct {
	Application
	Registration
	Logger
	Mysql
	SqlXml
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
	Address      string
	Port         int
	RegisterSelf bool `yaml:"register-self"`
	Enabled      bool
}

type Logger struct {
	log.Type
	log.Level
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type Mysql struct {
	Username    string
	Password    string
	Path        string
	Dbname      string
	Config      string
	MaxIdleCons int
	MaxOpenCons int
	LogMode     bool
}

type SqlXml struct {
	FileUrl string
}

var TempestConfig Tempest
