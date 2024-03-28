package routes

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func NameHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "greeting.html")))

	name := r.FormValue("name")

	data := struct {
		Title    string
		Greeting string
	}{
		Title:    "Greeting",
		Greeting: "Hello, " + template.HTMLEscapeString(name) + "!",
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
