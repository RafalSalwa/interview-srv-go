package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RafalSalwa/interview-app-srv/cmd/auth_service/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/auth_service/internal/server"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
func run() error {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	l := logger.NewConsole()

	srv := server.NewServerGRPC(cfg, l)

	if errSrv := srv.Run(); errSrv != nil {
		l.Error().Err(err).Msg("srv:run")
	}
	return nil
}
