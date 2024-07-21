package handler

import (
	"net/http"
	"strconv"
	"todo-app/code"
	"todo-app/domain/entity"
	"todo-app/interface/response"
	"todo-app/usecase"

	"github.com/go-chi/chi/v5"
)

type UserHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
}

type UserHandlerImpl struct {
	user usecase.UserUseCase
}

func NewUserHandler(user usecase.UserUseCase) *UserHandlerImpl {
	return &UserHandlerImpl{
		user,
	}
}

func (h *UserHandlerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.user.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(w, r, users)
}

func (h *UserHandlerImpl) GetByID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		response.Error(w, r, code.Errorf(code.InvalidArgument, "invalid path parameter: %v", err))
		return
	}
	user, err := h.user.GetFindByID(userID)
	if err != nil {
		response.Error(w, r, err)
		return
	}
	response.Success(w, r, user)
}

func toUserResponse(user entity.User) response.User {
	return response.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func toUserListResponse(users []entity.User, total int) response.UserList {
	res := make([]response.User, 0, len(users))
	for i := range users {
		res = append(res, toUserResponse(users[i]))
	}
	return response.UserList{
		Users: res,
	}
}
