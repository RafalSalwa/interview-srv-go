package repository

import (
	"context"
	"encoding/json"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel"
	"strconv"
)

type Redis struct {
	log         *logger.Logger
	redisClient redis.UniversalClient
}

func NewRedisRepository(redisClient redis.UniversalClient, log *logger.Logger) *Redis {
	return &Redis{log: log, redisClient: redisClient}
}

func (r Redis) PutUser(ctx context.Context, user models.UserDBModel) error {
	ctx, span := otel.GetTracerProvider().Tracer("auth_service-redis").Start(ctx, "Redis PutUser")
	defer span.End()

	key := strconv.FormatInt(user.Id, 10)
	bytes, err := json.Marshal(user)

	if err != nil {
		r.log.Error().Err(err).Msg("redis:user:put:marshal")
		return err
	}

	if errR := r.redisClient.HSetNX(ctx, "users", key, bytes).Err(); errR != nil {
		r.log.Error().Err(errR).Msg("redis:user:put:HSetNX")
		return errR
	}
	return nil
}
