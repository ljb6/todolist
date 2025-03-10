package routes

import (
	"log"
	"net/http"

	"github.com/ljb6/todolist/internal/handlers"
)

func InitializeServer() {

	http.HandleFunc("/", handlers.PageHandler)
	http.HandleFunc("/submit-task", handlers.FormHandler)
	http.HandleFunc("/mark-as-done", handlers.MarkAsDoneHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}