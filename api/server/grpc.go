package server

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/sirupsen/logrus"

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
	var logrusLogger *logrus.Logger
	var customFunc grpc_logrus.CodeToLevel

	logrusEntry := logrus.NewEntry(logrusLogger)
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)
	opts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(customFunc),
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(logrusEntry, opts...),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(logrusEntry, opts...),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

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

	<-shutdown
	grpcServer.GracefulStop()
}
