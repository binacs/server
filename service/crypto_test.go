package service

import (
	"testing"

	"github.com/binacsgo/log"

	pb "github.com/BinacsLee/server/api/crypto"
	"github.com/BinacsLee/server/config"
)

func getNormalConfigWithCrypto() *config.Config {
	return &config.Config{
		WebConfig: config.WebConfig{
			K8sService: map[string]string{
				"CryptoBASE64": "cryptfunc-base64-svc.cryptfunc:8888",
				"CryptoAES":    "cryptfunc-aes-svc.cryptfunc:8888",
				"CryptoDES":    "cryptfunc-des-svc.cryptfunc:8888",
			},
		},
	}
}

func TestCryptoServiceImpl_AfterInject(t *testing.T) {
	type fields struct {
		Config       *config.Config
		Logger       log.Logger
		clientBASE64 pb.CryptoClient
		clientAES    pb.CryptoClient
		clientDES    pb.CryptoClient
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				Config: getNormalConfigWithCrypto(),
				Logger: log.NewNopLogger(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CryptoServiceImpl{
				Config:       tt.fields.Config,
				Logger:       tt.fields.Logger,
				clientBASE64: tt.fields.clientBASE64,
				clientAES:    tt.fields.clientAES,
				clientDES:    tt.fields.clientDES,
			}
			if err := cs.AfterInject(); (err != nil) != tt.wantErr {
				t.Errorf("CryptoServiceImpl.AfterInject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TODO mock clients
