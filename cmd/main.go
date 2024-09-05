package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"

	"todo-app/interface/handler"
)

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "db:3306",
		DBName:               "todo_app",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	app := handler.NewApp(db)
	router := app.Router()

	http.ListenAndServe(":8080", router)
}
