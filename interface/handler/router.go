package handler

import (
	"database/sql"
	"net/http"
	"todo-app/domain/service"
	"todo-app/infrastructure/persistent"
	"todo-app/usecase"

	"github.com/go-chi/chi/v5"
)

type App struct {
	user UserHandler
}

func NewApp(db *sql.DB) *App {
	userRepository := persistent.NewUserInfrastructure(db)

	userService := service.NewUserRepository(userRepository)

	userUseCase := usecase.NewUserUseCase(userService)

	return &App{
		user: NewUserHandler(userUseCase),
	}
}

func (a *App) Router() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	router.Route("/users", func(r chi.Router) {
		r.Get("/", a.user.GetAll)
		r.Get("/{userID}", a.user.GetByID)
	})

	return router
}
