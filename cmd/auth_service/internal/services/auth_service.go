package services

import (
	"context"
	"fmt"

	"github.com/RafalSalwa/interview-app-srv/pkg/encdec"
	"github.com/RafalSalwa/interview-app-srv/pkg/tracing"

	"github.com/RafalSalwa/interview-app-srv/cmd/auth_service/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/auth_service/internal/repository"
	"github.com/RafalSalwa/interview-app-srv/pkg/generator"
	"github.com/RafalSalwa/interview-app-srv/pkg/hashing"
	"github.com/RafalSalwa/interview-app-srv/pkg/jwt"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
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

func (a *AuthServiceImpl) EncryptEmails() {
	//email := "interview@interview.com"
	//enc, _ := encdec.Encrypt(email)
	//enc2, _ := encdec.Encrypt(email)
	//enc3, _ := encdec.Encrypt(email)
	//fmt.Println(enc, enc2, enc3)
	//dec, _ := encdec.Decrypt(enc)
	//dec2, _ := encdec.Decrypt(enc2)
	//dec3, _ := encdec.Decrypt(enc3)
	//fmt.Println(dec, dec2, dec3)

	//ctx := context.Background()
	//udb := &models.UserDBModel{}
	//users, err := a.repository.Find(ctx, udb)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, user := range users {
	//	enc, _ := encdec.Encrypt(user.Email)
	//	user.Email = enc
	//	hash, _ := hashing.Argon2ID("interview")
	//	user.Password = hash
	//	fmt.Println(user)
	//	_ = a.repository.Save(ctx, user)
	//}
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
	um.VerificationCode = vcode

	hash, err := hashing.Argon2ID(um.Password)
	if err != nil {
		return nil, err
	}
	um.Password = hash

	cipherText, err := encdec.Encrypt(cur.Email)
	if err != nil {
		return nil, err
	}
	um.Email = cipherText

	if errDB := a.repository.SignUp(ctx, um); errDB != nil {
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

func (a *AuthServiceImpl) SignInUser(ctx context.Context, reqUser *models.SignInUserRequest) (*models.UserResponse, error) {
	ctx, span := tracing.InitSpan(ctx, "auth_service-rpc", "AuthService SignInUser")
	defer span.End()

	enc, err := encdec.Encrypt(reqUser.Email)
	if err != nil {
		tracing.RecordError(span, err)
		return nil, err
	}
	fmt.Println("email", reqUser.Email, "enc", enc)
	udb := &models.UserDBModel{
		Email: enc,
	}
	udb, err = a.repository.Load(ctx, udb)
	if err != nil {
		tracing.RecordError(span, err)
		return nil, err
	}

	if err = hashing.Argon2IDComparePasswordAndHash(reqUser.Password, udb.Password); err != nil {
		tracing.RecordError(span, err)
		return nil, err
	}

	ur := &models.UserResponse{}
	err = ur.FromDBModel(udb)
	if err != nil {
		tracing.RecordError(span, err)
		return nil, err
	}

	tp, err := jwt.GenerateTokenPair(a.config, udb.Id)
	if err != nil {
		tracing.RecordError(span, err)
		return nil, err
	}

	ur.AssignTokenPair(tp)
	return ur, nil
}

func (a *AuthServiceImpl) GetVerificationKey(ctx context.Context, email string) (*models.UserResponse, error) {

	enc, _ := encdec.Encrypt(email)
	user := &models.UserDBModel{
		Email: enc,
	}
	dbUser, err := a.repository.Load(ctx, user)
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

func (a *AuthServiceImpl) Find(ctx context.Context, user *models.UserDBModel) (*models.UserResponse, error) {
	dbUser, err := a.repository.Load(ctx, user)
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

func (a *AuthServiceImpl) Load(ctx context.Context, user *models.UserDBModel) (*models.UserResponse, error) {
	dbUser, err := a.repository.Load(ctx, user)
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
	dbUser, err := a.repository.Load(ctx, udb)
	if err != nil {
		return err
	}

	if errV := a.repository.ConfirmVerify(ctx, udb); errV != nil {
		return errV
	}
	ur := &models.UserResponse{}

	return ur.FromDBModel(dbUser)
}

func (a *AuthServiceImpl) FindUserByID(uid int64) (*models.UserDBModel, error) {
	ctx := context.Background()
	dbUser, err := a.repository.ById(ctx, uid)
	if err != nil {
		return nil, err
	}
	return dbUser, nil
}
