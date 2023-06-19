package server

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc_config "github.com/RafalSalwa/interview-app-srv/pkg/grpc"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/rpc_api"
	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"

	"github.com/RafalSalwa/interview-app-srv/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	pb.UnimplementedUserServiceServer
	config      grpc_config.Config
	logger      *logger.Logger
	authService services.AuthService
	userService services.UserSqlService
}

func NewGrpcServer(config grpc_config.Config,
	logger *logger.Logger,
	authService services.AuthService,
	userService services.UserSqlService) (*Server, error) {

	server := &Server{
		config:      config,
		logger:      logger,
		authService: authService,
		userService: userService,
	}

	return server, nil
}

func (srv Server) Run() {
	logEntry := logger.NewGRPCLogger()
	grpc_logrus.ReplaceGrpcLogger(logEntry)

	opts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(func(code codes.Code) logrus.Level {
			if code == codes.OK {
				return logrus.InfoLevel
			}
			return logrus.ErrorLevel
		}),

		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ms", duration.Milliseconds()
		}),
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(logEntry, opts...),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(logEntry, opts...),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	authServer, err := rpc_api.NewGrpcAuthServer(srv.config, srv.authService)
	if err != nil {
		srv.logger.Error().Err(err)
	}
	userServer, err := rpc_api.NewGrpcUserServer(srv.config, srv.userService)
	if err != nil {
		srv.logger.Error().Err(err)
	}

	pb.RegisterAuthServiceServer(grpcServer, authServer)
	pb.RegisterUserServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	listener, err := net.Listen("tcp", srv.config.Addr)
	if err != nil {
		srv.logger.Error().Err(err)
	}

	srv.logger.Info().Msgf("Starting gRPC server on: %s", srv.config.Addr)
	err = grpcServer.Serve(listener)
	if err != nil {
		srv.logger.Error().Err(err)
	}

	<-shutdown
	grpcServer.GracefulStop()
}
