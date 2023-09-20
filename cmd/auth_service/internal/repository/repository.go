package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/RafalSalwa/interview-app-srv/pkg/models"
)

type UserRepository interface {
	SingUp(ctx context.Context, user *models.UserDBModel) error
	Load(user *models.UserDBModel) (*models.UserDBModel, error)
	ById(ctx context.Context, id int64) (*models.UserDBModel, error)
	ByLogin(ctx context.Context, user *models.SignInUserRequest) (*models.UserDBModel, error)
	ConfirmVerify(ctx context.Context, udb *models.UserDBModel) error
	UpdateLastLogin(ctx context.Context, u *models.UserDBModel) (*models.UserDBModel, error)
	FindUserById(uid int64) (*models.UserDBModel, error)
	GetConnection() *gorm.DB
}
