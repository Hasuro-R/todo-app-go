package repository

import (
	"todo-app/domain/entity"
)

type UserRepository interface {
	FindAll() (users []entity.User, err error)
	FindByID(id int) (user entity.User, err error)
}
