package config

import (
	"reflect"
	"testing"
)

func TestMysqlConfig_GenerateDSN(t *testing.T) {
	type fields struct {
		Conns        []SigCon
		DSN          []string
		MaxIdleConns int
		MaxOpenConns int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "normal",
			fields: fields{
				Conns: []SigCon{
					{
						User:     "root1",
						Password: "password",
						Host:     "mysql_dc",
						Port:     "3306",
						DB:       "testdb",
					},
					{
						User:     "root2",
						Password: "password",
						Host:     "mysql_dc",
						Port:     "3306",
						DB:       "testdb",
					},
				},
			},
			want: []string{
				"root1:password@tcp(mysql_dc:3306)/testdb?charset=utf8&parseTime=True&loc=Local",
				"root2:password@tcp(mysql_dc:3306)/testdb?charset=utf8&parseTime=True&loc=Local",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := MysqlConfig{
				Conns:        tt.fields.Conns,
				DSN:          tt.fields.DSN,
				MaxIdleConns: tt.fields.MaxIdleConns,
				MaxOpenConns: tt.fields.MaxOpenConns,
			}
			if got := mc.GenerateDSN(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MysqlConfig.GenerateDSN() = %v, want %v", got, tt.want)
			}
		})
	}
}
