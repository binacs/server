package config

type GRPCConfig struct {
	HttpPort string `toml:"HttpPort"`
	CertPath string `toml:"CertPath"`
	KeyPath  string `toml:"KeyPath"`
	Host     string `toml:"Host"`
}

func defaultGRPCConfig() GRPCConfig {
	return GRPCConfig{
		HttpPort: "80",
		CertPath: "server.crt",
		KeyPath:  "server.key",
		Host:     "binacs.cn",
	}
}
