package config

// GRPCConfig the config of grpc service
type GRPCConfig struct {
	HttpPort string `toml:"HttpPort"`
	CertPath string `toml:"CertPath"`
	KeyPath  string `toml:"KeyPath"`
	Host     string `toml:"Host"`
}

func defaultGRPCConfig() GRPCConfig {
	return GRPCConfig{
		HttpPort: "9500",
		CertPath: "api.server.crt",
		KeyPath:  "api.server.key",
		Host:     "api.binacs.cn",
	}
}
