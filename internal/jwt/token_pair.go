package jwt

import (
	"github.com/RafalSalwa/interview-app-srv/config"
)

func GenerateTokenPair(c config.ConfToken, uid int64, username string) (*TokenPair, error) {

	accessClaims := UserClaims{
		ID:       uid,
		Username: username,
	}

	t, err := CreateToken(c.AccessTokenExpiresIn, accessClaims, c.AccessTokenPrivateKey)
	if err != nil {
		return nil, err
	}

	refreshClaims := UserClaims{
		ID: uid,
	}
	rt, err := CreateToken(c.RefreshTokenExpiresIn, refreshClaims, c.RefreshTokenPrivateKey)
	if err != nil {
		return nil, err
	}

	tp := &TokenPair{AccessToken: t, RefreshToken: rt}
	return tp, nil
}
