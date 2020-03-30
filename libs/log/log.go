package log

import (
	"go.uber.org/zap"
)

type ZapLoggerWapper struct {
	logger *zap.SugaredLogger
}

func NewZapLoggerWapper(logger *zap.SugaredLogger) *ZapLoggerWapper {
	return &ZapLoggerWapper{
		logger: logger,
	}
}
func (z *ZapLoggerWapper) With(ctx ...interface{}) Logger {
	return NewZapLoggerWapper(z.logger.With(ctx...))
}

func (z *ZapLoggerWapper) Debug(msg string, ctx ...interface{}) {
	z.logger.Debugw(msg, ctx...)
}
func (z *ZapLoggerWapper) Info(msg string, ctx ...interface{}) {
	z.logger.Infow(msg, ctx...)
}
func (z *ZapLoggerWapper) Warn(msg string, ctx ...interface{}) {
	z.logger.Warnw(msg, ctx...)
}
func (z *ZapLoggerWapper) Error(msg string, ctx ...interface{}) {
	z.logger.Errorw(msg, ctx...)
}
