package services

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/internal/mapper"
	"github.com/RafalSalwa/interview-app-srv/internal/repository"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/RafalSalwa/interview-app-srv/util/jwt"
)

type AuthServiceImpl struct {
	ctx        context.Context
	repository repository.UserRepository
	logger     *logger.Logger
	config     config.ConfToken
}

type AuthService interface {
	SignUpUser(request *models.CreateUserRequest) (*models.UserDBResponse, error)
	Load(request *models.LoginUserRequest) (*models.UserResponse, error)
	FindUserById(uid int64) (*models.UserDBModel, error)
}

func NewAuthService(ctx context.Context, r repository.UserRepository, l *logger.Logger, c config.ConfToken) AuthService {
	return &AuthServiceImpl{ctx, r, l, c}
}

func (s *AuthServiceImpl) SignUpUser(user *models.CreateUserRequest) (*models.UserDBResponse, error) {
	return nil, nil
}

func (s *AuthServiceImpl) Load(user *models.LoginUserRequest) (*models.UserResponse, error) {
	dbUser, err := s.repository.ByLogin(s.ctx, user)
	if err != nil {
		return nil, err
	}
	if dbUser == nil {
		return nil, nil
	}
	loginTime, err := s.repository.UpdateLastLogin(s.ctx, dbUser.Id)
	if err != nil {
		return nil, err
	}
	dbUser.LastLogin = loginTime
	ur := mapper.MapUserDBModelToUserResponse(dbUser)

	accessToken, err := jwt.CreateToken(s.config.AccessTokenExpiresIn, dbUser.Id, s.config.RefreshTokenPrivateKey)
	if err != nil {
		s.logger.Error().Err(err).Msg("access_token")
		return nil, err
	}
	ur.Token = accessToken
	refreshToken, err := jwt.CreateToken(s.config.RefreshTokenExpiresIn, dbUser.Id, s.config.RefreshTokenPrivateKey)
	if err != nil {
		s.logger.Error().Err(err).Msg("refresh_token")
		return nil, err
	}
	ur.RefreshToken = refreshToken

	return ur, nil
}

func (s *AuthServiceImpl) FindUserById(uid int64) (*models.UserDBModel, error) {
	dbUser, err := s.repository.ById(s.ctx, uid)
	if err != nil {
		return nil, err
	}
	return dbUser, nil
}
