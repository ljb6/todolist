package handlers

import (
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/ljb6/todolist/internal/database"
	"github.com/ljb6/todolist/internal/models"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/main.html"))
	tmpl.Execute(w, nil)
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	task := r.FormValue("task")
	
	db, err := database.OpenDatabase()
	if err != nil {
		log.Fatal(err)
	}

	database.AddTask(db, models.Task{Text: task, Done: false, Time: time.Now()})
}