package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/server"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
)

// @title           gateway service
// @version         1.0
// @description     This is a sample server celler server.

// @license.name  MiT
// @license.url   https://opensource.org/license/mit/

// @host      host.docker.internal:8080
// @BasePath  /auth

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
