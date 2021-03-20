package config

// WebConfig the config of web service
type WebConfig struct {
	HttpPort    string            `toml:"HttpPort"`
	HttpsPort   string            `toml:"HttpsPort"`
	SSLRedirect bool              `toml:"SSLRedirect"`
	TmplPath    string            `toml:"TmplPath"` // end with '/'
	CertPath    string            `toml:"CertPath"`
	KeyPath     string            `toml:"KeyPath"`
	Host        string            `toml:"Host"`
	K8sService  map[string]string `toml:"K8sService"`
}

func defaultWebConfig() WebConfig {
	return WebConfig{
		HttpPort:    "80",
		HttpsPort:   "443",
		SSLRedirect: true,
		TmplPath:    "./",
		CertPath:    "server.crt",
		KeyPath:     "server.key",
		Host:        "binacs.cn",
	}
}
