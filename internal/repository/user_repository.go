package repository

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"time"
)

type UserRepository interface {
	ById(ctx context.Context, id int64) (*models.UserDBModel, error)
	ByLogin(ctx context.Context, user *models.LoginUserRequest) (*models.UserDBModel, error)
	UpdateLastLogin(ctx context.Context, uid int64) (time.Time, error)
	FindUserById(uid int64) (*models.UserDBModel, error)
}
