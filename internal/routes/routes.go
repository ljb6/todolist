package routes

import (
	"log"
	"net/http"

	"github.com/ljb6/todolist/internal/handlers"
)

func InitializeServer() {

	http.HandleFunc("/", handlers.PageHandler)
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}