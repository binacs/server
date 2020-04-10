package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"xorm.io/core"
)

var coreLevelToBeegoLevel = map[core.LogLevel]zapcore.Level{
	core.LOG_WARNING: zap.WarnLevel,
	core.LOG_INFO:    zap.InfoLevel,
	core.LOG_DEBUG:   zap.DebugLevel,
	core.LOG_ERR:     zap.ErrorLevel,
	core.LOG_OFF:     zap.WarnLevel,
	core.LOG_UNKNOWN: zap.InfoLevel,
}

var _ core.ILogger = &MysqlLogger{}

type MysqlLogger struct {
	w       *zap.Logger
	showSQL bool
}

func NewMysqlLogger(w *zap.Logger) *MysqlLogger {
	return &MysqlLogger{w: w}
}

func (s *MysqlLogger) Debug(v ...interface{}) {
	s.w.Debug(fmt.Sprint(v...))
}

func (s *MysqlLogger) Debugf(format string, v ...interface{}) {
	s.w.Debug(fmt.Sprintf(format, v...))
}

func (s *MysqlLogger) Error(v ...interface{}) {
	s.w.Error(fmt.Sprint(v...))
}

func (s *MysqlLogger) Errorf(format string, v ...interface{}) {
	s.w.Error(fmt.Sprintf(format, v...))
}

func (s *MysqlLogger) Info(v ...interface{}) {
	s.w.Info(fmt.Sprint(v...))
}

func (s *MysqlLogger) Infof(format string, v ...interface{}) {
	s.w.Info(fmt.Sprintf(format, v...))
}

func (s *MysqlLogger) Warn(v ...interface{}) {
	s.w.Warn(fmt.Sprint(v...))
}

func (s *MysqlLogger) Warnf(format string, v ...interface{}) {
	s.w.Warn(fmt.Sprintf(format, v...))
}

func (s *MysqlLogger) Level() core.LogLevel {
	return core.LOG_UNKNOWN
}

func (s *MysqlLogger) SetLevel(l core.LogLevel) {
}

func (s *MysqlLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		s.showSQL = true
		return
	}
	s.showSQL = show[0]
}

func (s *MysqlLogger) IsShowSQL() bool {
	return s.showSQL
}
