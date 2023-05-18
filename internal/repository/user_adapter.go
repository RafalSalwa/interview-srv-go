package repository

import (
	"context"
	"time"

	"github.com/RafalSalwa/interview-app-srv/internal/password"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"gorm.io/gorm"
)

type UserAdapter struct {
	DB *gorm.DB
}

func NewUserAdapter(db *gorm.DB) UserRepository {
	return &UserAdapter{DB: db}
}

func (r *UserAdapter) ById(ctx context.Context, id int64) (*models.UserDBModel, error) {
	var user models.UserDBModel
	r.DB.First(&user, "id = ?", id)
	return &user, nil
}

func (r *UserAdapter) ByLogin(ctx context.Context, user *models.LoginUserRequest) (*models.UserDBModel, error) {
	var dbUser models.UserDBModel

	r.DB.First(&dbUser, "username = ? OR email = ?", user.Username, user.Email)
	if password.CheckPasswordHash(user.Password, dbUser.Password) {
		return &dbUser, nil
	}
	return nil, nil
}

func (r *UserAdapter) UpdateLastLogin(ctx context.Context, u *models.UserDBModel) (*models.UserDBModel, error) {
	now := time.Now()
	r.DB.Model(u).Update("LastLogin", now)
	u.LastLogin = now
	return u, nil
}

func (r *UserAdapter) FindUserById(uid int64) (*models.UserDBModel, error) {

	return nil, nil
}
