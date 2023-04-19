package jwt

type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Username    string `json:"username"`
	TokenString string `json:"token"`
}
