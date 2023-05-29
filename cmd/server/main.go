package main

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/internal/repository"
	_ "net/http/pprof"

	apiHandler "github.com/RafalSalwa/interview-app-srv/api/handler"
	apiRouter "github.com/RafalSalwa/interview-app-srv/api/router"
	apiServer "github.com/RafalSalwa/interview-app-srv/api/server"
	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/internal/services"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/sql"
)

var (
	conf *config.Conf
	ctx  context.Context

	userHandler apiHandler.UserHandler
	authHandler apiHandler.IAuthHandler

	userService services.UserSqlService
	authService services.AuthService
)

func main() {
	ctx = context.TODO()
	conf = config.New()
	l := logger.NewConsole(conf.App.Debug)

	db := sql.NewUsersDB(conf.DB, l)
	ormDB := sql.NewUsersDBGorm(conf.DB, l)

	userRepository := repository.NewUserAdapter(ormDB)
	userService = services.NewMySqlService(db, l)
	authService = services.NewAuthService(ctx, userRepository, l, conf.Token)

	r := apiRouter.NewApiRouter(l, conf)

	userHandler = apiHandler.NewUserHandler(r, userService, l)
	authHandler = apiHandler.NewAuthHandler(r, authService, l)

	apiRouter.RegisterUserRouter(r, userHandler, conf)
	apiRouter.RegisterAuthRouter(r, authHandler)

	srv := apiServer.NewServer(conf, r, l)
	srv.Run()

	grpcServer, _ := apiServer.NewGrpcServer(conf.GRPC, l, authService, userService)
	grpcServer.Run()
}
