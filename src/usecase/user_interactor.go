package usecase

import "clean-architecture-go-sample/entities"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u entities.User) (err error) {
	_, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Users() (user entities.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user entities.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}
