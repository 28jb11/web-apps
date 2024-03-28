package handlers

import (
	"html/template"
	"net/http"
)

type ContactHandler struct {
	Tmpl *template.Template
}

func (t ContactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type PageData struct {
		Title string
		Name  string
	}

	data := PageData{
		Title: "Contact",
		Name:  "",
	}

	err := t.Tmpl.ExecuteTemplate(w, "contact", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
