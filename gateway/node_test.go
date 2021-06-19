package gateway

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/binacsgo/log"

	"github.com/BinacsLee/server/config"
	mock_gateway "github.com/BinacsLee/server/mock/gateway"
)

var (
	someWebSvcErr  = errors.New("some WebSvc err")
	someGRPCSvcErr = errors.New("some GRPCSvc err")
)

func getMockWebSvcRetNil(ctrl *gomock.Controller) *mock_gateway.MockWebService {
	mock := mock_gateway.NewMockWebService(ctrl)
	mock.EXPECT().Serve().Return(nil)
	return mock
}

func getMockWebSvcRetErr(ctrl *gomock.Controller) *mock_gateway.MockWebService {
	mock := mock_gateway.NewMockWebService(ctrl)
	mock.EXPECT().Serve().Return(someWebSvcErr)
	return mock
}

func getMockGRPCSvcRetNil(ctrl *gomock.Controller) *mock_gateway.MockGRPCService {
	mock := mock_gateway.NewMockGRPCService(ctrl)
	mock.EXPECT().Serve().Return(nil)
	return mock
}

func getMockGRPCSvcRetErr(ctrl *gomock.Controller) *mock_gateway.MockGRPCService {
	mock := mock_gateway.NewMockGRPCService(ctrl)
	mock.EXPECT().Serve().Return(someGRPCSvcErr)
	return mock
}

func TestNodeServiceImpl_OnStart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		Config  *config.Config
		Logger  log.Logger
		WebSvc  WebService
		GRPCSvc GRPCService
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "modeAll",
			fields: fields{
				Config: &config.Config{
					Mode: config.ALL,
				},
				Logger:  log.NewNopLogger(),
				WebSvc:  getMockWebSvcRetNil(ctrl),
				GRPCSvc: getMockGRPCSvcRetNil(ctrl),
			},
			wantErr: false,
		},
		{
			name: "modeWebErr",
			fields: fields{
				Config: &config.Config{
					Mode: config.WEB,
				},
				Logger: log.NewNopLogger(),
				WebSvc: getMockWebSvcRetErr(ctrl),
			},
			// TODO catch error by channel
			// wantErr: true,
			wantErr: false,
		},
		{
			name: "modeGRPCErr",
			fields: fields{
				Config: &config.Config{
					Mode: config.GRPC,
				},
				Logger:  log.NewNopLogger(),
				GRPCSvc: getMockGRPCSvcRetErr(ctrl),
			},
			// TODO catch error by channel
			// wantErr: true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := &NodeServiceImpl{
				Config:  tt.fields.Config,
				Logger:  tt.fields.Logger,
				WebSvc:  tt.fields.WebSvc,
				GRPCSvc: tt.fields.GRPCSvc,
			}
			if err := ns.OnStart(); (err != nil) != tt.wantErr {
				t.Errorf("NodeServiceImpl.OnStart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
