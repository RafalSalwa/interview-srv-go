package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/RafalSalwa/interview-app-srv/internal/generator"
	"github.com/RafalSalwa/interview-app-srv/internal/password"

	"github.com/RafalSalwa/interview-app-srv/internal/mapper"

	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/internal/jwt"
	"github.com/RafalSalwa/interview-app-srv/internal/repository"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
)

type AuthServiceImpl struct {
	ctx        context.Context
	repository repository.UserRepository
	logger     *logger.Logger
	config     config.ConfToken
}

type AuthService interface {
	SignUpUser(request *models.CreateUserRequest) (*models.UserResponse, error)
	Load(request *models.LoginUserRequest) (*models.UserResponse, error)
	FindUserById(uid int64) (*models.UserDBModel, error)
}

func NewAuthService(ctx context.Context, r repository.UserRepository, l *logger.Logger, c config.ConfToken) AuthService {
	return &AuthServiceImpl{ctx, r, l, c}
}

func (s *AuthServiceImpl) SignUpUser(cur *models.CreateUserRequest) (*models.UserResponse, error) {
	um := mapper.UserCreateRequestToDBModel(cur)

	roles, _ := json.Marshal(models.Roles{Roles: []string{"ROLE_USER"}})
	vcode, _ := generator.RandomString(6)
	um.Password, _ = password.HashPassword(um.Password)

	um.Roles = roles
	um.VerificationCode = *vcode
	um.CreatedAt = time.Now()

	s.repository.SingUp(um)

	ur := mapper.MapUserDBModelToUserResponse(um)
	return ur, nil
}

func (s *AuthServiceImpl) Load(user *models.LoginUserRequest) (*models.UserResponse, error) {
	dbUser, err := s.repository.ByLogin(s.ctx, user)
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

	tp, err := jwt.GenerateTokenPair(s.config, dbUser.Id, dbUser.Username)
	_, _ = jwt.DecodeToken(tp.AccessToken, s.config.AccessTokenPublicKey)
	// s.logger.Info().Msgf("%v", v)
	if err != nil {
		s.logger.Error().Err(err).Msg("token_pair")
		return nil, err
	}

	ur := mapper.MapUserDBModelToUserResponse(dbUser)
	ur.Token = tp.AccessToken
	ur.RefreshToken = tp.RefreshToken

	return ur, nil
}

func (s *AuthServiceImpl) FindUserById(uid int64) (*models.UserDBModel, error) {
	dbUser, err := s.repository.ById(s.ctx, uid)
	if err != nil {
		return nil, err
	}
	return dbUser, nil
}
