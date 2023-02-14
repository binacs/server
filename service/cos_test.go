package service

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/binacs/server/api/cos"
	"github.com/binacs/server/config"
	"github.com/binacsgo/log"
)

func getNormalConfigWithCos() *config.Config {
	return &config.Config{
		CosConfig: config.CosConfig{
			BucketURL: "https://examplebucket-12345.cos.COS_REGION.myqcloud.com",
			SecretID:  "SecretID",
			SecretKey: "SecretKey",
			PassKey:   "PassKey",
		},
	}
}

func TestCosServiceImpl_AfterInject(t *testing.T) {
	type fields struct {
		Config       *config.Config
		Logger       log.Logger
		timeTemplate string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				Config: getNormalConfigWithCos(),
				Logger: log.NewNopLogger(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CosServiceImpl{
				Config:       tt.fields.Config,
				Logger:       tt.fields.Logger,
				timeTemplate: tt.fields.timeTemplate,
			}
			if err := cs.AfterInject(); (err != nil) != tt.wantErr {
				t.Errorf("CosServiceImpl.AfterInject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCosServiceImpl_generateFileName(t *testing.T) {
	type fields struct {
		Config       *config.Config
		Logger       log.Logger
		timeTemplate string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "empty",
			fields: fields{
				Config: getNormalConfigWithCos(),
				Logger: log.NewNopLogger(),
			},
			args: args{
				name: "",
			},
			want: "_.",
		},
		{
			name: "filename",
			fields: fields{
				Config: getNormalConfigWithCos(),
				Logger: log.NewNopLogger(),
			},
			args: args{
				name: "filename",
			},
			want: "filename_",
		},
		{
			name: "complex filename",
			fields: fields{
				Config: getNormalConfigWithCos(),
				Logger: log.NewNopLogger(),
			},
			args: args{
				name: "some_filename%abc^123",
			},
			want: "some_filename%abc^123_",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CosServiceImpl{
				Config:       tt.fields.Config,
				Logger:       tt.fields.Logger,
				timeTemplate: tt.fields.timeTemplate,
			}
			if got := cs.generateFileName(tt.args.name); got != tt.want {
				t.Errorf("CosServiceImpl.generateFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCosServiceImpl_generateCosURI(t *testing.T) {
	type fields struct {
		Config       *config.Config
		Logger       log.Logger
		timeTemplate string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "filename",
			fields: fields{
				Config: getNormalConfigWithCos(),
				Logger: log.NewNopLogger(),
			},
			args: args{
				name: "filename",
			},
			want: "https://examplebucket-12345.cos.COS_REGION.myqcloud.com/filename",
		},
		{
			name: "filename with blank",
			fields: fields{
				Config: getNormalConfigWithCos(),
				Logger: log.NewNopLogger(),
			},
			args: args{
				name: "some filename with blank",
			},
			want: "https://examplebucket-12345.cos.COS_REGION.myqcloud.com/some%20filename%20with%20blank",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CosServiceImpl{
				Config:       tt.fields.Config,
				Logger:       tt.fields.Logger,
				timeTemplate: tt.fields.timeTemplate,
			}
			if got := cs.generateCosURI(tt.args.name); got != tt.want {
				t.Errorf("CosServiceImpl.generateCosURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCosServiceImpl_processCosURI(t *testing.T) {
	type fields struct {
		Config       *config.Config
		Logger       log.Logger
		timeTemplate string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "filename",
			fields: fields{
				Config: getNormalConfigWithCos(),
				Logger: log.NewNopLogger(),
			},
			args: args{
				name: "https://examplebucket-12345.cos.COS_REGION.myqcloud.com/filename",
			},
			want: "filename",
		},
		{
			name: "filename with blank",
			fields: fields{
				Config: getNormalConfigWithCos(),
				Logger: log.NewNopLogger(),
			},
			args: args{
				name: "https://examplebucket-12345.cos.COS_REGION.myqcloud.com/some%20filename%20with%20blank",
			},
			want: "some%20filename%20with%20blank",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CosServiceImpl{
				Config:       tt.fields.Config,
				Logger:       tt.fields.Logger,
				timeTemplate: tt.fields.timeTemplate,
			}
			if got := cs.processCosURI(tt.args.name); got != tt.want {
				t.Errorf("CosServiceImpl.processCosURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCosServiceImpl_CosBucketURL(t *testing.T) {
	type fields struct {
		Config       *config.Config
		Logger       log.Logger
		timeTemplate string
	}
	type args struct {
		ctx context.Context
		req *pb.CosBucketURLReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CosBucketURLResp
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				Config: getNormalConfigWithCos(),
				Logger: log.NewNopLogger(),
			},
			args: args{
				ctx: context.Background(),
				req: &pb.CosBucketURLReq{},
			},
			want: &pb.CosBucketURLResp{
				Data: &pb.CosBucketURLResObj{
					BucketURL: getNormalConfigWithCos().CosConfig.BucketURL,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CosServiceImpl{
				Config:       tt.fields.Config,
				Logger:       tt.fields.Logger,
				timeTemplate: tt.fields.timeTemplate,
			}
			got, err := cs.CosBucketURL(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CosServiceImpl.CosBucketURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CosServiceImpl.CosBucketURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO Can't mock github.com/tencentyun/cos-go-sdk-v5 cli
