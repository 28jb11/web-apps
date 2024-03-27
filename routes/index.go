package routes

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var tmpl = template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
