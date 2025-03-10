package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/ljb6/todolist/internal/database"
	"github.com/ljb6/todolist/internal/models"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := database.GetTasks(database.DB)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tasks)

	tmpl := template.Must(template.ParseFiles("web/main.html"))
	tmpl.Execute(w, struct{ Tasks []models.Task }{Tasks: tasks})
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	task := r.FormValue("task")

	formattedTime := time.Now().String()[11:16]

	database.AddTask(database.DB, models.Task{Text: task, Done: false, Time: formattedTime})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func MarkAsDoneHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	database.MarkTaskAsDone(database.DB, id)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	database.DeleteTask(database.DB, id)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}