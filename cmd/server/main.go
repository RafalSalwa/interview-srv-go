package main

import (
	"context"
	apiGrpc "github.com/RafalSalwa/interview-app-srv/api/grpc"
	apiHandler "github.com/RafalSalwa/interview-app-srv/api/handler"
	apiRouter "github.com/RafalSalwa/interview-app-srv/api/router"
	apiServer "github.com/RafalSalwa/interview-app-srv/api/server"
	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/internal/cqrs"
	"github.com/RafalSalwa/interview-app-srv/services"
	"github.com/RafalSalwa/interview-app-srv/sql"
	"github.com/RafalSalwa/interview-app-srv/util/logger"
	"github.com/RafalSalwa/interview-app-srv/util/validator"
	"github.com/gorilla/mux"
	"net/http"
)

type Application struct {
	Commands cqrs.Commands
	Queries  cqrs.Queries
}

var (
	conf   *config.Conf
	server *http.Server
	ctx    context.Context
	//redisclient *redis.Client

	userHandler apiHandler.IUserHandler
	authHandler apiHandler.AuthHandler

	userService services.UserSqlService

	authService services.AuthService
)

func main() {
	ctx = context.TODO()

	conf = config.New()
	l := logger.NewConsole(conf.App.Debug)
	v := validator.New()
	r := mux.NewRouter()

	db := sql.NewUsersDB(conf.DB)
	userService = services.NewMySqlService(db)
	authService = services.NewAuthService(ctx)

	userHandler = apiHandler.NewUserHandler(r, userService, l)
	authHandler = apiHandler.NewAuthHandler(r, authService, l)
	router := apiRouter.NewUserRouter(userHandler, v)
	server = apiServer.NewServer(conf, router)

	l.Info().Msgf("Starting REST server %v", server.Addr)
	apiServer.Run(server, conf)

	grpcServer, _ := apiGrpc.NewGrpcServer(conf.GRPC, authService, userService)
	l.Info().Msgf("Starting gRPC server %v", server.Addr)
	grpcServer.Run()

}
