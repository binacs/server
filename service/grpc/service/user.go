package service

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	pb "github.com/BinacsLee/server/api/user/v1"
	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/base"
	"github.com/BinacsLee/server/libs/errcode"
	"github.com/BinacsLee/server/libs/log"
	"github.com/BinacsLee/server/service/db"
	"github.com/BinacsLee/server/types"
	"github.com/BinacsLee/server/types/table"
)

// GRPCUserServiceImpl GRPC user service implement
type GRPCUserServiceImpl struct {
	Config   *config.Config  `inject-name:"Config"`
	Logger   log.Logger      `inject-name:"GRPCLogger"`
	RedisSvc db.RedisService `inject-name:"RedisService"`
	MysqlSvc db.MysqlService `inject-name:"MysqlService"`
}

// Register a new user
func (us *GRPCUserServiceImpl) Register(ctx context.Context, gsrv *grpc.Server, gwmux *runtime.ServeMux) error {
	if gsrv != nil {
		pb.RegisterUserServer(gsrv, us)
	}

	if gwmux != nil {
		creds, err := credentials.NewClientTLSFromFile(us.Config.GRPCConfig.CertPath, us.Config.GRPCConfig.Host)
		if err != nil {
			return fmt.Errorf("NewClientTLSFromFile: %v", err)
		}
		opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
		if err := pb.RegisterUserHandlerFromEndpoint(ctx, gwmux, ":"+us.Config.GRPCConfig.HttpPort, opts); err != nil {
			return fmt.Errorf("RegisterUserHandlerFromEndpoint: %v", err)
		}
	}
	return nil
}

// interface...

// UserTest for test
func (us *GRPCUserServiceImpl) UserTest(ctx context.Context, req *pb.UserTestReq) (*pb.UserTestResp, error) {
	us.Logger.Info("UserTest", "req", req)
	return &pb.UserTestResp{Code: 200, Msg: "test"}, nil
}

// UserRegister register
func (us *GRPCUserServiceImpl) UserRegister(ctx context.Context, req *pb.UserRegisterReq) (*pb.UserRegisterResp, error) {
	us.Logger.Info("UserRegister", "req", req)
	rsp := &pb.UserRegisterResp{
		Code: errcode.ErrGrpcSuccess.Code(), // 0
		Msg:  errcode.ErrGrpcSuccess.Error(),
	}
	userID := req.GetId()
	userPWD := req.GetPwd()
	user := &table.User{
		UIN: userID,
		Pwd: userPWD,
	}
	affected, err := us.MysqlSvc.GetEngineG().Master().Insert(user)
	if err != nil || affected == 0 {
		us.Logger.Error("UserRegister MysqlSvc Insert", "affected", affected, "err", err)
		return nil, errcode.ErrGrpcAppExecute.AppendMsg("MysqlSvc Insert| err: %s, affected: %s", err, affected)
	}
	rsp.Data = &pb.UserRegisterDataObj{
		Id:  userID,
		Pwd: userPWD,
	}
	us.Logger.Info("UserRegister Success", "userID", userID, "UID", user.UID)
	return rsp, nil
}

// UserAuth auth
func (us *GRPCUserServiceImpl) UserAuth(ctx context.Context, req *pb.UserAuthReq) (*pb.UserAuthResp, error) {
	us.Logger.Info("UserAuth", "req", req)
	rsp := &pb.UserAuthResp{
		Code: errcode.ErrGrpcSuccess.Code(),
		Msg:  errcode.ErrGrpcSuccess.Error(),
	}
	tokenInfo, Err := us.doAuthToken(ctx, req)
	if Err != nil {
		rsp.Code = Err.Code()
		rsp.Msg = Err.Error()
		us.Logger.Error("UserAuth doAuthToken", "errCode", rsp.Code, "errMsg", rsp.Msg)
		grpc.SendHeader(ctx, metadata.Pairs("result", "0"))
		types.GrpcMatadataSetResultFail(ctx, rsp)
		return rsp, nil
	}
	rsp.Data = tokenInfo
	us.Logger.Info("UserAuth Success")
	//model.GetComponentMgr().Heartbeat(req.ClientId)
	return rsp, nil
}

