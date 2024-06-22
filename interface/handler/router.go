package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type App struct{}

func (a *App) Router() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	return router
}
