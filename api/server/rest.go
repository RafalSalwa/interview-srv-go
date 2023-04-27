package server

import (
	"context"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func NewServer(c *config.Conf, handler http.Handler) *http.Server {
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      handler,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}
	return s
}

func Run(s *http.Server) {

	_ = s.ListenAndServe()

	closed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		ctx, cancel := context.WithTimeout(context.Background(), c.Server.TimeoutIdle)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {

		}
		close(closed)
	}()
}
