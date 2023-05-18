package repository

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
)

type UserRepository interface {
	ById(ctx context.Context, id int64) (*models.UserDBModel, error)
	ByLogin(ctx context.Context, user *models.LoginUserRequest) (*models.UserDBModel, error)
	UpdateLastLogin(ctx context.Context, u *models.UserDBModel) (*models.UserDBModel, error)
	FindUserById(uid int64) (*models.UserDBModel, error)
}
