package handlers

import (
	"html/template"
	"net/http"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/contact.html",
	))
	tmpl.ExecuteTemplate(w, "layout", nil)
}
