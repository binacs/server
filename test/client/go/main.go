package main

import (
	"context"
	"crypto/tls"

	//"bytes"
	"log"
	//"encoding/json"
	"net/http"

	"github.com/BinacsLee/server/test/client/go/swagger"
)

var (
	cli *swagger.APIClient
	ctx context.Context
)

func main() {
	// skip verify
	ts := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpCli := &http.Client{Transport: ts}
	cfg := &swagger.Configuration{
		BasePath:      "https://localhost:80",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
		HTTPClient:    httpCli,
	}
	cli = swagger.NewAPIClient(cfg)
	ctx = context.Background()
	// set ctx
	ctx = context.WithValue(ctx, swagger.ContextAccessToken, "aaaaaaaa")

	// test APIs
	//testRegister("swagger_test_reg_id", "swagger_test_reg_pwd")
	//testAuth("swagger_test_reg_id", "swagger_test_reg_pwd")
	//testRefresh("15867690970949d8595d7c76d759a5c85b8d9b7c51idswagger_test_reg_id")
	testInfo("1586686656c2c874b6a3089a5032208fdcd4d1d996idswagger_test_reg_idF15867690970949d8595d7c76d759a5c85b8d9b7c51idswagger_test_reg_id")

	return
}

func testTest() {
	// UserTest
	log.Printf("========== UserTest ==========")
	req := swagger.BinacsApiUserV1UserTestReq{}
	resp, httpResp, err := cli.UserApi.UserTest(ctx, req)
	if err != nil {
		log.Printf("UserTest err = %s", err)
	}
	log.Printf("UserTest resp = %s", resp)
	log.Printf("UserTest httpResp = %+v", httpResp)
}

func testRegister(id, pwd string) {
	log.Printf("========== UserRegister ==========")
	req := swagger.BinacsApiUserV1UserRegisterReq{
		Id:  id,
		Pwd: pwd,
	}
	resp, httpResp, err := cli.UserApi.UserRegister(ctx, req)
	if err != nil {
		log.Printf("UserTest err = %s", err)
	}
	log.Printf("UserTest resp = %s", resp)
	log.Printf("UserTest httpResp = %+v", httpResp)
}

func testAuth(id, pwd string) {
	log.Printf("========== UserAuth ==========")
	req := swagger.BinacsApiUserV1UserAuthReq{
		Id:  id,
		Pwd: pwd,
	}
	resp, httpResp, err := cli.UserApi.UserAuth(ctx, req)
	if err != nil {
		log.Printf("UserAuth err = %s", err)
	}
	log.Printf("UserAuth resp = %s", resp)
	log.Printf("UserAuth httpResp = %+v", httpResp)
}

func testRefresh(refToken string) {
	log.Printf("========== UserRefresh ==========")
	req := swagger.BinacsApiUserV1UserRefreshReq{
		RefreshToken: refToken,
	}
	resp, httpResp, err := cli.UserApi.UserRefresh(ctx, req)
	if err != nil {
		log.Printf("UserRefresh err = %s", err)
	}
	log.Printf("UserRefresh resp = %s", resp)
	log.Printf("UserRefresh httpResp = %+v", httpResp)
}

func testInfo(accToken string) {
	log.Printf("========== UserInfo ==========")
	req := swagger.BinacsApiUserV1UserInfoReq{
		AccessToken: accToken,
	}
	resp, httpResp, err := cli.UserApi.UserInfo(ctx, req)
	if err != nil {
		log.Printf("UserInfo err = %s", err)
	}
	log.Printf("UserInfo resp = %s", resp)
	log.Printf("UserInfo httpResp = %+v", httpResp)
}
