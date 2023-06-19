package jwt

func GenerateTokenPair(c JWTConfig, uid int64) (*TokenPair, error) {

	accessClaims := UserClaims{
		ID: uid,
	}

	t, err := CreateToken(c.Access.ExpiresIn, accessClaims, c.Access.PrivateKey)
	if err != nil {
		return nil, err
	}

	refreshClaims := UserClaims{
		ID: uid,
	}
	rt, err := CreateToken(c.Refresh.ExpiresIn, refreshClaims, c.Refresh.PrivateKey)
	if err != nil {
		return nil, err
	}

	tp := &TokenPair{AccessToken: t, RefreshToken: rt}
	return tp, nil
}
