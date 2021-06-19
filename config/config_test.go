package config

import (
	"reflect"
	"sync"
	"testing"
)

func getDefaultConfig() *Config {
	return &Config{
		WorkSpace:   "./",
		File:        "./testdata/default_config.toml",
		Mode:        "all",
		WebConfig:   defaultWebConfig(),
		GRPCConfig:  defaultGRPCConfig(),
		TraceConfig: defaultTraceConfig(),
		LogConfig:   defaultLogConfig(),
		RedisConfig: defaultRedisConfig(),
		MysqlConfig: defaultMysqlConfig(),
		CosConfig:   defaultCosConfig(),
		rwmtx:       &sync.RWMutex{},
	}
}

func getNormalConfig() *Config {
	return &Config{
		WorkSpace: "./",
		File:      "./testdata/normal_config.toml",
		Mode:      "all",
		WebConfig: WebConfig{
			HttpPort:    "80",
			HttpsPort:   "443",
			SSLRedirect: true,
			TmplPath:    "./test/template/",
			CertPath:    "./test/tls/server.crt",
			KeyPath:     "./test/tls/server.key",
			Host:        "server.grpc.io",
			K8sService: map[string]string{
				"CryptoBASE64": "cryptfunc-base64-svc.cryptfunc:8888",
				"CryptoAES":    "cryptfunc-aes-svc.cryptfunc:8888",
				"CryptoDES":    "cryptfunc-des-svc.cryptfunc:8888",
			},
		},
		GRPCConfig: GRPCConfig{
			HttpPort: "9500",
			CertPath: "./test/tls/api.server.crt",
			KeyPath:  "./test/tls/api.server.key",
			Host:     "server.grpc.io",
		},
		TraceConfig: TraceConfig{
			ServiceName:   "binacs-cn",
			AgentHostPort: "jaeger_dc:6831",
		},
		LogConfig: LogConfig{
			File:       "./server.log",
			Level:      "debug",
			Maxsize:    500,
			MaxBackups: 100,
			Maxage:     1000,
		},
		RedisConfig: RedisConfig{
			Network:      "tcp",
			Addr:         "redis_dc:6379",
			Password:     "password",
			DB:           0,
			PoolSize:     10,
			MinIdleConns: 3,
		},
		MysqlConfig: MysqlConfig{
			Conns: []SigCon{
				{
					User:     "root",
					Password: "password",
					Host:     "mysql_dc",
					Port:     "3306",
					DB:       "testdb",
				},
			},
			DSN:          []string{"user:psw@tcp(localhost:3306)/db"},
			MaxIdleConns: 33,
			MaxOpenConns: 100,
		},
		CosConfig: CosConfig{
			BucketURL: "https://examplebucket-12345.cos.COS_REGION.myqcloud.com",
			SecretID:  "SecretID",
			SecretKey: "SecretKey",
			PassKey:   "PassKey",
		},
		rwmtx: &sync.RWMutex{},
	}
}

func getNoFileConfig() *Config {
	cfg := defaultConfig()
	return &cfg
}

func TestLoadFromFile(t *testing.T) {
	type args struct {
		configFile string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "default",
			args: args{
				configFile: "./testdata/default_config.toml",
			},
			want:    getDefaultConfig(),
			wantErr: false,
		},
		{
			name: "normal",
			args: args{
				configFile: "./testdata/normal_config.toml",
			},
			want:    getNormalConfig(),
			wantErr: false,
		},
		{
			name: "nofile",
			args: args{
				configFile: "./testdata/something.toml",
			},
			want:    getNoFileConfig(),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadFromFile(tt.args.configFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFromFile() error:%v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadFromFile():%v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Reload(t *testing.T) {
	type fields struct {
		WorkSpace   string
		File        string
		Mode        string
		WebConfig   WebConfig
		GRPCConfig  GRPCConfig
		TraceConfig TraceConfig
		LogConfig   LogConfig
		RedisConfig RedisConfig
		MysqlConfig MysqlConfig
		CosConfig   CosConfig
		rwmtx       *sync.RWMutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "empty_to_normal",
			fields: fields{
				File:  "./testdata/normal_config.toml",
				rwmtx: &sync.RWMutex{},
			},
			wantErr: false,
		},
		{
			name: "nofile",
			fields: fields{
				WorkSpace:   ".",
				File:        "./nofile.toml",
				Mode:        "all",
				WebConfig:   defaultWebConfig(),
				GRPCConfig:  defaultGRPCConfig(),
				TraceConfig: defaultTraceConfig(),
				LogConfig:   defaultLogConfig(),
				RedisConfig: defaultRedisConfig(),
				MysqlConfig: defaultMysqlConfig(),
				CosConfig:   defaultCosConfig(),
				rwmtx:       &sync.RWMutex{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				WorkSpace:   tt.fields.WorkSpace,
				File:        tt.fields.File,
				Mode:        tt.fields.Mode,
				WebConfig:   tt.fields.WebConfig,
				GRPCConfig:  tt.fields.GRPCConfig,
				TraceConfig: tt.fields.TraceConfig,
				LogConfig:   tt.fields.LogConfig,
				RedisConfig: tt.fields.RedisConfig,
				MysqlConfig: tt.fields.MysqlConfig,
				CosConfig:   tt.fields.CosConfig,
				rwmtx:       tt.fields.rwmtx,
			}
			if err := cfg.Reload(); (err != nil) != tt.wantErr {
				t.Errorf("Config.Reload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