// UserRefresh refresh token
func (us *GRPCUserServiceImpl) UserRefresh(ctx context.Context, req *pb.UserRefreshReq) (*pb.UserRefreshResp, error) {
	us.Logger.Info("UserRefresh", "req", req)
	rsp := &pb.UserRefreshResp{
		Code: errcode.ErrGrpcSuccess.Code(),
		Msg:  errcode.ErrGrpcSuccess.Error(),
	}
	//tokenSt := ctx.Value(types.RefreshTokenContextKey) // "RefreshToken"
	tokenSt := req.GetRefreshToken()
	//if tokenSt == nil {
	if tokenSt == "" {
		rsp.Code = errcode.ErrGrpcAuth.Code()
		rsp.Msg = errcode.ErrGrpcAuth.Error()
		us.Logger.Error("UserRefresh req ctx", "errCode", rsp.Code, "errMsg", rsp.Msg)
		types.GrpcMatadataSetResultFail(ctx, rsp)
		return rsp, nil
	}
	newTokenInfo, Err := us.doRefreshToken(ctx, tokenSt, req)
	if Err != nil {
		rsp.Code = Err.Code()
		rsp.Msg = Err.Error()
		us.Logger.Error("UserRefresh doRefreshToken", "errCode", rsp.Code, "errMsg", rsp.Msg)
		types.GrpcMatadataSetResultFail(ctx, rsp)
		return rsp, nil
	}
	rsp.Data = newTokenInfo
	us.Logger.Info("UserRefresh Success", "code", rsp.Code, "msg", rsp.Msg)
	// model.GetComponentMgr().Heartbeat(req.ClientId)
	return rsp, nil
}

// UserInfo user information
func (us *GRPCUserServiceImpl) UserInfo(ctx context.Context, req *pb.UserInfoReq) (*pb.UserInfoResp, error) {
	us.Logger.Info("UserInfo", "req", req)
	rsp := &pb.UserInfoResp{
		Code: errcode.ErrGrpcSuccess.Code(),
		Msg:  errcode.ErrGrpcSuccess.Error(),
	}
	data, Err := us.doUserInfo(ctx, req)
	if Err != nil {
		rsp.Code = Err.Code()
		rsp.Msg = Err.Error()
		us.Logger.Error("UserInfo doUserInfo", "errCode", rsp.Code, "errMsg", rsp.Msg)
		types.GrpcMatadataSetResultFail(ctx, rsp)
		return rsp, nil
	}
	rsp.Data = data
	us.Logger.Info("UserInfo Success")
	return rsp, nil
}

// ---------------------------------------------------------------
// internal basic function

// Auth Token
func (us *GRPCUserServiceImpl) doAuthToken(ctx context.Context, req *pb.UserAuthReq) (*pb.UserTokenObj, *errcode.Error) {
	id := req.GetId()
	pwd := req.GetPwd()
	//grantType := req.GetGrantType()
	if len(id) == 0 || len(pwd) == 0 {
		//|| grantType != types.GrantType_ClientCredentials {
		return nil, errcode.ErrGrpcAppExecute.AppendMsg("doAuthToken| id:%s,pwd:%s", id, pwd)
	}

	user := &table.User{}
	// todo 为访问mysql 验证登录信息
	exsit, err := us.MysqlSvc.GetEngineG().Where("id=?", id).Get(user)
	if err != nil {
		return nil, errcode.ErrGrpcSysMysqlErr.AppendMsg("doAuthToken| Mysql Get err: %v", err)
	} else if !exsit {
		return nil, errcode.ErrGrpcSysMysqlErr.AppendMsg("doAuthToken| user not exist")
	}
	tokenRsp, err := us.newAuthToken(ctx, id, pwd)
	if err != nil {
		return nil, errcode.ErrGrpcSysRedisErr.AppendMsg("doAuthToken| newAuthToken %v", err)
	}
	return tokenRsp, nil
}

