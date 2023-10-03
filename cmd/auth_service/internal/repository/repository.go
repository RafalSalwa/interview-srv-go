package repository

import (
	"context"

	"github.com/RafalSalwa/interview-app-srv/pkg/models"
)

type UserRepository interface {
	FindOne(ctx context.Context, user *models.UserDBModel) (*models.UserDBModel, error)
	FindAll(ctx context.Context, user *models.UserDBModel) ([]models.UserDBModel, error)
	Save(ctx context.Context, user models.UserDBModel) error
	Update(ctx context.Context, user models.UserDBModel) error
	Confirm(ctx context.Context, udb *models.UserDBModel) error
	GetOrCreate(ctx context.Context, id int64) (*models.UserDBModel, error)

	//UpdateLastLogin(ctx context.Context, u *models.UserDBModel) error
	//ByLogin(ctx context.Context, user *models.SignInUserRequest) (*models.UserDBModel, error)
	//FindUserById(uid int64) (*models.UserDBModel, error)

	//SignUp(ctx context.Context, user models.UserDBModel) error
	//GetConnection() *gorm.DB
}
