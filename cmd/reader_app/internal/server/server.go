package server

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/cmd/reader_app/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/reader_app/internal/services"
	"github.com/RafalSalwa/interview-app-srv/internal/repository"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	redisClient "github.com/RafalSalwa/interview-app-srv/pkg/redis"
	"github.com/RafalSalwa/interview-app-srv/pkg/sql"
	"github.com/RafalSalwa/interview-app-srv/pkg/tracing"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	log         *logger.Logger
	cfg         *config.Config
	v           *validator.Validate
	mongoClient *mongo.Client
	redisClient redis.UniversalClient
}

func NewServerGRPC(cfg *config.Config, log *logger.Logger) *server {
	return &server{log: log, cfg: cfg, v: validator.New()}
}

func (srv *server) Run() error {
	if srv.cfg.Jaeger.Enable {
		tracer, closer, err := tracing.NewJaegerTracer(*srv.cfg.Jaeger)
		if err != nil {
			return err
		}
		defer closer.Close() // nolint: errcheck
		opentracing.SetGlobalTracer(tracer)
	}

	ctx := context.TODO()
	//mongoClient, err := apiMongo.NewClient(ctx, srv.cfg.Mongo)
	//if err != nil {
	//	srv.log.Error().Err(err).Msg("grpc:run:mongo")
	//	return errors.Wrap(err, "NewMongoDBConn")
	//}
	//srv.mongoClient = mongoClient
	//defer mongoClient.Disconnect(ctx) // nolint: errcheck

	client, err := redisClient.NewUniversalRedisClient(srv.cfg.Redis)
	if err != nil {
		srv.log.Error().Err(err).Msg("redis")
	}
	srv.redisClient = client
	defer srv.redisClient.Close() // nolint: errcheck

	db, _ := sql.NewMySQLConnection(srv.cfg.MySQL)
	ormDB, _ := sql.NewGormConnection(srv.cfg.MySQL)

	userRepository := repository.NewUserAdapter(ormDB)

	userService := services.NewMySqlService(db, srv.log)
	authService := services.NewAuthService(ctx, userRepository, srv.log, srv.cfg.JWTToken)

	grpcServer, err := NewGrpcServer(srv.cfg.GRPC, srv.log, authService, userService)
	if err != nil {
		srv.log.Error().Err(err)
	}
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		grpcServer.Run()
	}()
	<-shutdown
	return nil
}
