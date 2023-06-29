package auth

import (
	"errors"
	"github.com/RafalSalwa/interview-app-srv/config"
	simpleHandler "github.com/RafalSalwa/interview-app-srv/pkg/simple_handler"
	"net/http"
)

type IAuthType interface {
	middleware(h simpleHandler.HandlerFunc) http.HandlerFunc
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

func NewAuthMethod(h simpleHandler.HandlerFunc, method string) (IAuthType, error) {
	val, ok := types[method]
	if !ok {
		return nil, errors.New("wrong auth type")
	}
	c := config.New()
	switch val {
	case apiKey:
		return newApiKeyMiddleware(h, c.Server.APIKey), nil
	case basic:
		return newBasicAuthMiddleware(h, c.Server.BasicAuth.Username, c.Server.BasicAuth.Password), nil
	case bearerToken:
		return newBearerTokenMiddleware(h, c.Server.BearerToken), nil
	}
	return nil, nil
}
