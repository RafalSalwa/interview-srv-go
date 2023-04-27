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
	userService services.UserService
}

func NewGrpcServer(config config.ConfGRPC, authService services.AuthService,
	userService services.UserService, userCollection *mongo.Collection) (*Server, error) {

	server := &Server{
		config:         config,
		authService:    authService,
		userService:    userService,
		userCollection: userCollection,
	}

	return server, nil
}

func startGrpcServer(config config.Config) {
	server, err := gapi.NewGrpcServer(config, authService, userService, authCollection)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}
}
