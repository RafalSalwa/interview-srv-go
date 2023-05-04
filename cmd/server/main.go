package main

import (
	"context"
	"net/http"

	apiHandler "github.com/RafalSalwa/interview-app-srv/api/handler"
	apiRouter "github.com/RafalSalwa/interview-app-srv/api/router"
	apiServer "github.com/RafalSalwa/interview-app-srv/api/server"
	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/services"
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
	//v := validator.New()
	r := apiRouter.NewApiRouter(l)

	db := sql.NewUsersDB(conf.DB)
	userService = services.NewMySqlService(db)
	authService = services.NewAuthService(ctx)

	userHandler = apiHandler.NewUserHandler(r, userService, l)
	authHandler = apiHandler.NewAuthHandler(r, authService, l)

	apiRouter.RegisterUserRouter(r, userHandler)
	apiRouter.RegisterAuthRouter(r, authHandler)

	//grpcServer, _ := apiGrpc.NewGrpcServer(conf.GRPC, authService, userService)
	//l.Info().Msg("Starting gRPC server.")
	//grpcServer.Run()

	server = apiServer.NewServer(conf, r)
	l.Info().Msgf("Starting REST server %v", server.Addr)

	apiRouter.GetRoutesList(r)
	apiServer.Run(server, conf)
	l.Info().Msg("Started")

}
