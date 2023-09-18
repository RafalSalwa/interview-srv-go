package repository

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/pkg/mongo"
	redisClient "github.com/RafalSalwa/interview-app-srv/pkg/redis"
	"github.com/RafalSalwa/interview-app-srv/pkg/sql"
	"github.com/pkg/errors"
)

const (
	MySQL   string = "MySQL"
	MongoDB string = "MongoDB"
	Redis   string = "RedisRepository"
)

func NewUserRepository(ctx context.Context, dbType string, params interface{}) (UserRepository, error) {

	switch dbType {

	case MySQL:
		mysqlParams, ok := params.(*sql.MySQL)
		if !ok {
			return nil, errors.New("Missing parameters provided for mySQL connection")
		}
		con, err := sql.NewGormConnection(*mysqlParams)
		if err != nil {
			return nil, err
		}
		return newMySQLUserRepository(con), nil

	case MongoDB:
		mongoParams, ok := params.(*mongo.Config)
		if !ok {
			return nil, errors.New("Missing parameters for MongoDB connection")
		}
		mongoClient, err := mongo.NewClient(ctx, *mongoParams)
		if err != nil {
			return nil, err
		}
		return newMongoDBUserRepository(mongoClient), nil

	case Redis:
		redisParams, ok := params.(*redisClient.Config)
		if !ok {
			panic("Invalid Redis parameters")
		}
		universalRedisClient, err := redisClient.NewUniversalRedisClient(redisParams)
		if err != nil {
			return nil, err
		}

		return newRedisUserRepository(&universalRedisClient), nil
	default:
		panic("Unsupported database type")
	}
}
