package repository

import "clean-architecture-with-go/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
