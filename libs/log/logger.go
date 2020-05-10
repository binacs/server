package log

// Logger interface
type Logger interface {
	With(ctx ...interface{}) Logger
	Debug(msg string, ctx ...interface{})
	Info(msg string, ctx ...interface{})
	Warn(msg string, ctx ...interface{})
	Error(msg string, ctx ...interface{})
}
