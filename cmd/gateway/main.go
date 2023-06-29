package main

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/server"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	l := logger.NewConsole(cfg.App.Debug)
	srv := server.NewRESTServer(cfg, l)
	srv.Run(ctx)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint
}
