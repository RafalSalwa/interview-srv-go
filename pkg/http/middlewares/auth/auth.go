package auth

import (
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/handler"
	"github.com/RafalSalwa/interview-app-srv/pkg/auth"
)

func NewAuthorizer(h handler.AuthHandler, cfg *auth.Auth) (IAuthType, error) {
	at, err := NewAuthMethod(h, cfg)
	if err != nil {
		return nil, err
	}
	return at, nil
}
