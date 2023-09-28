package services

import (
	"context"

	"github.com/RafalSalwa/interview-app-srv/cmd/user_service/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/user_service/internal/repository"
	"github.com/RafalSalwa/interview-app-srv/pkg/jwt"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	apiMongo "github.com/RafalSalwa/interview-app-srv/pkg/mongo"
	"github.com/RafalSalwa/interview-app-srv/pkg/rabbitmq"
	redisClient "github.com/RafalSalwa/interview-app-srv/pkg/redis"
	"github.com/RafalSalwa/interview-app-srv/pkg/sql"
)

type UserServiceImpl struct {
	repository      repository.UserRepository
	mongoRepo       *repository.Mongo
	redisRepo       *repository.Redis
	rabbitPublisher *rabbitmq.Publisher
	logger          *logger.Logger
	config          jwt.JWTConfig
}

type UserService interface {
	GetUser(ctx context.Context, user *models.SignInUserRequest) (*models.UserDBModel, error)
	GetById(ctx context.Context, id int64) (*models.UserDBModel, error)
	UsernameInUse(user *models.UserDBModel) (bool, error)
	StoreVerificationData(ctx context.Context, vCode string) error
	UpdateUser(user *models.UpdateUserRequest) (err error)
	LoginUser(user *models.SignInUserRequest) (*models.UserResponse, error)
	UpdateUserPassword(userid int64, password string) error
	CreateUser(user *models.SignUpUserRequest) (*models.UserResponse, error)
}

func NewUserService(ctx context.Context, cfg *config.Config, log *logger.Logger) UserServiceImpl {
	mongoClient, err := apiMongo.NewClient(ctx, cfg.Mongo)
	if err != nil {
		log.Error().Err(err).Msg("grpc:run:mongo")
	}

	universalRedisClient, err := redisClient.NewUniversalRedisClient(ctx, cfg.Redis)
	if err != nil {
		log.Error().Err(err).Msg("redis")
	}

	publisher, errP := rabbitmq.NewPublisher(cfg.Rabbit)
	if errP != nil {
		log.Error().Err(err).Msg("rabbitmq")
	}

	ormDB, err := sql.NewGormConnection(cfg.MySQL)
	if err != nil {
		log.Error().Err(err).Msg("gorm")
	}
	userRepository := repository.NewUserAdapter(ormDB)
	mongoRepo := repository.NewMongoRepository(mongoClient, log)
	redisRepo := repository.NewRedisRepository(universalRedisClient, log)

	return UserServiceImpl{
		repository:      userRepository,
		mongoRepo:       mongoRepo,
		redisRepo:       redisRepo,
		rabbitPublisher: publisher,
		logger:          log,
		config:          cfg.JWTToken,
	}
}

func (s *UserServiceImpl) GetUser(ctx context.Context, user *models.SignInUserRequest) (*models.UserDBModel, error) {
	userDbModel := &models.UserDBModel{}
	userDbModel.Username = user.Username
	userDbModel.Email = user.Email
	userDbModel.Password = user.Password

	ur, err := s.repository.Load(userDbModel)
	if err != nil {
		return nil, err
	}
	return ur, nil
}
func (s *UserServiceImpl) GetById(ctx context.Context, id int64) (*models.UserDBModel, error) {
	user, err := s.repository.ById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) UsernameInUse(user *models.UserDBModel) (bool, error) {
	ur, err := s.repository.Load(user)
	if err != nil {
		return false, err
	}
	return ur == nil, nil
}

func (s *UserServiceImpl) StoreVerificationData(ctx context.Context, vCode string) error {
	err := s.repository.ConfirmVerify(ctx, vCode)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) UpdateUser(user *models.UpdateUserRequest) (err error) {
	// TODO implement me
	panic("implement me")
}

func (s *UserServiceImpl) LoginUser(user *models.SignInUserRequest) (*models.UserResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (s *UserServiceImpl) UpdateUserPassword(userid int64, password string) error {
	err := s.repository.ChangePassword(userid, password)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) CreateUser(user *models.SignUpUserRequest) (*models.UserResponse, error) {
	panic("implement me")
}
