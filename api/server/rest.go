package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/pkg/tracing"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"

	apiConfig "github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"

	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

type REST struct {
	srv *http.Server
	log *logger.Logger
	cfg *apiConfig.Conf
}

func NewRESTServer(c *apiConfig.Conf, r *mux.Router, l *logger.Logger) *REST {
	tlsConf := new(tls.Config)
	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
		TLSConfig:    tlsConf,
	}

	return &REST{
		srv: s,
		log: l,
		cfg: c,
	}
}

func (s REST) Run() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	if s.cfg.App.JaegerEnabled {
		tracer, closer, err := tracing.NewJaegerTracer(s.cfg.Jaeger)
		if err != nil {
			s.log.Error().Err(err).Msg("REST:run:jaeger")
		}
		defer closer.Close() // nolint: errcheck
		opentracing.SetGlobalTracer(tracer)
	}

	go func() {
		s.log.Info().Msgf("Starting REST server on: %v", s.srv.Addr)
		if err := s.srv.ListenAndServeTLS("/etc/ssl/certs/server.crt", "/etc/ssl/private/server.key"); err != nil && err != http.ErrServerClosed {
			s.log.Error().Err(err)
		}
	}()

	closed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		ctx, cancel := context.WithTimeout(context.Background(), s.srv.IdleTimeout)
		defer cancel()

		if err := s.srv.Shutdown(ctx); err != nil {

		}

		close(closed)
	}()
}
