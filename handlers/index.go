package handlers

import (
	"html/template"
	"net/http"
)

type IndexHandler struct {
	Tmpl *template.Template
}

func (t IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type PageData struct {
		Title string
		Name  string
	}

	data := PageData{
		Title: "Index",
		Name:  "",
	}

	err := t.Tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
