package rpc_api

import (
	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/intrvproto"
	"github.com/RafalSalwa/interview-app-srv/services"
)

type AuthServer struct {
	intrvproto.UnimplementedAuthServiceServer
	config      config.Config
	authService services.AuthService
	userService services.UserService
}

func NewGrpcAuthServer(config config.ConfGRPC, authService services.AuthService,
	userService services.UserService) (*AuthServer, error) {

	authServer := &AuthServer{
		config:      config,
		authService: authService,
		userService: userService,
	}

	return authServer, nil
}
