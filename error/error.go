package error

type Code string

type UnifiedErr interface {
    error
    ErrCode() string
}

type SystemErr struct {
    Code
    Message string
}

func (e SystemErr) Error() string {
    return e.Message
}

func (e SystemErr) ErrCode() string {
    return string(e.Code)
}

type ApplicationErr struct {
    Code
    Message string
}

func (e ApplicationErr) Error() string {
    return e.Message
}

func (e ApplicationErr) ErrCode() string {
    return string(e.Code)
}
