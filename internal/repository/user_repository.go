package repository

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
)

type UserRepository interface {
	ById(ctx context.Context, id string) (*models.UserDBModel, error)
	ByLogin(ctx context.Context, id string) (*models.UserDBModel, error)
}
