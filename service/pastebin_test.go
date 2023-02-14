package service

import (
	"testing"

	"github.com/binacsgo/pastebin"

	"github.com/binacsgo/log"

	"github.com/binacs/server/config"
)

func getNormalConfigWithPastebin() *config.Config {
	return &config.Config{
		WebConfig: config.WebConfig{
			SSLRedirect: false,
			Host:        "test.binacs.cn",
		},
	}
}

func TestPastebinServiceImpl_AfterInject(t *testing.T) {
	type fields struct {
		Config     *config.Config
		Logger     log.Logger
		RedisSvc   RedisService
		MysqlSvc   MysqlService
		TinyURLSvc TinyURLService
		instance   pastebin.PasteBin
		prefix     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				Config: getNormalConfigWithPastebin(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &PastebinServiceImpl{
				Config:     tt.fields.Config,
				Logger:     tt.fields.Logger,
				RedisSvc:   tt.fields.RedisSvc,
				MysqlSvc:   tt.fields.MysqlSvc,
				TinyURLSvc: tt.fields.TinyURLSvc,
				instance:   tt.fields.instance,
				prefix:     tt.fields.prefix,
			}
			if err := ps.AfterInject(); (err != nil) != tt.wantErr {
				t.Errorf("PastebinServiceImpl.AfterInject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPastebinServiceImpl_Parse(t *testing.T) {
	type fields struct {
		Config     *config.Config
		Logger     log.Logger
		RedisSvc   RedisService
		MysqlSvc   MysqlService
		TinyURLSvc TinyURLService
		instance   pastebin.PasteBin
		prefix     string
	}
	type args struct {
		content string
		syntax  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "markdown",
			fields: fields{
				Config:   getNormalConfigWithPastebin(),
				Logger:   log.NewNopLogger(),
				instance: new(pastebin.MarkDownImpl),
				prefix:   "test.binacs.cn",
			},
			args: args{
				content: "something",
				syntax:  "markdown",
			},
			want: "<p>something</p>\n",
		},
		{
			name: "raw",
			fields: fields{
				Config:   getNormalConfigWithPastebin(),
				Logger:   log.NewNopLogger(),
				instance: new(pastebin.MarkDownImpl),
				prefix:   "test.binacs.cn",
			},
			args: args{
				content: "something",
				syntax:  "raw",
			},
			want: "something",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &PastebinServiceImpl{
				Config:     tt.fields.Config,
				Logger:     tt.fields.Logger,
				RedisSvc:   tt.fields.RedisSvc,
				MysqlSvc:   tt.fields.MysqlSvc,
				TinyURLSvc: tt.fields.TinyURLSvc,
				instance:   tt.fields.instance,
				prefix:     tt.fields.prefix,
			}
			if got := ps.Parse(tt.args.content, tt.args.syntax); got != tt.want {
				t.Errorf("PastebinServiceImpl.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
