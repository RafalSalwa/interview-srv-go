package services

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/cmd/auth_service/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/auth_service/internal/repository"
	"github.com/RafalSalwa/interview-app-srv/pkg/generator"
	"github.com/RafalSalwa/interview-app-srv/pkg/hashing"
	"github.com/RafalSalwa/interview-app-srv/pkg/jwt"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/RafalSalwa/interview-app-srv/pkg/query"
	"github.com/RafalSalwa/interview-app-srv/pkg/rabbitmq"
	"go.opentelemetry.io/otel"
)

type AuthServiceImpl struct {
	repository      repository.UserRepository
	rabbitPublisher *rabbitmq.Publisher
	logger          *logger.Logger
	config          jwt.JWTConfig
}

func NewAuthService(ctx context.Context, cfg *config.Config, log *logger.Logger) AuthService {

	publisher, errP := rabbitmq.NewPublisher(cfg.Rabbit)
	if errP != nil {
		log.Error().Err(errP).Msg("auth:service:publisher")
	}

	userRepository, errR := repository.NewUserRepository(ctx, cfg.App.RepositoryType, cfg)
	if errR != nil {
		log.Error().Err(errR).Msg("auth:service:repository")
	}

	return &AuthServiceImpl{
		repository:      userRepository,
		rabbitPublisher: publisher,
		logger:          log,
		config:          cfg.JWTToken,
	}
}

func (a *AuthServiceImpl) SignUpUser(ctx context.Context, cur *models.SignUpUserRequest) (*models.UserResponse, error) {
	ctx, span := otel.GetTracerProvider().Tracer("auth_service-service").Start(ctx, "Service SignUpUser")
	defer span.End()
	if err := hashing.Validate(cur.Password, cur.PasswordConfirm); err != nil {
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
	um.Password, err = hashing.Argon2ID(um.Password)
	if err != nil {
		return nil, err
	}

	um.VerificationCode = *vcode
	if errDB := a.repository.SingUp(ctx, um); errDB != nil {
		return nil, errDB
	}
	if err = a.rabbitPublisher.Publish(ctx, um.AMQP()); err != nil {
		return nil, err
	}
	ur := &models.UserResponse{}
	err = ur.FromDBModel(um)
	if err != nil {
		return nil, err
	}
	return ur, nil
}

func (a *AuthServiceImpl) SignInUser(user *models.SignInUserRequest) (*models.UserResponse, error) {
	q := query.Use(a.repository.GetConnection()).UserDBModel
	dbu, errDB := q.FilterWithUsernameOrEmail(user.Username, user.Email)
	if errDB != nil {
		return nil, errDB
	}

	ur := &models.UserResponse{}
	err := ur.FromDBModel(dbu)
	if err != nil {
		return nil, err
	}

	tp, err := jwt.GenerateTokenPair(a.config, dbu.Id)
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

func (a *AuthServiceImpl) Find(user *models.UserDBModel) (*models.UserResponse, error) {
	dbUser, err := a.repository.Load(user)
	if err != nil {
		return nil, err
	}
	if dbUser == nil {
		return nil, nil
	}

	ur := &models.UserResponse{}
	err = ur.FromDBModel(dbUser)
	if err != nil {
		return nil, err
	}

	return ur, nil
}

func (a *AuthServiceImpl) Load(user *models.UserDBModel) (*models.UserResponse, error) {
	ctx := context.Background()
	dbUser, err := a.repository.Load(user)
	if err != nil {
		return nil, err
	}
	if dbUser == nil {
		return nil, nil
	}
	dbUser, err = a.repository.UpdateLastLogin(ctx, dbUser)
	if err != nil {
		return nil, err
	}

	tp, err := jwt.GenerateTokenPair(a.config, dbUser.Id)
	_, _ = jwt.DecodeToken(tp.AccessToken, a.config.Access.PublicKey)
	if err != nil {
		a.logger.Error().Err(err).Msg("token_pair")
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

func (a *AuthServiceImpl) Verify(ctx context.Context, vCode string) error {
	udb := &models.UserDBModel{
		VerificationCode: vCode,
	}
	dbUser, err := a.repository.Load(udb)
	if err != nil {
		return err
	}

	if errV := a.repository.ConfirmVerify(ctx, udb); errV != nil {
		return errV
	}
	ur := &models.UserResponse{}

	if errM := ur.FromDBModel(dbUser); errM != nil {
		return errM
	}
	return nil
}

func (a *AuthServiceImpl) FindUserById(uid int64) (*models.UserDBModel, error) {
	ctx := context.Background()
	dbUser, err := a.repository.ById(ctx, uid)
	if err != nil {
		return nil, err
	}
	return dbUser, nil
}
