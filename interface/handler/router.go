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
	user      UserHandler
	workspace WorkspaceHandler
}

func NewApp(db *sql.DB) *App {
	userRepository := persistent.NewUserInfrastructure(db)
	workspaceRepository := persistent.NewWorkspacePersistent(db)

	userService := service.NewUserRepository(userRepository)
	workspaceService := service.NewWorkspaceService(workspaceRepository)

	userUseCase := usecase.NewUserUseCase(userService)
	workspaceUseCase := usecase.NewWorkspaceUseCase(workspaceService)

	return &App{
		user:      NewUserHandler(userUseCase),
		workspace: NewWorkspaceHandler(workspaceUseCase),
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

	router.Route("/workspaces", func(r chi.Router) {
		r.Get("/", a.workspace.GetAll)
		r.Get("/{userID}", a.workspace.GetByUserID)
	})

	return router
}
