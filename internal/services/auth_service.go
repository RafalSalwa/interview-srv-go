package services

import (
	"context"
	"encoding/json"

	"github.com/RafalSalwa/interview-app-srv/internal/generator"
	"github.com/RafalSalwa/interview-app-srv/internal/mapper"
	"github.com/RafalSalwa/interview-app-srv/internal/password"
	"github.com/RafalSalwa/interview-app-srv/internal/repository"
	"github.com/RafalSalwa/interview-app-srv/pkg/jwt"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/RafalSalwa/interview-app-srv/pkg/query"
)

type AuthServiceImpl struct {
	ctx        context.Context
	repository repository.UserRepository
	logger     *logger.Logger
	config     jwt.JWTConfig
}

type AuthService interface {
	SignUpUser(ctx context.Context, request *models.CreateUserRequest) (*models.UserResponse, error)
	SignInUser(request *models.LoginUserRequest) (*models.UserResponse, error)
	GetVerificationKey(ctx context.Context, email string) (*models.UserResponse, error)
	Verify(ctx context.Context, vCode string) error
	Load(request *models.UserDBModel) (*models.UserResponse, error)
	FindUserById(uid int64) (*models.UserDBModel, error)
}

func NewAuthService(ctx context.Context, r repository.UserRepository, l *logger.Logger, c jwt.JWTConfig) AuthService {
	return &AuthServiceImpl{ctx, r, l, c}
}

func (s *AuthServiceImpl) SignUpUser(ctx context.Context, cur *models.CreateUserRequest) (*models.UserResponse, error) {
	if err := password.Validate(cur.Password, cur.PasswordConfirm); err != nil {
		return nil, err
	}

	um := &models.UserDBModel{}
	if err := um.FromCreateUserReq(cur); err != nil {
		return nil, err
	}

	roles, err := json.Marshal(models.Roles{Roles: []string{"ROLE_USER"}})
	if err != nil {
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

	um.Roles = roles
	um.VerificationCode = *vcode

	if errDB := s.repository.SingUp(um); errDB != nil {
		return nil, errDB
	}
	ur := &models.UserResponse{}
	err = ur.FromDBModel(um)
	if err != nil {
		return nil, err
	}

	return ur, nil
}
func (s *AuthServiceImpl) GetVerificationKey(ctx context.Context, email string) (*models.UserResponse, error) {
	user := &models.UserDBModel{
		Email: email,
	}
	dbUser, err := s.repository.Load(user)
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

func (s *AuthServiceImpl) Load(user *models.UserDBModel) (*models.UserResponse, error) {
	dbUser, err := s.repository.Load(user)
	if err != nil {
		return nil, err
	}
	if dbUser == nil {
		return nil, nil
	}
	dbUser, err = s.repository.UpdateLastLogin(s.ctx, dbUser)
	if err != nil {
		return nil, err
	}

	tp, err := jwt.GenerateTokenPair(s.config, dbUser.Id)
	_, _ = jwt.DecodeToken(tp.AccessToken, s.config.Access.PublicKey)
	if err != nil {
		s.logger.Error().Err(err).Msg("token_pair")
		return nil, err
	}

	ur := mapper.MapUserDBModelToUserResponse(dbUser)
	ur.Token = tp.AccessToken
	ur.RefreshToken = tp.RefreshToken

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
	dbUser, err := s.repository.ById(s.ctx, uid)
	if err != nil {
		return nil, err
	}
	return dbUser, nil
}
