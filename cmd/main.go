package main

import (
	"net/http"

	"todo-app/interface/handler"
)

func main() {
	app := handler.App{}
	router := app.Router()

	http.ListenAndServe(":8080", router)
}
