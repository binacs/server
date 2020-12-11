package config

// TraceConfig the config of jaeger tracing
type TraceConfig struct {
	ServiceName   string `toml:"ServiceName"`
	AgentHostPort string `toml:"AgentHostPort"`
}

func defaultTraceConfig() TraceConfig {
	return TraceConfig{
		ServiceName:   "biancs-cn-trace-svc",
		AgentHostPort: "127.0.0.1:6831",
	}
}
