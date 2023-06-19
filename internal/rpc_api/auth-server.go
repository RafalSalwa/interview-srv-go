package rpc_api

import (
	"github.com/RafalSalwa/interview-app-srv/internal/services"
	grpc_config "github.com/RafalSalwa/interview-app-srv/pkg/grpc"
	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	config      grpc_config.Config
	authService services.AuthService
}

func NewGrpcAuthServer(config grpc_config.Config, authService services.AuthService) (*AuthServer, error) {
	authServer := &AuthServer{
		config:      config,
		authService: authService,
	}

	return authServer, nil
}
