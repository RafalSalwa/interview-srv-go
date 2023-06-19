package rpc_api

import (
	"github.com/RafalSalwa/interview-app-srv/internal/services"
	readerGRPC "github.com/RafalSalwa/interview-app-srv/pkg/grpc"
	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	config      readerGRPC.Config
	authService services.AuthService
}

func NewGrpcAuthServer(config readerGRPC.Config, authService services.AuthService) (*AuthServer, error) {
	authServer := &AuthServer{
		config:      config,
		authService: authService,
	}

	return authServer, nil
}
