package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/server"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	l := logger.NewConsole(cfg.App.Debug)
	srv := server.NewRESTServer(cfg, l)
	srv.Run()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint
}
