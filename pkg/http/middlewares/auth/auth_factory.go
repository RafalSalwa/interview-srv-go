package auth

import (
	"errors"
	"github.com/RafalSalwa/interview-app-srv/pkg/auth"
	"net/http"

	apiHandler "github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/handler"
)

type IAuthType interface {
	Middleware(h apiHandler.HandlerFunc) http.HandlerFunc
}

type AuthType int

const (
	apiKey AuthType = 1 << iota
	basic
	bearerToken
)

var types = map[string]interface{}{
	"key":          apiKey,
	"basic":        basic,
	"bearer_token": bearerToken,
}

func NewAuthMethod(h apiHandler.AuthHandler, cfg *auth.Auth) (IAuthType, error) {
	val, ok := types[cfg.AuthMethod]
	if !ok {
		return nil, errors.New("wrong auth type")
	}
	switch val {
	case apiKey:
		return newApiKeyMiddleware(h, cfg.APIKey), nil
	case basic:
		return newBasicAuthMiddleware(h, cfg.BasicAuth.Username, cfg.BasicAuth.Password), nil
	case bearerToken:
		return newBearerTokenMiddleware(h, cfg.BearerToken), nil
	}
	return nil, nil
}
