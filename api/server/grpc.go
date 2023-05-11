package server

import (
	"net"

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
	authService services.AuthService
	userService services.UserSqlService
}

func NewGrpcServer(config config.ConfGRPC, authService services.AuthService,
	userService services.UserSqlService) (*Server, error) {

	server := &Server{
		config:      config,
		authService: authService,
		userService: userService,
	}

	return server, nil
}

func (server Server) Run() error {

	grpcServer := grpc.NewServer()

	authServer, err := rpc_api.NewGrpcAuthServer(server.config, server.authService, server.userService)
	if err != nil {
		return err
	}

	userServer, err := rpc_api.NewGrpcUserServer(server.config, server.userService)
	if err != nil {
		return err
	}

	pb.RegisterAuthServiceServer(grpcServer, authServer)
	pb.RegisterUserServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	listener, _ := net.Listen("tcp", server.config.GrpcServerAddress)

	_ = grpcServer.Serve(listener)

	return nil
}
