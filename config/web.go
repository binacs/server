package config

// WebConfig the config of web service
type WebConfig struct {
	HttpPort     string            `toml:"HttpPort"`
	HttpsPort    string            `toml:"HttpsPort"`
	StaticPath   string            `toml:"StaticPath"` // end with '/'
	CertPath     string            `toml:"CertPath"`
	KeyPath      string            `toml:"KeyPath"`
	Host         string            `toml:"Host"`
	ReverseProxy map[string]string `toml:"ReverseProxy"`
	K8sService	 map[string]string `toml:"K8sService"`
}

func defaultWebConfig() WebConfig {
	return WebConfig{
		HttpPort:   "80",
		HttpsPort:  "443",
		StaticPath: "./",
		CertPath:   "server.crt",
		KeyPath:    "server.key",
		Host:       "binacs.cn",
	}
}
