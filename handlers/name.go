package handlers

import (
	"html/template"
	"net/http"
)

type NameHandler struct {
	Tmpl *template.Template
}

func (t NameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	greetingName := r.FormValue("name")

	type PageData struct {
		Title string
		Name  string
	}
	data := PageData{
		Title: "Name",
		Name:  greetingName,
	}

	err := t.Tmpl.ExecuteTemplate(w, "greeting", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
