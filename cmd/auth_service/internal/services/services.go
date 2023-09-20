package services

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
)

type AuthService interface {
	SignUpUser(ctx context.Context, request *models.SignUpUserRequest) (*models.UserResponse, error)
	SignInUser(request *models.SignInUserRequest) (*models.UserResponse, error)
	GetVerificationKey(ctx context.Context, email string) (*models.UserResponse, error)
	Verify(ctx context.Context, vCode string) error
	Load(request *models.UserDBModel) (*models.UserResponse, error)
	Find(request *models.UserDBModel) (*models.UserResponse, error)
	FindUserById(uid int64) (*models.UserDBModel, error)
}
