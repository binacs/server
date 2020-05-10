package log

import (
	"go.uber.org/zap"
)

// ZapLoggerWapper zap logger wapper
type ZapLoggerWapper struct {
	logger *zap.SugaredLogger
}

// NewZapLoggerWapper return a pointer to ZapLoggerWapper
func NewZapLoggerWapper(logger *zap.SugaredLogger) *ZapLoggerWapper {
	return &ZapLoggerWapper{
		logger: logger,
	}
}

// With return a new Logger with context
func (z *ZapLoggerWapper) With(ctx ...interface{}) Logger {
	return NewZapLoggerWapper(z.logger.With(ctx...))
}

// Debug ...
func (z *ZapLoggerWapper) Debug(msg string, ctx ...interface{}) {
	z.logger.Debugw(msg, ctx...)
}

// Info ...
func (z *ZapLoggerWapper) Info(msg string, ctx ...interface{}) {
	z.logger.Infow(msg, ctx...)
}

// Warn ...
func (z *ZapLoggerWapper) Warn(msg string, ctx ...interface{}) {
	z.logger.Warnw(msg, ctx...)
}

// Error ...
func (z *ZapLoggerWapper) Error(msg string, ctx ...interface{}) {
	z.logger.Errorw(msg, ctx...)
}
