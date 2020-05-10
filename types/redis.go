package types

import (
	"fmt"
	"strings"
	"time"
)

const (
	TokenSalt          = "Binacs_Token_Salt" // TokenSalt used in user create
	AccessTokenExpire  = int64(1 * 60)       // AccessTokenExpire time
	RefreshTokenExpire = int64(1 * 24 * 60)  // RefreshTokenExpire time

	TokenType_Bearer = "Bearer" // TokenType_Bearer

	AccessTokenContextKey  = "access_token"  // AccessTokenContextKey context key
	RefreshTokenContextKey = "refresh_token" // RefreshTokenContextKey context key
	AuthMethodName         = "MethodName"    // AuthMethodName

	//GrantType_ClientCredentials = "client_credentials"
	//GrantType_RefreshToken      = "refresh_token"

	Prefix_CredentialKey   = "/SVR/USR/CRED/%s/%s"   // Prefix_CredentialKey used in db service
	Prefix_AccessTokenKey  = "/SVR/USR/TOKEN/ACC/%s" // Prefix_AccessTokenKey used in db service
	Prefix_RefreshTokenKey = "/SVR/USR/TOKEN/REF/%s" // Prefix_RefreshTokenKey used in db service

	Size_RefreshToken = int(6) // Size_RefreshToken
	Size_AccessToken  = int(6) // Size_AccessToken
)

// AccessTokenExpireDuration return the access token expire duration (time.Duration)
func AccessTokenExpireDuration() time.Duration {
	return time.Duration(AccessTokenExpire) * time.Minute
}

// RefreshTokenExpireDuration return the refresh token expire duration (time.Duration)
func RefreshTokenExpireDuration() time.Duration {
	return time.Duration(RefreshTokenExpire) * time.Minute
}

// RedisGetRefreshTokenFromAccessToken return the access token by refresh token
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

// RedisAccessTokenKey return the access token key
func RedisAccessTokenKey(accessToken string) string {
	return fmt.Sprintf(Prefix_AccessTokenKey, accessToken)
}

// RedisRefreshTokenKey return the refresh token key
func RedisRefreshTokenKey(refreshToken string) string {
	return fmt.Sprintf(Prefix_RefreshTokenKey, refreshToken)
}

// RedisCredentialsKey return the credentials key
func RedisCredentialsKey(id, pwd string) string {
	return fmt.Sprintf(Prefix_CredentialKey, id, pwd)
}
