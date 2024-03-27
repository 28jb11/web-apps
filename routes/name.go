package routes

import (
	"html/template"
	"net/http"
)

func NameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	response := "Hello, " + template.HTMLEscapeString(name) + "!"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
