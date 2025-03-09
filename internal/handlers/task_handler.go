package handlers

import (
	"net/http"
	"text/template"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/main.html"))
	tmpl.Execute(w, nil)
}