package config

type GRPCConfig struct {
	HttpPort string `toml:"HttpPort"`
}

func defaultGRPCConfig() GRPCConfig {
	return GRPCConfig{
		HttpPort: "80",
	}
}