func (us *GRPCUserServiceImpl) newAuthToken(ctx context.Context, id, pwd string) (*pb.UserTokenObj, error) {
	acTokenGen := base.GenToken(id, types.AccessTokenExpireDuration())
	refTokenGen := base.GenToken(id, types.RefreshTokenExpireDuration())
	//credkey := types.RedisCredentialsKey(id, pwd)
	userToken := fmt.Sprintf("%sF%s", acTokenGen, refTokenGen)
	refToken := fmt.Sprintf("%s", refTokenGen)
	if err := us.RedisSvc.Set(types.RedisAccessTokenKey(userToken), id, types.AccessTokenExpireDuration()); err != nil {
		return nil, err
	}
	if err := us.RedisSvc.Set(types.RedisRefreshTokenKey(refToken), id, types.RefreshTokenExpireDuration()); err != nil {
		return nil, err
	}
	return &pb.UserTokenObj{
		AccessToken:  userToken,
		RefreshToken: refToken,
		ExpireTime:   types.AccessTokenExpire - 2,
	}, nil
}

// Refresh Token
func (us *GRPCUserServiceImpl) doRefreshToken(ctx context.Context, tokenSt interface{}, req *pb.UserRefreshReq) (*pb.UserTokenObj, *errcode.Error) {
	refreshToken, ok := tokenSt.(string)
	if !ok {
		return nil, errcode.ErrGrpcAppExecute.AppendMsg("doRefreshToken| error on token usset")
	}
	// check refresh token
	//refreshToken := types.RedisGetRefreshTokenFromAccessToken(useressToken)
	// id check
	id, err := us.RedisSvc.Get(types.RedisRefreshTokenKey(refreshToken))
	if err != nil {
		return nil, errcode.ErrGrpcAppExecute.AppendMsg("doRefreshToken| error on refreshToken check, err: %s", err)
	}
	// refresh
	acTokenGen := base.GenToken(id, types.AccessTokenExpireDuration())
	newAccToken := fmt.Sprintf("%sF%s", acTokenGen, refreshToken)
	if err := us.RedisSvc.Set(types.RedisAccessTokenKey(newAccToken), id, types.AccessTokenExpireDuration()); err != nil {
		return nil, errcode.ErrGrpcAppExecute.AppendMsg("doRefreshToken| error on userToken set, err: %s", err)
	}
	return &pb.UserTokenObj{
		AccessToken:  newAccToken,
		RefreshToken: refreshToken,
		ExpireTime:   types.AccessTokenExpire - 2,
	}, nil

}

func (us *GRPCUserServiceImpl) doUserInfo(ctx context.Context, req *pb.UserInfoReq) (*pb.UserInfoObj, *errcode.Error) {
	userToken := req.GetAccessToken()
	id, err := us.RedisSvc.Get(types.RedisAccessTokenKey(userToken))
	if err != nil {
		return nil, errcode.ErrGrpcSysRedisErr.AppendMsg("UserInfo| Redis Get AccessToken err: %v", err)
	}
	user := &table.User{}
	exsit, err := us.MysqlSvc.GetEngineG().Where("id=?", id).Get(user)
	if err != nil {
		return nil, errcode.ErrGrpcSysMysqlErr.AppendMsg("UserInfo| Mysql Get err: %v", err)
	} else if !exsit {
		return nil, errcode.ErrGrpcSysMysqlErr.AppendMsg("UserInfo| Invalid? user not exist")
	}
	return &pb.UserInfoObj{
		Id:    id,
		Role:  user.Role,
		Desc:  user.Desc,
		Ctime: user.CreatedAt.String(), // todo panic
	}, nil
}
