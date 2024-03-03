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

func (wc *WebConfig) GetDomain() string {
	var domain string
	if wc.SSLRedirect {
		domain = "https://"
	} else {
		domain = "http://"
	}
	return domain + wc.Host
}

func defaultWebConfig() WebConfig {
	return WebConfig{
		HttpPort:    "80",
		HttpsPort:   "443",
		SSLRedirect: true,
		TmplPath:    "./",
		CertPath:    "server.crt",
		KeyPath:     "server.key",
		Host:        "binacs.space",
	}
}
