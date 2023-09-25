package server

import (
	"context"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/pkg/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"os"
	"os/signal"
	"syscall"

	"github.com/RafalSalwa/interview-app-srv/cmd/auth_service/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/auth_service/internal/services"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
)

type Server struct {
	log *logger.Logger
	cfg *config.Config
}

func NewGRPC(cfg *config.Config, log *logger.Logger) *Server {
	return &Server{log: log, cfg: cfg}
}

func (srv *Server) Run() error {
	ctx, rejectContext := context.WithCancel(NewContextCancellableByOsSignals(context.Background()))

	authService := services.NewAuthService(ctx, srv.cfg, srv.log)
	s := NewGrpcServer(srv.cfg.GRPC, srv.log, srv.cfg.Probes, authService)

	go func() {
		s.Run()
	}()

	if srv.cfg.Jaeger.Enable {
		tp, err := tracing.NewJaegerTracer(srv.cfg.Jaeger)
		if err != nil {
			srv.log.Error().Err(err).Msg("Auth:jaeger:register")
		}
		otel.SetTracerProvider(tp)
		otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

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
