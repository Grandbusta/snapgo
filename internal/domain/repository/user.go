package repository

import "github.com/Grandbusta/snapgo/internal/domain/entity"

type IUserRepository interface {
	CreateUser(*entity.User) (*entity.User, error)
}

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) CreateUser(*entity.User) (*entity.User, error) {
	return &entity.User{}, nil
}
