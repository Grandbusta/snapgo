package repository

import "github.com/Grandbusta/snapgo/internal/domain/entity"

type UserRepository interface {
	CreateUser(*entity.User) (*entity.User, error)
}

func CreateUser(*entity.User) (*entity.User, error) {
	return &entity.User{}, nil
}
