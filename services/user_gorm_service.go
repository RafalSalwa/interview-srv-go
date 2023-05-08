package services

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/internal/repository"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
)

type UserService interface {
	Load(ctx context.Context, id string) (*models.UserDBModel, error)
}

func NewORMUserService(repository repository.UserRepository) UserService {
	return &userService{repository: repository}
}

type userService struct {
	repository repository.UserRepository
}

func (s *userService) Load(ctx context.Context, id string) (*models.UserDBModel, error) {
	res, err := s.repository.ById(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, err
}
