package consts

import c "github.com/go-tempest/tempest/error/code"

var NoServiceInstancesFound = &c.Code{
    Code:           "100001",
    DefaultMessage: "The service instance was not found",
}

var IllegalArgument = &c.Code{
    Code:           "100002",
    DefaultMessage: "Illegal argument",
}