package config

type WebConfig struct {
	HttpPort string `toml:"HttpPort"`
	CertPath string `toml:"CertPath"`
	KeyPath  string `toml:"KeyPath"`
	Host     string `toml:"Host"`
}

func defaultWebConfig() WebConfig {
	return WebConfig{
		HttpPort: "80",
		CertPath: "server.crt",
		KeyPath:  "server.key",
		Host:     "binacs.cn",
	}
}
