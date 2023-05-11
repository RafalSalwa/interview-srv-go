package server

import (
	"context"
	"fmt"

	"github.com/gorilla/mux"

	apiConfig "github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"

	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Handler struct {
	handler http.Handler
	logger  *logger.Logger
}

func NewServer(c *apiConfig.Conf, r *mux.Router) *http.Server {
	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}
	return s
}

func Run(s *http.Server, conf *apiConfig.Conf) {

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		}
	}()

	closed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		ctx, cancel := context.WithTimeout(context.Background(), conf.Server.TimeoutIdle)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {

		}

		close(closed)
	}()
	//if conf.App.Env == "production" {
	//	s.ListenAndServeTLS("/etc/ssl/private/unohouse.com.pl.crt", "/etc/ssl/private/unohouse.com.pl.key")
	//} else {
	//	s.ListenAndServe()
	//}
}
