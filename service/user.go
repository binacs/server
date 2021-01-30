package service

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/binacsgo/log"
	"github.com/binacsgo/token"

	pb "github.com/BinacsLee/server/api/user"
	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/types"
	"github.com/BinacsLee/server/types/table"
)

// UserServiceImpl GRPC user service implement
type UserServiceImpl struct {
	Config   *config.Config `inject-name:"Config"`
	Logger   log.Logger     `inject-name:"GRPCLogger"`
	RedisSvc RedisService   `inject-name:"RedisService"`
	MysqlSvc MysqlService   `inject-name:"MysqlService"`
}

// AfterInject do inject
func (us *UserServiceImpl) AfterInject() error {
	return nil
}

// Register the service
func (us *UserServiceImpl) Register(ctx context.Context, gsrv *grpc.Server, gwmux *runtime.ServeMux, opts []grpc.DialOption) error {
	if gsrv != nil {
		pb.RegisterUserServer(gsrv, us)
	}

	if gwmux != nil {
		if err := pb.RegisterUserHandlerFromEndpoint(ctx, gwmux, ":"+us.Config.GRPCConfig.HttpPort, opts); err != nil {
			return fmt.Errorf("RegisterUserHandlerFromEndpoint: %v", err)
		}
	}
	return nil
}

// ------------------------------- interface -------------------------------

// UserTest for test
func (us *UserServiceImpl) UserTest(ctx context.Context, req *pb.UserTestReq) (*pb.UserTestResp, error) {
	return &pb.UserTestResp{Code: 200, Msg: "test"}, nil
}

// UserRegister register
func (us *UserServiceImpl) UserRegister(ctx context.Context, req *pb.UserRegisterReq) (*pb.UserRegisterResp, error) {
	userID, userPWD := req.GetId(), req.GetPwd()
	user := &table.User{
		UIN: userID,
		Pwd: userPWD,
	}
	affected, err := us.MysqlSvc.GetEngineG().Master().Insert(user)
	if err != nil || affected == 0 {
		return nil, nil
	}

	us.Logger.Info("UserRegister Success", "userID", userID, "UID", user.UID)
	return &pb.UserRegisterResp{
		Data: &pb.UserRegisterDataObj{
			Id:  userID,
			Pwd: userPWD,
		},
	}, nil
}

// UserAuth auth
func (us *UserServiceImpl) UserAuth(ctx context.Context, req *pb.UserAuthReq) (*pb.UserAuthResp, error) {
	id, pwd := req.GetId(), req.GetPwd()
	if len(id) == 0 || len(pwd) == 0 {
		return nil, fmt.Errorf("ID or PWD empty")
	}

	user := &table.User{}
	if exsit, err := us.MysqlSvc.GetEngineG().Where("id=?", id).Get(user); err != nil || !exsit || pwd != user.Pwd {
		return nil, err
	}

	token, refresh := token.GenTokenAndRefresh(id)
	if err := us.redisSetTokenAndRefresh(id, token, refresh); err != nil {
		return nil, err
	}

	us.Logger.Info("UserAuth Success")
	return &pb.UserAuthResp{
		Data: &pb.UserTokenObj{
			AccessToken:  token,
			RefreshToken: refresh,
			ExpireTime:   types.AccessTokenExpire - 2,
		},
	}, nil
}

// UserRefresh refresh token
func (us *UserServiceImpl) UserRefresh(ctx context.Context, req *pb.UserRefreshReq) (*pb.UserRefreshResp, error) {
	refresh := req.GetRefreshToken()
	if len(refresh) == 0 {
		return nil, fmt.Errorf("Token empty")
	}

	id, err := us.RedisSvc.Get(types.RedisRefreshTokenKey(refresh))
	if err != nil {
		return nil, err
	}

	token := token.GenTokenWithRefresh(id, refresh)
	if err := us.redisSetTokenAndRefresh(id, token, refresh); err != nil {
		return nil, err
	}

	us.Logger.Info("UserRefresh Success")
	return &pb.UserRefreshResp{
		Data: &pb.UserTokenObj{
			AccessToken:  token,
			RefreshToken: refresh,
			ExpireTime:   types.AccessTokenExpire - 2,
		},
	}, nil
}

// UserInfo user information
func (us *UserServiceImpl) UserInfo(ctx context.Context, req *pb.UserInfoReq) (*pb.UserInfoResp, error) {
	token := req.GetAccessToken()
	id, err := us.RedisSvc.Get(types.RedisAccessTokenKey(token))
	if err != nil {
		return nil, err
	}

	user := &table.User{}
	if exsit, err := us.MysqlSvc.GetEngineG().Where("id=?", id).Get(user); err != nil || !exsit {
		return nil, err
	}

	us.Logger.Info("UserInfo Success")
	return &pb.UserInfoResp{
		Data: &pb.UserInfoObj{
			Id:    id,
			Role:  user.Role,
			Desc:  user.Desc,
			Ctime: user.CreatedAt.String(), // todo panic
		},
	}, nil
}

func (us *UserServiceImpl) redisSetTokenAndRefresh(id, token, refresh string) error {
	if err := us.RedisSvc.Set(types.RedisAccessTokenKey(token), id, types.AccessTokenExpireDuration()); err != nil {
		return err
	}
	if err := us.RedisSvc.Set(types.RedisRefreshTokenKey(refresh), id, types.RefreshTokenExpireDuration()); err != nil {
		return err
	}
	return nil
}
