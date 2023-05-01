package services

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
)

type AuthServiceImpl struct {
	ctx context.Context
}

type AuthService interface {
	SignUpUser(*models.SignUpInput) (*models.UserDBResponse, error)
	SignInUser(*models.SignInInput) (*models.UserDBResponse, error)
	Token()
}

func NewAuthService(ctx context.Context) AuthService {
	return &AuthServiceImpl{ctx}
}

func (uc *AuthServiceImpl) Token() {

}

func (uc *AuthServiceImpl) SignUpUser(user *models.SignUpInput) (*models.UserDBResponse, error) {
	return nil, nil
}

func (uc *AuthServiceImpl) SignInUser(*models.SignInInput) (*models.UserDBResponse, error) {
	return nil, nil
}
