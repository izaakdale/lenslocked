package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/izaakdale/lenslocked/views"
)

func main() {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	mux.Get("/", homeHandler)
	mux.Get("/contact", contactHandler)
	mux.Get("/faq", faqHandler)
	mux.Get("/galleries/{id}", galleriesHandler)
	mux.NotFound(notFoundHandler)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe("localhost:3000", mux)
}

func executeTemplate(w http.ResponseWriter, path string) {
	tpl, err := views.Parse(path)
	if err != nil {
		log.Printf("error parsing template: %v", err)
		http.Error(w, "template parse error", http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "home.gohtml"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "contact.gohtml"))
}
func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

func galleriesHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "<h1>your id is: %s</h1", id)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Oops! Page not found", http.StatusNotFound)
}
