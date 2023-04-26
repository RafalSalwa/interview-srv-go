package main

import (
	"context"
	"fmt"
	apiHandler "github.com/RafalSalwa/interview-app-srv/api/handler"
	apiRouter "github.com/RafalSalwa/interview-app-srv/api/router"
	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/service"
	"github.com/RafalSalwa/interview-app-srv/sql"
	"github.com/RafalSalwa/interview-app-srv/util/logger"
	"github.com/RafalSalwa/interview-app-srv/util/validator"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func main() {
	c := config.New()
	l := logger.NewConsole(c.App.Debug)
	v := validator.New()

	db := sql.NewUsersDB(c.DB)
	us := service.NewMySqlService(db)

	r := mux.NewRouter()
	h := apiHandler.NewHandler(r, us, l)
	router := apiRouter.NewRouter(h, v)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      router,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}
	closed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		l.Info().Msgf("Shutting down server %v", s.Addr)

		ctx, cancel := context.WithTimeout(context.Background(), c.Server.TimeoutIdle)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			l.Error().Err(err).Msg("Server shutdown failure")
		}

		if err := db.Close(); err != nil {
			l.Error().Err(err).Msg("DB connection closing failure")
		}
		close(closed)
	}()

	l.Info().Msgf("Starting server %v", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		l.Fatal().Err(err).Msg("Server startup failure")
	}
}
