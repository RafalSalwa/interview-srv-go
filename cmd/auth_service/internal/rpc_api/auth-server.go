package rpc_api

import (
	"github.com/RafalSalwa/interview-app-srv/cmd/auth_service/internal/services"
	grpcconfig "github.com/RafalSalwa/interview-app-srv/pkg/grpc"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	config      grpcconfig.Config
	logger      *logger.Logger
	authService services.AuthService
}

func NewGrpcAuthServer(config grpcconfig.Config, logger *logger.Logger, authService services.AuthService) (*AuthServer, error) {
	authServer := &AuthServer{
		config:      config,
		logger:      logger,
		authService: authService,
	}

	return authServer, nil
}
