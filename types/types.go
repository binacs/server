package types

const (
	AccessTokenExpire  = int64(1 * 60)
	RefreshTokenExpire = int64(1 * 60 * 60)

	TokenType_Bearer = "Bearer"

	TokenContextKey = "TokenInfo"
	AuthMethodName  = "MethodName"

	GrantType_ClientCredentials = "client_credentials"
	GrantType_RefreshToken      = "refresh_token"

	Size_RefreshToken = int(6)
	Size_AccessToken  = int(6)
)
