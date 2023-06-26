package server

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/cqrs"
	"go.opentelemetry.io/otel"
	"time"

	gatewayConfig "github.com/RafalSalwa/interview-app-srv/cmd/gateway/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/handler"
	apiRouter "github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/router"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/tracing"
	"github.com/gorilla/mux"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

type REST struct {
	srv         *http.Server
	userHandler handler.UserHandler
	authHandler handler.AuthHandler
	router      *mux.Router
	cqrs        *cqrs.Application
	log         *logger.Logger
	cfg         *gatewayConfig.Config
}

func NewRESTServer(c *gatewayConfig.Config, l *logger.Logger) *REST {
	tlsConf := new(tls.Config)
	r := apiRouter.NewApiRouter(c, l)

	s := &http.Server{
		Addr:         c.Http.Addr,
		Handler:      r,
		ReadTimeout:  c.Http.TimeoutRead,
		WriteTimeout: c.Http.TimeoutWrite,
		IdleTimeout:  c.Http.TimeoutIdle,
		TLSConfig:    tlsConf,
	}

	return &REST{
		srv:    s,
		router: r,
		log:    l,
		cfg:    c,
	}
}

func (s *REST) Run(ctx context.Context) {
	s.SetupCQRS(ctx)
	s.SetupHandlers()
	s.SetupRoutes()
	go func() {
		s.log.Info().Msgf("Starting REST server on: %v", s.srv.Addr)
		if s.cfg.App.Env == "dev" {
			if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				s.log.Error().Err(err).Msg("REST:Listen")
			}
		} else {
			if err := s.srv.ListenAndServeTLS(
				"/etc/ssl/certs/server.crt",
				"/etc/ssl/private/server.key"); err != nil && !errors.Is(err, http.ErrServerClosed) {
				s.log.Error().Err(err).Msg("REST:ListenTLS")
			}
		}
	}()

	if s.cfg.Jaeger.Enable {
		tp, err := tracing.NewJaegerTracer(*s.cfg.Jaeger)
		if err != nil {
			s.log.Error().Err(err).Msg("REST:jaeger:register")
		}
		otel.SetTracerProvider(tp)

		ctxj, cancel := context.WithCancel(context.Background())
		defer cancel()

		defer func(ctxj context.Context) {
			ctxj, cancel = context.WithTimeout(ctxj, time.Second*5)
			defer cancel()
			if err := tp.Shutdown(ctxj); err != nil {
				s.log.Fatal().Err(err).Msg("REST:jaeger:shutdown")
			}
		}(ctxj)
	}
	closed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		ctx, cancel := context.WithTimeout(context.Background(), s.srv.IdleTimeout)
		defer cancel()

		if err := s.srv.Shutdown(ctx); err != nil {
			s.log.Error().Err(err).Msg("REST:shutdown")
		}

		close(closed)
	}()
}

func (s *REST) SetupHandlers() {
	s.userHandler = handler.NewUserHandler(s.router, s.cqrs, s.log)
	s.authHandler = handler.NewAuthHandler(s.router, s.cqrs, s.log)
}

func (s *REST) SetupRoutes() {
	apiRouter.RegisterUserRouter(s.router, s.userHandler, s.cfg.Auth.JWTToken)
	apiRouter.RegisterAuthRouter(s.router, s.authHandler)
}

func (s *REST) SetupCQRS(ctx context.Context) {
	s.cqrs, _ = cqrs.NewCQRSService(ctx, s.cfg)
}
