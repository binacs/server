package types

import (
	"fmt"
	"strings"
	"time"
)

const (
	TokenSalt          = "Binacs_Token_Salt"
	AccessTokenExpire  = int64(1 * 60)
	RefreshTokenExpire = int64(1 * 24 * 60)

	TokenType_Bearer = "Bearer"

	AccessTokenContextKey  = "access_token"
	RefreshTokenContextKey = "refresh_token"
	AuthMethodName         = "MethodName"

	//GrantType_ClientCredentials = "client_credentials"
	//GrantType_RefreshToken      = "refresh_token"

	Prefix_CredentialKey   = "/SVR/USR/CRED/%s/%s"
	Prefix_AccessTokenKey  = "/SVR/USR/TOKEN/ACC/%s"
	Prefix_RefreshTokenKey = "/SVR/USR/TOKEN/REF/%s"

	Size_RefreshToken = int(6)
	Size_AccessToken  = int(6)
)

func AccessTokenExpireDuration() time.Duration {
	return time.Duration(AccessTokenExpire) * time.Minute
}

func RefreshTokenExpireDuration() time.Duration {
	return time.Duration(RefreshTokenExpire) * time.Minute
}

func RedisGetRefreshTokenFromAccessToken(key string) string {
	if len(key) == 0 {
		return ""
	}
	slic := strings.Split(key, "F")
	if len(slic) != 2 {
		return ""
	}
	return slic[0]
}

func RedisAccessTokenKey(accessToken string) string {
	return fmt.Sprintf(Prefix_AccessTokenKey, accessToken)
}

func RedisRefreshTokenKey(refreshToken string) string {
	return fmt.Sprintf(Prefix_RefreshTokenKey, refreshToken)
}

func RedisCredentialsKey(id, pwd string) string {
	return fmt.Sprintf(Prefix_CredentialKey, id, pwd)
}
