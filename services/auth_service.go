package services

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/internal/repository"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
)

type AuthServiceImpl struct {
	ctx        context.Context
	repository repository.UserRepository
	logger     *logger.Logger
}

type AuthService interface {
	SignUpUser(request *models.CreateUserRequest) (*models.UserDBResponse, error)
	SignInUser(request *models.LoginUserRequest) (*models.UserDBResponse, error)
	Token()
}

func NewAuthService(ctx context.Context, r repository.UserRepository, l *logger.Logger) AuthService {
	return &AuthServiceImpl{ctx, r, l}
}

func (uc *AuthServiceImpl) Token() {

}

func (uc *AuthServiceImpl) SignUpUser(user *models.CreateUserRequest) (*models.UserDBResponse, error) {
	return nil, nil
}

func (uc *AuthServiceImpl) SignInUser(*models.LoginUserRequest) (*models.UserDBResponse, error) {
	return nil, nil
}
