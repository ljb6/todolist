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
	http.HandleFunc("/delete-task", handlers.DeleteTaskHandler)
	http.HandleFunc("/delete-all", handlers.DeleteAllTasksHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}