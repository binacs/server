package service

import (
	"context"

	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"

	//pb "github.com/BinacsLee/server/api/crypto/v1"
	pb "github.com/BinacsLee/Cryptology/api/cryptofunc"
	"github.com/BinacsLee/server/config"
	//"github.com/BinacsLee/server/libs/errcode"
	"github.com/BinacsLee/server/libs/log"
	//"github.com/BinacsLee/server/service/db"
)

// WebCryptoServiceImpl Web crypto service implement
type WebCryptoServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"WebLogger"`
	// 这里不同加密算法的客户端可以合并(把类型作为参数) 但为了显示不同服务将其单独列出
	clientBASE64 pb.CryptoFuncClient
	clientAES    pb.CryptoFuncClient
	clientDES    pb.CryptoFuncClient
}

// AfterInject do inject
func (cs *WebCryptoServiceImpl) AfterInject() error {
	if addr, ok := cs.Config.WebConfig.K8sService["CryptoBASE64"]; ok {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			return err
		}
		cs.clientBASE64 = pb.NewCryptoFuncClient(conn)
	}
	if addr, ok := cs.Config.WebConfig.K8sService["CryptoAES"]; ok {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			return err
		}
		cs.clientAES = pb.NewCryptoFuncClient(conn)
	}
	if addr, ok := cs.Config.WebConfig.K8sService["CryptoDES"]; ok {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			return err
		}
		cs.clientDES = pb.NewCryptoFuncClient(conn)
	}
	return nil
}

// CryptoEncrypt encrypt
func (cs *WebCryptoServiceImpl) CryptoEncrypt(ctx *gin.Context, text, tp string) (string, error) {
	cs.Logger.Info("WebCryptoServiceImpl CryptoEncrypt Start", "text", text, "tpye", tp)
	var client pb.CryptoFuncClient
	switch tp {
	case "BASE64":
		client = cs.clientBASE64
	case "AES":
		client = cs.clientAES
	case "DES":
		client = cs.clientDES
	default:
		cs.Logger.Error("CryptoEncrypt: NOT BASE64/AES/DES", "type", tp)
		return "TypeNotSupport", nil
	}

	req := &pb.EncryptReq{
		Src: text,
	}
	resp, err := client.Encrypt(context.Background(), req)
	if err != nil {
		cs.Logger.Error("CryptoEncrypt client.Encrypt", "err", err)
		return "client.Encrypt Error", err
	}
	return resp.GetRes(), nil
}

// CryptoDecrypt decrypt
func (cs *WebCryptoServiceImpl) CryptoDecrypt(ctx *gin.Context, text, tp string) (string, error) {
	cs.Logger.Info("WebCryptoServiceImpl CryptoDecrypt Start", "text", text, "tpye", tp)
	var client pb.CryptoFuncClient
	switch tp {
	case "BASE64":
		client = cs.clientBASE64
	case "AES":
		client = cs.clientAES
	case "DES":
		client = cs.clientDES
	default:
		cs.Logger.Error("CryptoDecrypt: NOT BASE64/AES/DES", "type", tp)
		return "TypeNotSupport", nil
	}

	req := &pb.DecryptReq{
		Src: text,
	}
	resp, err := client.Decrypt(context.Background(), req)
	if err != nil {
		cs.Logger.Error("CryptoDecrypt client.Encrypt", "err", err)
		return "client.Decrypt Error", err
	}
	return resp.GetRes(), nil
}
