package repository

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"gorm.io/gorm"
)

type UserAdapter struct {
	DB *gorm.DB
}

func NewUserAdapter(db *gorm.DB) UserRepository {
	return &UserAdapter{DB: db}
}

func (r *UserAdapter) ById(ctx context.Context, id string) (*models.UserDBModel, error) {
	var user models.UserDBModel
	r.DB.First(&user, "id = ?", id)
	return &user, nil
}

func (r *UserAdapter) ByLogin(ctx context.Context, id string) (*models.UserDBModel, error) {
	var user models.UserDBModel
	r.DB.First(&user, "username = ? OR email = ?", id)
	return &user, nil
}
