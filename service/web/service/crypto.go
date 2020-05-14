package service

import (
	"fmt"
	"github.com/gin-gonic/gin"

	pb "github.com/BinacsLee/server/api/crypto/v1"
	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/errcode"
	"github.com/BinacsLee/server/libs/log"
	"github.com/BinacsLee/server/service/db"
)
// WebCryptoServiceImpl Web crypto service implement
type WebCryptoServiceImpl struct {
	Config   *config.Config  `inject-name:"Config"`
	Logger   log.Logger      `inject-name:"WebLogger"`
	RedisSvc db.RedisService `inject-name:"RedisService"`
	MysqlSvc db.MysqlService `inject-name:"MysqlService"`
}

// CryptoEncrypt encrypt
func (cs *WebCryptoServiceImpl) CryptoEncrypt(ctx *gin.Context, text, tp string) (*pb.CryptoEncryptResp, error) {
	cs.Logger.Info("WebCryptoServiceImpl CryptoEncrypt Start", "text", text, "tpye", tp)
	// req useless for now
	req := &pb.CryptoEncryptReq{
		Algorithm: tp,
		PlainText: text,
	}
	rsp := &pb.CryptoEncryptResp{
		Code: errcode.ErrGrpcSuccess.Code(),
		Msg:  errcode.ErrGrpcSuccess.Error(),
	}
	data, appErr := cs.doCryptoEncrypt(ctx, req)
	if appErr != nil {
		rsp.Code = appErr.Code()
		rsp.Msg = appErr.Error()
		cs.Logger.Error("CryptoEncrypt error", "errCode", rsp.Code, "errMsg", rsp.Msg)
		return rsp, nil
	}
	rsp.Data = data
	return rsp, nil
}

// CryptoDecrypt decrypt
func (cs *WebCryptoServiceImpl) CryptoDecrypt(ctx *gin.Context, text, tp string) (*pb.CryptoDecryptResp, error) {
	cs.Logger.Info("WebCryptoServiceImpl CryptoDecrypt Start", "text", text, "tpye", tp)
	// req useless for now
	req := &pb.CryptoDecryptReq{
		Algorithm:   tp,
		EncryptText: text,
	}
	rsp := &pb.CryptoDecryptResp{
		Code: errcode.ErrGrpcSuccess.Code(),
		Msg:  errcode.ErrGrpcSuccess.Error(),
	}
	data, appErr := cs.doCryptoDecrypt(ctx, req)
	if appErr != nil {
		rsp.Code = appErr.Code()
		rsp.Msg = appErr.Error()
		cs.Logger.Error("WebCryptoServiceImpl CryptoDecrypt error", "errCode", rsp.Code, "errMsg", rsp.Msg)
		return rsp, nil
	}
	rsp.Data = data
	return rsp, nil
}

func (cs *WebCryptoServiceImpl) doCryptoEncrypt(ctx *gin.Context, req *pb.CryptoEncryptReq) (*pb.CryptoEncryptResObj, *errcode.Error) {
	algs := req.GetAlgorithm()
	plainText := req.GetPlainText()
	fmt.Println("algs=",algs, " plainText=",plainText)
	var encryptText string
	switch algs {
	case "BASE64":
	case "AES":
	case "DES":
	case "RSA":
	default:
		cs.Logger.Error("doCryptoEncrypt: NOT BASE64/AES/DES/RSA", "type", algs)
		return nil, errcode.ErrGrpcAppExecute.AppendMsg("doCryptoEncrypt: NOT AES/RSA/ECC/SMx, type: %s", algs)
	}
	cs.Logger.Info("doCryptoEncrypt: Encrypt succeed")
	return &pb.CryptoEncryptResObj{
		EncryptText: encryptText,
	}, nil
}

func (cs *WebCryptoServiceImpl) doCryptoDecrypt(ctx *gin.Context, req *pb.CryptoDecryptReq) (*pb.CryptoDecryptResObj, *errcode.Error) {
	algs := req.GetAlgorithm()
	encryptText := req.GetEncryptText()
	fmt.Println("algs=",algs, " encryptText=",encryptText)
	var plainText string
	switch algs {
	case "BASE64":
	case "AES":
	case "DES":
	case "RSA":
	default:
		cs.Logger.Error("doCryptoDecrypt: NOT BASE64/AES/DES/RSA", "type", algs)
		return nil, errcode.ErrGrpcAppExecute.AppendMsg("doCryptoDecrypt: NOT AES/RSA/ECC/SMx, type: %s", algs)
	}
	cs.Logger.Info("doCryptoDecrypt: Encrypt succeed")
	return &pb.CryptoDecryptResObj{
		PlainText: plainText,
	}, nil
}
