package main

import (
	"log"

	"github.com/RafalSalwa/interview-app-srv/cmd/auth_service/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/auth_service/internal/server"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err) 
	}

	l := logger.NewConsole(cfg.App.Debug)
	srv := server.NewServerGRPC(cfg, l)
 
	if errSrv := srv.Run(); errSrv != nil {
		l.Error().Err(err).Msg("srv:run") 
	}
} 
