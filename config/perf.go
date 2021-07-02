package config

// PerfConfig the config of web service
type PerfConfig struct {
	// HttpPort of prrof, zero means disable the performance
	HttpPort string `toml:"HttpPort"`

	// Plugins...
	// TODO
}

func defaultPprofConfig() PerfConfig {
	return PerfConfig{
		HttpPort: NoPerf,
	}
}
