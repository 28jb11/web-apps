package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/28jb11/calculator/handlers"
)

func main() {
	// Parse templates
	templates, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	// Serve static files from the "src" directory
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/src/", http.StripPrefix("/src/", fs))

	// Initialize handlers with templates.
	baseplateHandler := &handlers.BaseplateHandler{Tmpl: templates}
	indexHandler := &handlers.IndexHandler{Tmpl: templates}
	nameHandler := &handlers.NameHandler{Tmpl: templates}
	contactHandler := &handlers.ContactHandler{Tmpl: templates}
	aboutHandler := &handlers.AboutHandler{Tmpl: templates}
	calcHandler := &handlers.CalcHandler{Tmpl: templates}

	// Setup handlers.
	http.Handle("/", baseplateHandler)
	http.Handle("/index", indexHandler)
	http.Handle("/name", nameHandler)
	http.Handle("/contact", contactHandler)
	http.Handle("/about", aboutHandler)
	http.Handle("/calculate", calcHandler)
	http.Handle("/update-display", handlers.DisplayHandler{})
	// Start server.
	log.Println("Listening on :8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
