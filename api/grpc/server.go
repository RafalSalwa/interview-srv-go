package grpc

import (
	"github.com/RafalSalwa/interview-app-srv/config"
	pb "github.com/RafalSalwa/interview-app-srv/grpc"
	"github.com/RafalSalwa/interview-app-srv/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
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

func (server Server) Run() {

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, _ := net.Listen("tcp", server.config.GrpcServerAddress)

	_ = grpcServer.Serve(listener)

}
