package services

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
)

type AuthServiceImpl struct {
	ctx context.Context
}

type AuthService interface {
	SignUpUser(request *models.CreateUserRequest) (*models.UserDBResponse, error)
	SignInUser(request *models.LoginUserRequest) (*models.UserDBResponse, error)
	Token()
}

func NewAuthService(ctx context.Context) AuthService {
	return &AuthServiceImpl{ctx}
}

func (uc *AuthServiceImpl) Token() {

}

func (uc *AuthServiceImpl) SignUpUser(user *models.CreateUserRequest) (*models.UserDBResponse, error) {
	return nil, nil
}

func (uc *AuthServiceImpl) SignInUser(*models.LoginUserRequest) (*models.UserDBResponse, error) {
	return nil, nil
}
