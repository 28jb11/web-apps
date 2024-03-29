package handlers

import (
	"html/template"
	"net/http"
)

type BaseplateHandler struct {
	Tmpl *template.Template
}

func (t BaseplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type PageData struct {
		Title string
		Name  string
	}

	data := PageData{
		Title: "Index",
		Name:  "",
	}

	err := t.Tmpl.ExecuteTemplate(w, "baseplate", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
