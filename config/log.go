package config

// LogConfig config of log
type LogConfig struct {
	File       string `toml:"File"`
	Level      string `toml:"Level"`
	MaxSize    int    `toml:"MaxSize"` //unit is M
	MaxBackups int    `toml:"MaxBackups"`
	MaxAge     int    `toml:"MaxAge"`
}

func defaultLogConfig() LogConfig {
	return LogConfig{
		File:       "log/server.log",
		MaxSize:    100, // 100MB
		MaxBackups: 10,  // backups
		MaxAge:     30,  // 30 days
	}
}
