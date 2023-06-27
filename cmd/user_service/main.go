package main

import (
	"github.com/RafalSalwa/interview-app-srv/cmd/user_service/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/user_service/internal/server"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"log"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	l := logger.NewConsole(cfg.App.Debug)
	srv := server.NewServerGRPC(cfg, l)

	if errSrv := srv.Run(); errSrv != nil {
		log.Fatal(err)
	}
}
