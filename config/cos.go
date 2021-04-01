package config

// CosConfig config of COS
type CosConfig struct {
	BucketURL string `toml:"BucketURL"`
	SecretID  string `toml:"SecretID"`
	SecretKey string `toml:"SecretKey"`
	PassKey   string `toml:"PassKey"`
}

func defaultCosConfig() CosConfig {
	return CosConfig{
		BucketURL: "https://examplebucket-12345.cos.COS_REGION.myqcloud.com",
		SecretID:  "SecretID",
		SecretKey: "SecretKey",
		PassKey:   "PassKey",
	}
}
