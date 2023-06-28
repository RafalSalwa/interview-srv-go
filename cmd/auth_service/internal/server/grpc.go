package server

import (
	"github.com/RafalSalwa/interview-app-srv/cmd/auth_service/internal/rpc_api"
	"github.com/RafalSalwa/interview-app-srv/cmd/auth_service/internal/services"
	grpcconfig "github.com/RafalSalwa/interview-app-srv/pkg/grpc"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpclogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type GRPC struct {
	pb.UnimplementedAuthServiceServer
	config      grpcconfig.Config
	logger      *logger.Logger
	authService services.AuthService
}

func NewGrpcServer(config grpcconfig.Config,
	logger *logger.Logger,
	authService services.AuthService) (*GRPC, error) {

	srv := &GRPC{
		config:      config,
		logger:      logger,
		authService: authService,
	}

	return srv, nil
}

func (s GRPC) Run() {
	logEntry := logger.NewGRPCLogger()
	grpclogrus.ReplaceGrpcLogger(logEntry)

	opts := []grpclogrus.Option{
		grpclogrus.WithLevels(func(code codes.Code) logrus.Level {
			if code == codes.OK {
				return logrus.InfoLevel
			}
			return logrus.ErrorLevel
		}),

		grpclogrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ms", duration.Milliseconds()
		}),
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
			grpcctxtags.StreamServerInterceptor(),
			otelgrpc.StreamServerInterceptor(),
			grpclogrus.StreamServerInterceptor(logEntry, opts...),
			grpcrecovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
			grpcctxtags.UnaryServerInterceptor(),
			otelgrpc.UnaryServerInterceptor(),
			grpclogrus.UnaryServerInterceptor(logEntry, opts...),
			grpcrecovery.UnaryServerInterceptor(),
		)),
	)

	authServer, err := rpc_api.NewGrpcAuthServer(s.config, s.logger, s.authService)
	if err != nil {
		s.logger.Error().Err(err)
	}

	pb.RegisterAuthServiceServer(grpcServer, authServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", s.config.Addr)
	if err != nil {
		s.logger.Error().Err(err)
	}

	s.logger.Info().Msgf("Starting gRPC server on: %s", s.config.Addr)
	err = grpcServer.Serve(listener)
	if err != nil {
		s.logger.Error().Err(err)
	}
	grpcServer.GracefulStop()
}
