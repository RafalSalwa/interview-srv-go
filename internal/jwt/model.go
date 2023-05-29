package jwt

type Token struct {
	Username    string `json:"username"`
	TokenString string `json:"token"`
}
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Config struct {
}

type UserClaims struct {
	ID       int64
	Username string
}
