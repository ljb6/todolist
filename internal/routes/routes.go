package routes

import (
	"log"
	"net/http"

	"github.com/ljb6/todolist/internal/handlers"
)

func InitializeServer() {

	http.HandleFunc("/", handlers.PageHandler)
	http.HandleFunc("/submit-task", handlers.FormHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}