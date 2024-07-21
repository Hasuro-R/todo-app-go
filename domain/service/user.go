package service

import (
	"todo-app/domain/entity"
	"todo-app/domain/repository"
)

type UserService interface {
	GetAll() (users []entity.User, err error)
	GetByID(id int) (user entity.User, err error)
}

type UserServiceImpl struct {
	user repository.UserRepository
}

func NewUserRepository(user repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		user,
	}
}

func (srv *UserServiceImpl) GetAll() ([]entity.User, error) {
	users, err := srv.user.FindAll()
	if err != nil {
		return nil, err
	}

	return users, err
}

func (srv *UserServiceImpl) GetByID(id int) (entity.User, error) {
	user, err := srv.user.FindByID(id)
	if err != nil {
		return entity.User{}, err
	}

	return user, err
}
