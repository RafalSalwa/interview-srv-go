package main

import (
	"context"
	apiGrpc "github.com/RafalSalwa/interview-app-srv/api/grpc"
	apiHandler "github.com/RafalSalwa/interview-app-srv/api/handler"
	apiRouter "github.com/RafalSalwa/interview-app-srv/api/router"
	apiServer "github.com/RafalSalwa/interview-app-srv/api/server"
	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/services"
	"github.com/RafalSalwa/interview-app-srv/sql"
	"github.com/RafalSalwa/interview-app-srv/util/logger"
	"github.com/RafalSalwa/interview-app-srv/util/validator"
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

var (
	conf   *config.Conf
	server *http.Server
	ctx    context.Context
	//redisclient *redis.Client

	userHandler apiHandler.UserHandler

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
	server = apiServer.NewServer(c, router)
	apiServer.Run(server)
	l.Info().Msgf("Starting REST server %v", server.Addr)

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		l.Error().Err(err).Msgf("failed to listen: %v", err)
	}
	grpcServer, _ := apiGrpc.NewGrpcServer(c)
	l.Info().Msgf("Starting gRPC server %v", s.Addr)
	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			return
		}

	}()

}
