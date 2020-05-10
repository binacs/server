package log

type nopLogger struct{}

var _ Logger = (*nopLogger)(nil)

// NewNopLogger ...
func NewNopLogger() Logger { return &nopLogger{} }

// With ...
func (l *nopLogger) With(...interface{}) Logger {
	return l
}

// Debug ...
func (nopLogger) Debug(string, ...interface{}) {}

// Info ...
func (nopLogger) Info(string, ...interface{}) {}

// Warn ...
func (nopLogger) Warn(string, ...interface{}) {}

// Error ...
func (nopLogger) Error(string, ...interface{}) {}
