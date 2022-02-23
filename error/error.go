package error

import "github.com/go-tempest/tempest/error/code"

type UnifiedErr interface {
    error
    Code() string
}

type SystemErr struct {
    C             *code.Code
    CustomMessage string
}

func (e SystemErr) Error() string {
    return getMessage(e.CustomMessage, e.C.DefaultMessage)
}

func (e SystemErr) Code() string {
    return e.C.Code
}

type ApplicationErr struct {
    C             *code.Code
    CustomMessage string
}

func (e ApplicationErr) Error() string {
    return getMessage(e.CustomMessage, e.C.DefaultMessage)
}

func (e ApplicationErr) Code() string {
    return e.C.Code
}

func getMessage(customMessage, defaultMessage string) string {
    if customMessage == "" {
        return defaultMessage
    }
    return customMessage
}
