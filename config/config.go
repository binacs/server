package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

// Config ...
type Config struct {
	Host    string `toml:host`
	Port    string `toml:port`
	ImgPath string `toml:"imgpath"`
}

var config Config
var configFile = "./config.toml"

// SetPath
func SetConfPath(path string) {
	configFile = path
}

// GetConfig config struct from config file
func GetConfig() Config {
	if config.Host == "" {
		// default value
		if _, err := toml.DecodeFile(configFile, &config); err != nil {
			log.Fatal("read the struct of config failed !", err)
		}
	}
	return config
}
