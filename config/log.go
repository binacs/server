package config

// LogConfig config of log
type LogConfig struct {
	File       string `toml:"File"`
	Level      string `toml:"Level"`
	Maxsize    int    `toml:"Maxsize"` //unit is M
	MaxBackups int    `toml:"MaxBackups"`
	Maxage     int    `toml:"Maxage"`
}

func defaultLogConfig() LogConfig {
	return LogConfig{
		File:       "log/server.log",
		Maxsize:    500,
		MaxBackups: 100,
		Maxage:     1000,
	}
}
