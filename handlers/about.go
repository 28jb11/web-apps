package handlers

import (
	"html/template"
	"net/http"
)

type AboutHandler struct {
	Tmpl *template.Template
}

func (t AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// type PageData struct {
	// 	Title string
	// 	Name  string
	// }
	//
	// data := PageData{
	// 	Title: "About",
	// 	Name:  "",
	// }

	err := t.Tmpl.ExecuteTemplate(w, "about", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
