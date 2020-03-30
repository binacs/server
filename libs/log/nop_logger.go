package log

type nopLogger struct{}

var _ Logger = (*nopLogger)(nil)

func NewNopLogger() Logger { return &nopLogger{} }

func (l *nopLogger) With(...interface{}) Logger {
	return l
}
func (nopLogger) Debug(string, ...interface{}) {}
func (nopLogger) Info(string, ...interface{})  {}
func (nopLogger) Warn(string, ...interface{})  {}
func (nopLogger) Error(string, ...interface{}) {}
