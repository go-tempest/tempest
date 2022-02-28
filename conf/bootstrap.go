package conf

import "github.com/go-tempest/tempest/env"

type Bootstrap struct {
    Profiles
}

type Profiles struct {
    Active env.Env
}