package config

import "fmt"

type SigCon struct {
	User     string `toml:"User"`
	Password string `toml:"Password"`
	Host     string `toml:"Host"`
	Port     string `toml:"Port"`
	DB       string `toml:"DB"`
}

type MysqlConfig struct {
	Conns        []SigCon `toml:"Conns"`
	DSN          []string
	MaxIdleConns int `toml:"MaxIdleConns"`
	MaxOpenConns int `toml:"MaxOpenConns"`
}

func defaultMysqlConfig() MysqlConfig {
	return MysqlConfig{
		DSN:          []string{"user:psw@tcp(localhost:3306)/db"},
		MaxIdleConns: 3,
		MaxOpenConns: 10,
	}
}

func (mc MysqlConfig) GenerateDSN() []string {
	num := len(mc.Conns)
	DSN := make([]string, num, num)
	for idx, sigCon := range mc.Conns {
		DSN[idx] = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			sigCon.User, sigCon.Password, sigCon.Host, sigCon.Port, sigCon.DB)
	}
	return DSN
}
