package server

import (
	"context"
	"fmt"

	"github.com/gorilla/mux"

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
	l   *logger.Logger
}

func NewRESTServer(c *apiConfig.Conf, r *mux.Router, l *logger.Logger) *REST {
	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	return &REST{
		srv: s,
		l:   l,
	}
}

func (s REST) Run() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		}
	}()

	s.l.Info().Msgf("Starting REST server %v", s.srv.Addr)

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
