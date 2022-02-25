package log

import "go.uber.org/zap"

type Log interface {
	With(args ...interface{}) *zap.SugaredLogger
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}
