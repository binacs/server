package config

import (
	"reflect"
	"testing"
)

func Test_defaultPprofConfig(t *testing.T) {
	tests := []struct {
		name string
		want PerfConfig
	}{
		{
			name: "normal",
			want: PerfConfig{
				HttpPort: "0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defaultPprofConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultPprofConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
