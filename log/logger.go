package log

type Log interface {
	BaseLog
	Base
}

type BaseLog interface {
	With(args ...interface{}) Base
}

type Base interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}
