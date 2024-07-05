package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"

	"todo-app/interface/handler"
)

func main() {
	app := handler.App{}
	router := app.Router()
	cfg := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "go_todo_app",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	http.ListenAndServe(":8080", router)
}
