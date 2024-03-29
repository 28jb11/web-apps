package handlers

import (
	"html/template"
	"net/http"
)

type CalcHandler struct {
	Tmpl *template.Template
}

func (t CalcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := t.Tmpl.ExecuteTemplate(w, "calculator", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type DisplayHandler struct{}

func (h DisplayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	button := r.FormValue("button")
	if button == "" {
		button = "0" // Default display value
	}

	// Just send back the button value for the display update
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(button))
}
