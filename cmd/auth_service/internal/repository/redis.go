package repository

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel"
	otelcodes "go.opentelemetry.io/otel/codes"
	"gorm.io/gorm"
)

type RedisRepository struct {
	log         *logger.Logger
	redisClient redis.UniversalClient
}

type RedisAdapter struct {
	DB *redis.UniversalClient
}

func (r RedisAdapter) Update(ctx context.Context, user models.UserDBModel) error {
	//TODO implement me
	panic("implement me")
}

func (r RedisAdapter) Save(ctx context.Context, user models.UserDBModel) error {
	//TODO implement me
	panic("implement me")
}

func (r RedisAdapter) Find(ctx context.Context, user *models.UserDBModel) ([]models.UserDBModel, error) {
	//TODO implement me
	panic("implement me")
}

func newRedisUserRepository(db *redis.UniversalClient) UserRepository {
	return &RedisAdapter{DB: db}
}

func NewRedisRepository(redisClient redis.UniversalClient, log *logger.Logger) *RedisRepository {
	return &RedisRepository{log: log, redisClient: redisClient}
}

func (r RedisAdapter) SignUp(ctx context.Context, user *models.UserDBModel) error {
	// TODO implement me
	panic("implement me")
}

func (r RedisAdapter) Load(ctx context.Context, user *models.UserDBModel) (*models.UserDBModel, error) {
	// TODO implement me
	panic("implement me")
}

func (r RedisAdapter) ById(ctx context.Context, id int64) (*models.UserDBModel, error) {
	// TODO implement me
	panic("implement me")
}

func (r RedisAdapter) ByLogin(ctx context.Context, user *models.SignInUserRequest) (*models.UserDBModel, error) {
	// TODO implement me
	panic("implement me")
}

func (r RedisAdapter) ConfirmVerify(ctx context.Context, udb *models.UserDBModel) error {
	// TODO implement me
	panic("implement me")
}

func (r RedisAdapter) UpdateLastLogin(ctx context.Context, u *models.UserDBModel) (*models.UserDBModel, error) {
	// TODO implement me
	panic("implement me")
}

func (r RedisAdapter) FindUserById(uid int64) (*models.UserDBModel, error) {
	// TODO implement me
	panic("implement me")
}

func (r RedisAdapter) GetConnection() *gorm.DB {
	// TODO implement me
	panic("implement me")
}

func (r RedisRepository) PutUser(ctx context.Context, user models.UserDBModel) error {
	ctx, span := otel.GetTracerProvider().Tracer("auth_service-redis").Start(ctx, "RedisRepository PutUser")
	defer span.End()

	key := strconv.FormatInt(user.Id, 10)
	bytes, err := json.Marshal(user)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelcodes.Error, err.Error())
		return err
	}

	if errR := r.redisClient.HSetNX(ctx, "users", key, bytes).Err(); errR != nil {
		span.RecordError(err)
		span.SetStatus(otelcodes.Error, err.Error())
		return errR
	}
	return nil
}
