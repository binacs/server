package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	pb "github.com/binacs/server/api/tinyurl"
	"github.com/binacs/server/config"
	mock_service "github.com/binacs/server/mock/service"
	"github.com/binacsgo/log"
	"github.com/binacsgo/tinyurl"
	"github.com/golang/mock/gomock"
)

var (
	foobar  = "foobar"
	someErr = errors.New("some err")
)

func getNormalConfigWithTinyURL() *config.Config {
	return &config.Config{
		WebConfig: config.WebConfig{
			SSLRedirect: false,
			Host:        "test.binacs.space",
		},
	}
}

func getMockRedisSvcGetRetFoobar(ctrl *gomock.Controller) *mock_service.MockRedisService {
	mock := mock_service.NewMockRedisService(ctrl)
	mock.EXPECT().Get(gomock.Any()).Return(foobar, nil)
	return mock
}

func getMockRedisSvcGetRetErr(ctrl *gomock.Controller) *mock_service.MockRedisService {
	mock := mock_service.NewMockRedisService(ctrl)
	mock.EXPECT().Get(gomock.Any()).Return("", someErr)
	return mock
}

func getMockRedisSvcSetRetNil(ctrl *gomock.Controller) *mock_service.MockRedisService {
	mock := mock_service.NewMockRedisService(ctrl)
	mock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	return mock
}

func getMockRedisSvcSetRetErr(ctrl *gomock.Controller) *mock_service.MockRedisService {
	mock := mock_service.NewMockRedisService(ctrl)
	mock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(someErr)
	return mock
}

