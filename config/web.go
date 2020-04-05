package config

type WebConfig struct {
	HttpPort string `toml:"HttpPort"`
}

func defaultWebConfig() WebConfig {
	return WebConfig{
		HttpPort: "80",
	}
}
