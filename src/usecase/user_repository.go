package usecase

import "clean-architecture-go-sample/entities"

type UserRepository interface {
	Store(entities.User) (int, error)
	FindById(int) (entities.User, error)
	FindAll() (entities.Users, error)
}
