package auth

import (
	"errors"
	"net/http"

	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/config"
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

func NewAuthMethod(h apiHandler.AuthHandler, cfg *config.Config) (IAuthType, error) {
	val, ok := types[cfg.Auth.AuthMethod]
	if !ok {
		return nil, errors.New("wrong auth type")
	}
	switch val {
	case apiKey:
		return newApiKeyMiddleware(h, cfg.Auth.APIKey), nil
	case basic:
		return newBasicAuthMiddleware(h, cfg.Auth.BasicAuth.Username, cfg.Auth.BasicAuth.Password), nil
	case bearerToken:
		return newBearerTokenMiddleware(h, cfg.Auth.BearerToken), nil
	}
	return nil, nil
}
