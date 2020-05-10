package config

import (
	"crypto/tls"
)

// RedisConfig the config of redis
type RedisConfig struct {
	Network      string `toml:"Network"`
	Addr         string `toml:"Addr"`
	Password     string `toml:"Password"`
	DB           int    `toml:"DB"`
	PoolSize     int    `toml:"PoolSize"`
	MinIdleConns int    `toml:"MinIdleConns"`
	TLSConfig    *tls.Config
	//Limiter Limiter

	// HA
	SentinelMaster    string   `toml:"SentinelMaster"`
	SentinelEndpoints []string `toml:"SentinelEndpoints"`
}

func defaultRedisConfig() RedisConfig {
	return RedisConfig{
		Network:      "tcp",
		Addr:         "127.0.0.1:6379",
		Password:     "",
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 3,
	}
}
