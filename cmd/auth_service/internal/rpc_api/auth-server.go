package rpc_api

import (
	"github.com/RafalSalwa/interview-app-srv/cmd/auth_service/internal/services"
	grpcconfig "github.com/RafalSalwa/interview-app-srv/pkg/grpc"
	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	config      grpcconfig.Config
	authService services.AuthService
}

func NewGrpcAuthServer(config grpcconfig.Config, authService services.AuthService) (*AuthServer, error) {
	authServer := &AuthServer{
		config:      config,
		authService: authService,
	}

	return authServer, nil
}