func TestTinyURLServiceImpl_AfterInject(t *testing.T) {
	type fields struct {
		Config   *config.Config
		Logger   log.Logger
		RedisSvc RedisService
		MysqlSvc MysqlService
		instance tinyurl.TinyURL
		prefix   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				Config: getNormalConfigWithTinyURL(),
				Logger: log.NewNopLogger(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TinyURLServiceImpl{
				Config:   tt.fields.Config,
				Logger:   tt.fields.Logger,
				RedisSvc: tt.fields.RedisSvc,
				MysqlSvc: tt.fields.MysqlSvc,
				instance: tt.fields.instance,
				prefix:   tt.fields.prefix,
			}
			if err := ts.AfterInject(); (err != nil) != tt.wantErr {
				t.Errorf("TinyURLServiceImpl.AfterInject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTinyURLServiceImpl_TinyURLEncode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		Config   *config.Config
		Logger   log.Logger
		RedisSvc RedisService
		MysqlSvc MysqlService
		instance tinyurl.TinyURL
		prefix   string
	}
	type args struct {
		ctx context.Context
		req *pb.TinyURLEncodeReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.TinyURLEncodeResp
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				Config:   getNormalConfigWithTinyURL(),
				Logger:   log.NewNopLogger(),
				RedisSvc: getMockRedisSvcSetRetNil(ctrl),
				instance: tinyurl.GetMD5Impl(),
				prefix:   getNormalConfigWithTinyURL().WebConfig.GetDomain() + "/r/",
			},
			args: args{
				ctx: context.Background(),
				req: &pb.TinyURLEncodeReq{
					Url: "http://test.binacs.space",
				},
			},
			want: &pb.TinyURLEncodeResp{
				Data: &pb.TinyURLEncodeResObj{
					Turl: "http://test.binacs.space/r/vMN7vi",
				},
			},
			wantErr: false,
		},
		{
			name: "illegal prefix",
			fields: fields{
				Config: getNormalConfigWithTinyURL(),
				Logger: log.NewNopLogger(),
			},
			args: args{
				ctx: context.Background(),
				req: &pb.TinyURLEncodeReq{
					Url: "test.binacs.space",
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "redis set err",
			fields: fields{
				Config:   getNormalConfigWithTinyURL(),
				Logger:   log.NewNopLogger(),
				RedisSvc: getMockRedisSvcSetRetErr(ctrl),
				instance: tinyurl.GetMD5Impl(),
				prefix:   getNormalConfigWithTinyURL().WebConfig.GetDomain() + "/r/",
			},
			args: args{
				ctx: context.Background(),
				req: &pb.TinyURLEncodeReq{
					Url: "https://test.binacs.space",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TinyURLServiceImpl{
				Config:   tt.fields.Config,
				Logger:   tt.fields.Logger,
				RedisSvc: tt.fields.RedisSvc,
				MysqlSvc: tt.fields.MysqlSvc,
				instance: tt.fields.instance,
				prefix:   tt.fields.prefix,
			}
			got, err := ts.TinyURLEncode(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TinyURLServiceImpl.TinyURLEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TinyURLServiceImpl.TinyURLEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTinyURLServiceImpl_TinyURLDecode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		Config   *config.Config
		Logger   log.Logger
		RedisSvc RedisService
		MysqlSvc MysqlService
		instance tinyurl.TinyURL
		prefix   string
	}
	type args struct {
		ctx context.Context
		req *pb.TinyURLDecodeReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.TinyURLDecodeResp
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				Config:   getNormalConfigWithTinyURL(),
				Logger:   log.NewNopLogger(),
				RedisSvc: getMockRedisSvcGetRetFoobar(ctrl),
				instance: tinyurl.GetMD5Impl(),
				prefix:   getNormalConfigWithTinyURL().WebConfig.GetDomain() + "/r/",
			},
			args: args{
				ctx: context.Background(),
				req: &pb.TinyURLDecodeReq{
					Turl: "http://test.binacs.space/r/vMN7vi",
				},
			},
			want: &pb.TinyURLDecodeResp{
				Data: &pb.TinyURLDecodeResObj{
					Url: foobar,
				},
			},
			wantErr: false,
		},
		{
			name: "illegal prefix",
			fields: fields{
				Config: getNormalConfigWithTinyURL(),
				Logger: log.NewNopLogger(),
				prefix: getNormalConfigWithTinyURL().WebConfig.GetDomain() + "/r/",
			},
			args: args{
				ctx: context.Background(),
				req: &pb.TinyURLDecodeReq{
					Turl: "test.binacs.space/r/vMN7vi",
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "redis get err",
			fields: fields{
				Config:   getNormalConfigWithTinyURL(),
				Logger:   log.NewNopLogger(),
				RedisSvc: getMockRedisSvcGetRetErr(ctrl),
				instance: tinyurl.GetMD5Impl(),
				prefix:   getNormalConfigWithTinyURL().WebConfig.GetDomain() + "/r/",
			},
			args: args{
				ctx: context.Background(),
				req: &pb.TinyURLDecodeReq{
					Turl: "http://test.binacs.space/r/vMN7vi",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TinyURLServiceImpl{
				Config:   tt.fields.Config,
				Logger:   tt.fields.Logger,
				RedisSvc: tt.fields.RedisSvc,
				MysqlSvc: tt.fields.MysqlSvc,
				instance: tt.fields.instance,
				prefix:   tt.fields.prefix,
			}
			got, err := ts.TinyURLDecode(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TinyURLServiceImpl.TinyURLDecode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TinyURLServiceImpl.TinyURLDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTinyURLServiceImpl_Encode(t *testing.T) {
	type fields struct {
		Config   *config.Config
		Logger   log.Logger
		RedisSvc RedisService
		MysqlSvc MysqlService
		instance tinyurl.TinyURL
		prefix   string
	}
	type args struct {
		url string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "illegal prefix",
			fields: fields{
				instance: tinyurl.GetMD5Impl(),
			},
			args: args{
				url: "http://test.binacs.space",
			},
			want: "vMN7vi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TinyURLServiceImpl{
				Config:   tt.fields.Config,
				Logger:   tt.fields.Logger,
				RedisSvc: tt.fields.RedisSvc,
				MysqlSvc: tt.fields.MysqlSvc,
				instance: tt.fields.instance,
				prefix:   tt.fields.prefix,
			}
			if got := ts.Encode(tt.args.url); got != tt.want {
				t.Errorf("TinyURLServiceImpl.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTinyURLServiceImpl_URLSearch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		Config   *config.Config
		Logger   log.Logger
		RedisSvc RedisService
		MysqlSvc MysqlService
		instance tinyurl.TinyURL
		prefix   string
	}
	type args struct {
		turl string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				Config:   getNormalConfigWithTinyURL(),
				Logger:   log.NewNopLogger(),
				RedisSvc: getMockRedisSvcGetRetFoobar(ctrl),
				instance: tinyurl.GetMD5Impl(),
				prefix:   getNormalConfigWithTinyURL().WebConfig.GetDomain() + "/r/",
			},
			args: args{
				turl: "someturl",
			},
			want:    foobar,
			wantErr: false,
		},
		{
			name: "redis get err",
			fields: fields{
				Config:   getNormalConfigWithTinyURL(),
				Logger:   log.NewNopLogger(),
				RedisSvc: getMockRedisSvcGetRetErr(ctrl),
				instance: tinyurl.GetMD5Impl(),
				prefix:   getNormalConfigWithTinyURL().WebConfig.GetDomain() + "/r/",
			},
			args: args{
				turl: "someturl",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TinyURLServiceImpl{
				Config:   tt.fields.Config,
				Logger:   tt.fields.Logger,
				RedisSvc: tt.fields.RedisSvc,
				MysqlSvc: tt.fields.MysqlSvc,
				instance: tt.fields.instance,
				prefix:   tt.fields.prefix,
			}
			got, err := ts.URLSearch(tt.args.turl)
			if (err != nil) != tt.wantErr {
				t.Errorf("TinyURLServiceImpl.URLSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TinyURLServiceImpl.URLSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
