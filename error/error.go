package error

import (
    "fmt"
    "github.com/go-tempest/tempest/error/code"
)

type UnifiedErr interface {
    error
    Code() string
}

type SystemErr struct {
    C             *code.Code
    CustomMessage string
}

func (e SystemErr) Error() string {
    return e.getMessage()
}

func (e SystemErr) Code() string {
    return e.C.Code
}

func (e SystemErr) getMessage() string {
    return getMessageWithCode(e.Code(), e.CustomMessage, e.C.DefaultMessage)
}

type ApplicationErr struct {
    C             *code.Code
    CustomMessage string
}

func (e ApplicationErr) Error() string {
    return e.getMessage()
}

func (e ApplicationErr) Code() string {
    return e.C.Code
}

func (e ApplicationErr) getMessage() string {
    return getMessageWithCode(e.Code(), e.CustomMessage, e.C.DefaultMessage)
}

func getMessageWithCode(code, customMessage, defaultMessage string) (message string) {
    if customMessage == "" {
        message = fmt.Sprintf("[%s] %s", code, defaultMessage)
    } else {
        message = fmt.Sprintf("[%s] %s", code, customMessage)
    }
    return
}
