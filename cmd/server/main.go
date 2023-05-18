package main

import (
	"context"
	"net/http"

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
	conf   *config.Conf
	server *http.Server
	ctx    context.Context

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

	r := apiRouter.NewApiRouter(l, conf.Token)
	userHandler = apiHandler.NewUserHandler(r, userService, l)
	authHandler = apiHandler.NewAuthHandler(r, authService, l)

	apiRouter.RegisterUserRouter(r, userHandler)
	apiRouter.RegisterAuthRouter(r, authHandler)

	server = apiServer.NewServer(conf, r)
	l.Info().Msgf("Starting REST server %v", server.Addr)

	apiRouter.GetRoutesList(r)
	apiServer.Run(server, conf)

	grpcServer, _ := apiServer.NewGrpcServer(conf.GRPC, authService, userService)
	l.Info().Msg("Starting gRPC server.")
	grpcServer.Run()
}
