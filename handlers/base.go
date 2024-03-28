package handlers

import (
	"html/template"
	"net/http"
)

type BaseHandler struct {
	Tmpl *template.Template
}

func (t BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type PageData struct {
		Title string
		Name  string
	}

	data := PageData{
		Title: "Web Server",
		Name:  "",
	}

	err := t.Tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
