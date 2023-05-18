package cqrs

import (
	"github.com/RafalSalwa/interview-app-srv/internal/repository"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
)

type Reader interface {
	Get(id int64) (*models.User, error)
	Search(query string) ([]*models.User, error)
	List() ([]*models.User, error)
}

// Writer user writer
type Writer interface {
	Create(e *models.User) (int64, error)
	Update(e *models.User) error
	Delete(id int64) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

type UserUsecase interface {
	Get(id int64) (*models.User, error)
	Search(query string) ([]*models.User, error)
	List() ([]*models.User, error)
	Create(email, password, firstName, lastName string) (int64, error)
	Update(e *models.User) error
	Delete(id int64) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func (u userUsecase) Get(id int64) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userUsecase) Search(query string) ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userUsecase) List() ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userUsecase) Create(email, password, firstName, lastName string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (u userUsecase) Update(e *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userUsecase) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: ur}
}
