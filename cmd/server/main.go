package main

import (
	"context"
	_ "net/http/pprof"

	apiHandler "github.com/RafalSalwa/interview-app-srv/api/handler"
	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/internal/services"
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
	//ctx = context.TODO()
	//conf = config.NewViper()
	//l := logger.NewConsole(conf.App.Debug)
	//
	//db := sql.NewUsersDB(conf.DB, l)
	//ormDB := sql.NewUsersDBGorm(conf.DB, l)
	//userRepository := repository.NewUserAdapter(ormDB)

	//userService = services.NewMySqlService(db, l)
	//authService = services.NewAuthService(ctx, userRepository, l, conf.Token)

	//r := apiRouter.NewApiRouter(l, conf)
	//userHandler = apiHandler.NewUserHandler(r, userService, l)
	//authHandler = apiHandler.NewAuthHandler(r, authService, l)
	//
	//apiRouter.RegisterUserRouter(r, userHandler, conf)
	//apiRouter.RegisterAuthRouter(r, authHandler)
	//
	//srv := apiServer.NewRESTServer(conf, r, l)
	//srv.Run()
	//grpcServer, err := apiServer.NewGrpcServer(conf.GRPC, l, authService, userService)
	//if err != nil {
	//	l.Error().Err(err)
	//}
	//grpcServer.Run()
}
