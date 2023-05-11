package rpc_api

import (
	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/internal/services"
	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	config      config.ConfGRPC
	authService services.AuthService
	userService services.UserSqlService
}

func NewGrpcAuthServer(config config.ConfGRPC, authService services.AuthService,
	userService services.UserSqlService) (*AuthServer, error) {

	authServer := &AuthServer{
		config:      config,
		authService: authService,
		userService: userService,
	}

	return authServer, nil
}
