package services

import (
    "context"
    "github.com/RafalSalwa/interview-app-srv/cmd/auth_service/config"
    "github.com/RafalSalwa/interview-app-srv/cmd/auth_service/internal/repository"
    "github.com/RafalSalwa/interview-app-srv/internal/generator"
    "github.com/RafalSalwa/interview-app-srv/internal/password"
    "github.com/RafalSalwa/interview-app-srv/pkg/jwt"
    "github.com/RafalSalwa/interview-app-srv/pkg/logger"
    "github.com/RafalSalwa/interview-app-srv/pkg/models"
    apiMongo "github.com/RafalSalwa/interview-app-srv/pkg/mongo"
    "github.com/RafalSalwa/interview-app-srv/pkg/query"
    "github.com/RafalSalwa/interview-app-srv/pkg/rabbitmq"
    redisClient "github.com/RafalSalwa/interview-app-srv/pkg/redis"
    "github.com/RafalSalwa/interview-app-srv/pkg/sql"
)

type AuthServiceImpl struct {
	repository      repository.UserRepository
	mongoRepo       *repository.Mongo
	redisRepo       *repository.Redis
	rabbitPublisher *rabbitmq.Publisher
	logger          *logger.Logger
	config          jwt.JWTConfig
}

type AuthService interface {
	SignUpUser(ctx context.Context, request *models.CreateUserRequest) (*models.UserResponse, error)
	SignInUser(request *models.LoginUserRequest) (*models.UserResponse, error)
	GetVerificationKey(ctx context.Context, email string) (*models.UserResponse, error)
	Verify(ctx context.Context, vCode string) error
	Load(request *models.UserDBModel) (*models.UserResponse, error)
	FindUserById(uid int64) (*models.UserDBModel, error)
}

func NewAuthService(ctx context.Context, cfg *config.Config, log *logger.Logger) AuthService {
	mongoClient, err := apiMongo.NewClient(ctx, cfg.Mongo)
	if err != nil {
		log.Error().Err(err).Msg("grpc:run:mongo")
	}

	universalRedisClient, err := redisClient.NewUniversalRedisClient(cfg.Redis)
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

	return &AuthServiceImpl{
		repository:      userRepository,
		mongoRepo:       mongoRepo,
		redisRepo:       redisRepo,
		rabbitPublisher: publisher,
		logger:          log,
		config:          cfg.JWTToken,
	}
}

func (s *AuthServiceImpl) SignUpUser(ctx context.Context, cur *models.CreateUserRequest) (*models.UserResponse, error) {
	if err := password.Validate(cur.Password, cur.PasswordConfirm); err != nil {
		return nil, err
	}

	um := &models.UserDBModel{}
	if err := um.FromCreateUserReq(cur); err != nil {
		return nil, err
	}

	vcode, err := generator.RandomString(12)
	if err != nil {
		return nil, err
	}

	um.Password, err = password.HashPassword(um.Password)
	if err != nil {
		return nil, err
	}

	um.VerificationCode = *vcode

	if errDB := s.repository.SingUp(um); errDB != nil {
		s.logger.Error().Err(errDB).Msg("repo:user:signUp")
		return nil, errDB
	}

	if errR := s.redisRepo.PutUser(ctx, *um); errR != nil {
		s.logger.Error().Err(err).Msg("redis:user:set")
		return nil, errR
	}
	if err = s.rabbitPublisher.Publish(um.AMQP()); err != nil {
		s.logger.Error().Err(err).Msg("rabbit:publish")
		return nil, err
	}
	ur := &models.UserResponse{}
	err = ur.FromDBModel(um)
	if err != nil {
		return nil, err
	}

	return ur, nil
}

func (s *AuthServiceImpl) SignInUser(user *models.LoginUserRequest) (*models.UserResponse, error) {
	q := query.Use(s.repository.GetConnection()).UserDBModel
	dbu, errDB := q.FilterWithUsernameOrEmail(user.Username, user.Email)
	if errDB != nil {
		return nil, errDB
	}

	ur := &models.UserResponse{}
	err := ur.FromDBModel(dbu)
	if err != nil {
		return nil, err
	}

	tp, err := jwt.GenerateTokenPair(s.config, dbu.Id)
	if err != nil {
		return nil, err
	}
	ur.AssignTokenPair(tp)
	return ur, nil
}

func (a *AuthServiceImpl) GetVerificationKey(ctx context.Context, email string) (*models.UserResponse, error) {
	user := &models.UserDBModel{
		Email: email,
	}
	dbUser, err := a.repository.Load(user)
	if err != nil {
		return nil, err
	}
	ur := &models.UserResponse{}
	err = ur.FromDBModel(dbUser)
	if err != nil {
		return nil, err
	}

	return ur, nil
}

func (s *AuthServiceImpl) Load(user *models.UserDBModel) (*models.UserResponse, error) {
	ctx := context.Background()
	dbUser, err := s.repository.Load(user)
	if err != nil {
		return nil, err
	}
	if dbUser == nil {
		return nil, nil
	}
	dbUser, err = s.repository.UpdateLastLogin(ctx, dbUser)
	if err != nil {
		return nil, err
	}

	tp, err := jwt.GenerateTokenPair(s.config, dbUser.Id)
	_, _ = jwt.DecodeToken(tp.AccessToken, s.config.Access.PublicKey)
	if err != nil {
		s.logger.Error().Err(err).Msg("token_pair")
		return nil, err
	}

	ur := &models.UserResponse{}
	err = ur.FromDBModel(dbUser)
	if err != nil {
		return nil, err
	}
	ur.AssignTokenPair(tp)

	return ur, nil
}

func (s *AuthServiceImpl) Verify(ctx context.Context, vCode string) error {
	udb := &models.UserDBModel{
		VerificationCode: vCode,
	}
	dbUser, err := s.repository.Load(udb)
	if err != nil {
		return err
	}

	if errV := s.repository.ConfirmVerify(ctx, udb); errV != nil {
		return errV
	}
	ur := &models.UserResponse{}

	if errM := ur.FromDBModel(dbUser); errM != nil {
		return errM
	}
	return nil
}

func (s *AuthServiceImpl) FindUserById(uid int64) (*models.UserDBModel, error) {
	ctx := context.Background()
	dbUser, err := s.repository.ById(ctx, uid)
	if err != nil {
		return nil, err
	}
	return dbUser, nil
}
