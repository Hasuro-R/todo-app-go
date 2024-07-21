package usecase

import (
	"todo-app/domain/entity"
	"todo-app/domain/service"
)

type UserUseCase interface {
	GetAll() (users []entity.User, err error)
	GetFindByID(id int) (user entity.User, err error)
}

type UserUseCaseImpl struct {
	user service.UserService
}

func NewUserUseCase(user service.UserService) *UserUseCaseImpl {
	return &UserUseCaseImpl{
		user,
	}
}

func (uc *UserUseCaseImpl) GetAll() ([]entity.User, error) {
	users, err := uc.user.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *UserUseCaseImpl) GetFindByID(id int) (entity.User, error) {
	user, err := uc.user.GetByID(id)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
