package auth

import (
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/handler"
)

func NewAuthorizer(h handler.AuthHandler, cfg *config.Config) (IAuthType, error) {
	at, err := NewAuthMethod(h, cfg)
	if err != nil {
		return nil, err
	}
	return at, nil
}
