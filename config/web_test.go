package config

import "testing"

func TestWebConfig_GetDomain(t *testing.T) {
	type fields struct {
		HttpPort    string
		HttpsPort   string
		SSLRedirect bool
		TmplPath    string
		CertPath    string
		KeyPath     string
		Host        string
		K8sService  map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "noSSLRedirect",
			fields: fields{
				SSLRedirect: false,
				Host:        "test.binacs.cn",
			},
			want: "http://test.binacs.cn",
		},
		{
			name: "doSSLRedirect",
			fields: fields{
				SSLRedirect: true,
				Host:        "test.binacs.cn",
			},
			want: "https://test.binacs.cn",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := &WebConfig{
				HttpPort:    tt.fields.HttpPort,
				HttpsPort:   tt.fields.HttpsPort,
				SSLRedirect: tt.fields.SSLRedirect,
				TmplPath:    tt.fields.TmplPath,
				CertPath:    tt.fields.CertPath,
				KeyPath:     tt.fields.KeyPath,
				Host:        tt.fields.Host,
				K8sService:  tt.fields.K8sService,
			}
			if got := wc.GetDomain(); got != tt.want {
				t.Errorf("WebConfig.GetDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
