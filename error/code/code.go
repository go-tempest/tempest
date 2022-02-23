package code

type Code struct {
    Code           string
    DefaultMessage string
}

func New(code, defaultMessage string) *Code {
    return &Code{
        Code:           code,
        DefaultMessage: defaultMessage,
    }
}
