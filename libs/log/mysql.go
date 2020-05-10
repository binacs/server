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

// MysqlLogger
type MysqlLogger struct {
	w       *zap.Logger
	showSQL bool
}

// NewMysqlLogger return a pointer to MysqlLogger
func NewMysqlLogger(w *zap.Logger) *MysqlLogger {
	return &MysqlLogger{w: w}
}

// Debug ...
func (s *MysqlLogger) Debug(v ...interface{}) {
	s.w.Debug(fmt.Sprint(v...))
}

// Debugf ...
func (s *MysqlLogger) Debugf(format string, v ...interface{}) {
	s.w.Debug(fmt.Sprintf(format, v...))
}

// Error ...
func (s *MysqlLogger) Error(v ...interface{}) {
	s.w.Error(fmt.Sprint(v...))
}

// Errorf ...
func (s *MysqlLogger) Errorf(format string, v ...interface{}) {
	s.w.Error(fmt.Sprintf(format, v...))
}

// Info ...
func (s *MysqlLogger) Info(v ...interface{}) {
	s.w.Info(fmt.Sprint(v...))
}

// Infof ...
func (s *MysqlLogger) Infof(format string, v ...interface{}) {
	s.w.Info(fmt.Sprintf(format, v...))
}

// Warn ...
func (s *MysqlLogger) Warn(v ...interface{}) {
	s.w.Warn(fmt.Sprint(v...))
}

// Warnf ...
func (s *MysqlLogger) Warnf(format string, v ...interface{}) {
	s.w.Warn(fmt.Sprintf(format, v...))
}

// Level get level
func (s *MysqlLogger) Level() core.LogLevel {
	return core.LOG_UNKNOWN
}

// SetLevel set level
func (s *MysqlLogger) SetLevel(l core.LogLevel) {
}

// ShowSQL show sql
func (s *MysqlLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		s.showSQL = true
		return
	}
	s.showSQL = show[0]
}

// IsShowSQL sql showed or not
func (s *MysqlLogger) IsShowSQL() bool {
	return s.showSQL
}
