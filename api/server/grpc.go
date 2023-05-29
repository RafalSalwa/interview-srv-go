package server

import (
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/RafalSalwa/interview-app-srv/internal/rpc_api"

	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"

	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	pb.UnimplementedUserServiceServer
	config      config.ConfGRPC
	logger      *logger.Logger
	authService services.AuthService
	userService services.UserSqlService
}

func NewGrpcServer(config config.ConfGRPC, logger *logger.Logger, authService services.AuthService,
	userService services.UserSqlService) (*Server, error) {
	server := &Server{
		config:      config,
		logger:      logger,
		authService: authService,
		userService: userService,
	}

	return server, nil
}

func (server Server) Run() {
	grpcServer := grpc.NewServer()

	authServer, _ := rpc_api.NewGrpcAuthServer(server.config, server.authService, server.userService)
	userServer, _ := rpc_api.NewGrpcUserServer(server.config, server.userService)

	pb.RegisterAuthServiceServer(grpcServer, authServer)
	pb.RegisterUserServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	listener, err := net.Listen("tcp", server.config.GrpcServerAddress)
	if err != nil {
		server.logger.Error().Err(err)
	}
	server.logger.Info().Msgf("Starting gRPC server %s", server.config.GrpcServerAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		server.logger.Error().Err(err)
	}

	select {
	case <-shutdown:
		grpcServer.GracefulStop()
	}

}
