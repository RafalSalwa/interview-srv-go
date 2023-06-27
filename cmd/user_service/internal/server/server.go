package server

import (
	"context"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/cmd/user_service/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/user_service/internal/services"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	log         *logger.Logger
	cfg         *config.Config
	v           *validator.Validate
	mongoClient *mongo.Client
	redisClient redis.UniversalClient
}

func NewServerGRPC(cfg *config.Config, log *logger.Logger) *server {
	return &server{log: log, cfg: cfg, v: validator.New()}
}

func (srv *server) Run() error {
	ctx, rejectContext := context.WithCancel(NewContextCancellableByOsSignals(context.Background()))

	userService := services.NewUserService(ctx, srv.cfg, srv.log)

	grpcServer, err := NewGrpcServer(srv.cfg.GRPC, srv.log, &userService)
	if err != nil {
		srv.log.Error().Err(err)
	}
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		grpcServer.Run()
	}()
	<-shutdown
	rejectContext()
	return nil
}

func NewContextCancellableByOsSignals(parent context.Context) context.Context {
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	newCtx, cancel := context.WithCancel(parent)

	go func() {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			fmt.Println("Received Interrupt signal")
			cancel()
		case syscall.SIGTERM:
			fmt.Println("Received sigterm signal")
			cancel()
		}
	}()

	return newCtx
}
