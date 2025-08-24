package service

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/binacsgo/log"

	pb "github.com/binacs/server/api/crypto"
	"github.com/binacs/server/config"
)

// CryptoServiceImpl Web crypto service implement
type CryptoServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"WebLogger"`
	// 这里不同加密算法的客户端可以合并(把类型作为参数) 但为了显示不同服务将其单独列出
	clientBASE64 pb.CryptoClient
	clientAES    pb.CryptoClient
	clientDES    pb.CryptoClient
}

// AfterInject do inject
func (cs *CryptoServiceImpl) AfterInject() error {
	go cs.connect()
	return nil
}

// Register the service
func (cs *CryptoServiceImpl) Register(ctx context.Context, gsrv *grpc.Server, gwmux *runtime.ServeMux, opts []grpc.DialOption) error {
	if gsrv != nil {
		pb.RegisterCryptoServer(gsrv, cs)
	}

	if gwmux != nil {
		if err := pb.RegisterCryptoHandlerFromEndpoint(ctx, gwmux, ":"+cs.Config.GRPCConfig.HttpPort, opts); err != nil {
			return fmt.Errorf("RegisterUserHandlerFromEndpoint: %v", err)
		}
	}
	return nil
}

func (cs *CryptoServiceImpl) connect() {
	for svc, addr := range cs.Config.WebConfig.K8sService {
		conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			cs.Logger.Error("CryptoServiceImpl: Dial", "svc", svc, "err", err)
			continue
		}
		cli := pb.NewCryptoClient(conn)
		switch svc {
		case "CryptoBASE64":
			cs.clientBASE64 = cli
		case "CryptoAES":
			cs.clientAES = cli
		case "CryptoDES":
			cs.clientDES = cli
		default:
			cs.Logger.Error("CryptoServiceImpl: Client svc type error", "svc", svc, "err", err)
		}
	}
}

// CryptoEncrypt encrypt
func (cs *CryptoServiceImpl) CryptoEncrypt(ctx context.Context, req *pb.CryptoEncryptReq) (*pb.CryptoEncryptResp, error) {
	cs.Logger.Info("CryptoServiceImpl CryptoEncrypt Start", "algorithm", req.GetAlgorithm(), "plainText", req.GetPlainText())
	var client pb.CryptoClient
	switch req.GetAlgorithm() {
	case "BASE64":
		client = cs.clientBASE64
	case "AES":
		client = cs.clientAES
	case "DES":
		client = cs.clientDES
	default:
		cs.Logger.Error("CryptoEncrypt: NOT BASE64/AES/DES")
		return nil, nil
	}
	if client == nil {
		cs.Logger.Error("CryptoEncrypt client nil")
		// TODO return nil for now
		return nil, nil
	}
	return client.CryptoEncrypt(context.Background(), req)
}

// CryptoDecrypt decrypt
func (cs *CryptoServiceImpl) CryptoDecrypt(ctx context.Context, req *pb.CryptoDecryptReq) (*pb.CryptoDecryptResp, error) {
	cs.Logger.Info("CryptoServiceImpl CryptoEncrypt Start", "algorithm", req.GetAlgorithm(), "encryptText", req.GetEncryptText())
	var client pb.CryptoClient
	switch req.GetAlgorithm() {
	case "BASE64":
		client = cs.clientBASE64
	case "AES":
		client = cs.clientAES
	case "DES":
		client = cs.clientDES
	default:
		cs.Logger.Error("CryptoEncrypt: NOT BASE64/AES/DES")
		return nil, nil
	}
	if client == nil {
		cs.Logger.Error("CryptoEncrypt client nil")
		// TODO return nil for now
		return nil, nil
	}
	return client.CryptoDecrypt(context.Background(), req)
}
